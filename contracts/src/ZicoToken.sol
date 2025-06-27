// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "chainlink/contracts/src/v0.8/ccip/applications/CCIPReceiver.sol";
import "chainlink/contracts/src/v0.8/ccip/interfaces/IRouterClient.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "chainlink/contracts/src/v0.8/vrf/VRFConsumerBaseV2.sol";
import "chainlink/contracts/src/v0.8/vrf/interfaces/VRFCoordinatorV2Interface.sol";
import {FunctionsClient} from "../lib/chainlink/contracts/src/v0.8/functions/v1_0_0/FunctionsClient.sol";
import {FunctionsRequest} from "../lib/chainlink/contracts/src/v0.8/functions/v1_0_0/libraries/FunctionsRequest.sol";

contract ZicoToken is ERC20, CCIPReceiver, VRFConsumerBaseV2, FunctionsClient {
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

    address public treasuryVault;
    uint16 public crossChainFeeBps = 50;
    uint16 public constant MAX_FEE_BPS = 1000;

    address public timelock;
    uint64 public functionsSubscriptionId;
    bytes32 public functionsDonId;
    uint32 public functionsCallbackGasLimit = 100000;
    uint256 public aprStaking;

    event APRUpdated(uint256 newApr);

    // Taxas de serviços premium
    mapping(bytes32 => uint256) public premiumServiceFees; // taxa em wei por serviço
    event PremiumServiceFeeSet(bytes32 indexed serviceId, uint256 fee);
    event PremiumServicePaid(address indexed user, bytes32 indexed serviceId, uint256 fee);

    // Integração com staking shares
    address public stakingContract;
    event StakingContractSet(address indexed stakingContract);

    // Pausa automática
    bool public paused;
    event Paused();
    event Unpaused();
    address public proofOfReserve;
    event ProofOfReserveSet(address indexed proofOfReserve);

    IERC20 public zicoToken;

    modifier onlyTimelock() {
        require(msg.sender == timelock, "Not timelock");
        _;
    }

    modifier whenNotPaused() {
        require(!paused, "Paused");
        _;
    }

    event CrossChainSend(uint64 indexed toChain, address indexed to, uint256 amount);
    event CrossChainReceive(address indexed to, uint256 amount);
    event Staked(address indexed user, uint256 amount, uint256 sharesMinted);
    event Unstaked(address indexed user, uint256 amount);
    event RewardDistributed(address indexed user, uint256 reward);
    event RandomRewardRequested(uint256 requestId, uint256 rewardAmount);
    event RandomRewardGranted(address indexed winner, uint256 rewardAmount);
    event TreasuryVaultUpdated(address indexed newVault);
    event CrossChainFeeUpdated(uint16 newFeeBps);

    constructor(
        address _router,
        address _linkToken,
        address vrfCoordinator,
        bytes32 _keyHash,
        uint64 _subscriptionId,
        address _timelock,
        address _functionsRouter,
        address _zicoToken
    )
        ERC20("ZICOAI", "ZICOAI")
        CCIPReceiver(_router)
        VRFConsumerBaseV2(vrfCoordinator)
        FunctionsClient(_functionsRouter)
    {
        router = IRouterClient(_router);
        linkToken = _linkToken;
        COORDINATOR = VRFCoordinatorV2Interface(vrfCoordinator);
        keyHash = _keyHash;
        subscriptionId = _subscriptionId;
        timelock = _timelock;
        zicoToken = IERC20(_zicoToken);
        _mint(msg.sender, 1_000_000_000 ether);
    }

    function transferTimelock(address newTimelock) external onlyTimelock {
        require(newTimelock != address(0), "Zero address");
        timelock = newTimelock;
    }

    function stake(uint256 zicoAmount) external whenNotPaused returns (uint256 sharesMinted) {
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

    function unstake(uint256 shareAmount) external whenNotPaused returns (uint256 zicoReturned) {
        require(shareAmount > 0 && shareAmount <= balanceOf(msg.sender), "Invalid amount");
        uint256 totalShares = totalSupply();
        uint256 totalZico = zicoTokenBalance();
        zicoReturned = (shareAmount * totalZico) / totalShares;
        require(zicoReturned > 0, "Zero ZICO");
        _burn(msg.sender, shareAmount);
        require(zicoToken.transfer(msg.sender, zicoReturned), "Transfer failed");
        emit Unstaked(msg.sender, shareAmount);
    }

    function zicoTokenBalance() public view returns (uint256) {
        return IERC20(address(zicoToken)).balanceOf(address(this));
    }

    function distributeRewards() external onlyTimelock {
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

    function sendCrossChain(uint64 destChain, uint256 amount) external whenNotPaused {
        address remote = remotes[destChain];
        require(remote != address(0), "Unknown destination");
        require(treasuryVault != address(0), "TreasuryVault not set");
        require(amount > 0, "Amount must be > 0");

        // Calcula a taxa
        uint256 feeAmount = (amount * crossChainFeeBps) / 10000;
        uint256 netAmount = amount - feeAmount;
        require(netAmount > 0, "Net amount must be > 0");

        _transfer(msg.sender, treasuryVault, feeAmount);
        _burn(msg.sender, netAmount);

        bytes memory data = abi.encode(msg.sender, netAmount);
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
        emit CrossChainSend(destChain, msg.sender, netAmount);
    }

    function _ccipReceive(Client.Any2EVMMessage memory message) internal override onlyRouter {
        address expectedRemote = remotes[message.sourceChainSelector];
        require(expectedRemote != address(0), "Remote not set for chain");
        address sender = abi.decode(message.sender, (address));
        require(sender == expectedRemote, "Unauthorized CCIP sender");
        (address to, uint256 amount) = abi.decode(message.data, (address, uint256));
        _mint(to, amount);
        emit CrossChainReceive(to, amount);
    }

    function getAllStakers() internal view returns (address[] memory) {
        return stakerList;
    }

    function requestRandomReward(uint256 rewardAmount) external onlyTimelock {
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

    function setRemote(uint64 chainId, address remoteAddress) external onlyTimelock {
        remotes[chainId] = remoteAddress;
    }

    function setTreasuryVault(address _treasuryVault) external onlyTimelock {
        require(_treasuryVault != address(0), "Zero address");
        treasuryVault = _treasuryVault;
        emit TreasuryVaultUpdated(_treasuryVault);
    }

    function setCrossChainFeeBps(uint16 _feeBps) external onlyTimelock {
        require(_feeBps <= MAX_FEE_BPS, "Fee too high");
        crossChainFeeBps = _feeBps;
        emit CrossChainFeeUpdated(_feeBps);
    }

    function setFunctionsConfig(uint64 _subId, bytes32 _donId, uint32 _callbackGas) external onlyTimelock {
        functionsSubscriptionId = _subId;
        functionsDonId = _donId;
        functionsCallbackGasLimit = _callbackGas;
    }

    using FunctionsRequest for FunctionsRequest.Request;

    function requestAprUpdate(string calldata source, string[] calldata args) external onlyTimelock returns (bytes32) {
        FunctionsRequest.Request memory req;
        req.initializeRequestForInlineJavaScript(source);
        if (args.length > 0) {
            req.setArgs(args);
        }
        bytes memory requestData = FunctionsRequest.encodeCBOR(req);
        return _sendRequest(requestData, functionsSubscriptionId, functionsCallbackGasLimit, functionsDonId);
    }

    function fulfillRequest(bytes32, bytes memory response, bytes memory err) internal override {
        require(err.length == 0, "Chainlink Functions error");
        uint256 newApr = abi.decode(response, (uint256));
        aprStaking = newApr;
        emit APRUpdated(newApr);
    }

    function setPremiumServiceFee(bytes32 serviceId, uint256 fee) external onlyTimelock {
        premiumServiceFees[serviceId] = fee;
        emit PremiumServiceFeeSet(serviceId, fee);
    }

    // Exemplo: exigir 1000 shares para acessar serviços premium
    uint256 public constant PREMIUM_SERVICE_MIN_SHARES = 1000 ether;
    function payPremiumService(bytes32 serviceId) external onlyStakingLevel(PREMIUM_SERVICE_MIN_SHARES) whenNotPaused {
        uint256 fee = premiumServiceFees[serviceId];
        require(fee > 0, "Service not available");
        require(treasuryVault != address(0), "TreasuryVault not set");
        _transfer(msg.sender, treasuryVault, fee);
        emit PremiumServicePaid(msg.sender, serviceId, fee);
    }

    function setStakingContract(address _stakingContract) external onlyTimelock {
        require(_stakingContract != address(0), "Zero address");
        stakingContract = _stakingContract;
        emit StakingContractSet(_stakingContract);
    }

    function hasStakingLevel(address user, uint256 minShares) public view returns (bool) {
        if (stakingContract == address(0)) return false;
        return IERC20(stakingContract).balanceOf(user) >= minShares;
    }

    modifier onlyStakingLevel(uint256 minShares) {
        require(hasStakingLevel(msg.sender, minShares), "Insufficient staking level");
        _;
    }

    function pause() external onlyTimelock {
        paused = true;
        emit Paused();
    }

    function unpause() external onlyTimelock {
        paused = false;
        emit Unpaused();
    }

    function setProofOfReserve(address _proofOfReserve) external onlyTimelock {
        require(_proofOfReserve != address(0), "Zero address");
        proofOfReserve = _proofOfReserve;
        emit ProofOfReserveSet(_proofOfReserve);
    }

    // Função para Chainlink Functions pausar automaticamente se supplyOK for falso
    function autoPauseIfUnhealthy() external {
        require(proofOfReserve != address(0), "ProofOfReserve not set");
        (bool success, bytes memory data) = proofOfReserve.staticcall(abi.encodeWithSignature("isSupplyHealthy()"));
        require(success, "ProofOfReserve call failed");
        bool supplyOK = abi.decode(data, (bool));
        if (!supplyOK && !paused) {
            paused = true;
            emit Paused();
        }
    }
}
