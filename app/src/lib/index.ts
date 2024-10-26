import { Contract } from "ethers";

import { web3Config } from "./smartcontract";
import { BrowserProvider } from "ethers";

const contractAddress = "0x3fb7950daC347a4C80081a1fB2E9D50e17CB255f";
const provider = new BrowserProvider(window.ethereum);

const signer = await provider.getSigner();

const signerContract = new Contract(contractAddress, web3Config.abi, signer);

export const getAGYSOContract = () => {
  return signerContract;
};
