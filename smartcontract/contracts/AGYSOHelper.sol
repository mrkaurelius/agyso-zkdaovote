// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.12;

import "hardhat/console.sol";

/**
 * TODO Vote authentication
 */

contract AGYSOHelper {
    address public constant PROPOSER_ADD = 0xbb9aE6C54376DEe7a3e41BCeFEAb456866ab25a7;

    mapping(address => bytes) public encBallots;

    address public alignedServiceManager;
    address public paymentServiceAddr;

    event Proof(bool verified);
    event BallotBox(address voter, bytes32 proofCommitment);

    // Dummy constructor
    constructor() {}

    // constructor(address _alignedServiceManager, address _paymentServiceAddr) {
    //     alignedServiceManager = _alignedServiceManager;
    //     paymentServiceAddr = _paymentServiceAddr;
    // }

    function setBallotBox(bytes memory encVotes) private {
        encBallots[PROPOSER_ADD] = encVotes;
    }

    function getBallotBox() public view returns (bytes memory) {
        return encBallots[PROPOSER_ADD];
    }

    function finishElection() public {
        // require(msg.sender == PROPOSER_ADD, "asdf");
        setBallotBox(bytes(""));
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
