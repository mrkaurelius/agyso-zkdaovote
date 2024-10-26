# agyso-zkdaovote

Privacy preserving DAO voting using aligned network and PLONK proofs.

## Deployments

Network: Ethereum Holesky.
Contract: TODO

## Running agyso-daovote

Tested with Debian 11, Ubuntu 22, go1.23, rust 1.82.0

0. `git clone --recursive https://github.com/mrkaurelius/agyso-zkdaovote`
1. Copy `vault` to `/var/tmp/agyso-daovote/vault`
2. Build `sdk/daovote-rs` and add to PATH
3. Initialise circuit and protocol with `cd ./zk && make cli`
4. Start `zk/cmd/daovote-zk-service` service which generates proof and submit proofs to aligned layer
5. Start dapp frontend with `cd ./app && npm run`
6. profit 

## Directories

### zk

- `zk/cmd/zkdaovote-cli` handles protocol and circuit initialisation.
- `zk/cmd/daovote-zk-service` generates proof, submit proof using `sdk/daovote-rs` (aligned sdk)

### sdk

TODO

### smartcontract

TODO

### app

TODO

## TODO 

- daovote-zk-cli
    - ~~initialise circiut. generate and persist ccs, pk, vk.~~
- daovote-zk-service
    - Generate proof
    - Submit proof to aligned
    - Return aligned verification data to frontend
- daovote-rs: 
    ~~generate json in daovote-rs and put to proper place for aligned.ts usage~~
- validator-contract
    - ~~Submit proof using own address~~
    - ~~On chain zk proof verification~~
- app
    - ...
- docs
    - Architecture diagram

## Aligned Docs

- [Validating public input](https://github.com/yetanotherco/aligned_layer/tree/testnet/examples/validating-public-input)
- [Gnark guide](https://github.com/yetanotherco/aligned_layer/blob/testnet/docs/3_guides/3.2_generate_gnark_proof.md)
- [Hackathon](https://mirror.xyz/0x7794D1c55568270A81D8Bf39e1bcE96BEaC10901/_ia8GvSKS6bxU7YV8otdlIomtqWgSLef-lVl887O86U)
- [Hackathon Judging Criteria](https://mirror.xyz/0x7794D1c55568270A81D8Bf39e1bcE96BEaC10901/JnG4agqhW0oiskZJgcFdi9SLKvqkTBrbXkuk1nT6lxk)
