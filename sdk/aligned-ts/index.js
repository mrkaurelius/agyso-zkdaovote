import { getAligned, ProvingSystemId, Option } from "aligned-ts";

import { ethers } from "ethers";
import { verify } from "node:crypto";

import { readFileSync } from "node:fs";

// aligned submit \
// --proving_system GnarkPlonkBn254 \
// --proof "$PROOF_BASE_PATH/plonk.proof" \
// --public_input "$PROOF_BASE_PATH/plonk_pub_input.pub" \
// --vk "$PROOF_BASE_PATH/plonk.vk" \
// --batcher_url wss://batcher.alignedlayer.com \
// --network holesky \
// --rpc_url https://ethereum-holesky-rpc.publicnode.com

const main = async () => {
  const aligned = getAligned();

  const privateKey =
    "63dc97fe651de68a37a0fe8b2c28c5e56fc4f699d3d352bf0be72210017febe4";

  const wallet = new ethers.Wallet(privateKey);
  const address = await wallet.getAddress();
  console.log(address);

  // PROOF_BASE_PATH="/var/tmp/agyso-daovote/proof/plonk"
  const basePath = "/var/tmp/agyso-daovote/proof/plonk";

  const proof = readFileSync(`${basePath}/plonk.proof`, null);
  const publicInput = readFileSync(`${basePath}/plonk_pub_input.pub`, null);
  const vk = readFileSync(`${basePath}/plonk.vk`, null);
  const proofGeneratorAddress = address;

  const resp = await aligned.submit(
    {
      provingSystem: ProvingSystemId.Groth16Bn254,
      proof: proof,
      publicInput: Option.from(publicInput),
      verificationKey: Option.from(vk),
      vmProgramCode: Option.None,
      proofGeneratorAddress: proofGeneratorAddress,
    },
    wallet
  );

  console.log(resp);
};

main();
