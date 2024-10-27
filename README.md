# agyso-zkdaovote

Privacy preserving DAO voting using **AlignedLayer** and PLONK proofs.

## Deployments

Network: Ethereum Holesky  
[AGYSO DAO Vote Validator Contract: 0x2DF98167F436d676B57614892Fb490D749b9c43b](https://holesky.etherscan.io/address/0x2df98167f436d676b57614892fb490d749b9c43b)

## Protocol

TL;DR
A blockchain-based voting system designed to preserve voter privacy and ensure accurate vote weighting based on each participant’s token balance (vote power). The system achieves confidentiality by storing encrypted votes on-chain, leveraging homomorphic encryption for private vote aggregation, and using zero-knowledge proofs (ZKPs) for verification. 

### Homomorphic ElGamal Encryption
In Protocol, all ciphertexts proceed on ElGamal Homomorphic Encryption. First, the person who initiates the vote generates a private-public key pair and shares the public key. Voters then cast their votes using this public key. Votes are collected by homomorphic addition and decrypted when the voting process is finished. 
The ElGamal encryption implementation outside and inside the circuit is available in the  `zk/zk` folder. Reduced form of Twisted Edwards]of [BN254](https://iden3-docs.readthedocs.io/en/latest/_downloads/33717d75ab84e11313cc0d8a090b636f/Baby-Jubjub.pdf) was used as the elliptic curve.

Assume $(x, Y) = (x,\hspace{0.1in} x\cdot G)$ is a private-public key pair. The encryption and decryption as follows:

$$ Enc(m,Y) = (C_1, \hspace{0.1in} C_2) = (r \cdot  G, \hspace{0.1in} m \cdot G + r \cdot Y ) $$

$$ Dec(C_1,C_2, x ) = (C_2 - x \cdot C_1) = (m \cdot G)$$

After brute-force (or [using Shanks Algorithm](https://www.mat.uniroma2.it/~geatti/HCMC2023/Lecture4.pdf)) we can get message m. 

Let's say there is another encryption with $Y$

$$ Enc(m',Y) = (C'_1, \hspace{0.1in} C'_2) = (r' \cdot  G, \hspace{0.1in} m' \cdot G + r' \cdot Y ) $$

We can add these two encryptions.

$$ (C_1, \hspace{0.1in} C_2)+ (C'_1, \hspace{0.1in} C'_2) = [(r' + r) \cdot  G, \hspace{0.1in} (m'+m) \cdot G + (r' +r) \cdot Y ] $$


Then, by decrypting it, we obtain the sum of the two plaintexts $m' + m$.


### Zero-Knowledge Proof (ZKP)

Plonk on BN254 was used in the project. Since the protocol is based on encrypted votes, it must be proven that the encrypted votes are created properly. ZKP verify the following:

1. **Vote Power Validation**: ZKP confirms the encrypted votes accurately represents the voter’s token-based power (No more votes were given than token power) .
2. **Non-Negative Vote Check**: ZKP ensures no negative values are encrypted.
3. **Homomorphic Summation Correctness**: ZKP confirms that homomorphic encryption and summation were correctly applied, enabling accurate aggregation without revealing individual votes.

Gnark was used for proof generation and its circuit structure is as follows. (The number of ballots is fixed at 4.)

 ```
type CircuitMain struct {
	VoteWeight   frontend.Variable    `gnark:",public"` //
	MasterPubKey twistededwards.Point `gnark:",public"`
	Vote         [4]frontend.Variable
	Randoms      [4]frontend.Variable
	EncVoteOld   VotesCircuit `gnark:",public"`
	EncVoteNew   VotesCircuit `gnark:",public"`
}
 ```

* **VoteWeight** represents vote power.
* **MasterPubKey** represents encryption public key.
* **Vote** 
represents votes.
* **Randoms** represents randoms used in encryption.
* **EncVoteOld**  represents the encrypted votes to be exchanged in the chain.
* **EncVoteNew**  represents new encrypted votes added to EncVoteOld.

## Running agyso-daovote

Tested with Debian 11, Ubuntu 22, go1.23, rust 1.82.0

0. `git clone --recursive https://github.com/mrkaurelius/agyso-zkdaovote`
1. Copy `vault` to `/var/tmp/agyso-daovote/vault` TODO automation
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

aligned_layer verison: `v0.9.2`

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
