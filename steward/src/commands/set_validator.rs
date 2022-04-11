use std::{convert::TryFrom, path, sync::Arc, time::Duration};

use abscissa_core::{clap::Parser, Command, Runnable};
use ethers::prelude::*;
use signatory::FsKeyStore;

use crate::{allocation, cellars::uniswapv3::UniswapV3CellarState, gas::CellarGas, prelude::*};

/// Set-validator subcommand
#[derive(Command, Debug, Parser)]
pub struct SetValidatorCmd {
    /// Validator Address
    #[clap(short = 'v', long)]
    validator: H160,

    /// Boolean, true if you want to set validator and false if you don't want to.
    #[clap(short = 'u', long)]
    value: bool,

    /// Cellar Address
    cellar_address: H160,
}

impl Runnable for SetValidatorCmd {
    fn run(&self) {
        let config = APP.config();
        let cellar = allocation::get_cellar(self.cellar_address).unwrap();

        let keystore = path::Path::new(&config.keystore);
        let keystore = FsKeyStore::create_or_open(keystore).expect("Could not open keystore");

        let name = &config
            .keys
            .delegate_key
            .parse()
            .expect("Could not parse name");
        let key = keystore.load(name).expect("Could not load key");

        let key = key
            .to_pem()
            .parse::<k256::elliptic_curve::SecretKey<k256::Secp256k1>>()
            .expect("Could not parse key");

        let wallet: LocalWallet = Wallet::from(key);

        let eth_host = config.ethereum.rpc.clone();
        let _address = wallet.address();

        abscissa_tokio::run(&APP, async {
            let client = Provider::<Http>::try_from(eth_host)
                .unwrap()
                .interval(Duration::from_secs(3000u64));

            let client = SignerMiddleware::new(client, wallet);
            let gas = CellarGas::etherscan_standard().await.unwrap();

            // MyContract expects Arc, create with client
            let client = Arc::new(client);

            let mut contract_state =
                UniswapV3CellarState::new(cellar.cellar_address, client.clone());
            contract_state.gas_price = Some(gas);

            contract_state
                .set_validator(self.validator, self.value)
                .await
                .unwrap();
        })
        .unwrap_or_else(|e| {
            status_err!("executor exited with error: {}", e);
            std::process::exit(1);
        });
    }
}
