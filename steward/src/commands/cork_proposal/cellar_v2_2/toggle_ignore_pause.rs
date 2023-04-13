use crate::{application::APP, cellars, commands::cork_proposal::print_proposal, prelude::*};
use abscissa_core::{clap::Parser, Command, Runnable};
use steward_proto::steward::{
    cellar_v2dot2_governance::{Function, ToggleIgnorePause},
    governance_call::Call,
    CellarV2dot2Governance, GovernanceCall,
};

/// Shutdown subcommand
#[derive(Command, Debug, Parser)]
#[clap(
    long_about = "DESCRIPTION\n\nCalls toggleIgnorePause() on the target V2.2 cellar contract at the specified block height.\nFor more information see https://github.com/PeggyJV/cellar-contracts/blob/main/src/base/Cellar.sol"
)]
pub struct ToggleIgnorePauseCmd {
    /// Target contract for scheduled cork.
    #[clap(short, long)]
    cellar_id: String,

    /// Block height to schedule cork.
    #[clap(short, long)]
    block_height: u64,

    /// Whether to ignore the pause
    #[clap(short, long)]
    ignore_pause: bool,

    /// Only print JSON output, omitting explanatory text
    #[clap(short, long)]
    quiet: bool,
}

impl Runnable for ToggleIgnorePauseCmd {
    fn run(&self) {
        abscissa_tokio::run_with_actix(&APP, async {
            cellars::validate_cellar_id(&self.cellar_id)
                .await
                .unwrap_or_else(|err| {
                    status_err!("invalid cellar ID: {}", err);
                    std::process::exit(1);
                });

            let governance_call = GovernanceCall {
                call: Some(Call::CellarV2dot2Governance(CellarV2dot2Governance {
                    function: Some(Function::ToggleIgnorePause(ToggleIgnorePause {
                        ignore_pause: self.ignore_pause,
                    })),
                })),
            };

            print_proposal(
                self.block_height,
                self.cellar_id.clone(),
                governance_call,
                self.quiet,
            )
        })
        .unwrap_or_else(|e| {
            status_err!("executor exited with error: {}", e);
            std::process::exit(1);
        });
    }
}
