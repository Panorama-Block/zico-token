// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {OVM_GasPriceOracle} from
    "../../vendor/@eth-optimism/contracts/v0.8.9/contracts/L2/predeploys/OVM_GasPriceOracle.sol";
import {ChainModuleBase} from "./ChainModuleBase.sol";

/**
 * @notice This contract is deprecated. Please use OptimismModuleV2 which utilizes the most recent offerings from OP
 * and can estimate L1 fee with much lower cost.
 */
contract OptimismModule is ChainModuleBase {
    /// @dev OP_L1_DATA_FEE_PADDING includes 80 bytes for L1 data padding for Optimism and BASE
    bytes private constant OP_L1_DATA_FEE_PADDING =
        hex"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff";
    /// @dev OVM_GASPRICEORACLE_ADDR is the address of the OVM_GasPriceOracle precompile on Optimism.
    /// @dev reference: https://community.optimism.io/docs/developers/build/transaction-fees/#estimating-the-l1-data-fee
    address private constant OVM_GASPRICEORACLE_ADDR = 0x420000000000000000000000000000000000000F;
    OVM_GasPriceOracle private constant OVM_GASPRICEORACLE = OVM_GasPriceOracle(OVM_GASPRICEORACLE_ADDR);

    uint256 private constant FIXED_GAS_OVERHEAD = 60_000;
    uint256 private constant PER_CALLDATA_BYTE_GAS_OVERHEAD = 270;

    // @dev This will be updated to use the new function introduced by OP team
    function getCurrentL1Fee(uint256 dataSize) external view override returns (uint256) {
        return _getL1Fee(dataSize);
    }

    function getMaxL1Fee(uint256 dataSize) external view override returns (uint256) {
        return _getL1Fee(dataSize);
    }

    /* @notice this function provides an estimation for L1 fee incurred by calldata of a certain size
    * @dev this function uses the getL1Fee function in OP gas price oracle. it estimates the exact L1 fee but it costs
    * a lot of gas to call.
    * @param dataSize the size of calldata
    * @return l1Fee the L1 fee
    */
    function _getL1Fee(uint256 dataSize) internal view returns (uint256) {
        // fee is 4 per 0 byte, 16 per non-zero byte. Worst case we can have all non zero-bytes.
        // Instead of setting bytes to non-zero, we initialize 'new bytes' of length 4*dataSize to cover for zero bytes.
        bytes memory txCallData = new bytes(4 * dataSize);
        return OVM_GASPRICEORACLE.getL1Fee(bytes.concat(txCallData, OP_L1_DATA_FEE_PADDING));
    }

    function getGasOverhead()
        external
        view
        override
        returns (uint256 chainModuleFixedOverhead, uint256 chainModulePerByteOverhead)
    {
        return (FIXED_GAS_OVERHEAD, PER_CALLDATA_BYTE_GAS_OVERHEAD);
    }
}
