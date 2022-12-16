use std::time::Duration;

use abscissa_core::{
    tracing::{debug, error, info, log::warn},
    Application,
};
use gravity_bridge::gravity_proto::{
    cosmos_sdk_proto::cosmos::base::tendermint::v1beta1::GetLatestBlockRequest,
    gravity::DelegateKeysByOrchestratorRequest,
};
use somm_proto::{
    cork::ScheduledCorkProposal, cosmos_sdk_proto::cosmos::gov::v1beta1::QueryProposalRequest,
};
use steward_proto::steward::{governance_call::Call, GovernanceCall};
use tokio::task::JoinHandle;
use tonic::{transport::Channel, Code};

use crate::{
    cellars::{aave_v2_stablecoin, cellar_v1},
    config::DELEGATE_ADDRESS,
    cork::schedule_cork,
    error::{Error, ErrorKind},
    prelude::APP,
};

use super::{client::CorkQueryClient, id_hash};

type GovQueryClient = gravity_bridge::gravity_proto::cosmos_sdk_proto::cosmos::gov::v1beta1::query_client::QueryClient<Channel>;
type GravityQueryClient =
    gravity_bridge::gravity_proto::gravity::query_client::QueryClient<Channel>;
type TendermintQueryClient = gravity_bridge::gravity_proto::cosmos_sdk_proto::cosmos::base::tendermint::v1beta1::service_client::ServiceClient<Channel>;

pub async fn start_scheduled_cork_proposal_polling_thread() -> JoinHandle<()> {
    debug!("starting cork proposal polling thread");
    let config = APP.config();
    let mut state = ProposalThreadState::default();
    let query_period = Duration::from_secs(config.cork.proposal_poll_period);

    // get validator address that corresponds to delegate for cork scheduling confirmation
    let mut client = GravityQueryClient::connect(config.cosmos.grpc.clone())
        .await
        .expect("failed to connect gravity query client at startup");
    let request = DelegateKeysByOrchestratorRequest {
        orchestrator_address: DELEGATE_ADDRESS.to_string(),
    };
    state.validator_address = client
        .delegate_keys_by_orchestrator(request)
        .await
        .expect("failed to get validator address at startup.")
        .into_inner()
        .validator_address;

    tokio::spawn(async move {
        let mut fail_count = 0;
        loop {
            tokio::time::sleep(query_period).await;

            debug!("polling for new approved scheduled cork proposals");
            state
                .update_last_observed_height(config.cosmos.grpc.clone())
                .await;
            if let Err(err) = poll_approved_cork_proposals(&mut state).await {
                fail_count += 1;
                error!(
                    "proposal {} failed to process {} time(s). latest error: {}",
                    state.last_processed_proposal_id + 1,
                    fail_count,
                    err
                );
            } else {
                fail_count = 0;
            }
        }
    })
}

#[derive(Debug, Default)]
struct ProposalThreadState {
    last_observed_height: u64,
    last_processed_proposal_id: u64,
    validator_address: String,
}

impl ProposalThreadState {
    async fn update_last_observed_height(&mut self, endpoint: String) {
        match TendermintQueryClient::connect(endpoint).await {
            Ok(mut client) => {
                match client.get_latest_block(GetLatestBlockRequest {}).await {
                    Ok(r) => {
                        let r = r.into_inner();
                        if r.block.is_some() {
                            let block = r.block.unwrap();
                            if block.header.is_some() {
                                self.last_observed_height = block.header.unwrap().height as u64;
                            }
                        }
                    }
                    Err(err) => warn!("failed to query latest block height. this could eventually result in degraded scheduled cork proposal query performance: {}", err),
                };
            }
            Err(err) => warn!("failed to query latest block height. this could eventually result in degraded scheduled cork proposal query performance: {}", err),
        };
    }

    fn increment_proposal_id(&mut self) {
        self.last_processed_proposal_id += 1;
    }
}

async fn poll_approved_cork_proposals(state: &mut ProposalThreadState) -> Result<(), Error> {
    let config = APP.config();
    let mut gov_client = GovQueryClient::connect(config.cosmos.grpc.clone()).await?;

    loop {
        // Proposal IDs start at 1, so this should be ok even for the first query after startup.
        let proposal_id = state.last_processed_proposal_id + 1;
        debug!("querying proposal {}", proposal_id);
        let proposal = match gov_client
            .proposal(tonic::Request::new(QueryProposalRequest { proposal_id }))
            .await
        {
            Ok(r) => r.into_inner().proposal,
            Err(err) => {
                if err.code() == Code::NotFound {
                    info!(
                        "no new proposals. last processed proposal ID: {}",
                        proposal_id
                    );

                    break;
                } else {
                    return Err(proposal_processing_error(format!(
                        "error querying proposal {}: {}",
                        proposal_id, err
                    )));
                }
            }
        };

        // Unsure if this can ever happen but needs to be handled.
        if proposal.is_none() {
            error!(
                "proposal {} was None even though the query status code indicates it was found.",
                proposal_id
            );
            state.increment_proposal_id();
            continue;
        }

        let proposal = proposal.unwrap();
        let content = match proposal.content {
            Some(c) => c,
            None => {
                warn!(
                    "ignoring proposal of ID {} because of empty content",
                    proposal.proposal_id
                );
                state.increment_proposal_id();
                continue;
            }
        };
        if content.type_url != "/cork.v2.ScheduledCorkProposal" {
            debug!("proposal {} not a ScheduledCorkProposal", proposal_id);
            state.increment_proposal_id();
            continue;
        }

        info!("processing scheduled cork proposal of ID {}", proposal_id);
        let cork_proposal: ScheduledCorkProposal =
            match prost::Message::decode(content.value.as_slice()) {
                Ok(c) => c,
                Err(err) => {
                    error!(
                        "failed to decode ScheduledCorkProposal {}: {}",
                        proposal_id, err
                    );
                    state.increment_proposal_id();
                    continue;
                }
            };
        if cork_proposal.block_height <= state.last_observed_height {
            info!(
                "proposal {} block height {} has already passed.",
                proposal_id, cork_proposal.block_height
            );
            state.increment_proposal_id();
            continue;
        }

        let json = cork_proposal.contract_call_proto_json;
        let cellar_id = cork_proposal.target_contract_address;
        let block_height = cork_proposal.block_height;
        debug!(
            "proposal {} contract call proto JSON: {}",
            proposal_id, json
        );
        let governance_call = match serde_json::from_str::<GovernanceCall>(&json) {
            Ok(c) => c,
            Err(err) => {
                error!("failed to decode GovernanceCall JSON {}: {}", json, err);
                state.increment_proposal_id();
                continue;
            }
        };
        if governance_call.call.is_none() {
            warn!(
                "governance call for proposal {} is empty and will be ignored: {:?}",
                proposal_id, governance_call
            );
            state.increment_proposal_id();
            continue;
        }
        let encoded_call: Vec<u8> = match governance_call.call.unwrap() {
            Call::AaveV2Stablecoin(data) => {
                if data.function.is_none() {
                    warn!(
                        "scheduled cork proposal {} call data contains no function data and will be ignored: {:?}",
                        proposal.proposal_id, data,
                    );
                    state.increment_proposal_id();
                    continue;
                }
                let function = data.function.unwrap();
                aave_v2_stablecoin::get_encoded_governance_call(
                    function,
                    &cellar_id,
                    proposal.proposal_id,
                )?
            }
            Call::CellarV1(data) => {
                if data.function.is_none() {
                    warn!(
                        "scheduled cork proposal {} call data contains no function data and will be ignored: {:?}",
                        proposal.proposal_id, data,
                    );
                    state.increment_proposal_id();
                    continue;
                }
                let function = data.function.unwrap();
                match cellar_v1::get_encoded_governance_call(
                    function,
                    &cellar_id,
                    proposal.proposal_id,
                ) {
                    Ok(d) => d,
                    // this is likely a bug in steward
                    Err(err) => {
                        error!(
                            "failed to get encoded governance call data for proposal {}: {}",
                            proposal.proposal_id, err
                        );
                        state.increment_proposal_id();
                        continue;
                    }
                }
            }
        };

        debug!(
            "proposal {} encoded call: {:?}",
            proposal.proposal_id, encoded_call
        );

        if let Err(err) = schedule_cork(&cellar_id, encoded_call.clone(), block_height).await {
            match confirm_sheduling(state, &cellar_id, encoded_call, block_height).await {
                Ok(confirmed) => {
                    if confirmed {
                        info!(
                            "call for cellar {} scheduled for block {} by proposal {}",
                            cellar_id, block_height, proposal.proposal_id
                        );
                        state.increment_proposal_id();
                        continue;
                    }
                }
                Err(err) => {
                    return Err(proposal_processing_error(format!(
                        "failed to confirm cork scheduling for proposal {} reason: {}",
                        proposal.proposal_id, err
                    )))
                }
            };

            return Err(proposal_processing_error(format!(
                "failed to schedule cork for proposal {} reason: {}",
                proposal.proposal_id, err
            )));
        } else {
            info!(
                "call for cellar {} scheduled for block {} by proposal {}",
                cellar_id, block_height, proposal.proposal_id
            );
            state.increment_proposal_id();
        }
    }

    Ok(())
}

async fn confirm_sheduling(
    state: &mut ProposalThreadState,
    cellar_id: &str,
    encoded_call: Vec<u8>,
    block_height: u64,
) -> Result<bool, Error> {
    let id = id_hash(block_height, cellar_id, encoded_call);
    let mut client = CorkQueryClient::new().await?;

    Ok(client
        .get_scheduled_corks_by_id(&id)
        .await?
        .into_inner()
        .corks
        .iter()
        .any(|c| c.validator == state.validator_address))
}

fn proposal_processing_error(message: String) -> Error {
    ErrorKind::ProposalProcessingError.context(message).into()
}