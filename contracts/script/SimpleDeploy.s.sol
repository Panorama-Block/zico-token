// // SPDX-License-Identifier: MIT
// pragma solidity ^0.8.24;

// import "forge-std/Script.sol";

// contract SimpleDeployScript is Script {
//     function run() external {
//         vm.startBroadcast();

//         // Deploy basic contracts first to test
//         console.log("=========================================");
//         console.log("         SIMPLE ZICO TOKEN DEPLOY");
//         console.log("=========================================");
//         console.log("Deployer:", msg.sender);
//         console.log("");

//         // ===== MOCK ADDRESSES FOR LOCAL ANVIL TESTING =====
//         address router = 0x0BF3dE8c5D3e8A2B34D2BEeB17ABfCeBaf363A59; 
//         address linkToken = 0x514910771AF9Ca656af840dff83E8264EcF986CA;
//         address vrfCoordinator = 0x8103B0A8A00be2DDC778e6e7eaa21791Cd364625;
//         bytes32 keyHash = 0x79d3d8832d904592c0bf9818b621522c988bb8b0c05cdc3b15aea1b6e8db0c15;
//         uint64 subscriptionId = 1;
//         address functionsRouter = 0x6E2dc0F9DB014aE19888F539E59285D2Ea04244C;
//         address timelock = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266; // Anvil account[0]

//         console.log("Mock addresses set for local testing");
//         console.log("All contracts will be deployed sequentially...");
//         console.log("");

//         vm.stopBroadcast();

//         console.log("=========================================");
//         console.log("         DEPLOY READY!");
//         console.log("=========================================");
//         console.log("Run the following commands to deploy each contract:");
//         console.log("");
//         console.log("1. forge create src/TreasuryVault.sol:TreasuryVault \\");
//         console.log("   --constructor-args <zicoToken> <linkToken> <timelock>");
//         console.log("");
//         console.log("2. forge create src/ZicoStaking.sol:ZicoStakingShares \\");
//         console.log("   --constructor-args <zicoToken> <timelock>");
//         console.log("");
//         console.log("3. forge create src/ZicoRaffle.sol:ZicoRaffle \\");
//         console.log("   --constructor-args <vrfCoordinator> <zicoToken> <treasury> <keyHash> <subscriptionId>");
//         console.log("=========================================");
//     }
// } 