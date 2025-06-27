// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "chainlink/contracts/src/v0.8/ccip/applications/CCIPReceiver.sol";
import "chainlink/contracts/src/v0.8/ccip/interfaces/IRouterClient.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "chainlink/contracts/src/v0.8/vrf/VRFConsumerBaseV2.sol";
import "chainlink/contracts/src/v0.8/vrf/interfaces/VRFCoordinatorV2Interface.sol";

contract ZicoToken is ERC20, Ownable, CCIPReceiver, VRFConsumerBaseV2 {
    mapping(uint64 => address) public remotes;
    address public immutable linkToken;
    IRouterClient private router;
    address[] public stakerList;

    uint256 public totalStaked;
    mapping(address => uint256) public stakes;
    mapping(address => uint256) public rewards;

    VRFCoordinatorV2Interface COORDINATOR;
    uint64 public subscriptionId;
    bytes32 public keyHash;
    uint32 public callbackGasLimit = 100000;
    uint16 public requestConfirmations = 3;
    uint32 public numWords = 1;

    mapping(uint256 => uint256) public requestIdToReward;

    event CrossChainSend(uint64 indexed toChain, address indexed to, uint256 amount);
    event CrossChainReceive(address indexed to, uint256 amount);
    event Staked(address indexed user, uint256 amount);
    event Unstaked(address indexed user, uint256 amount);
    event RewardDistributed(address indexed user, uint256 reward);
    event RandomRewardRequested(uint256 requestId, uint256 rewardAmount);
    event RandomRewardGranted(address indexed winner, uint256 rewardAmount);

    constructor(address _router, address _linkToken, address vrfCoordinator, bytes32 _keyHash, uint64 _subscriptionId)
        ERC20("Zico Token", "ZICO")
        CCIPReceiver(_router)
        Ownable(msg.sender)
        VRFConsumerBaseV2(vrfCoordinator)
    {
        router = IRouterClient(_router);
        linkToken = _linkToken;
        COORDINATOR = VRFCoordinatorV2Interface(vrfCoordinator);
        keyHash = _keyHash;
        subscriptionId = _subscriptionId;

        _mint(msg.sender, 1_000_000_000 ether);
    }

    function stake(uint256 amount) external {
        require(amount > 0, "Cannot stake 0");
        _transfer(msg.sender, address(this), amount);

        if (stakes[msg.sender] == 0) {
            stakerList.push(msg.sender);
        }

        stakes[msg.sender] += amount;
        totalStaked += amount;
        emit Staked(msg.sender, amount);
    }

    function unstake(uint256 amount) external {
        require(amount > 0 && amount <= stakes[msg.sender], "Invalid unstake");
        stakes[msg.sender] -= amount;
        totalStaked -= amount;
        _transfer(address(this), msg.sender, amount);
        emit Unstaked(msg.sender, amount);
    }

    function distributeRewards() external onlyOwner {
        uint256 rewardPool = balanceOf(address(this)) / 100;
        address[] memory stakers = getAllStakers();
        for (uint256 i = 0; i < stakers.length; i++) {
            address stakerAddr = stakers[i];
            uint256 reward = (stakes[stakerAddr] * rewardPool) / totalStaked;
            if (reward > 0) {
                rewards[stakerAddr] += reward;
                emit RewardDistributed(stakerAddr, reward);
            }
        }
    }

    function claimReward() external {
        uint256 reward = rewards[msg.sender];
        require(reward > 0, "No rewards");
        rewards[msg.sender] = 0;
        _transfer(address(this), msg.sender, reward);
    }

    function sendCrossChain(uint64 destChain, uint256 amount) external {
        address remote = remotes[destChain];
        require(remote != address(0), "Unknown destination");

        _burn(msg.sender, amount);
        bytes memory data = abi.encode(msg.sender, amount);

        Client.EVMTokenAmount[] memory emptyTokenAmounts;

        Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
            receiver: abi.encode(remote),
            data: data,
            tokenAmounts: emptyTokenAmounts,
            extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: 200_000})),
            feeToken: linkToken
        });

        uint256 fee = router.getFee(destChain, message);
        require(IERC20(linkToken).allowance(msg.sender, address(this)) >= fee, "Insufficient LINK allowance");

        IERC20(linkToken).transferFrom(msg.sender, address(this), fee);
        IERC20(linkToken).approve(address(router), fee);

        router.ccipSend(destChain, message);
        emit CrossChainSend(destChain, msg.sender, amount);
    }

    function _ccipReceive(Client.Any2EVMMessage memory message) internal override onlyRouter {
        (address to, uint256 amount) = abi.decode(message.data, (address, uint256));
        _mint(to, amount);
        emit CrossChainReceive(to, amount);
    }

    function getAllStakers() internal view returns (address[] memory) {
        return stakerList;
    }

    function requestRandomReward(uint256 rewardAmount) external onlyOwner {
        require(stakerList.length > 0, "No stakers to reward");
        require(rewardAmount > 0 && rewardAmount <= balanceOf(address(this)), "Invalid reward amount");

        uint256 requestId =
            COORDINATOR.requestRandomWords(keyHash, subscriptionId, requestConfirmations, callbackGasLimit, numWords);

        requestIdToReward[requestId] = rewardAmount;

        emit RandomRewardRequested(requestId, rewardAmount);
    }

    function fulfillRandomWords(uint256 requestId, uint256[] memory randomWords) internal override {
        uint256 rewardAmount = requestIdToReward[requestId];
        require(rewardAmount > 0, "Reward not found");

        uint256 winnerIndex = randomWords[0] % stakerList.length;
        address winner = stakerList[winnerIndex];

        rewards[winner] += rewardAmount;

        delete requestIdToReward[requestId];

        emit RandomRewardGranted(winner, rewardAmount);
    }

    function setRemote(uint64 chainId, address remoteAddress) external onlyOwner {
        remotes[chainId] = remoteAddress;
    }
}
