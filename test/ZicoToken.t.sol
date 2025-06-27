// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Test.sol";
import "../src/ZicoToken.sol";

contract ZicoTokenTest is Test {
    ZicoToken public token;
    address public owner;
    address public user1;
    address public user2;
    
    // Mock addresses for testing
    address mockRouter = address(0x1);
    address mockLink = address(0x2);
    address mockVrfCoordinator = address(0x3);
    bytes32 mockKeyHash = bytes32(uint256(1));
    uint64 mockSubscriptionId = 1;

    function setUp() public {
        owner = address(this);
        user1 = makeAddr("user1");
        user2 = makeAddr("user2");
        
        // Deploy the contract with mock addresses
        token = new ZicoToken(
            mockRouter,
            mockLink,
            mockVrfCoordinator,
            mockKeyHash,
            mockSubscriptionId
        );
    }

    function test_InitialSupply() public {
        assertEq(token.totalSupply(), 1_000_000 ether);
        assertEq(token.balanceOf(owner), 1_000_000 ether);
    }

    function test_Transfer() public {
        uint256 amount = 1000 ether;
        token.transfer(user1, amount);
        
        assertEq(token.balanceOf(user1), amount);
        assertEq(token.balanceOf(owner), 1_000_000 ether - amount);
    }

    function test_Staking() public {
        uint256 stakeAmount = 100 ether;
        
        // Transfer tokens to user1 first
        token.transfer(user1, stakeAmount);
        
        // User1 stakes tokens
        vm.prank(user1);
        token.stake(stakeAmount);
        
        assertEq(token.stakes(user1), stakeAmount);
        assertEq(token.totalStaked(), stakeAmount);
        assertEq(token.balanceOf(address(token)), stakeAmount);
    }

    function test_Unstaking() public {
        uint256 stakeAmount = 100 ether;
        uint256 unstakeAmount = 50 ether;
        
        // Setup staking first
        token.transfer(user1, stakeAmount);
        vm.prank(user1);
        token.stake(stakeAmount);
        
        // Unstake partial amount
        vm.prank(user1);
        token.unstake(unstakeAmount);
        
        assertEq(token.stakes(user1), stakeAmount - unstakeAmount);
        assertEq(token.totalStaked(), stakeAmount - unstakeAmount);
    }

    function test_OnlyOwnerFunctions() public {
        // Test that non-owner cannot call distributeRewards
        vm.prank(user1);
        vm.expectRevert();
        token.distributeRewards();
        
        // Test that owner can call distributeRewards
        token.distributeRewards(); // Should not revert
    }
} 