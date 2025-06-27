// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IZicoToken {
    function transferFrom(address from, address to, uint256 amount) external returns (bool);
    function transfer(address to, uint256 amount) external returns (bool);
}

contract ZicoStakingShares is ERC20, Ownable {
    IZicoToken public zicoToken;
    address public timelock;

    event Staked(address indexed user, uint256 zicoAmount, uint256 sharesMinted);
    event Unstaked(address indexed user, uint256 zicoAmount, uint256 sharesBurned);
    event TimelockTransferred(address indexed newTimelock);
    event ZicoTokenSet(address indexed newToken);

    modifier onlyTimelock() {
        require(msg.sender == timelock, "Not timelock");
        _;
    }

    constructor(address _zicoToken, address _timelock) ERC20("ZICO Staking Share", "zSHARE") Ownable(msg.sender) {
        require(_zicoToken != address(0) && _timelock != address(0), "Zero address");
        zicoToken = IZicoToken(_zicoToken);
        timelock = _timelock;
    }

    function transferTimelock(address newTimelock) external onlyTimelock {
        require(newTimelock != address(0), "Zero address");
        timelock = newTimelock;
        emit TimelockTransferred(newTimelock);
    }

    function setZicoToken(address newToken) external onlyTimelock {
        require(newToken != address(0), "Zero address");
        zicoToken = IZicoToken(newToken);
        emit ZicoTokenSet(newToken);
    }

    function stake(uint256 zicoAmount) external returns (uint256 sharesMinted) {
        require(zicoAmount > 0, "Cannot stake 0");
        uint256 totalZico = zicoTokenBalance();
        uint256 totalShares = totalSupply();
        if (totalShares == 0 || totalZico == 0) {
            sharesMinted = zicoAmount;
        } else {
            sharesMinted = (zicoAmount * totalShares) / totalZico;
        }
        require(sharesMinted > 0, "Zero shares");
        require(zicoToken.transferFrom(msg.sender, address(this), zicoAmount), "Transfer failed");
        _mint(msg.sender, sharesMinted);
        emit Staked(msg.sender, zicoAmount, sharesMinted);
    }

    function unstake(uint256 shareAmount) external returns (uint256 zicoReturned) {
        require(shareAmount > 0 && shareAmount <= balanceOf(msg.sender), "Invalid amount");
        uint256 totalShares = totalSupply();
        uint256 totalZico = zicoTokenBalance();
        zicoReturned = (shareAmount * totalZico) / totalShares;
        require(zicoReturned > 0, "Zero ZICO");
        _burn(msg.sender, shareAmount);
        require(zicoToken.transfer(msg.sender, zicoReturned), "Transfer failed");
        emit Unstaked(msg.sender, zicoReturned, shareAmount);
    }

    function zicoTokenBalance() public view returns (uint256) {
        return IERC20(address(zicoToken)).balanceOf(address(this));
    }
}
