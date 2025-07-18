// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.19;

import {BaseTestWithConfiguredVerifierAndFeeManager} from "./BaseVerifierTest.t.sol";
import {VerifierProxy} from "../../../v0.3.0/VerifierProxy.sol";
import {FeeManager} from "../../../v0.3.0/FeeManager.sol";

contract VerifierProxyInitializeVerifierTest is BaseTestWithConfiguredVerifierAndFeeManager {
    function test_setFeeManagerZeroAddress() public {
        vm.expectRevert(abi.encodeWithSelector(VerifierProxy.ZeroAddress.selector));
        s_verifierProxy.setFeeManager(FeeManager(address(0)));
    }

    function test_setFeeManagerWhichDoesntHonourInterface() public {
        vm.expectRevert(abi.encodeWithSelector(VerifierProxy.FeeManagerInvalid.selector));
        s_verifierProxy.setFeeManager(FeeManager(address(s_verifier)));
    }

    function test_setFeeManagerWhichDoesntHonourIERC165Interface() public {
        vm.expectRevert();
        s_verifierProxy.setFeeManager(FeeManager(address(1)));
    }
}
