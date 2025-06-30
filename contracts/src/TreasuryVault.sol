// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/governance/TimelockController.sol";

interface IRewardManager {
    function updateReward(address user, uint256 rewardPool) external;
}

contract TreasuryVault {
    IERC20 public immutable zicoToken;
    IERC20 public immutable linkToken;
    IRewardManager public rewardManager;

    uint256 public zicoCollected;
    uint256 public linkCollected;

    event FeeReceived(address indexed payer, uint256 zicoAmount, uint256 linkAmount);
    event Harvested(uint256 zicoAmount, uint256 linkAmount);
    event RewardManagerSet(address rewardManager);
    event TimelockTransferred(address indexed newTimelock);
    event OperationalWalletSet(address indexed newWallet);
    event FeeDistributionUpdated(uint16 stakingPercent, uint16 operationsPercent);

    address public timelock;
    address public operationalWallet;
    uint16 public stakingPercent = 8000; // 80% (em basis points)
    uint16 public operationsPercent = 2000; // 20% (em basis points)
    uint16 public constant MAX_BPS = 10000;

    modifier onlyTimelock() {
        require(msg.sender == timelock, "Not timelock");
        _;
    }

    constructor(address _zicoToken, address _linkToken, address _timelock) {
        require(_zicoToken != address(0) && _linkToken != address(0) && _timelock != address(0), "Zero address");
        zicoToken = IERC20(_zicoToken);
        linkToken = IERC20(_linkToken);
        timelock = _timelock;
        operationalWallet = _timelock;
    }

    function transferTimelock(address newTimelock) external onlyTimelock {
        require(newTimelock != address(0), "Zero address");
        timelock = newTimelock;
        emit TimelockTransferred(newTimelock);
    }

    function setRewardManager(address _rewardManager) external onlyTimelock {
        rewardManager = IRewardManager(_rewardManager);
        emit RewardManagerSet(_rewardManager);
    }

    function setOperationalWallet(address _wallet) external onlyTimelock {
        require(_wallet != address(0), "Zero address");
        operationalWallet = _wallet;
        emit OperationalWalletSet(_wallet);
    }

    function setFeeDistribution(uint16 _stakingPercent, uint16 _operationsPercent) external onlyTimelock {
        require(_stakingPercent + _operationsPercent == MAX_BPS, "Percents must sum 10000");
        stakingPercent = _stakingPercent;
        operationsPercent = _operationsPercent;
        emit FeeDistributionUpdated(_stakingPercent, _operationsPercent);
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

    function harvest() external onlyTimelock {
        uint256 zicoAmount = zicoCollected;
        uint256 linkAmount = linkCollected;

        zicoCollected = 0;
        linkCollected = 0;

        uint256 stakingAmount = (zicoAmount * stakingPercent) / MAX_BPS;
        uint256 operationsAmount = zicoAmount - stakingAmount;

        if (address(rewardManager) != address(0) && stakingAmount > 0) {
            zicoToken.approve(address(rewardManager), stakingAmount);
            rewardManager.updateReward(msg.sender, stakingAmount);
        }
        if (operationalWallet != address(0) && operationsAmount > 0) {
            require(zicoToken.transfer(operationalWallet, operationsAmount), "Ops transfer failed");
        }

        emit Harvested(zicoAmount, linkAmount);
    }

    function withdrawTokens(address token, address to, uint256 amount) external onlyTimelock {
        require(IERC20(token).transfer(to, amount), "Withdraw failed");
    }
}
