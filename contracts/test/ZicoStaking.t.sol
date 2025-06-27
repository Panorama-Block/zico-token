// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Test.sol";
import "../src/ZicoStaking.sol";
import "../src/ZicoRaffle.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MockZico is ERC20 {
    constructor() ERC20("ZICO", "ZICO") {
        _mint(msg.sender, 1_000_000_000 ether);
    }
}

contract ZicoStakingSharesTest is Test {
    ZicoStakingShares public staking;
    MockZico public zicoToken;
    address public timelock = address(0x123);
    address public user1 = address(0x456);
    address public user2 = address(0x789);

    function setUp() public {
        zicoToken = new MockZico();
        staking = new ZicoStakingShares(address(zicoToken), timelock);
        zicoToken.transfer(user1, 1_000_000 ether);
        zicoToken.transfer(user2, 1_000_000 ether);
    }

    function testStakeAndUnstake() public {
        vm.startPrank(user1);
        zicoToken.approve(address(staking), 1000 ether);
        uint256 shares = staking.stake(1000 ether);
        assertEq(staking.balanceOf(user1), shares);
        uint256 zicoReturned = staking.unstake(shares);
        assertEq(zicoReturned, 1000 ether);
        assertEq(staking.balanceOf(user1), 0);
        vm.stopPrank();
    }

    function testProportionalShares() public {
        // user1 faz stake
        vm.startPrank(user1);
        zicoToken.approve(address(staking), 1000 ether);
        uint256 shares1 = staking.stake(1000 ether);
        vm.stopPrank();
        // user2 faz stake depois
        vm.startPrank(user2);
        zicoToken.approve(address(staking), 1000 ether);
        uint256 shares2 = staking.stake(1000 ether);
        vm.stopPrank();
        // Como não houve rendimento, shares devem ser iguais
        assertEq(shares1, shares2);
        // user1 resgata tudo
        vm.startPrank(user1);
        uint256 zicoReturned = staking.unstake(shares1);
        assertEq(zicoReturned, 1000 ether);
        vm.stopPrank();
    }

    function testTransferShares() public {
        vm.startPrank(user1);
        zicoToken.approve(address(staking), 1000 ether);
        uint256 shares = staking.stake(1000 ether);
        staking.transfer(user2, shares);
        assertEq(staking.balanceOf(user2), shares);
        vm.stopPrank();
    }

    function testOnlyTimelock() public {
        vm.expectRevert("Not timelock");
        staking.setZicoToken(address(zicoToken));
        vm.expectRevert("Not timelock");
        staking.transferTimelock(user2);
        vm.prank(timelock);
        staking.setZicoToken(address(zicoToken));
        vm.prank(timelock);
        staking.transferTimelock(user2);
        assertEq(staking.timelock(), user2);
    }

    function testZeroStake() public {
        vm.startPrank(user1);
        zicoToken.approve(address(staking), 0);
        vm.expectRevert("Cannot stake 0");
        staking.stake(0);
        vm.stopPrank();
    }

    function testZeroUnstake() public {
        vm.startPrank(user1);
        zicoToken.approve(address(staking), 1000 ether);
        staking.stake(1000 ether);
        vm.expectRevert("Invalid amount");
        staking.unstake(0);
        vm.stopPrank();
    }
}

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