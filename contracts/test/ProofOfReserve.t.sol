// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Test.sol";
import "../src/ProofOfReserve.sol";

contract ProofOfReserveTest is Test {
    ProofOfReserve public proof;
    address public oracle = address(0x123);
    address public notOracle = address(0x456);
    uint64[] public selectors;

    function setUp() public {
        selectors = new uint64[](2);
        selectors[0] = 1;
        selectors[1] = 2;
        proof = new ProofOfReserve(oracle, selectors);
    }

    function testReportSupplyAndTotal() public {
        vm.prank(oracle);
        proof.reportSupply(1, 100);
        vm.prank(oracle);
        proof.reportSupply(2, 200);
        assertEq(proof.getReportedSupply(1), 100);
        assertEq(proof.getReportedSupply(2), 200);
        assertEq(proof.totalGlobalSupply(), 300);
    }

    function testOnlyOracle() public {
        vm.expectRevert("Not authorized");
        proof.reportSupply(1, 100);
        vm.expectRevert("Not authorized");
        proof.setTolerancePercent(2);
        vm.expectRevert("Not authorized");
        proof.updateOracle(notOracle);
    }

    function testUpdateOracle() public {
        vm.prank(oracle);
        proof.updateOracle(notOracle);
        vm.prank(notOracle);
        proof.reportSupply(1, 123);
        assertEq(proof.getReportedSupply(1), 123);
    }

    function testSetTolerancePercent() public {
        vm.prank(oracle);
        proof.setTolerancePercent(5);
        assertEq(proof.tolerancePercent(), 5);
        vm.prank(oracle);
        vm.expectRevert("Excessive tolerance");
        proof.setTolerancePercent(11);
    }

    function testGetAllChains() public view {
        uint64[] memory chains = proof.getAllChains();
        assertEq(chains.length, 2);
        assertEq(chains[0], 1);
        assertEq(chains[1], 2);
    }

    function testEvents() public {
        vm.prank(oracle);
        vm.expectEmit(true, true, false, false);
        emit ProofOfReserve.SupplyReported(1, 100);
        proof.reportSupply(1, 100);
    }
}
