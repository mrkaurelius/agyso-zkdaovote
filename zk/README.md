# agyso-daovote


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