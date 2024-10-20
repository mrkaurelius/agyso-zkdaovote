# agyso-daovote zk

TODO

- use wallet 

## daovote-service

Proof generation request

### Generate Proof Request

POST /proof/vote

#### Req

```json
{
    "voteWeight": "", // Accounts vote weight, hex
    "masterPubKey": "", // Election result publishing key, hex
    "encryptedVote": "" // Current state of encrypted vote, hex
}
```

#### Res

```json
{
    "proof": "", //  , tbd convert to aligned layer format.
    "witness": "" // , tbd convert to aligned layer format.
}
```

## Rem

```json
{
  "verification_data_commitment": {
    "proof_commitment": [ 7, 163, 216, 197, 48, 138, 170, 147, 215, 56, 242, 139, 192, 11, 184, 97, 173, 241, 125, 91, 113, 167, 103, 17, 62, 58, 93, 67, 242, 147, 26, 255 ], "pub_input_commitment": [ 129, 76, 162, 227, 171, 199, 231, 231, 226, 213, 180, 135, 67, 197, 194, 248, 53, 57, 254, 66, 57, 74, 111, 127, 16, 234, 39, 219, 89, 17, 225, 201 ], "proving_system_aux_data_commitment": [ 176, 131, 203, 41, 179, 229, 93, 173, 78, 85, 57, 10, 14, 244, 8, 241, 93, 148, 248, 9, 154, 140, 202, 183, 187, 195, 178, 135, 149, 30, 101, 63 ], "proof_generator_addr": [ 243, 159, 214, 229, 26, 173, 136, 246, 244, 206, 106, 184, 130, 114, 121, 207, 255, 185, 34, 102 ] }, "batch_merkle_root": [ 140, 171, 215, 18, 65, 35, 254, 247, 148, 21, 44, 99, 242, 241, 172, 104, 185, 124, 7, 152, 125, 237, 123, 204, 148, 131, 33, 20, 34, 241, 51, 48 ], "batch_inclusion_proof": { "merkle_path": [ [ 40, 209, 128, 68, 90, 247, 191, 191, 16, 98, 193, 44, 227, 147, 51, 245, 241, 93, 70, 241, 30, 142, 156, 14, 191, 250, 141, 199, 248, 50, 231, 179 ], [ 22, 63, 192, 55, 198, 75, 113, 160, 152, 142, 134, 40, 60, 245, 135, 209, 255, 234, 211, 118, 35, 33, 52, 85, 236, 202, 50, 222, 74, 21, 103, 32 ] ] }, "index_in_batch": 1
}
```