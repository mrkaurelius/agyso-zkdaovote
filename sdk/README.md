# agyso-zkdaovote SDK

DAO Vote aligned sdk tools.

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
proof_commitment_hex: c8e5930a4f7dc527599d386f9c809c3da8c6bc4b1da2eb506bd9265ef1b42e17
public_input_commitment_hex 695a27b7a1c0481ee771c8dc6f737d08839cfd7662bde404eb6690cdd82258ae
proving_system_aux_data_commitment_hex: b083cb29b3e55dad4e55390a0ef408f15d94f8099a8ccab7bbc3b287951e653f
proof_generator_addr_hex: 66f9664f97f2b50f62d13ea064982f936de76657
batch_merkle_root_hex: 977c4faa94afc8857065c1fe1739e3cfc9bd36513daaa393a0cbc6aec35a8b7e
batch_inclusion_proof: 42d052bb9ad42c82bb3c64a34eed7e94d0024ebe7b6e652ca6da56b2d7656003b6c4b3938babaf4aa90da44b04d9d702efe2958938939f56c5be3055ae7ce754
index_in_batch_hex: 3
```

docs

- [SDK Intro](https://docs.alignedlayer.com/guides/1_sdk_how_to)