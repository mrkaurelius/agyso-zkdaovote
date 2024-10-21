// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

contract AGYSODaoVoteValidator {
    address public alignedServiceManager;
    address public paymentServiceAddr;

    constructor(address _alignedServiceManager, address _paymentServiceAddr) {
        alignedServiceManager = _alignedServiceManager;
        paymentServiceAddr = _paymentServiceAddr;
    }

    function verifyBatchInclusion(
        bytes32 proofCommitment,
        bytes32 pubInputCommitment,
        bytes32 programIdCommitment,
        bytes20 proofGeneratorAddr,
        bytes32 batchMerkleRoot,
        bytes memory merkleProof,
        uint256 verificationDataBatchIndex,
        bytes memory pubInputBytes
    ) public returns (bool) {


        require(
            pubInputCommitment == keccak256(abi.encodePacked(pubInputBytes)),
            "public input enayiligi"
        );

        (bool callWasSuccessful, bytes memory proofIsIncluded) = alignedServiceManager.staticcall(
            abi.encodeWithSignature(
                "verifyBatchInclusion(bytes32,bytes32,bytes32,bytes20,bytes32,bytes,uint256,address)",
                proofCommitment,
                pubInputCommitment,
                programIdCommitment,
                proofGeneratorAddr,
                batchMerkleRoot,
                merkleProof,
                verificationDataBatchIndex,
                paymentServiceAddr
            )
        );

        require(callWasSuccessful, "static_call failed");

        return abi.decode(proofIsIncluded, (bool));
    }

    // function bytesToTwoUint32(bytes memory data) public pure returns (uint32, uint32) {
    //     require(data.length >= 8, "Input bytes must be at least 8 bytes long");

    //     uint32 first = uint32(uint8(data[0])) |
    //         (uint32(uint8(data[1])) << 8) |
    //         (uint32(uint8(data[2])) << 16) |
    //         (uint32(uint8(data[3])) << 24);

    //     uint32 second = uint32(uint8(data[4])) |
    //         (uint32(uint8(data[5])) << 8) |
    //         (uint32(uint8(data[6])) << 16) |
    //         (uint32(uint8(data[7])) << 24);

    //     return (first, second);
    // }
}
