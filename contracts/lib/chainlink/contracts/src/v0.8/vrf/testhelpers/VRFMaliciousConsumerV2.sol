// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {LinkTokenInterface} from "../../shared/interfaces/LinkTokenInterface.sol";
import {VRFCoordinatorV2Interface} from "../interfaces/VRFCoordinatorV2Interface.sol";
import {VRFConsumerBaseV2} from "../VRFConsumerBaseV2.sol";

contract VRFMaliciousConsumerV2 is VRFConsumerBaseV2 {
    uint256[] public s_randomWords;
    uint256 public s_requestId;
    VRFCoordinatorV2Interface internal COORDINATOR;
    LinkTokenInterface internal LINKTOKEN;
    uint64 public s_subId;
    uint256 public s_gasAvailable;
    bytes32 internal s_keyHash;

    constructor(address vrfCoordinator, address link) VRFConsumerBaseV2(vrfCoordinator) {
        COORDINATOR = VRFCoordinatorV2Interface(vrfCoordinator);
        LINKTOKEN = LinkTokenInterface(link);
    }

    function fulfillRandomWords(uint256 requestId, uint256[] memory randomWords) internal override {
        s_gasAvailable = gasleft();
        s_randomWords = randomWords;
        s_requestId = requestId;
        // Should revert
        COORDINATOR.requestRandomWords(s_keyHash, s_subId, 1, 200000, 1);
    }

    function createSubscriptionAndFund(uint96 amount) external {
        if (s_subId == 0) {
            s_subId = COORDINATOR.createSubscription();
            COORDINATOR.addConsumer(s_subId, address(this));
        }
        // Approve the link transfer.
        LINKTOKEN.transferAndCall(address(COORDINATOR), amount, abi.encode(s_subId));
    }

    function updateSubscription(address[] memory consumers) external {
        require(s_subId != 0, "subID not set");
        for (uint256 i = 0; i < consumers.length; i++) {
            COORDINATOR.addConsumer(s_subId, consumers[i]);
        }
    }

    function requestRandomness(bytes32 keyHash) external returns (uint256) {
        s_keyHash = keyHash;
        return COORDINATOR.requestRandomWords(keyHash, s_subId, 1, 500000, 1);
    }
}
