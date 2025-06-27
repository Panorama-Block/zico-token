// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract ZICOStaking is ERC20, Ownable {
    uint256 public totalStaked;
    mapping(address => uint256) public stakes;
    mapping(address => uint256) public rewards;

    constructor() ERC20("ZICOAI", "ZICOAI") Ownable(msg.sender) {
        _mint(msg.sender, 1_000_000_000 ether);
    }

    function stake(uint256 amount) external {
        require(amount > 0, "Cannot stake 0");
        _transfer(msg.sender, address(this), amount);

        stakes[msg.sender] += amount;
        totalStaked += amount;
    }

    function unstake(uint256 amount) external {
        require(amount > 0 && amount <= stakes[msg.sender], "Invalid unstake");
        stakes[msg.sender] -= amount;
        totalStaked -= amount;
        _transfer(address(this), msg.sender, amount);
    }

    function calculateReward(address user, uint256 rewardPool) public view returns (uint256) {
        if (totalStaked == 0 || stakes[user] == 0) return 0;
        return (stakes[user] * rewardPool) / totalStaked;
    }

    function updateReward(address user, uint256 rewardPool) external onlyOwner {
        uint256 reward = calculateReward(user, rewardPool);
        if (reward > 0) {
            rewards[user] += reward;
        }
    }

    function claimReward() external {
        uint256 reward = rewards[msg.sender];
        require(reward > 0, "No rewards");
        rewards[msg.sender] = 0;
        _transfer(address(this), msg.sender, reward);
    }
}
