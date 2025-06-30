// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Test.sol";
import "../src/ZicoToken.sol";
import "../src/TreasuryVault.sol" as LocalTreasuryVault;
import "../src/ProofOfReserve.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MockRouter {}

contract MockVRFCoordinator {}

contract MockFunctionsRouter {}

contract MockZico is ERC20 {
    constructor() ERC20("ZICO", "ZICO") {
        _mint(msg.sender, 1_000_000_000 ether);
    }
}

contract ZicoTokenTest is Test {
    ZicoToken public zico;
    MockZico public zicoToken;
    LocalTreasuryVault.TreasuryVault public treasury;
    ProofOfReserve public proof;
    address public timelock = address(0x123);
    address public user = address(0x456);
    address public premium = address(0x789);
    address public router = address(new MockRouter());
    address public vrf = address(new MockVRFCoordinator());
    address public functionsRouter = address(new MockFunctionsRouter());
    bytes32 public keyHash = bytes32(uint256(1));
    uint64 public subId = 1;

    function setUp() public {
        zicoToken = new MockZico();
        treasury = new LocalTreasuryVault.TreasuryVault(address(zicoToken), address(zicoToken), timelock);
        uint64[] memory selectors = new uint64[](1);
        selectors[0] = 1;
        proof = new ProofOfReserve(timelock, selectors);
        vm.prank(timelock);
        zico = new ZicoToken(
            router, address(zicoToken), vrf, keyHash, subId, timelock, functionsRouter, address(zicoToken)
        );
        vm.prank(timelock);
        zico.setTreasuryVault(address(treasury));
        vm.prank(timelock);
        zico.setProofOfReserve(address(proof));
        zicoToken.transfer(user, 1_000_000 ether);
    }

    function testStakeAndUnstake() public {
        vm.startPrank(user);
        zicoToken.approve(address(zico), 1000 ether);
        uint256 shares = zico.stake(1000 ether);
        assertEq(zico.balanceOf(user), shares);

        uint256 totalShares = zico.totalSupply();
        uint256 totalZico = zico.zicoTokenBalance();
        uint256 expectedReturn = (shares * totalZico) / totalShares;

        uint256 zicoReturned = zico.unstake(shares);
        assertEq(zicoReturned, expectedReturn);
        assertEq(zico.balanceOf(user), 0);
        vm.stopPrank();
    }

    function test_RevertWhen_GatingWithoutStaking() public {
        bytes32 serviceId = keccak256("AI_AGENT");
        vm.prank(timelock);
        zico.setPremiumServiceFee(serviceId, 10 ether);
        vm.prank(timelock);
        zico.setStakingContract(address(zico));
        vm.startPrank(user);
        zicoToken.approve(address(zico), 500 ether);
        zico.stake(500 ether);
        vm.expectRevert("Insufficient staking level");
        zico.payPremiumService(serviceId);
        vm.stopPrank();
    }

    function testPauseAndUnpause() public {
        vm.prank(timelock);
        zico.pause();
        vm.startPrank(user);
        zicoToken.approve(address(zico), 1000 ether);
        vm.expectRevert("Paused");
        zico.stake(1000 ether);
        vm.stopPrank();
        vm.prank(timelock);
        zico.unpause();
        vm.startPrank(user);
        zico.stake(1000 ether);
        assertGt(zico.balanceOf(user), 0);
        vm.stopPrank();
    }

    function testAutoPauseIfUnhealthy() public {
        vm.prank(timelock);
        zico.unpause();
        vm.etch(address(proof), hex"00");
        vm.prank(user);
        vm.expectRevert();
        zico.autoPauseIfUnhealthy();
    }
}
