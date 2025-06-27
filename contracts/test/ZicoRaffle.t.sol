// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Test.sol";
import "../src/ZicoRaffle.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MockZico is ERC20 {
    constructor() ERC20("ZICO", "ZICO") {
        _mint(msg.sender, 1_000_000_000 ether);
    }
}

contract MockVRFCoordinatorV2 {}

contract ZicoRaffleMock is ZicoRaffle {
    constructor(
        address vrfCoordinator,
        address zicoToken,
        address treasuryVault,
        bytes32 keyHash,
        uint64 subscriptionId
    ) ZicoRaffle(vrfCoordinator, zicoToken, treasuryVault, keyHash, subscriptionId) {}

    function callFulfillRandomWords(uint256 requestId, uint256[] memory randomWords) public {
        fulfillRandomWords(requestId, randomWords);
    }
}

contract ZicoRaffleTest is Test {
    ZicoRaffleMock public raffle;
    MockZico public zicoToken;
    address public owner = address(this);
    address public treasury = address(0x123);
    address public user1 = address(0x456);
    address public user2 = address(0x789);
    address public vrf = address(new MockVRFCoordinatorV2());
    bytes32 public keyHash = bytes32(uint256(1));
    uint64 public subId = 1;

    function setUp() public {
        zicoToken = new MockZico();
        raffle = new ZicoRaffleMock(
            vrf,
            address(zicoToken),
            treasury,
            keyHash,
            subId
        );
        zicoToken.transfer(treasury, 1_000_000 ether);
    }

    function testAddAndClearParticipants() public {
        address[] memory participants = new address[](2);
        participants[0] = user1;
        participants[1] = user2;
        raffle.addParticipants(participants);
        address[] memory got = raffle.getParticipants();
        assertEq(got.length, 2);
        assertEq(got[0], user1);
        assertEq(got[1], user2);
        raffle.clearParticipants();
        assertEq(raffle.getParticipants().length, 0);
    }

    function testRequestRaffleNoParticipants() public {
        uint256 prize = 100 ether;
        vm.expectRevert("No participants");
        raffle.requestRaffle(prize);
    }

    function testRequestRaffleInsufficientAllowance() public {
        address[] memory participants = new address[](1);
        participants[0] = user1;
        raffle.addParticipants(participants);
        uint256 prize = 100 ether;
        // Não aprova o prêmio
        vm.expectRevert("Insufficient allowance from treasury");
        raffle.requestRaffle(prize);
    }
} 