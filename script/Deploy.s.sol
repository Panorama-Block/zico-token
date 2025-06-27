// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "forge-std/Script.sol";
import "../src/ZicoToken.sol";

contract DeployScript is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        // Mock addresses for local testing (Anvil)
        address router = 0x5FbDB2315678afecb367f032d93F642f64180aa3; // Mock CCIP Router
        address linkToken = 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512; // Mock LINK Token
        address vrfCoordinator = 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0; // Mock VRF Coordinator
        bytes32 keyHash = 0xd89b2bf150e3b9e13446986e571fb9cab24b13cea0a43ea20a6049a85cc807cc; // Mock key hash
        uint64 subscriptionId = 1; // Mock subscription ID

        ZicoToken zicoToken = new ZicoToken(
            router,
            linkToken,
            vrfCoordinator,
            keyHash,
            subscriptionId
        );

        console.log("ZicoToken deployed to:", address(zicoToken));
        console.log("Initial supply minted to:", msg.sender);

        vm.stopBroadcast();
    }
}   