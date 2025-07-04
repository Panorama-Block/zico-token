// SPDX-License-Identifier: MIT
// solhint-disable-next-line one-contract-per-file
pragma solidity 0.8.6;

import {VRFTypes} from "./VRFTypes.sol";

/**
 * @title BatchVRFCoordinatorV2
 * @notice The BatchVRFCoordinatorV2 contract acts as a proxy to write many random responses to the
 *   provided VRFCoordinatorV2 contract efficiently in a single transaction.
 */
contract BatchVRFCoordinatorV2 {
    // solhint-disable-next-line chainlink-solidity/prefix-immutable-variables-with-i
    VRFCoordinatorV2 public immutable COORDINATOR;

    event ErrorReturned(uint256 indexed requestId, string reason);
    event RawErrorReturned(uint256 indexed requestId, bytes lowLevelData);

    constructor(address coordinatorAddr) {
        COORDINATOR = VRFCoordinatorV2(coordinatorAddr);
    }

    /**
     * @notice fulfills multiple randomness requests with the provided proofs and commitments.
     * @param proofs the randomness proofs generated by the VRF provider.
     * @param rcs the request commitments corresponding to the randomness proofs.
     */
    function fulfillRandomWords(VRFTypes.Proof[] calldata proofs, VRFTypes.RequestCommitment[] memory rcs) external {
        // solhint-disable-next-line gas-custom-errors
        require(proofs.length == rcs.length, "input array arg lengths mismatch");
        for (uint256 i = 0; i < proofs.length; i++) {
            try COORDINATOR.fulfillRandomWords(proofs[i], rcs[i]) returns (uint96 /* payment */ ) {
                continue;
            } catch Error(string memory reason) {
                uint256 requestId = _getRequestIdFromProof(proofs[i]);
                emit ErrorReturned(requestId, reason);
            } catch (bytes memory lowLevelData) {
                uint256 requestId = _getRequestIdFromProof(proofs[i]);
                emit RawErrorReturned(requestId, lowLevelData);
            }
        }
    }

    /**
     * @notice Returns the proving key hash associated with this public key.
     * @param publicKey the key to return the hash of.
     */
    function _hashOfKey(uint256[2] memory publicKey) internal pure returns (bytes32) {
        return keccak256(abi.encode(publicKey));
    }

    /**
     * @notice Returns the request ID of the request associated with the given proof.
     * @param proof the VRF proof provided by the VRF oracle.
     */
    function _getRequestIdFromProof(VRFTypes.Proof memory proof) internal pure returns (uint256) {
        bytes32 keyHash = _hashOfKey(proof.pk);
        return uint256(keccak256(abi.encode(keyHash, proof.seed)));
    }
}

// solhint-disable-next-line interface-starts-with-i
interface VRFCoordinatorV2 {
    function fulfillRandomWords(VRFTypes.Proof calldata proof, VRFTypes.RequestCommitment memory rc)
        external
        returns (uint96);
}
