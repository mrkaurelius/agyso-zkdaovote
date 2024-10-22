import { HardhatRuntimeEnvironment } from "hardhat/types";

import { readFileSync } from "node:fs";

interface CallData {
  pub_input_hex: string;
  proof_commitment_hex: string;
  public_input_commitment_hex: string;
  proving_system_aux_data_commitment_hex: string;
  proof_generator_addr_hex: string;
  batch_merkle_root_hex: string;
  batch_inclusion_proof_hex: string;
  index_in_batch: number;
}

const getBalance = async (hre: HardhatRuntimeEnvironment) => {
  const ethers = hre.ethers;

  const signer = (await ethers.getSigners())[0];
  const signerAddress = await signer.getAddress();
  const balance = await ethers.provider.getBalance(signerAddress);

  console.log(balance);
};

const onChainVerify = async (hre: HardhatRuntimeEnvironment) => {
  const ethers = hre.ethers;

  const signer = (await ethers.getSigners())[0];

  const contractAddress = "0xa3656Ea3B745cEE21D2E761F222fb3eb08387342";

  const contract = await hre.ethers.getContractAt("AGYSODaoVoteValidator", contractAddress, signer);

  const jsonData = readFileSync("/var/tmp/agyso-daovote/proof/plonk/calldata.json", "utf-8");
  const callData: CallData = JSON.parse(jsonData);

  const contractResp = await contract.verifyBatchInclusion(
    callData.proof_commitment_hex,
    callData.public_input_commitment_hex,
    callData.proving_system_aux_data_commitment_hex,
    callData.proof_generator_addr_hex,
    callData.batch_merkle_root_hex,
    callData.batch_inclusion_proof_hex,
    callData.index_in_batch,
    callData.pub_input_hex
  );

  console.log(contractResp);
};

const aligned = async (args: any, hre: HardhatRuntimeEnvironment) => {
  await onChainVerify(hre);
};

export { aligned };
