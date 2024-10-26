import { HardhatUserConfig, task } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";

import { aligned } from "./tasks/aligned";

task("aligned", "aligned task", async (taskArgs, hre) => {
  await aligned(taskArgs, hre);
});

const config: HardhatUserConfig = {
  solidity: "0.8.27",
  networks: {
    holesky: {
      url: "https://ethereum-holesky-rpc.publicnode.com",
      accounts: [
        "63dc97fe651de68a37a0fe8b2c28c5e56fc4f699d3d352bf0be72210017febe4",
        "c04071acdaa1832de5b702340ad37f3e14b02dde6e8a621cee23b0ed39089225",
      ],
    },
  },
};

export default config;
