// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Script.sol";
import {ZicoToken} from "../src/ZicoToken.sol";
import {ZicoStakingShares} from "../src/ZicoStaking.sol";
import {TreasuryVault} from "../src/TreasuryVault.sol"; 
import {ZicoRaffle} from "../src/ZicoRaffle.sol";
import {ProofOfReserve} from "../src/ProofOfReserve.sol";
import "forge-std/console.sol";

contract DeployScript is Script {
    function run() external {
        vm.startBroadcast();

        address router        = _getEnvAddress("CCIP_ROUTER", 0xF694E193200268f9a4868e4Aa017A0118C9a8177 );
        address linkToken     = _getEnvAddress("LINK_TOKEN", 0x0b9d5D9136855f6FEc3c0993feE6E9CE8a297846 );
        address vrfCoordinator= _getEnvAddress("VRF_COORDINATOR", 0x2eD832Ba664535e5886b75D64C46EB9a228C2610 );
        address functionsRouter=_getEnvAddress("FUNCTIONS_ROUTER", 0xA9d587a00A31A52Ed70D6026794a8FC5E2F5dCb0);
        bytes32 keyHash       = _getEnvBytes32("VRF_KEYHASH", 0x354d2f95da55398f44b7cff77da56283d9c6c829a4bdf1bbcaf2ad6a4d081f61 );
        uint64  subscriptionId= _getEnvUint64("VRF_SUB_ID", 1);

  
        address timelock      = _getEnvAddress("TIMELOCK", msg.sender);

        console.log("=========================================");
        console.log("         ZICO TOKEN ECOSYSTEM DEPLOY");
        console.log("=========================================");
        console.log("Deployer:", msg.sender);
        console.log("");

        console.log("1. Deploying ZicoToken...");
        ZicoToken zicoToken = new ZicoToken(
            router,
            linkToken,
            vrfCoordinator,
            keyHash,
            subscriptionId,
            timelock,
            functionsRouter,
            address(0) // temp
        );
       
      
        console.log("   ZicoToken deployed:", address(zicoToken));
    console.log("2. Deploying TreasuryVault...");
        TreasuryVault treasury = new TreasuryVault(
            address(zicoToken),
            linkToken,
            timelock
        );
        console.log("   TreasuryVault deployed:", address(treasury));

        console.log("3. Deploying ZicoStaking...");
        ZicoStakingShares staking = new ZicoStakingShares(
            address(zicoToken),
            timelock
        );
        console.log("   ZicoStaking deployed:", address(staking));

        console.log("4. Deploying ZicoRaffle...");
        ZicoRaffle raffle = new ZicoRaffle(
            vrfCoordinator,
            address(zicoToken),
            address(treasury),
            keyHash,
            subscriptionId
        );
        console.log("   ZicoRaffle deployed:", address(raffle));

        console.log("5. Deploying ProofOfReserve...");
        uint64[] memory chainSelectors = new uint64[](3);
        chainSelectors[0] = 5009297550715157269; // Ethereum Sepolia
        chainSelectors[1] = 14767482510784806043; // Avalanche Fuji
        chainSelectors[2] = 3478487238524512106;  // Arbitrum Sepolia

        
        ProofOfReserve proofOfReserve = new ProofOfReserve(
            timelock, 
            chainSelectors
        );
        console.log("   ProofOfReserve deployed:", address(proofOfReserve));

        vm.stopBroadcast();

        console.log("");
        console.log("=========================================");
        console.log("         DEPLOYMENT COMPLETE! ");
        console.log("=========================================");
        console.log(" Contract Addresses:");
        console.log("");
        console.log("  ZicoToken:      ", address(zicoToken));
        console.log("  TreasuryVault:  ", address(treasury));
        console.log("  ZicoStaking:    ", address(staking));
        console.log("  ZicoRaffle:     ", address(raffle));
        console.log("  ProofOfReserve: ", address(proofOfReserve));
        console.log("");
        console.log(" Copy these addresses to your frontend!");
        console.log("=========================================");
    }

    function _getEnvAddress(string memory key, address fallbackAddr) internal view returns (address) {
        try vm.envAddress(key) returns (address a) { return a; } catch { return fallbackAddr; }
    }
    function _getEnvBytes32(string memory key, bytes32 fallbackVal) internal view returns (bytes32) {
        try vm.envBytes32(key) returns (bytes32 v) { return v; } catch { return fallbackVal; }
    }
    function _getEnvUint64(string memory key, uint64 fallbackVal) internal view returns (uint64) {
        try vm.envUint(key) returns (uint256 v) { return uint64(v); } catch { return fallbackVal; }
    }
} 