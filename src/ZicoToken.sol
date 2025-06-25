// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "chainlink/contracts/src/v0.8/ccip/applications/CCIPReceiver.sol";
import "chainlink/contracts/src/v0.8/ccip/interfaces/IRouterClient.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract ZICOToken is ERC20, Ownable, CCIPReceiver {
    mapping(uint64 => address) public remotes;
    address public immutable linkToken;
    IRouterClient private router;

    event CrossChainSend(uint64 indexed toChain, address indexed to, uint256 amount);
    event CrossChainReceive(address indexed to, uint256 amount);

    constructor(
        address _router,
        address _linkToken
    )
        ERC20("ZICOAI", "ZICOAI")
        CCIPReceiver(_router)
        Ownable(msg.sender)
    {
        router = IRouterClient(_router);
        linkToken = _linkToken;

        _mint(msg.sender, 1_000_000 ether);
    }

    function setRemote(uint64 chainId, address remoteAddress) external onlyOwner {
        remotes[chainId] = remoteAddress;
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
            extraArgs: Client._argsToBytes(
                Client.EVMExtraArgsV1({gasLimit: 200_000})
            ),
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
}
