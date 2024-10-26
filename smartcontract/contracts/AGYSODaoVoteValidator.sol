// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.12;

contract AGYSODaoVoteValidator {
    address public constant PROPOSER_ADDRESS = 0xbb9aE6C54376DEe7a3e41BCeFEAb456866ab25a7;

    mapping(address => bytes) public encBallots;

    address public alignedServiceManager;
    address public paymentServiceAddr;

    event Proof(bool verified);
    event BallotBox(address voter, bytes32 proofCommitment);

    constructor(address _alignedServiceManager, address _paymentServiceAddr) {
        alignedServiceManager = _alignedServiceManager;
        paymentServiceAddr = _paymentServiceAddr;
    }

    function setBallotBox(bytes memory encVotes) private {
        encBallots[PROPOSER_ADDRESS] = encVotes;
    }

    function getBallotBox() public view returns (bytes memory) {
        return encBallots[PROPOSER_ADDRESS];
    }

    function finishElection() public {
        require(msg.sender == PROPOSER_ADDRESS);
        setBallotBox(bytes(""));
    }

    function castZKVote(
        bytes32 proofCommitment,
        bytes32 pubInputCommitment,
        bytes32 provingSystemAuxDataCommitment,
        bytes20 proofGeneratorAddr,
        bytes32 batchMerkleRoot,
        bytes memory merkleProof,
        uint256 verificationDataBatchIndex,
        bytes memory pubInputBytes
    ) public {
        require(pubInputCommitment == keccak256(abi.encodePacked(pubInputBytes)), "pubinp comm. don't match");

        (bool callWasSuccessful, bytes memory proofIsIncluded) = alignedServiceManager.staticcall(
            abi.encodeWithSignature(
                "verifyBatchInclusion(bytes32,bytes32,bytes32,bytes20,bytes32,bytes,uint256,address)",
                proofCommitment,
                pubInputCommitment,
                provingSystemAuxDataCommitment,
                proofGeneratorAddr,
                batchMerkleRoot,
                merkleProof,
                verificationDataBatchIndex,
                paymentServiceAddr
            )
        );

        require(callWasSuccessful, "static_call failed");

        bool isVerified = abi.decode(proofIsIncluded, (bool));
        require(isVerified, "on chain verification failed");

        parseAndSetBallots(pubInputBytes, 1240 / 2, (1240 + 16 * 64) / 2);

        emit Proof(isVerified);
        emit BallotBox(msg.sender, proofCommitment);
    }

    function parseAndSetBallots(bytes memory publicInput, uint256 start, uint256 end) public {
        bytes memory encVotes = new bytes(512);

        uint j = 0;
        for (uint256 i = start; i < end; i++) {
            encVotes[j++] = publicInput[i];
        }

        setBallotBox(encVotes);
    }
}
