// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {ChainlinkRequestInterface} from "../../interfaces/ChainlinkRequestInterface.sol";
import {OracleInterface} from "../../interfaces/OracleInterface.sol";

/* solhint-disable no-empty-blocks */
contract EmptyOracle is ChainlinkRequestInterface, OracleInterface {
    function cancelOracleRequest(bytes32, uint256, bytes4, uint256) external override {}
    function fulfillOracleRequest(bytes32, uint256, address, bytes4, uint256, bytes32)
        external
        override
        returns (bool)
    {}

    function getAuthorizationStatus(address) external pure returns (bool) {
        return false;
    }

    function onTokenTransfer(address, uint256, bytes calldata) external pure {}
    function oracleRequest(address, uint256, bytes32, address, bytes4, uint256, uint256, bytes calldata)
        external
        override
    {}
    function setFulfillmentPermission(address, bool) external {}
    function withdraw(address, uint256) external override {}
    function withdrawable() external view override returns (uint256) {}
}
