// SPDX-License-Identifier: GPL-3.0	
pragma solidity ^0.8.12;

contract AGYSODaoVoteValidator {

    address public constant PROPOSER_ADD = 0xbb9aE6C54376DEe7a3e41BCeFEAb456866ab25a7;
 

    mapping(address => bytes) public encVotes;
    

    address public alignedServiceManager;
    address public paymentServiceAddr;

    event ProofValidation(bool verified);

    constructor(address _alignedServiceManager, address _paymentServiceAddr) {
        alignedServiceManager = _alignedServiceManager;
        paymentServiceAddr = _paymentServiceAddr;
    }



    function setVotes(bytes memory encVote)  {
        encVotes[PROPOSER_ADD] = encVote;
    }

    function getVotes() public view returns (bytes memory) {
        return encVotes[PROPOSER_ADD];
    }

    function finishElection() public  {
        require(msg.sender == PROPOSER_ADD)
        setVotes(bytes(""));
    }



    function verifyBatchInclusion(
        bytes32 proofCommitment,
        bytes32 pubInputCommitment,
        bytes32 provingSystemAuxDataCommitment,
        bytes20 proofGeneratorAddr,
        bytes32 batchMerkleRoot,
        bytes memory merkleProof,
        uint256 verificationDataBatchIndex,
        bytes memory pubInputBytes
    ) public {
        require(pubInputCommitment == keccak256(abi.encodePacked(pubInputBytes)), "public input enayiligi");

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

        parseAndSetVote(pubInputBytes, 1240, 1240 + 16*64);

        emit ProofValidation(isVerified);
    }

    function parseAndSetVote(bytes memory pubInput, uint256 start, uint256 end) public {

        bytes memory result = new bytes(end - start);

        for (uint256 i = start; i < end; i++) {
            result[i - start] = pubInput[i];
        }
        
        setVotes(result);
    }


}
