//! Time independent bollinger ranges
/// This is a Rust type for the JSON data from time independent bollinger ranges.
use crate::prelude::*;
use crate::{config::TokenInfo, cosmos_somm::send};
use chrono::DateTime;
use deep_space::address::Address;
use deep_space::private_key::PrivateKey as CosmosPrivateKey;
use deep_space::{Coin, Contact};
use ethers::prelude::*;
use futures::TryStreamExt;
use mongodb::{bson::doc, options::FindOptions};
use num_bigint::ToBigInt;
use num_rational::BigRational;
use serde::{Deserialize, Serialize};
use somm_proto::somm as proto;
use uniswap_v3_sdk::{Price, Token};

// Struct TimeRange for time independent bollinger ranges
#[derive(Serialize, Deserialize, Clone)]
pub struct TimeRange {
    pub time: Option<DateTime<chrono::Utc>>,
    pub previous_update: Option<DateTime<chrono::Utc>>,
    pub pair_id: U256,
    pub token_info: (TokenInfo, TokenInfo),
    pub weight_factor: u32,
    pub tick_weights: Vec<TickWeight>,
    pub tick_spacing: i32,
}

impl Default for TimeRange {
    fn default() -> Self {
        TimeRange {
            time: None,
            previous_update: None,
            pair_id: U256::zero(),
            tick_weights: Vec::new(),
            weight_factor: 100,
            token_info: (TokenInfo::default(), TokenInfo::default()),
            tick_spacing: 10,
        }
    }
}

/// Implementation for TimeRange field format
impl std::fmt::Debug for TimeRange {
    // Implement TimeRange field format for time, previous_update, pair_id and tick_weight
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let mut fields = f.debug_struct("TimeRange");
        fields
            .field("time", &self.time)
            .field("previous_update", &self.previous_update)
            .field("pair_id", &self.pair_id)
            .field("token_info_0", &self.token_info.0)
            .field("token_info_1", &self.token_info.1);
        for (i, tick) in self.tick_weights.iter().enumerate() {
            fields.field(&format!("tick_weight #:{}", i), tick);
        }
        fields.finish()
    }
}

/// Struct TickWeights for time independent bollinger ranges
#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct TickWeight {
    pub upper_bound: i32,
    pub lower_bound: i32,
    pub weight: u32,
}

impl TickWeight {
    pub fn valid(&self) -> bool {
        if self.upper_bound > self.lower_bound {
            true
        } else {
            false
        }
    }

    pub fn within(&self, tick: &i32) -> bool {
        if &self.lower_bound <= tick && tick <= &self.upper_bound {
            return true;
        }
        return false;
    }
}

impl std::cmp::PartialOrd<i32> for TickWeight {
    fn partial_cmp(&self, other: &i32) -> Option<std::cmp::Ordering> {
        if other < &self.lower_bound {
            Some(std::cmp::Ordering::Less)
        } else if other > &self.upper_bound {
            Some(std::cmp::Ordering::Greater)
        } else if self.within(other) {
            Some(std::cmp::Ordering::Equal)
        } else {
            None
        }
    }
}
impl std::cmp::PartialEq<i32> for TickWeight {
    fn eq(&self, other: &i32) -> bool {
        self.within(other)
    }
}

impl std::cmp::PartialOrd for TickWeight {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        if self.lower_bound >= other.upper_bound {
            Some(std::cmp::Ordering::Greater)
        } else if self.upper_bound <= other.lower_bound {
            Some(std::cmp::Ordering::Less)
        } else if self.upper_bound == other.upper_bound
            && self.lower_bound == other.lower_bound
            && self.weight == other.weight
        {
            Some(std::cmp::Ordering::Equal)
        } else {
            None
        }
    }
}

impl std::cmp::Ord for TickWeight {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        if self.lower_bound >= other.upper_bound {
            std::cmp::Ordering::Greater
        } else if self.upper_bound <= other.lower_bound {
            std::cmp::Ordering::Less
        } else {
            std::cmp::Ordering::Equal
        }
    }
}

impl std::cmp::Eq for TickWeight {}

impl std::cmp::PartialEq for TickWeight {
    fn eq(&self, other: &Self) -> bool {
        if self.lower_bound == other.lower_bound
            && self.upper_bound == other.upper_bound
            && self.weight == other.weight
        {
            true
        } else {
            false
        }
    }
}

#[derive(Serialize, Deserialize, Debug)]
pub struct MongoData {
    pub _id: mongodb::bson::Bson,
    pub created_timestamp: mongodb::bson::Bson,
    pub pair_id: ethers::prelude::U256,
    pub symbol: String,
    pub tick_weights: Vec<MongoTickWeights>,
}
#[derive(Serialize, Deserialize, Debug)]
pub struct MongoTickWeights {
    pub lower: mongodb::bson::Bson,
    pub upper: mongodb::bson::Bson,
    pub weight: mongodb::bson::Bson,
}

// Implement TimeRange for time independent bollinger ranges
impl TimeRange {
    // Instantiate TimeRange for toime independent bollinger ranges with fn new.
    pub fn new(
        time: Option<DateTime<chrono::Utc>>,
        previous_update: Option<DateTime<chrono::Utc>>,
        pair_id: U256,
        weight_factor: u32,
        tick_weights: Vec<TickWeight>,
        token_0_info: TokenInfo,
        token_1_info: TokenInfo,
        tick_spacing: i32,
    ) -> Self {
        TimeRange {
            time,
            previous_update,
            pair_id,
            weight_factor,
            tick_weights,
            token_info: (token_0_info, token_1_info),
            tick_spacing,
        }
    }

    pub async fn poll(&mut self, contact: &Contact, database: &mongodb::Database) {
        // Get a handle to a collection in the database.
        let collection = database.collection::<MongoData>("tick_range_predictions");

        let find_options = FindOptions::builder()
            .sort(doc! { "created_timestamp": -1 })
            .build();

        let mut sorted_predictions = collection.find(None, find_options).await.unwrap();

        if let Some(latest_prediction) = sorted_predictions.try_next().await.unwrap() {
            info!("Latest prediction: {:?}", latest_prediction);
            self.previous_update = self.time;
            self.time = Some(
                latest_prediction
                    .created_timestamp
                    .as_datetime()
                    .unwrap()
                    .to_chrono(),
            );
            self.pair_id = latest_prediction.pair_id;
            self.tick_weights.clear();
            for tick_weight in latest_prediction.tick_weights {
                let upper_tick = tick_weight.upper.as_i32().unwrap();
                let upper_tick = self.tick_spacing - (upper_tick % self.tick_spacing) + upper_tick; //Normalize tick to tick_spacing
                let lower_tick = tick_weight.lower.as_i32().unwrap();
                let lower_tick = lower_tick - (lower_tick % self.tick_spacing); //Normalize tick to tick_spacing
                let weight: u32 =
                    (self.weight_factor as f64 * tick_weight.weight.as_f64().unwrap()) as u32;
                let tick_weight = TickWeight {
                    upper_bound: upper_tick,
                    lower_bound: lower_tick,
                    weight: weight,
                };
                if !tick_weight.valid() {
                    status_err!("Invalid tick {:?}", tick_weight);
                } else {
                    self.align_then_push(tick_weight);
                }
            }
        }
        self.tick_weights.sort_by(|a, b| a.weight.cmp(&b.weight));
        self.tick_weights.reverse();

        if self.tick_weights.len() > 7 {
            self.tick_weights.sort_by(|a, b| a.weight.cmp(&b.weight));
            self.tick_weights.reverse();
            let largest_weights: Vec<TickWeight> =
                self.tick_weights.iter().cloned().take(7).collect();
            self.tick_weights = largest_weights;
        }

        self.tick_weights.sort();
        self.tick_weights.reverse();

        // TODO(Levi) needs to be initialized (should it be derived from the cosmos_key??)
        let delegate_cosmos_address =
            Address::from_bech32("cosmos1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqnrql8a".to_string())
                .unwrap();

        // TODO(Levi) needs to be initialized:
        let cosmos_key = CosmosPrivateKey::from_phrase("purse sure leg gap above pull rescue glass circle attract erupt can sail gasp shy clarify inflict anger sketch hobby scare mad reject where", "").unwrap();

        // TODO(Levi) needs to be initialized
        let fee = "100footoken".parse().unwrap();

        // TODO(Levi) needs to be initialized
        let cellar_id = "TODO".to_owned();

        send::send_allocation(
            contact,
            delegate_cosmos_address,
            cosmos_key,
            fee,
            cellar_id,
            vec![self.to_allocation()],
        )
        .await
        .unwrap();
    }

    fn align_then_push(&mut self, mut tick_weight: TickWeight) {
        for t in self.tick_weights.iter() {
            if tick_weight.upper_bound < t.upper_bound && tick_weight.upper_bound >= t.lower_bound {
                tick_weight.upper_bound = t.lower_bound;
            }
            if tick_weight.lower_bound > t.lower_bound && tick_weight.lower_bound <= t.upper_bound {
                tick_weight.lower_bound = t.upper_bound;
            }
        }
        self.tick_weights.push(tick_weight);
    }

    fn to_allocation(&self) -> proto::Allocation {
        let tick_range: Vec<proto::TickRange> = self
            .tick_weights
            .iter()
            .map(|tick_weight| proto::TickRange {
                upper: tick_weight.upper_bound as u64,
                lower: tick_weight.lower_bound as u64,
                weight: tick_weight.weight as u64,
            })
            .collect();

        proto::Allocation {
            cellar: Some(proto::Cellar {
                id: self.pair_id.to_string(),
                tick_ranges: self
                    .tick_weights
                    .iter()
                    .map(|tick_weight| proto::TickRange {
                        upper: tick_weight.upper_bound as u64,
                        lower: tick_weight.lower_bound as u64,
                        weight: tick_weight.weight as u64,
                    })
                    .collect(),
            }),
            salt: "".to_string(), //TODO: Add salt,
        }
    }
}

fn f64_unit_to_price_for_stables(price: f64, token_0: &TokenInfo, token_1: &TokenInfo) -> Price {
    Price {
        token_0: Token {
            symbol: token_0.symbol.clone(),
            address: token_0.address.to_string(),
        },
        token_1: Token {
            symbol: token_1.symbol.clone(),
            address: token_1.address.to_string(),
        },
        amount_0: (BigRational::from_float(price).unwrap()
            * (BigRational::from_integer(10.to_bigint().unwrap()).pow(token_1.decimals.into())))
        .to_integer(), //TODO this is providing correct tick bounds but doesn't make sense. We should be putting in token_0 decimals here.
        amount_1: (1 * (10i32.to_bigint().unwrap().pow(token_0.decimals.into()))),
    }
}

#[cfg(test)]
mod test {
    use super::*;
    use uniswap_v3_sdk::priceToTick;
    #[test]
    fn test_64_to_price_for_stables() {
        let mut token_0 = TokenInfo::default();
        let mut token_1 = TokenInfo::default();
        token_0.symbol = "ETH".to_owned();
        token_0.address = "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"
            .parse()
            .unwrap();
        token_1.symbol = "USDT".to_owned();
        token_1.address = "0xdAC17F958D2ee523a2206206994597C13D831ec7"
            .parse()
            .unwrap();
        token_0.decimals = 18;
        token_1.decimals = 6;

        let price = f64_unit_to_price_for_stables(2000.0, &token_0, &token_1);

        let tick = priceToTick(price);

        assert_eq!(tick, -200312);
    }
}
