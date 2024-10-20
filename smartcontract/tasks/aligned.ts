import { HardhatRuntimeEnvironment } from "hardhat/types";

const getBalance = async (hre: HardhatRuntimeEnvironment) => {
  const ethers = hre.ethers;

  const signer = (await ethers.getSigners())[0];
  const signerAddress = await signer.getAddress();
  const balance = await ethers.provider.getBalance(signerAddress);

  console.log(balance);
};

const fibonacciValidatorVerify = async (hre: HardhatRuntimeEnvironment) => {
  const ethers = hre.ethers;

  const signer = (await ethers.getSigners())[0];

  const constractAddres = "0xB1d333976A7f76aae952490cB888047b504DC1f0";

  const contract = await hre.ethers.getContractAt(
    "FibonacciValidator",
    constractAddres,
    signer
  );

  // TODO BURADA KALDIM
  // contract.verifyBatchInclusion()

  console.log(contract);
};

const submitProof = async () => {};

const aligned = async (args: any, hre: HardhatRuntimeEnvironment) => {
  await submitProof();
};

export { aligned };
