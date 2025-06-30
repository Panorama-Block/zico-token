// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract ProofOfReserve {
    mapping(uint64 => uint256) public reportedSupplyByChain;
    uint64[] public chainSelectors;
    bytes32 public consolidatedHash;
    bool public supplyOK;
    uint256 public tolerancePercent = 1;
    address public oracle;

    event SupplyReported(uint64 indexed chainSelector, uint256 supply);
    event HashUpdated(bytes32 newHash, bool supplyOK);

    modifier onlyOracle() {
        require(msg.sender == oracle, "Not authorized");
        _;
    }

    constructor(address _oracle, uint64[] memory _chainSelectors) {
        oracle = _oracle;
        chainSelectors = _chainSelectors;
    }

    function setTolerancePercent(uint256 _percent) external onlyOracle {
        require(_percent <= 10, "Excessive tolerance");
        tolerancePercent = _percent;
    }

    function updateOracle(address newOracle) external onlyOracle {
        oracle = newOracle;
    }

    function reportSupply(uint64 chainSelector, uint256 supply) external onlyOracle {
        reportedSupplyByChain[chainSelector] = supply;
        emit SupplyReported(chainSelector, supply);
        _updateHashAndStatus();
    }

    function totalGlobalSupply() public view returns (uint256 total) {
        for (uint256 i = 0; i < chainSelectors.length; i++) {
            total += reportedSupplyByChain[chainSelectors[i]];
        }
    }

    function _updateHashAndStatus() internal {
        uint256 total = totalGlobalSupply();
        bytes32 newHash = keccak256(abi.encode(total));
        consolidatedHash = newHash;
        supplyOK = _isWithinTolerance();
        emit HashUpdated(newHash, supplyOK);
    }

    function _isWithinTolerance() internal view returns (bool) {
        uint256 total = totalGlobalSupply();
        for (uint256 i = 0; i < chainSelectors.length; i++) {
            uint256 s = reportedSupplyByChain[chainSelectors[i]];
            uint256 maxAllowed = (total * tolerancePercent) / 100;
            if (s > total + maxAllowed || s + maxAllowed < total) {
                return false;
            }
        }
        return true;
    }

    function isSupplyHealthy() external view returns (bool) {
        return supplyOK;
    }

    function getReportedSupply(uint64 chainSelector) external view returns (uint256) {
        return reportedSupplyByChain[chainSelector];
    }

    function getAllChains() external view returns (uint64[] memory) {
        return chainSelectors;
    }
}
