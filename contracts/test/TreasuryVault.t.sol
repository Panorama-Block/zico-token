// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Test.sol";
import "../src/TreasuryVault.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MockZico is ERC20 {
    constructor() ERC20("ZICO", "ZICO") {
        _mint(msg.sender, 1_000_000_000 ether);
    }
}

contract MockRewardManager is IRewardManager {
    address public lastUser;
    uint256 public lastReward;

    function updateReward(address user, uint256 rewardPool) external override {
        lastUser = user;
        lastReward = rewardPool;
    }
}

contract TreasuryVaultTest is Test {
    TreasuryVault public vault;
    MockZico public zicoToken;
    MockZico public linkToken;
    MockRewardManager public rewardManager;
    address public timelock = address(0x123);
    address public user = address(0x456);
    address public ops = address(0x789);

    function setUp() public {
        zicoToken = new MockZico();
        linkToken = new MockZico();
        vault = new TreasuryVault(address(zicoToken), address(linkToken), timelock);
        rewardManager = new MockRewardManager();
        vm.prank(timelock);
        vault.setRewardManager(address(rewardManager));
        vm.prank(timelock);
        vault.setOperationalWallet(ops);
        zicoToken.transfer(user, 1_000_000 ether);
    }

    function testPayFeeAndHarvest() public {
        vm.startPrank(user);
        zicoToken.approve(address(vault), 100 ether);
        vault.payFee(100 ether, 0);
        vm.stopPrank();
        assertEq(vault.zicoCollected(), 100 ether);
        // Harvest: 80 para staking, 20 para ops
        vm.prank(timelock);
        vault.harvest();
        assertEq(rewardManager.lastReward(), 80 ether);
        assertEq(zicoToken.balanceOf(ops), 20 ether);
        assertEq(vault.zicoCollected(), 0);
    }

    function testSetFeeDistribution() public {
        vm.prank(timelock);
        vault.setFeeDistribution(7000, 3000);
        assertEq(vault.stakingPercent(), 7000);
        assertEq(vault.operationsPercent(), 3000);
    }

    function testOnlyTimelock() public {
        vm.expectRevert("Not timelock");
        vault.setFeeDistribution(8000, 2000);
        vm.expectRevert("Not timelock");
        vault.setOperationalWallet(user);
        vm.expectRevert("Not timelock");
        vault.setRewardManager(address(rewardManager));
        vm.expectRevert("Not timelock");
        vault.harvest();
        vm.expectRevert("Not timelock");
        vault.withdrawTokens(address(zicoToken), user, 1 ether);
    }

    function testZeroAddressProtection() public {
        vm.prank(timelock);
        vm.expectRevert("Zero address");
        vault.setOperationalWallet(address(0));
    }

    function testHarvestNoBalance() public {
        vm.prank(timelock);
        vault.harvest(); // Não deve reverter
        assertEq(zicoToken.balanceOf(address(rewardManager)), 0);
        assertEq(zicoToken.balanceOf(ops), 0);
    }

    function testWithdrawTokens() public {
        vm.startPrank(user);
        zicoToken.approve(address(vault), 100 ether);
        vault.payFee(100 ether, 0);
        vm.stopPrank();
        vm.prank(timelock);
        vault.withdrawTokens(address(zicoToken), user, 50 ether);
        uint256 expected = 1_000_000 ether - 100 ether + 50 ether;
        assertEq(zicoToken.balanceOf(user), expected);
    }
}
