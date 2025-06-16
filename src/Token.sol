// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "chainlink/contracts/src/v0.8/ccip/applications/CCIPReceiver.sol";
import "chainlink/contracts/src/v0.8/ccip/interfaces/IRouterClient.sol";
// import "chainlink/contracts/src/v0.8/functions/dev/v1_X/FunctionsClient.sol";
// import "chainlink/contracts/src/v0.8/functions/dev/v1_X/libraries/FunctionsRequest.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract ZicoToken is ERC20, Ownable, CCIPReceiver {
    // using FunctionsRequest for FunctionsRequest.Request;

    mapping(uint64 => address) public remotes;
    // string public lastFunctionResponse;
    address public immutable linkToken;
    IRouterClient private router;

    constructor(
        address _router,
        address _linkToken
        // address _oracle
    ) 
        ERC20("Zico Token", "ZICO")
        Ownable(msg.sender)
        CCIPReceiver(_router)
        // FunctionsClient(_oracle)
    {
        router = IRouterClient(_router);
        linkToken = _linkToken;

        _mint(msg.sender, 1_000_000 ether);
    }

    function setRemote(uint64 chainSelector, address remote) external onlyOwner {
        remotes[chainSelector] = remote;
    }

    function sendCrossChain(
        uint64 destinationChainSelector,
        // address receiver,
        uint256 amount
    ) external {
        address remote = remotes[destinationChainSelector];
        require(remote != address(0), "Unknown destination");

        _burn(msg.sender, amount);

        // bytes memory data = abi.encode(receiver, amount);

        // Mensagem e envio com Chainlink CCIP (comentado por depender de tipos do Client)
        /*
        Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
            receiver: abi.encode(remote),
            data: data,
            tokenAmounts: new Client.EVMTokenAmount ,
            extraArgs: abi.encode(""),
            feeToken: linkToken
        });

        uint256 fee = router.getFee(destinationChainSelector, message);

        require(IERC20(linkToken).transferFrom(msg.sender, address(this), fee), "Fee transfer failed");
        IERC20(linkToken).approve(address(router), fee);

        router.ccipSend(destinationChainSelector, message);
        */
    }

    function _ccipReceive(Client.Any2EVMMessage memory /* message */) internal override onlyRouter {
    // Ainda não implementado
    }

    /*
    function executeFunction(bytes calldata sourceCode) external onlyOwner {
        FunctionsRequest.Request memory req;
        req.source = string(sourceCode);
        req.codeLocation = FunctionsRequest.Location.Inline;
        req.language = FunctionsRequest.CodeLanguage.JavaScript;

        bytes memory encodedRequest = req.encodeCBOR();
        _sendRequest(encodedRequest, 200_000, bytes32(0));
    }

    function _fulfillRequest(
        bytes32,
        bytes memory response,
        bytes memory
    ) internal override {
        lastFunctionResponse = string(response);
    }
    */
}
