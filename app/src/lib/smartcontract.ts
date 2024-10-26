export const web3Config = {
  abi: [
    {
      inputs: [
        {
          internalType: "address",
          name: "_alignedServiceManager",
          type: "address",
        },
        {
          internalType: "address",
          name: "_paymentServiceAddr",
          type: "address",
        },
      ],
      stateMutability: "nonpayable",
      type: "constructor",
    },
    {
      anonymous: false,
      inputs: [
        {
          indexed: false,
          internalType: "address",
          name: "voter",
          type: "address",
        },
        {
          indexed: false,
          internalType: "bytes32",
          name: "proofCommitment",
          type: "bytes32",
        },
      ],
      name: "BallotBox",
      type: "event",
    },
    {
      anonymous: false,
      inputs: [
        {
          indexed: false,
          internalType: "bool",
          name: "verified",
          type: "bool",
        },
      ],
      name: "Proof",
      type: "event",
    },
    {
      inputs: [],
      name: "PROPOSER_ADD",
      outputs: [
        {
          internalType: "address",
          name: "",
          type: "address",
        },
      ],
      stateMutability: "view",
      type: "function",
    },
    {
      inputs: [],
      name: "alignedServiceManager",
      outputs: [
        {
          internalType: "address",
          name: "",
          type: "address",
        },
      ],
      stateMutability: "view",
      type: "function",
    },
    {
      inputs: [
        {
          internalType: "bytes32",
          name: "proofCommitment",
          type: "bytes32",
        },
        {
          internalType: "bytes32",
          name: "pubInputCommitment",
          type: "bytes32",
        },
        {
          internalType: "bytes32",
          name: "provingSystemAuxDataCommitment",
          type: "bytes32",
        },
        {
          internalType: "bytes20",
          name: "proofGeneratorAddr",
          type: "bytes20",
        },
        {
          internalType: "bytes32",
          name: "batchMerkleRoot",
          type: "bytes32",
        },
        {
          internalType: "bytes",
          name: "merkleProof",
          type: "bytes",
        },
        {
          internalType: "uint256",
          name: "verificationDataBatchIndex",
          type: "uint256",
        },
        {
          internalType: "bytes",
          name: "pubInputBytes",
          type: "bytes",
        },
      ],
      name: "castZKVote",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    {
      inputs: [
        {
          internalType: "address",
          name: "",
          type: "address",
        },
      ],
      name: "encBallots",
      outputs: [
        {
          internalType: "bytes",
          name: "",
          type: "bytes",
        },
      ],
      stateMutability: "view",
      type: "function",
    },
    {
      inputs: [],
      name: "finishElection",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    {
      inputs: [],
      name: "getBallotBox",
      outputs: [
        {
          internalType: "bytes",
          name: "",
          type: "bytes",
        },
      ],
      stateMutability: "view",
      type: "function",
    },
    {
      inputs: [
        {
          internalType: "bytes",
          name: "publicInput",
          type: "bytes",
        },
        {
          internalType: "uint256",
          name: "start",
          type: "uint256",
        },
        {
          internalType: "uint256",
          name: "end",
          type: "uint256",
        },
      ],
      name: "parseAndSetBallots",
      outputs: [],
      stateMutability: "nonpayable",
      type: "function",
    },
    {
      inputs: [],
      name: "paymentServiceAddr",
      outputs: [
        {
          internalType: "address",
          name: "",
          type: "address",
        },
      ],
      stateMutability: "view",
      type: "function",
    },
  ],
};
