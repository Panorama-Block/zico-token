// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.24;

import {Pool} from "../../../libraries/Pool.sol";
import {RateLimiter} from "../../../libraries/RateLimiter.sol";
import {BurnWithFromMintTokenPool} from "../../../pools/BurnWithFromMintTokenPool.sol";
import {TokenPool} from "../../../pools/TokenPool.sol";
import {BurnMintSetup} from "../BurnMintTokenPool/BurnMintSetup.t.sol";

import {IERC20} from "../../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";

contract BurnWithFromMintTokenPoolSetup is BurnMintSetup {
    BurnWithFromMintTokenPool internal s_pool;

    function setUp() public virtual override {
        BurnMintSetup.setUp();

        s_pool = new BurnWithFromMintTokenPool(
            s_burnMintERC20, DEFAULT_TOKEN_DECIMALS, new address[](0), address(s_mockRMNRemote), address(s_sourceRouter)
        );
        s_burnMintERC20.grantMintAndBurnRoles(address(s_pool));

        _applyChainUpdates(address(s_pool));
    }
}

contract BurnWithFromMintTokenPool_lockOrBurn is BurnWithFromMintTokenPoolSetup {
    function test_Setup() public view {
        assertEq(address(s_burnMintERC20), address(s_pool.getToken()));
        assertEq(address(s_mockRMNRemote), s_pool.getRmnProxy());
        assertEq(false, s_pool.getAllowListEnabled());
        assertEq(type(uint256).max, s_burnMintERC20.allowance(address(s_pool), address(s_pool)));
        assertEq("BurnWithFromMintTokenPool 1.5.1", s_pool.typeAndVersion());
    }

    function test_PoolBurn() public {
        uint256 burnAmount = 20_000e18;

        deal(address(s_burnMintERC20), address(s_pool), burnAmount);
        assertEq(s_burnMintERC20.balanceOf(address(s_pool)), burnAmount);

        vm.startPrank(s_burnMintOnRamp);

        vm.expectEmit();
        emit RateLimiter.TokensConsumed(burnAmount);

        vm.expectEmit();
        emit IERC20.Transfer(address(s_pool), address(0), burnAmount);

        vm.expectEmit();
        emit TokenPool.Burned(address(s_burnMintOnRamp), burnAmount);

        bytes4 expectedSignature = bytes4(keccak256("burn(address,uint256)"));
        vm.expectCall(address(s_burnMintERC20), abi.encodeWithSelector(expectedSignature, address(s_pool), burnAmount));

        s_pool.lockOrBurn(
            Pool.LockOrBurnInV1({
                originalSender: OWNER,
                receiver: bytes(""),
                amount: burnAmount,
                remoteChainSelector: DEST_CHAIN_SELECTOR,
                localToken: address(s_burnMintERC20)
            })
        );

        assertEq(s_burnMintERC20.balanceOf(address(s_pool)), 0);
    }

    // Should not burn tokens if cursed.
    function test_RevertWhen_PoolBurnRevertNotHealthy() public {
        vm.mockCall(address(s_mockRMNRemote), abi.encodeWithSignature("isCursed(bytes16)"), abi.encode(true));
        uint256 before = s_burnMintERC20.balanceOf(address(s_pool));
        vm.startPrank(s_burnMintOnRamp);

        vm.expectRevert(TokenPool.CursedByRMN.selector);
        s_pool.lockOrBurn(
            Pool.LockOrBurnInV1({
                originalSender: OWNER,
                receiver: bytes(""),
                amount: 1e5,
                remoteChainSelector: DEST_CHAIN_SELECTOR,
                localToken: address(s_burnMintERC20)
            })
        );

        assertEq(s_burnMintERC20.balanceOf(address(s_pool)), before);
    }

    function test_RevertWhen_ChainNotAllowed() public {
        uint64 wrongChainSelector = 8838833;
        vm.expectRevert(abi.encodeWithSelector(TokenPool.ChainNotAllowed.selector, wrongChainSelector));
        s_pool.releaseOrMint(
            Pool.ReleaseOrMintInV1({
                originalSender: bytes(""),
                receiver: OWNER,
                amount: 1,
                localToken: address(s_burnMintERC20),
                remoteChainSelector: wrongChainSelector,
                sourcePoolAddress: _generateSourceTokenData().sourcePoolAddress,
                sourcePoolData: _generateSourceTokenData().extraData,
                offchainTokenData: ""
            })
        );
    }
}
