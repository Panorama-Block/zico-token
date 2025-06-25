// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IRewardManager {
    function updateReward(address user, uint256 rewardPool) external;
}

contract TreasuryVault is Ownable {
    IERC20 public immutable zicoToken;
    IERC20 public immutable linkToken;
    IRewardManager public rewardManager;

    uint256 public zicoCollected;
    uint256 public linkCollected;

    event FeeReceived(address indexed payer, uint256 zicoAmount, uint256 linkAmount);
    event Harvested(uint256 zicoAmount, uint256 linkAmount);
    event RewardManagerSet(address rewardManager);

    constructor(address _zicoToken, address _linkToken) Ownable(msg.sender) {
        require(_zicoToken != address(0) && _linkToken != address(0), "Zero address");
        zicoToken = IERC20(_zicoToken);
        linkToken = IERC20(_linkToken);
    }

    function setRewardManager(address _rewardManager) external onlyOwner {
        rewardManager = IRewardManager(_rewardManager);
        emit RewardManagerSet(_rewardManager);
    }

    function payFee(uint256 zicoAmount, uint256 linkAmount) external {
        if (zicoAmount > 0) {
            require(zicoToken.transferFrom(msg.sender, address(this), zicoAmount), "ZICO transfer failed");
            zicoCollected += zicoAmount;
        }
        if (linkAmount > 0) {
            require(linkToken.transferFrom(msg.sender, address(this), linkAmount), "LINK transfer failed");
            linkCollected += linkAmount;
        }
        emit FeeReceived(msg.sender, zicoAmount, linkAmount);
    }

    function harvest() external onlyOwner {
        uint256 zicoAmount = zicoCollected;
        uint256 linkAmount = linkCollected;

        zicoCollected = 0;
        linkCollected = 0;

        if (address(rewardManager) != address(0) && zicoAmount > 0) {
            zicoToken.approve(address(rewardManager), zicoAmount);
            rewardManager.updateReward(msg.sender, zicoAmount); // ou algum outro mecanismo de distribuição
        }

        emit Harvested(zicoAmount, linkAmount);
    }

    function withdrawTokens(address token, address to, uint256 amount) external onlyOwner {
        require(IERC20(token).transfer(to, amount), "Withdraw failed");
    }
}
