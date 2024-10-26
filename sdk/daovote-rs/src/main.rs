use std::env;
use std::fs::File;
use std::io::{Error, Write};
use std::path::PathBuf;
use std::str::FromStr;

use aligned_sdk::core::errors::{AlignedError, SubmitError};
use aligned_sdk::core::types::Network;
use aligned_sdk::core::types::{AlignedVerificationData, ProvingSystemId, VerificationData};
use aligned_sdk::sdk::{get_next_nonce, submit_and_wait_verification};

use env_logger::Env;
use log::error;
use log::info;

use ethers::signers::{LocalWallet, Signer};
use ethers::types::{Address, U256};
use ethers::utils::hex::{self};

use serde::{Deserialize, Serialize};

const BATCHER_URL: &str = "wss://batcher.alignedlayer.com";
const RPC_URL: &str = "https://ethereum-holesky-rpc.publicnode.com";

const PROOF_FILE_PATH: &str = "/var/tmp/agyso-daovote/proof/plonk/plonk.proof";
const PUB_INPUT_FILE_PATH: &str = "/var/tmp/agyso-daovote/proof/plonk/plonk_pub_input.pub";
const VK_FILE_PATH: &str = "/var/tmp/agyso-daovote/proof/plonk/plonk.vk";

const PROOF_GENERATOR_ADDRESS: &str = "0x66f9664f97F2b50F62D13eA064982f936dE76657";
const NETWORK: Network = Network::Holesky;

#[derive(Serialize, Deserialize)]
struct CallDataParams {
    pub_input_hex: String,
    proof_commitment_hex: String,
    public_input_commitment_hex: String,
    proving_system_aux_data_commitment_hex: String,
    proof_generator_addr_hex: String,
    batch_merkle_root_hex: String,
    batch_inclusion_proof_hex: String,
    index_in_batch: u64,
}

#[tokio::main]
async fn main() -> Result<(), SubmitError> {
    env_logger::Builder::from_env(Env::default().default_filter_or("debug")).init();

    let keystore_path = "/var/tmp/agyso-daovote/vault/agyso-aligned";
    let keystore_passwd = "";
    let wallet = LocalWallet::decrypt_keystore(keystore_path, keystore_passwd)
        .expect("Failed to decrypt keystore")
        .with_chain_id(17000u64);

    println!("wallet address: {}", wallet.address().to_string());

    let proof = read_file(PathBuf::from(PROOF_FILE_PATH)).unwrap_or_default();

    let pub_input = read_file(PathBuf::from(PUB_INPUT_FILE_PATH));

    let vk = read_file(PathBuf::from(VK_FILE_PATH));

    let mut pub_input_hex = hex::encode(pub_input.as_ref().unwrap());
    pub_input_hex.insert_str(0, "0x");

    info!("pub_input_hex: {}", pub_input_hex);

    let proof_generator_addr = Address::from_str(PROOF_GENERATOR_ADDRESS).unwrap();

    let verification_data = VerificationData {
        proving_system: ProvingSystemId::GnarkPlonkBn254,
        proof,
        pub_input,
        verification_key: vk,
        vm_program_code: None,
        proof_generator_addr,
    };

    // Set a fee of 0.1 Eth
    let max_fee = U256::from(5) * U256::from(100_000_000_000_000_000u128);

    let nonce = get_next_nonce(RPC_URL, wallet.address(), NETWORK)
        .await
        .expect("Failed to get next nonce");

    info!("nonce: {}", nonce.to_string());

    info!("Submitting vote proof to Aligned and waiting for verification...");
    let verification_data = submit_and_wait_verification(
        BATCHER_URL,
        RPC_URL,
        NETWORK,
        &verification_data,
        max_fee,
        wallet,
        nonce,
    )
    .await?;

    info!("Proof submitted to aligned. See the batch in the explorer:");

    let mut proof_commitment_hex = hex::encode(
        verification_data
            .verification_data_commitment
            .proof_commitment,
    );
    proof_commitment_hex.insert_str(0, "0x");
    info!("proof_commitment_hex: {}", proof_commitment_hex);

    let mut public_input_commitment_hex = hex::encode(
        verification_data
            .verification_data_commitment
            .pub_input_commitment,
    );
    public_input_commitment_hex.insert_str(0, "0x");
    info!(
        "public_input_commitment_hex {}",
        public_input_commitment_hex
    );

    let mut proving_system_aux_data_commitment_hex = hex::encode(
        verification_data
            .verification_data_commitment
            .proving_system_aux_data_commitment,
    );
    proving_system_aux_data_commitment_hex.insert_str(0, "0x");
    info!(
        "proving_system_aux_data_commitment_hex: {}",
        proving_system_aux_data_commitment_hex
    );

    // let PROOF_GENERATOR_ADDRESS
    let mut proof_generator_addr_hex = hex::encode(
        verification_data
            .verification_data_commitment
            .proof_generator_addr,
    );
    proof_generator_addr_hex.insert_str(0, "0x");
    info!("proof_generator_addr_hex: {}", proof_generator_addr_hex);

    let mut batch_merkle_root_hex = hex::encode(verification_data.batch_merkle_root);
    batch_merkle_root_hex.insert_str(0, "0x");
    info!("batch_merkle_root_hex: {}", batch_merkle_root_hex);

    // !!! could concationation create a problem?
    let mut batch_inclusion_proof_hex =
        hex::encode(verification_data.batch_inclusion_proof.merkle_path.concat());
    batch_inclusion_proof_hex.insert_str(0, "0x");
    info!("batch_inclusion_proof: {}", batch_inclusion_proof_hex);

    let index_in_batch = verification_data.index_in_batch as u64;
    info!("index_in_batch_hex: {}", index_in_batch);

    println!(
        "https://explorer.alignedlayer.com/batches/0x{}",
        hex::encode(verification_data.batch_merkle_root)
    );

    // info!(
    //     "https://explorer.alignedlayer.com/batches/0x{}",
    //     hex::encode(verification_data.batch_merkle_root)
    // );

    let json_obj = CallDataParams {
        pub_input_hex,
        proof_commitment_hex,
        public_input_commitment_hex,
        proving_system_aux_data_commitment_hex,
        proof_generator_addr_hex,
        batch_merkle_root_hex,
        batch_inclusion_proof_hex,
        index_in_batch,
    };

    let json_output = serde_json::to_string_pretty(&json_obj).unwrap();
    info!("Serialized JSON: {}", json_output);

    match File::create("/var/tmp/agyso-daovote/proof/plonk/calldata.json") {
        Ok(mut file) => match file.write_all(json_output.as_bytes()) {
            Ok(_) => println!(""),
            Err(e) => error!("Error writing to file: {}", e),
        },
        Err(e) => {
            error!("Error creating file: {}", e);
        }
    }

    Ok(())
}

fn read_file(file_name: PathBuf) -> Option<Vec<u8>> {
    std::fs::read(file_name).ok()
}

// fn save_response(
//     batch_inclusion_data_directory_path: PathBuf,
//     aligned_verification_data: &AlignedVerificationData,
// ) -> Result<(), SubmitError> {
//     std::fs::create_dir_all(&batch_inclusion_data_directory_path)
//         .map_err(|e| SubmitError::IoError(batch_inclusion_data_directory_path.clone(), e))?;

//     let batch_merkle_root = &hex::encode(aligned_verification_data.batch_merkle_root)[..8];
//     let batch_inclusion_data_file_name = batch_merkle_root.to_owned()
//         + "_"
//         + &aligned_verification_data.index_in_batch.to_string()
//         + ".json";

//     let batch_inclusion_data_path =
//         batch_inclusion_data_directory_path.join(batch_inclusion_data_file_name);

//     let data = serde_json::to_vec(&aligned_verification_data).unwrap();

//     let mut file = File::create(&batch_inclusion_data_path)
//         .map_err(|e| SubmitError::IoError(batch_inclusion_data_path.clone(), e))?;
//     file.write_all(data.as_slice())
//         .map_err(|e| SubmitError::IoError(batch_inclusion_data_path.clone(), e))?;

//     let current_dir = env::current_dir().expect("Failed to get current directory");

//     info!(
//         "Saved batch inclusion data to {:?}",
//         current_dir.join(batch_inclusion_data_path)
//     );

//     Ok(())
// }
