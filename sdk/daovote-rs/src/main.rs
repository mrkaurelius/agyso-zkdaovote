use std::env;
use std::fs::File;
use std::io::Write;
use std::path::PathBuf;
use std::str::FromStr;

use aligned_sdk::core::errors::{AlignedError, SubmitError};
use aligned_sdk::core::types::Network;
use aligned_sdk::core::types::{AlignedVerificationData, ProvingSystemId, VerificationData};
use aligned_sdk::sdk::{get_next_nonce, submit_and_wait_verification};

use clap::Parser;
use env_logger::Env;
use ethers::signers::{LocalWallet, Signer};
use ethers::types::{Address, U256};
use ethers::utils::hex::{self, ToHex};
use log::info;

const BATCHER_URL: &str = "wss://batcher.alignedlayer.com";
// const RPC_URL: &str = "https://ethereum-holesky-rpc.publicnode.com";
const RPC_URL: &str = "https://ethereum-holesky-rpc.publicnode.com";

const PROOF_FILE_PATH: &str = "/var/tmp/agyso-daovote/proof/plonk/plonk.proof";
const PUB_INPUT_FILE_PATH: &str = "/var/tmp/agyso-daovote/proof/plonk/plonk_pub_input.pub";
const VK_FILE_PATH: &str = "/var/tmp/agyso-daovote/proof/plonk/plonk.vk";

const PROOF_GENERATOR_ADDRESS: &str = "0x66f9664f97F2b50F62D13eA064982f936dE76657";
const NETWORK: Network = Network::Holesky;

#[tokio::main]
async fn main() -> Result<(), SubmitError> {
    env_logger::Builder::from_env(Env::default().default_filter_or("info")).init();

    // Our private key
    // let AGYSO_PRIVATE_KEY  = "63dc97fe651de68a37a0fe8b2c28c5e56fc4f699d3d352bf0be72210017febe4";

    let anvil_private_key: &str =
        "2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6"; // Anvil address 9

    let wallet = LocalWallet::from_str(anvil_private_key)
        .expect("wallet creation error")
        .with_chain_id(17000u64);

    let address_str = wallet.address().to_string();

    println!("wallet address: {}", address_str);

    let proof = read_file(PathBuf::from(PROOF_FILE_PATH)).unwrap_or_default();

    let pub_input = read_file(PathBuf::from(PUB_INPUT_FILE_PATH));

    let vk = read_file(PathBuf::from(VK_FILE_PATH));

    let pub_input_hex = hex::encode(pub_input.as_ref().unwrap());

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
    let ver_data = submit_and_wait_verification(
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

    let proof_commitment_hex = hex::encode(ver_data.verification_data_commitment.proof_commitment);
    info!("proof_commitment_hex: {}", proof_commitment_hex);

    let public_input_commitment_hex =
        hex::encode(ver_data.verification_data_commitment.pub_input_commitment);
    info!(
        "public_input_commitment_hex {}",
        public_input_commitment_hex
    );

    let proving_system_aux_data_commitment_hex = hex::encode(
        ver_data
            .verification_data_commitment
            .proving_system_aux_data_commitment,
    );
    info!(
        "proving_system_aux_data_commitment_hex: {}",
        proving_system_aux_data_commitment_hex
    );

    // let PROOF_GENERATOR_ADDRESS
    let proof_generator_addr_hex =
        hex::encode(ver_data.verification_data_commitment.proof_generator_addr);
    info!("proof_generator_addr_hex: {}", proof_generator_addr_hex);

    let batch_merkle_root_hex = hex::encode(ver_data.batch_merkle_root);
    info!("batch_merkle_root_hex: {}", batch_merkle_root_hex);

    // TODO could concationation create a problem?
    let batch_inclusion_proof = hex::encode(ver_data.batch_inclusion_proof.merkle_path.concat());
    info!("batch_inclusion_proof: {}", batch_inclusion_proof);

    info!("index_in_batch_hex: {}", ver_data.index_in_batch);


    info!(
        "https://explorer.alignedlayer.com/batches/0x{}",
        hex::encode(ver_data.batch_merkle_root)
    );

    Ok(())
}

fn read_file(file_name: PathBuf) -> Option<Vec<u8>> {
    std::fs::read(file_name).ok()
}

fn save_response(
    batch_inclusion_data_directory_path: PathBuf,
    aligned_verification_data: &AlignedVerificationData,
) -> Result<(), SubmitError> {
    std::fs::create_dir_all(&batch_inclusion_data_directory_path)
        .map_err(|e| SubmitError::IoError(batch_inclusion_data_directory_path.clone(), e))?;

    let batch_merkle_root = &hex::encode(aligned_verification_data.batch_merkle_root)[..8];
    let batch_inclusion_data_file_name = batch_merkle_root.to_owned()
        + "_"
        + &aligned_verification_data.index_in_batch.to_string()
        + ".json";

    let batch_inclusion_data_path =
        batch_inclusion_data_directory_path.join(batch_inclusion_data_file_name);

    let data = serde_json::to_vec(&aligned_verification_data).unwrap();

    let mut file = File::create(&batch_inclusion_data_path)
        .map_err(|e| SubmitError::IoError(batch_inclusion_data_path.clone(), e))?;
    file.write_all(data.as_slice())
        .map_err(|e| SubmitError::IoError(batch_inclusion_data_path.clone(), e))?;

    let current_dir = env::current_dir().expect("Failed to get current directory");

    info!(
        "Saved batch inclusion data to {:?}",
        current_dir.join(batch_inclusion_data_path)
    );

    Ok(())
}
