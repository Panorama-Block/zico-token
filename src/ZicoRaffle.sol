// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "chainlink/contracts/src/v0.8/vrf/VRFConsumerBaseV2.sol";
import "chainlink/contracts/src/v0.8/vrf/interfaces/VRFCoordinatorV2Interface.sol";

contract ZicoRaffle is VRFConsumerBaseV2, Ownable(msg.sender) {
    VRFCoordinatorV2Interface public immutable coordinator;
    IERC20 public immutable zicoToken;
    address public immutable treasuryVault;

    bytes32 public keyHash;
    uint64 public subscriptionId;
    uint32 public callbackGasLimit = 100000;
    uint16 public requestConfirmations = 3;
    uint32 public numWords = 1;

    address[] public participants;
    mapping(uint256 => uint256) public requestIdToPrize;

    event RaffleRequested(uint256 requestId, uint256 prizeAmount);
    event RaffleWinner(address indexed winner, uint256 prizeAmount);

    constructor(
        address _vrfCoordinator,
        address _zicoToken,
        address _treasuryVault,
        bytes32 _keyHash,
        uint64 _subscriptionId
    ) VRFConsumerBaseV2(_vrfCoordinator) {
        coordinator = VRFCoordinatorV2Interface(_vrfCoordinator);
        zicoToken = IERC20(_zicoToken);
        treasuryVault = _treasuryVault;
        keyHash = _keyHash;
        subscriptionId = _subscriptionId;
    }

    function addParticipants(address[] calldata _participants) external onlyOwner {
        for (uint256 i = 0; i < _participants.length; i++) {
            participants.push(_participants[i]);
        }
    }

    function clearParticipants() external onlyOwner {
        delete participants;
    }

    function requestRaffle(uint256 prizeAmount) external onlyOwner {
        require(participants.length > 0, "No participants");
        require(
            zicoToken.allowance(treasuryVault, address(this)) >= prizeAmount,
            "Insufficient allowance from treasury"
        );

        uint256 requestId = coordinator.requestRandomWords(
            keyHash,
            subscriptionId,
            requestConfirmations,
            callbackGasLimit,
            numWords
        );

        requestIdToPrize[requestId] = prizeAmount;
        emit RaffleRequested(requestId, prizeAmount);
    }

    function fulfillRandomWords(uint256 requestId, uint256[] memory randomWords) internal override {
        uint256 prize = requestIdToPrize[requestId];
        require(prize > 0, "Prize not set");

        uint256 winnerIndex = randomWords[0] % participants.length;
        address winner = participants[winnerIndex];

        delete requestIdToPrize[requestId];

        require(zicoToken.transferFrom(treasuryVault, winner, prize), "Transfer failed");
        emit RaffleWinner(winner, prize);
    }

    function getParticipants() external view returns (address[] memory) {
        return participants;
    }
}
