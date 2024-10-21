# AGYSO Dao Vote Aligned SDK

Aligned sdk experiments.

aligned tool.

```sh
aligned get-user-balance --network holesky --rpc_url https://ethereum-holesky-rpc.publicnode.com --user_addr 0xEbc82ed802F6309F4708De2eD6dD97E7a905Bfdf

aligned deposit-to-batcher --network holesky --rpc_url https://ethereum-holesky-rpc.publicnode.com  --keystore_path ../../vault/agyso-aligned --amount 00.29ether 
# [2024-10-21T21:02:04Z INFO  aligned] Payment sent to the batcher successfully. Tx: 0xecd2043e95524fa60029ff90314a489df115aedaeb842eaf38a8f659fd47423c

aligned get-user-balance --network holesky  --rpc_url https://ethereum-holesky-rpc.publicnode.com  --user_addr 0xEbc82ed802F6309F4708De2eD6dD97E7a905Bfdf
# [2024-10-21T21:02:12Z INFO  aligned] User 0xebc8…bfdf has 0.290000000000000000 ether in the batcher
```

proof response parsing

```
proof_commitment_hex: 07a3d8c5308aaa93d738f28bc00bb861adf17d5b71a767113e3a5d43f2931aff
public_input_commitment_hex 814ca2e3abc7e7e7e2d5b48743c5c2f83539fe42394a6f7f10ea27db5911e1c9
proving_system_aux_data_commitment_hex: b083cb29b3e55dad4e55390a0ef408f15d94f8099a8ccab7bbc3b287951e653f
proof_generator_addr_hex: 66f9664f97f2b50f62d13ea064982f936de76657
batch_merkle_root_hex: 4a1ec77aae2ab6f7aa427d1c890ac9e70ee651c2e33f3306bebe9c511588fd50
batch_inclusion_proof: 00a96779b5d2007e66e5795356f1c2dfe1ac6b224812661de450f7f163ab8a032c0581f24bb6168ad7dcdfe50d61201abd6c6c1c4a9acec3c1d6cc5747ae777fc7f817ce6a46700e040612163138a0db0a59b49040b6bf45a4a0a38be5582273
index_in_batch_hex: 4
```