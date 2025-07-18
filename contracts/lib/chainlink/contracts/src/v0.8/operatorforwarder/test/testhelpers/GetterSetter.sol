pragma solidity ^0.8.0;

// GetterSetter is a contract to aid debugging and testing during development.
// solhint-disable
contract GetterSetter {
    bytes32 private s_getBytes32;
    uint256 private s_getUint256;
    bytes32 private s_requestId;
    bytes private s_getBytes;

    event SetBytes32(address indexed from, bytes32 indexed value);
    event SetUint256(address indexed from, uint256 indexed value);
    event SetBytes(address indexed from, bytes value);

    event Output(bytes32 b32, uint256 u256, bytes32 b322);

    function setBytes32(bytes32 _value) public {
        s_getBytes32 = _value;
        emit SetBytes32(msg.sender, _value);
    }

    function requestedBytes32(bytes32 _requestId, bytes32 _value) public {
        s_requestId = _requestId;
        setBytes32(_value);
    }

    function setBytes(bytes memory _value) public {
        s_getBytes = _value;
        emit SetBytes(msg.sender, _value);
    }

    function getBytes() public view returns (bytes memory _value) {
        return s_getBytes;
    }

    function requestedBytes(bytes32 _requestId, bytes memory _value) public {
        s_requestId = _requestId;
        setBytes(_value);
    }

    function setUint256(uint256 _value) public {
        s_getUint256 = _value;
        emit SetUint256(msg.sender, _value);
    }

    function requestedUint256(bytes32 _requestId, uint256 _value) public {
        s_requestId = _requestId;
        setUint256(_value);
    }
}
