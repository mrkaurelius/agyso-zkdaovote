import { Contract } from "ethers";

import { web3Config } from "./smartcontract";
import { BrowserProvider } from "ethers";

const contractAddress = "0x2DF98167F436d676B57614892Fb490D749b9c43b";
const provider = new BrowserProvider(window.ethereum);

const signer = await provider.getSigner();

const signerContract = new Contract(contractAddress, web3Config.abi, signer);

export const getAGYSOContract = () => {
  return signerContract;
};
