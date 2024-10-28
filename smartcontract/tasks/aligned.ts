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

const contractAddress = "0x2DF98167F436d676B57614892Fb490D749b9c43b";

const getBalance = async (hre: HardhatRuntimeEnvironment) => {
  const ethers = hre.ethers;

  const signer = (await ethers.getSigners())[0];
  const signerAddress = await signer.getAddress();
  const balance = await ethers.provider.getBalance(signerAddress);

  console.log(balance);
};

const castZKVote = async (hre: HardhatRuntimeEnvironment) => {
  const ethers = hre.ethers;
  const signer = (await ethers.getSigners())[0];
  const contract = await hre.ethers.getContractAt("AGYSODaoVoteValidator", contractAddress, signer);

  const jsonData = readFileSync("/var/tmp/agyso-daovote/proof/plonk/calldata.json", "utf-8");
  const callData: CallData = JSON.parse(jsonData);

  const contractResp = await contract.castZKVote(
    callData.proof_commitment_hex,
    callData.public_input_commitment_hex,
    callData.proving_system_aux_data_commitment_hex,
    callData.proof_generator_addr_hex,
    callData.batch_merkle_root_hex,
    callData.batch_inclusion_proof_hex,
    callData.index_in_batch,
    callData.pub_input_hex
  );

  const receipt = await contractResp.wait(1);
  console.log(receipt);
};

const getBallotBox = async (hre: HardhatRuntimeEnvironment) => {
  const ethers = hre.ethers;
  const signer = (await ethers.getSigners())[0];
  const contract = await hre.ethers.getContractAt("AGYSODaoVoteValidator", contractAddress, signer);

  const contractResp = await contract.getBallotBox();
  console.log(contractResp);
};

const finishElection = async (hre: HardhatRuntimeEnvironment) => {
  const ethers = hre.ethers;
  const proposerSigner = (await ethers.getSigners())[1];
  const contract = await hre.ethers.getContractAt("AGYSODaoVoteValidator", contractAddress, proposerSigner);

  const contractResp = await contract.finishElection();

  const receipt = await contractResp.wait(1);
  console.log(receipt);
};

const createAccount = async (hre: HardhatRuntimeEnvironment) => {
  const wallet = hre.ethers.Wallet.createRandom();
  console.log(`address: ${wallet.address}\nprivate key: ${wallet.privateKey}`);
};

const aligned = async (args: any, hre: HardhatRuntimeEnvironment) => {
  // await createAccount(hre);

  // await finishElection(hre);
  await getBallotBox(hre);

  // await castZKVote(hre);
  // await getBallotBox(hre);
  // await finishElection(hre);
  // await getBallotBox(hre);
};

export { aligned };
