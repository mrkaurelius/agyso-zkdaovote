import { HardhatUserConfig } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

const config: HardhatUserConfig = {
  solidity: "0.8.27",
  networks: {
    holesky: {
      url: "https://ethereum-holesky-rpc.publicnode.com",
      accounts: ["63dc97fe651de68a37a0fe8b2c28c5e56fc4f699d3d352bf0be72210017febe4"]
    }
  
  }
};

// TODO add holesky network

export default config;