// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token_governor

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated"
)

var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

var TokenGovernorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"initialDefaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BRIDGE_MINTER_OR_BURNER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BURNER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CHECKER_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FREEZER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"RECOVERY_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UNFREEZER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UNPAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchDrainFrozenAccounts\",\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchFreeze\",\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchUnfreeze\",\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"beginDefaultAdminTransfer\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDefaultAdminTransfer\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDefaultAdminDelay\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"defaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultAdminDelayIncreaseWait\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"drainFrozenAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeTokenFunction\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"freeze\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAdmins\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBridgeMintersOrBurners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBurners\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChecker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCheckerAdmins\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFreezers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinters\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPausers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRecoveryManagers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMember\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMemberCount\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleMembers\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnfreezers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUnpausers\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingDefaultAdminDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"recoverGovernedTokenERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rollbackDefaultAdminDelay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setChecker\",\"inputs\":[{\"name\":\"newChecker\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unfreeze\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AccountFrozen\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AccountUnfrozen\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BridgeBurn\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BridgeMint\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CheckerUpdated\",\"inputs\":[{\"name\":\"previousChecker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newChecker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractPaused\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ContractUnpaused\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminDelayChangeScheduled\",\"inputs\":[{\"name\":\"newDelay\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"effectSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferCanceled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultAdminTransferScheduled\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"acceptSchedule\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FrozenAccountDrained\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovernedTokensRecovered\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NativeBurn\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NativeMint\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipAccepted\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenFunctionExecuted\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensRecovered\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminDelay\",\"inputs\":[{\"name\":\"schedule\",\"type\":\"uint48\",\"internalType\":\"uint48\"}]},{\"type\":\"error\",\"name\":\"AccessControlEnforcedDefaultAdminRules\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlInvalidDefaultAdmin\",\"inputs\":[{\"name\":\"defaultAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"BurnFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallFailed\",\"inputs\":[{\"name\":\"error\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InvalidFrom\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRecipient\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MintFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyBurnerOrBridge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyMinterOrBridge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddressNotAllowed\",\"inputs\":[]}]",
	Bin: "0x60a0346101bf57601f61363638819003918201601f19168301916001600160401b038311848410176101c4578084926060946040528339810103126101bf57610047816101da565b9060208101519065ffffffffffff821682036101bf57604061006991016101da565b6001600160a01b0381169182156101a957600180546001600160d01b031660d09290921b6001600160d01b031916919091179055600254906001600160a01b038216610198576001600160a01b031990911682176002556100c9906101ee565b61015e575b506001600160a01b0316801561014d5760805260405161331f90816102f78239608051818181610335015281816104f901528181610b5a01528181611090015281816111b8015281816112f40152818161184e01528181611cb7015281816124b801528181612574015281816126990152818161299c0152612bf90152f35b6342bcdf7f60e11b60005260046000fd5b600080526004602052610191907f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec61027c565b50386100ce565b631fe1e13d60e11b60005260046000fd5b636116401160e11b600052600060045260246000fd5b600080fd5b634e487b7160e01b600052604160045260246000fd5b51906001600160a01b03821682036101bf57565b6001600160a01b0381166000908152600080516020613616833981519152602052604090205460ff16610276576001600160a01b0316600081815260008051602061361683398151915260205260408120805460ff191660011790553391907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b50600090565b60018101908260005281602052604060002054156000146102ee578054680100000000000000008110156101c457600181018083558110156102d857839082600052602060002001555491600052602052604060002055600190565b634e487b7160e01b600052603260045260246000fd5b50505060009056fe608080604052600436101561001357600080fd5b600090813560e01c90816301ffc9a714611f4057508063022d63fb14611f225780630568b3da14611ee757806306a85f0f14611eac5780630aa6220b14611df05780631171bda914611d5c5780631e42b8b114611cdb57806321df0da714611c97578063248a9ca314611c61578063248e6c0814611c26578063282c51f314611beb5780632c92ba9c14611bc65780632dbc9db914611b885780632f2ff15d14611b3f578063312963e514611b0457806331993a1c14611ac957806331ae450b14611a6857806336568abe1461192d5780633f4ba83a146117eb57806340c10f19146117c557806342966c68146117a757806345c8b1a61461177d57806360ea92081461173f578063634e93da1461161f578063649a5ec7146114165780636b32810b1461139557806379ba5097146112d057806379cc6790146112aa5780638118efda1461122957806382b2e2571461116b5780638456cb591461102b57806384ef8ffc14610f5657806386fe8b4314610faa5780638d1fdf2f14610f7d5780638da5cb5b14610f565780639010d07c14610f1157806391d1485414610ec55780639ddc3f5c14610e44578063a0712d6814610e1b578063a1eda53c14610db8578063a217fddf14610d9c578063a3246ad314610d38578063a6cc8d4014610cb7578063c35f7c6214610c36578063c721e8c014610b37578063ca15c87314610b0d578063cc8463c814610ae2578063cefc1429146109ed578063cf6eefb71461099a578063cf880f4c146108a1578063d22224a914610859578063d53913931461081e578063d547741f146107a4578063d5c8514014610723578063d602b9fd146106a7578063e0c0906f14610680578063e63ab1e914610645578063e78a14f5146105b8578063f2fde38b146104cc578063fb1bb9de146104915763fcedb9e5146102c057600080fd5b602060031936011261048e576004359067ffffffffffffffff821161048e573660238301121561048e5781600401359167ffffffffffffffff831161048a576024810190602484369201011161048a57610318612267565b818060405185848237808681018381520390346001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165af1913d15610482573d9267ffffffffffffffff841161045557604051936103876020601f19601f84011601866121dc565b84523d82602086013e5b1561041657601f19601f857f3f341ef37380014c90cd66348f1e20455df1b642e6a3445f39dd66f88110cc659361041297604051966040885281604089015260608801376060828701015201168201916060818403016020820152806103fb606033950186612148565b0390a2604051918291602083526020830190612148565b0390f35b6040517fa5fa8d2b00000000000000000000000000000000000000000000000000000000815260206004820152806104516024820186612148565b0390fd5b6024827f4e487b710000000000000000000000000000000000000000000000000000000081526041600452fd5b606092610391565b5080fd5b80fd5b503461048e578060031936011261048e5760206040517f427da25fe773164f88948d3e215c94b6554e2ed5e5f203a821c9f2f6131cf75a8152f35b503461048e57602060031936011261048e576104e6612046565b6104ee612267565b816001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001691823b1561048a57816001600160a01b036024829360405194859384927ff2fde38b00000000000000000000000000000000000000000000000000000000845216978860048401525af180156105ad57610598575b5050337f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08380a380f35b816105a2916121dc565b61048a57813861056e565b6040513d84823e3d90fd5b503461048e578060031936011261048e577f3e02eaefb22229c9fa4ecb927d4a3b0bd2d30b1af650a970c5a95ba1f96906ea8152600460205260408120604051918260208354918281520192825260208220915b81811061062f5761041285610623818703826121dc565b604051918291826120b1565b825484526020909301926001928301920161060c565b503461048e578060031936011261048e5760206040517f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8152f35b503461048e578060031936011261048e5760206001600160a01b0360035416604051908152f35b503461048e578060031936011261048e576106c0612267565b600180547fffffffffffff0000000000000000000000000000000000000000000000000000811690915560a01c65ffffffffffff166106fc5780f35b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a96051098180a180f35b503461048e578060031936011261048e577fb9f997e9b8b4ceb22077d832bcbf6f751de391a9ad51c4c81e58db4b9214f8e38152600460205260408120604051918260208354918281520192825260208220915b81811061078e5761041285610623818703826121dc565b8254845260209093019260019283019201610777565b503461048e57604060031936011261048e576004356107c1612061565b9080156107f65790816107ed6107e86107f294600052600060205260016040600020015490565b612410565b612f45565b5080f35b6004837f3fc3c27a000000000000000000000000000000000000000000000000000000008152fd5b503461048e578060031936011261048e5760206040517f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a68152f35b503461048e57610868366120f4565b610870612267565b825b81811061087d578380f35b8061089b6108966108916001948688612189565b6121c8565b6124ae565b01610872565b503461048e57602060031936011261048e576108bb612046565b7fb9f997e9b8b4ceb22077d832bcbf6f751de391a9ad51c4c81e58db4b9214f8e3825281602052604082206001600160a01b03331660005260205260ff6040600020541615610963576001600160a01b0380600354921691827fffffffffffffffffffffffff0000000000000000000000000000000000000000821617600355167f35631e65f3bb46b57c2cfadf0d2ca3f08e6cbab52c2ff987b5328c9b37d45ea38380a380f35b60448263e2517d3f60e01b8152336004527fb9f997e9b8b4ceb22077d832bcbf6f751de391a9ad51c4c81e58db4b9214f8e3602452fd5b503461048e578060031936011261048e57604065ffffffffffff6109d46001549065ffffffffffff6001600160a01b0383169260a01c1690565b6001600160a01b03849392935193168352166020820152f35b503461048e578060031936011261048e576001546001600160a01b03163303610ab6576001546001600160a01b0381169060a01c65ffffffffffff1680158015610aac575b610a815750610a5590610a4f6001600160a01b0360025416612eb2565b50612d75565b507fffffffffffff00000000000000000000000000000000000000000000000000006001541660015580f35b7f19ca5ebb000000000000000000000000000000000000000000000000000000008352600452602482fd5b5042811015610a32565b807fc22c8022000000000000000000000000000000000000000000000000000000006024925233600452fd5b503461048e578060031936011261048e576020610afd61222e565b65ffffffffffff60405191168152f35b503461048e57602060031936011261048e5760406020916004358152600483522054604051908152f35b503461048e57610b4636612077565b90610b4f6122ba565b836001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b1561048a576040517f1171bda90000000000000000000000000000000000000000000000000000000081526001600160a01b03868116600483015284166024820152604481018590529082908290606490829084905af180156105ad57610c1d575b50506001600160a01b038091604051938452169216907f42f446384ff78c62bf2caa065e680c6010e2f54857d15bee828d6c7d5735389360203392a480f35b81610c27916121dc565b610c32578338610bde565b8380fd5b503461048e578060031936011261048e577fac950b4e2512fad4244af0109db80898a61d13325e9e992206b6f9cf76056ff98152600460205260408120604051918260208354918281520192825260208220915b818110610ca15761041285610623818703826121dc565b8254845260209093019260019283019201610c8a565b503461048e578060031936011261048e577f427da25fe773164f88948d3e215c94b6554e2ed5e5f203a821c9f2f6131cf75a8152600460205260408120604051918260208354918281520192825260208220915b818110610d225761041285610623818703826121dc565b8254845260209093019260019283019201610d0b565b503461048e57602060031936011261048e576004358152600460205260408120604051918260208354918281520192825260208220915b818110610d865761041285610623818703826121dc565b8254845260209093019260019283019201610d6f565b503461048e578060031936011261048e57602090604051908152f35b503461048e578060031936011261048e576002548060d01c9182151580610e11575b15610e08575060a01c65ffffffffffff165b6040805165ffffffffffff928316815292909116602083015290f35b91505080610dec565b5042831015610dda565b503461048e57602060031936011261048e57610e3960043533612626565b602060405160018152f35b503461048e578060031936011261048e577f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a8152600460205260408120604051918260208354918281520192825260208220915b818110610eaf5761041285610623818703826121dc565b8254845260209093019260019283019201610e98565b503461048e57604060031936011261048e576001600160a01b036040610ee9612061565b92600435815280602052209116600052602052602060ff604060002054166040519015158152f35b503461048e57604060031936011261048e576001600160a01b03610f4660209260043581526004845260406024359120613034565b90549060031b1c16604051908152f35b503461048e578060031936011261048e5760206001600160a01b0360025416604051908152f35b503461048e57602060031936011261048e57610fa7610f9a612046565b610fa261239e565b612bef565b80f35b503461048e578060031936011261048e577f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a8488152600460205260408120604051918260208354918281520192825260208220915b8181106110155761041285610623818703826121dc565b8254845260209093019260019283019201610ffe565b503461048e578060031936011261048e577f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a815280602052604081206001600160a01b03331660005260205260ff604060002054161561113457806001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b15611131578180916004604051809481937f8456cb590000000000000000000000000000000000000000000000000000000083525af180156105ad5761111c575b50337f81990fd9a5c552b8e3677917d8a03c07678f0d2cb68f88b634aca2022e9bd19f8280a280f35b81611126916121dc565b61048e5780386110f3565b50fd5b8063e2517d3f60e01b60449252336004527f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a602452fd5b503461048e578060031936011261048e576040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526020816024816001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000165afa9081156105ad5782916111f3575b602082604051908152f35b90506020813d602011611221575b8161120e602093836121dc565b8101031261048a576020915051386111e8565b3d9150611201565b503461048e578060031936011261048e577f92de27771f92d6942691d73358b3a4673e4880de8356f8f2cf452be87e02d3638152600460205260408120604051918260208354918281520192825260208220915b8181106112945761041285610623818703826121dc565b825484526020909301926001928301920161127d565b503461048e57604060031936011261048e57610e396112c7612046565b602435906128e3565b503461048e578060031936011261048e576112e9612267565b806001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b15611131578180916004604051809481937f79ba50970000000000000000000000000000000000000000000000000000000083525af180156105ad57611380575b50337fb27970c1714b28277b78cc17ac2fe9556e7f048cd48358cffe3dc7d547608fdc8280a280f35b8161138a916121dc565b61048e578038611357565b503461048e578060031936011261048e577f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a68152600460205260408120604051918260208354918281520192825260208220915b8181106114005761041285610623818703826121dc565b82548452602090930192600192830192016113e9565b503461048e57602060031936011261048e5760043565ffffffffffff811680820361161b57611443612267565b61144c42612fea565b9065ffffffffffff61145c61222e565b16808211156115b357507ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b929165ffffffffffff8262069780806114aa951091180262069780181690612cab565b906002548060d01c80611530575b5050600280546001600160a01b031660a083901b79ffffffffffff0000000000000000000000000000000000000000161760d084901b7fffffffffffff0000000000000000000000000000000000000000000000000000161790556040805165ffffffffffff9283168152919092166020820152a180f35b4211156115895779ffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffff00000000000000000000000000000000000000000000000000006001549260301b169116176001555b38806114b8565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec58480a1611582565b0365ffffffffffff81116115ee577ff1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b92916114aa9190612cab565b6024847f4e487b710000000000000000000000000000000000000000000000000000000081526011600452fd5b8280fd5b503461048e57602060031936011261048e57611639612046565b611641612267565b7f3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6602061167e61167042612fea565b61167861222e565b90612cab565b65ffffffffffff6001600160a01b036116ad6001549065ffffffffffff6001600160a01b0383169260a01c1690565b9690501694600154867fffffffffffff000000000000000000000000000000000000000000000000000079ffffffffffff00000000000000000000000000000000000000008660a01b169216171760015516611716575b65ffffffffffff60405191168152a280f35b7f8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a96051098580a1611704565b503461048e5761174e366120f4565b61175661239e565b825b818110611763578380f35b80611777610fa26108916001948688612189565b01611758565b503461048e57602060031936011261048e57610fa761179a612046565b6117a261232c565b61256a565b503461048e57602060031936011261048e57610e39600435336128e3565b503461048e57604060031936011261048e57610e396117e2612046565b60243590612626565b503461048e578060031936011261048e577f427da25fe773164f88948d3e215c94b6554e2ed5e5f203a821c9f2f6131cf75a815280602052604081206001600160a01b033316825260205260ff604082205416156118f657806001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016803b15611131578180916004604051809581937f3f4ba83a0000000000000000000000000000000000000000000000000000000083525af180156118e9576118d9575b337f5b65b0c1363b3003db9bcc5e1fd8805a6d6bf5bf6dc9d3431ee4494cd7d117668280a280f35b6118e2916121dc565b38816118b1565b50604051903d90823e3d90fd5b8063e2517d3f60e01b60449252336004527f427da25fe773164f88948d3e215c94b6554e2ed5e5f203a821c9f2f6131cf75a602452fd5b503461048e57604060031936011261048e5760043561194a612061565b90801580611a4b575b61199a575b336001600160a01b0383160361197257906107f291612f45565b6004837f6697b232000000000000000000000000000000000000000000000000000000008152fd5b60015465ffffffffffff60a082901c16906001600160a01b031615801590611a3b575b8015611a29575b6119f657507fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff60015416600155611958565b7f19ca5ebb00000000000000000000000000000000000000000000000000000000845265ffffffffffff16600452602483fd5b504265ffffffffffff821610156119c4565b5065ffffffffffff8116156119bd565b506001600160a01b03600254166001600160a01b03831614611953565b503461048e578060031936011261048e57808052600460205260408120604051918260208354918281520192825260208220915b818110611ab35761041285610623818703826121dc565b8254845260209093019260019283019201611a9c565b503461048e578060031936011261048e5760206040517f0acf805600123ef007091da3b3ffb39474074c656c127aa68cb0ffec232a8ff88152f35b503461048e578060031936011261048e5760206040517fac950b4e2512fad4244af0109db80898a61d13325e9e992206b6f9cf76056ff98152f35b503461048e57604060031936011261048e57600435611b5c612061565b9080156107f6579081611b836107e86107f294600052600060205260016040600020015490565b612e2f565b503461048e57611b97366120f4565b611b9f61232c565b825b818110611bac578380f35b80611bc06117a26108916001948688612189565b01611ba1565b503461048e57602060031936011261048e57610fa7611be3612046565b610896612267565b503461048e578060031936011261048e5760206040517f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a8488152f35b503461048e578060031936011261048e5760206040517f3e02eaefb22229c9fa4ecb927d4a3b0bd2d30b1af650a970c5a95ba1f96906ea8152f35b503461048e57602060031936011261048e576020611c8f600435600052600060205260016040600020015490565b604051908152f35b503461048e578060031936011261048e5760206040516001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461048e578060031936011261048e577f0acf805600123ef007091da3b3ffb39474074c656c127aa68cb0ffec232a8ff88152600460205260408120604051918260208354918281520192825260208220915b818110611d465761041285610623818703826121dc565b8254845260209093019260019283019201611d2f565b503461048e57611d6b36612077565b611d736122ba565b6001600160a01b038216928315611dc857816001600160a01b03611d9992169384612457565b6040519081527fa2231b10d9b4e4166c8a827c99f97691b05aa88fb04e009a4e499005b5c50fcc60203392a480f35b6004857f8579befe000000000000000000000000000000000000000000000000000000008152fd5b503461048e578060031936011261048e57611e09612267565b6002548060d01c80611e29575b826001600160a01b036002541660025580f35b421115611e825779ffffffffffffffffffffffffffffffffffffffffffffffffffff7fffffffffffff00000000000000000000000000000000000000000000000000006001549260301b169116176001555b3880611e16565b507f2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec58180a1611e7b565b503461048e578060031936011261048e5760206040517f92de27771f92d6942691d73358b3a4673e4880de8356f8f2cf452be87e02d3638152f35b503461048e578060031936011261048e5760206040517fb9f997e9b8b4ceb22077d832bcbf6f751de391a9ad51c4c81e58db4b9214f8e38152f35b503461048e578060031936011261048e576020604051620697808152f35b90503461048a57602060031936011261048a576004357fffffffff00000000000000000000000000000000000000000000000000000000811680910361161b57602092507f5a05180f000000000000000000000000000000000000000000000000000000008114908115611fb6575b5015158152f35b7f3149878600000000000000000000000000000000000000000000000000000000811491508115611fe9575b5038611faf565b7f7965db0b0000000000000000000000000000000000000000000000000000000081149150811561201c575b5038611fe2565b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501438612015565b600435906001600160a01b038216820361205c57565b600080fd5b602435906001600160a01b038216820361205c57565b600319606091011261205c576004356001600160a01b038116810361205c57906024356001600160a01b038116810361205c579060443590565b602060408183019282815284518094520192019060005b8181106120d55750505090565b82516001600160a01b03168452602093840193909201916001016120c8565b90602060031983011261205c5760043567ffffffffffffffff811161205c578260238201121561205c5780600401359267ffffffffffffffff841161205c5760248460051b8301011161205c576024019190565b919082519283825260005b848110612174575050601f19601f8460006020809697860101520116010190565b80602080928401015182828601015201612153565b91908110156121995760051b0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b356001600160a01b038116810361205c5790565b90601f601f19910116810190811067ffffffffffffffff8211176121ff57604052565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6002548060d01c801515908161225d575b50156122535760a01c65ffffffffffff1690565b5060015460d01c90565b905042113861223f565b3360009081527fad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5602052604090205460ff16156122a057565b63e2517d3f60e01b60005233600452600060245260446000fd5b3360009081527ff5d6910bc70e1ece69b520193ad380ecb31507820a100769882d4a52a371eb59602052604090205460ff16156122f357565b63e2517d3f60e01b600052336004527f0acf805600123ef007091da3b3ffb39474074c656c127aa68cb0ffec232a8ff860245260446000fd5b3360009081527fb9e6e2bd24b42e34daa2106ac2e8d73724be4745c0904ad7422f4ace20dc7eec602052604090205460ff161561236557565b63e2517d3f60e01b600052336004527f3e02eaefb22229c9fa4ecb927d4a3b0bd2d30b1af650a970c5a95ba1f96906ea60245260446000fd5b3360009081527f186963de052711d3805bfc5d362b0e04df96e6af4837bc69a56a0b6d203c7b6c602052604090205460ff16156123d757565b63e2517d3f60e01b600052336004527f92de27771f92d6942691d73358b3a4673e4880de8356f8f2cf452be87e02d36360245260446000fd5b80600052600060205260406000206001600160a01b03331660005260205260ff604060002054161561243f5750565b63e2517d3f60e01b6000523360045260245260446000fd5b6124ac926001600160a01b03604051937fa9059cbb0000000000000000000000000000000000000000000000000000000060208601521660248401526044830152604482526124a76064836121dc565b612cf8565b565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690813b1561205c576001600160a01b0390604051907fece531320000000000000000000000000000000000000000000000000000000082528160248160008096819516978860048401525af180156105ad57829061255a575b50507f8e535556a0a95ee3befe296cf986f7bf7d88881991e46f517c4b477c0ea69385339180a3565b612563916121dc565b3881612531565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690813b1561205c576001600160a01b0390604051907f45c8b1a60000000000000000000000000000000000000000000000000000000082528160248160008096819516978860048401525af180156105ad578290612616575b50507fe19c610e04dba2019efcfb0f9455fad3af646853bb02abad2a452db1fd47c327339180a3565b61261f916121dc565b38816125ed565b3360009081527f1fe763eeb6aa2988d40de1f3b4957809fae3f707a30732e8397ed70789991100602052604081205490929160ff9091169081158061289a575b612872576001600160a01b0381169330851461284a576001600160a01b0360035416806127c8575b506001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016906040517fa0712d6800000000000000000000000000000000000000000000000000000000815285600482015260208160248185875af19081156105ad578291612799575b50156127715750612710918491612457565b15612743576040519081527fc3b8ae385c02b938fbbbd694d1da0761f755ee2c942f26bbfc6723c986d31b9b60203392a3565b6040519081527fc8a07fbf83d1ff93486eea14a0676adef77c8c95fa79c85cd7bba0d10204102f60203392a3565b807f07637bd80000000000000000000000000000000000000000000000000000000060049252fd5b6127bb915060203d6020116127c1575b6127b381836121dc565b810190612fd2565b386126fe565b503d6127a9565b803b1561048a576040517ffa35dc460000000000000000000000000000000000000000000000000000000081526001600160a01b03841660048201526024810186905233604482015284151560648201529082908290608490829084905af180156105ad57829061283a575b5061268e565b612843916121dc565b3881612834565b807f9c8d2cd20000000000000000000000000000000000000000000000000000000060049252fd5b6004847fc7ccc1ca000000000000000000000000000000000000000000000000000000008152fd5b507f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a6845283602052604084206001600160a01b033316855260205260ff60408520541615612666565b3360009081527f1fe763eeb6aa2988d40de1f3b4957809fae3f707a30732e8397ed70789991100602052604081205491929160ff1690811580612ba6575b612b7e576001600160a01b03841693308514612b565781906001600160a01b036003541680612ac3575b506040517f23b872dd000000000000000000000000000000000000000000000000000000006020808301919091526001600160a01b03928316602483015230604483015260648083018890528252917f000000000000000000000000000000000000000000000000000000000000000016906129d2906129cc6084826121dc565b82612cf8565b6024604051809481937f42966c680000000000000000000000000000000000000000000000000000000083528960048401525af19081156105ad578291612aa4575b5015612a7c575015612a4e576040519081527fabf8a0bc0c6341b64dfa026a551cda9d3beb0e0525758303026bacbc11ad1d8c60203392a3565b6040519081527f76bf0a63dad9216ecf94a5b1fdefa2c44ed7c809fe0028c4c2ce064b151c903f60203392a3565b807f6f16aafc0000000000000000000000000000000000000000000000000000000060049252fd5b612abd915060203d6020116127c1576127b381836121dc565b38612a14565b809192503b1561161b576040517f3a6070230000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024810186905233604482015284151560648201529083908290608490829084905af18015612b4b57612b38575b90829161294b565b91612b45816020946121dc565b91612b30565b6040513d85823e3d90fd5b6004827f20994242000000000000000000000000000000000000000000000000000000008152fd5b807f0dcbb5220000000000000000000000000000000000000000000000000000000060049252fd5b507f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a848815280602052604081206001600160a01b033316825260205260ff60408220541615612921565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690813b1561205c576001600160a01b0390604051907f8d1fdf2f0000000000000000000000000000000000000000000000000000000082528160248160008096819516978860048401525af180156105ad578290612c9b575b50507f2a3de20682fb291f444b5c1469d7e0950c558ce3dadf97163687873e29bcf4ae339180a3565b612ca4916121dc565b3881612c72565b9065ffffffffffff8091169116019065ffffffffffff8211612cc957565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b906000602091828151910182855af115612d69576000513d612d6057506001600160a01b0381163b155b612d295750565b6001600160a01b03907f5274afe7000000000000000000000000000000000000000000000000000000006000521660045260246000fd5b60011415612d22565b6040513d6000823e3d90fd5b6002546001600160a01b038116612e05577fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03821690811760025590612dc490600061304c565b9081612dce575090565b600080526004602052612e01907f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec613104565b5090565b7f3fc3c27a0000000000000000000000000000000000000000000000000000000060005260046000fd5b8015612e6c575b612e40828261304c565b9182612e4b57505090565b612e019160005260046020526001600160a01b036040600020911690613104565b6002546001600160a01b038116612e05577fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03831617600255612e36565b612edb6001600160a01b0360025416916001600160a01b038116928314612f18575b600061317a565b9081612ee5575090565b600080526004602052612e01907f17ef568e3e12ab5b9c7254a8d58478811de00f9e6eb34345acd53bf8fd09d3ec613227565b7fffffffffffffffffffffffff000000000000000000000000000000000000000060025416600255612ed4565b801580612fb5575b612f88575b612f5c828261317a565b9182612f6757505090565b612e019160005260046020526001600160a01b036040600020911690613227565b7fffffffffffffffffffffffff000000000000000000000000000000000000000060025416600255612f52565b506001600160a01b03600254166001600160a01b03831614612f4d565b9081602091031261205c5751801515810361205c5790565b65ffffffffffff81116130025765ffffffffffff1690565b7f6dfcc65000000000000000000000000000000000000000000000000000000000600052603060045260245260446000fd5b80548210156121995760005260206000200190600090565b80600052600060205260406000206001600160a01b03831660005260205260ff60406000205416156000146130fd5780600052600060205260406000206001600160a01b038316600052602052604060002060017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008254161790556001600160a01b03339216907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d600080a4600190565b5050600090565b6001810190826000528160205260406000205415600014613172578054680100000000000000008110156121ff5761315d613146826001879401855584613034565b81939154906000199060031b92831b921b19161790565b90555491600052602052604060002055600190565b505050600090565b80600052600060205260406000206001600160a01b03831660005260205260ff604060002054166000146130fd5780600052600060205260406000206001600160a01b03831660005260205260406000207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541690556001600160a01b03339216907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b600080a4600190565b9060018201918160005282602052604060002054801515600014613309576000198101818111612cc9578254906000198201918211612cc9578181036132d2575b505050805480156132a35760001901906132828282613034565b60001982549160031b1b191690555560005260205260006040812055600190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6132f26132e26131469386613034565b90549060031b1c92839286613034565b905560005283602052604060002055388080613268565b5050505060009056fea164736f6c634300081a000aad3228b676f7d3cd4284a5443f17f1962b36e491b30a40b2405849e597ba5fb5",
}

var TokenGovernorABI = TokenGovernorMetaData.ABI

var TokenGovernorBin = TokenGovernorMetaData.Bin

func DeployTokenGovernor(auth *bind.TransactOpts, backend bind.ContractBackend, token common.Address, initialDelay *big.Int, initialDefaultAdmin common.Address) (common.Address, *types.Transaction, *TokenGovernor, error) {
	parsed, err := TokenGovernorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenGovernorBin), backend, token, initialDelay, initialDefaultAdmin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenGovernor{address: address, abi: *parsed, TokenGovernorCaller: TokenGovernorCaller{contract: contract}, TokenGovernorTransactor: TokenGovernorTransactor{contract: contract}, TokenGovernorFilterer: TokenGovernorFilterer{contract: contract}}, nil
}

type TokenGovernor struct {
	address common.Address
	abi     abi.ABI
	TokenGovernorCaller
	TokenGovernorTransactor
	TokenGovernorFilterer
}

type TokenGovernorCaller struct {
	contract *bind.BoundContract
}

type TokenGovernorTransactor struct {
	contract *bind.BoundContract
}

type TokenGovernorFilterer struct {
	contract *bind.BoundContract
}

type TokenGovernorSession struct {
	Contract     *TokenGovernor
	CallOpts     bind.CallOpts
	TransactOpts bind.TransactOpts
}

type TokenGovernorCallerSession struct {
	Contract *TokenGovernorCaller
	CallOpts bind.CallOpts
}

type TokenGovernorTransactorSession struct {
	Contract     *TokenGovernorTransactor
	TransactOpts bind.TransactOpts
}

type TokenGovernorRaw struct {
	Contract *TokenGovernor
}

type TokenGovernorCallerRaw struct {
	Contract *TokenGovernorCaller
}

type TokenGovernorTransactorRaw struct {
	Contract *TokenGovernorTransactor
}

func NewTokenGovernor(address common.Address, backend bind.ContractBackend) (*TokenGovernor, error) {
	abi, err := abi.JSON(strings.NewReader(TokenGovernorABI))
	if err != nil {
		return nil, err
	}
	contract, err := bindTokenGovernor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenGovernor{address: address, abi: abi, TokenGovernorCaller: TokenGovernorCaller{contract: contract}, TokenGovernorTransactor: TokenGovernorTransactor{contract: contract}, TokenGovernorFilterer: TokenGovernorFilterer{contract: contract}}, nil
}

func NewTokenGovernorCaller(address common.Address, caller bind.ContractCaller) (*TokenGovernorCaller, error) {
	contract, err := bindTokenGovernor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorCaller{contract: contract}, nil
}

func NewTokenGovernorTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenGovernorTransactor, error) {
	contract, err := bindTokenGovernor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorTransactor{contract: contract}, nil
}

func NewTokenGovernorFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenGovernorFilterer, error) {
	contract, err := bindTokenGovernor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorFilterer{contract: contract}, nil
}

func bindTokenGovernor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenGovernorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

func (_TokenGovernor *TokenGovernorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenGovernor.Contract.TokenGovernorCaller.contract.Call(opts, result, method, params...)
}

func (_TokenGovernor *TokenGovernorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenGovernor.Contract.TokenGovernorTransactor.contract.Transfer(opts)
}

func (_TokenGovernor *TokenGovernorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenGovernor.Contract.TokenGovernorTransactor.contract.Transact(opts, method, params...)
}

func (_TokenGovernor *TokenGovernorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenGovernor.Contract.contract.Call(opts, result, method, params...)
}

func (_TokenGovernor *TokenGovernorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenGovernor.Contract.contract.Transfer(opts)
}

func (_TokenGovernor *TokenGovernorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenGovernor.Contract.contract.Transact(opts, method, params...)
}

func (_TokenGovernor *TokenGovernorCaller) BRIDGEMINTERORBURNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "BRIDGE_MINTER_OR_BURNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) BRIDGEMINTERORBURNERROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.BRIDGEMINTERORBURNERROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) BRIDGEMINTERORBURNERROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.BRIDGEMINTERORBURNERROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) BURNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "BURNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) BURNERROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.BURNERROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) BURNERROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.BURNERROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) CHECKERADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "CHECKER_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) CHECKERADMINROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.CHECKERADMINROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) CHECKERADMINROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.CHECKERADMINROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.DEFAULTADMINROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.DEFAULTADMINROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) FREEZERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "FREEZER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) FREEZERROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.FREEZERROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) FREEZERROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.FREEZERROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) MINTERROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.MINTERROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) MINTERROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.MINTERROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) PAUSERROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.PAUSERROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) PAUSERROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.PAUSERROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) RECOVERYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "RECOVERY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) RECOVERYROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.RECOVERYROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) RECOVERYROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.RECOVERYROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) UNFREEZERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "UNFREEZER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) UNFREEZERROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.UNFREEZERROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) UNFREEZERROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.UNFREEZERROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) UNPAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "UNPAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) UNPAUSERROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.UNPAUSERROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) UNPAUSERROLE() ([32]byte, error) {
	return _TokenGovernor.Contract.UNPAUSERROLE(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) DefaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "defaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) DefaultAdmin() (common.Address, error) {
	return _TokenGovernor.Contract.DefaultAdmin(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) DefaultAdmin() (common.Address, error) {
	return _TokenGovernor.Contract.DefaultAdmin(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "defaultAdminDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) DefaultAdminDelay() (*big.Int, error) {
	return _TokenGovernor.Contract.DefaultAdminDelay(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) DefaultAdminDelay() (*big.Int, error) {
	return _TokenGovernor.Contract.DefaultAdminDelay(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "defaultAdminDelayIncreaseWait")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _TokenGovernor.Contract.DefaultAdminDelayIncreaseWait(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) DefaultAdminDelayIncreaseWait() (*big.Int, error) {
	return _TokenGovernor.Contract.DefaultAdminDelayIncreaseWait(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) GetAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "getAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) GetAdmins() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetAdmins(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) GetAdmins() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetAdmins(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) GetBridgeMintersOrBurners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "getBridgeMintersOrBurners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) GetBridgeMintersOrBurners() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetBridgeMintersOrBurners(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) GetBridgeMintersOrBurners() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetBridgeMintersOrBurners(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) GetBurners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "getBurners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) GetBurners() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetBurners(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) GetBurners() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetBurners(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) GetChecker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "getChecker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) GetChecker() (common.Address, error) {
	return _TokenGovernor.Contract.GetChecker(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) GetChecker() (common.Address, error) {
	return _TokenGovernor.Contract.GetChecker(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) GetCheckerAdmins(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "getCheckerAdmins")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) GetCheckerAdmins() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetCheckerAdmins(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) GetCheckerAdmins() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetCheckerAdmins(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) GetFreezers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "getFreezers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) GetFreezers() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetFreezers(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) GetFreezers() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetFreezers(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) GetMinters(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "getMinters")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) GetMinters() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetMinters(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) GetMinters() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetMinters(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) GetPausers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "getPausers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) GetPausers() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetPausers(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) GetPausers() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetPausers(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) GetRecoveryManagers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "getRecoveryManagers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) GetRecoveryManagers() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetRecoveryManagers(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) GetRecoveryManagers() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetRecoveryManagers(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TokenGovernor.Contract.GetRoleAdmin(&_TokenGovernor.CallOpts, role)
}

func (_TokenGovernor *TokenGovernorCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TokenGovernor.Contract.GetRoleAdmin(&_TokenGovernor.CallOpts, role)
}

func (_TokenGovernor *TokenGovernorCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TokenGovernor.Contract.GetRoleMember(&_TokenGovernor.CallOpts, role, index)
}

func (_TokenGovernor *TokenGovernorCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TokenGovernor.Contract.GetRoleMember(&_TokenGovernor.CallOpts, role, index)
}

func (_TokenGovernor *TokenGovernorCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TokenGovernor.Contract.GetRoleMemberCount(&_TokenGovernor.CallOpts, role)
}

func (_TokenGovernor *TokenGovernorCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TokenGovernor.Contract.GetRoleMemberCount(&_TokenGovernor.CallOpts, role)
}

func (_TokenGovernor *TokenGovernorCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _TokenGovernor.Contract.GetRoleMembers(&_TokenGovernor.CallOpts, role)
}

func (_TokenGovernor *TokenGovernorCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _TokenGovernor.Contract.GetRoleMembers(&_TokenGovernor.CallOpts, role)
}

func (_TokenGovernor *TokenGovernorCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) GetToken() (common.Address, error) {
	return _TokenGovernor.Contract.GetToken(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) GetToken() (common.Address, error) {
	return _TokenGovernor.Contract.GetToken(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) GetTokenBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "getTokenBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) GetTokenBalance() (*big.Int, error) {
	return _TokenGovernor.Contract.GetTokenBalance(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) GetTokenBalance() (*big.Int, error) {
	return _TokenGovernor.Contract.GetTokenBalance(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) GetUnfreezers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "getUnfreezers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) GetUnfreezers() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetUnfreezers(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) GetUnfreezers() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetUnfreezers(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) GetUnpausers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "getUnpausers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) GetUnpausers() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetUnpausers(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) GetUnpausers() ([]common.Address, error) {
	return _TokenGovernor.Contract.GetUnpausers(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TokenGovernor.Contract.HasRole(&_TokenGovernor.CallOpts, role, account)
}

func (_TokenGovernor *TokenGovernorCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TokenGovernor.Contract.HasRole(&_TokenGovernor.CallOpts, role, account)
}

func (_TokenGovernor *TokenGovernorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) Owner() (common.Address, error) {
	return _TokenGovernor.Contract.Owner(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) Owner() (common.Address, error) {
	return _TokenGovernor.Contract.Owner(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

	error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "pendingDefaultAdmin")

	outstruct := new(PendingDefaultAdmin)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewAdmin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_TokenGovernor *TokenGovernorSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _TokenGovernor.Contract.PendingDefaultAdmin(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) PendingDefaultAdmin() (PendingDefaultAdmin,

	error) {
	return _TokenGovernor.Contract.PendingDefaultAdmin(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

	error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "pendingDefaultAdminDelay")

	outstruct := new(PendingDefaultAdminDelay)
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewDelay = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Schedule = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

func (_TokenGovernor *TokenGovernorSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _TokenGovernor.Contract.PendingDefaultAdminDelay(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCallerSession) PendingDefaultAdminDelay() (PendingDefaultAdminDelay,

	error) {
	return _TokenGovernor.Contract.PendingDefaultAdminDelay(&_TokenGovernor.CallOpts)
}

func (_TokenGovernor *TokenGovernorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TokenGovernor.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

func (_TokenGovernor *TokenGovernorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenGovernor.Contract.SupportsInterface(&_TokenGovernor.CallOpts, interfaceId)
}

func (_TokenGovernor *TokenGovernorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TokenGovernor.Contract.SupportsInterface(&_TokenGovernor.CallOpts, interfaceId)
}

func (_TokenGovernor *TokenGovernorTransactor) AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "acceptDefaultAdminTransfer")
}

func (_TokenGovernor *TokenGovernorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _TokenGovernor.Contract.AcceptDefaultAdminTransfer(&_TokenGovernor.TransactOpts)
}

func (_TokenGovernor *TokenGovernorTransactorSession) AcceptDefaultAdminTransfer() (*types.Transaction, error) {
	return _TokenGovernor.Contract.AcceptDefaultAdminTransfer(&_TokenGovernor.TransactOpts)
}

func (_TokenGovernor *TokenGovernorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "acceptOwnership")
}

func (_TokenGovernor *TokenGovernorSession) AcceptOwnership() (*types.Transaction, error) {
	return _TokenGovernor.Contract.AcceptOwnership(&_TokenGovernor.TransactOpts)
}

func (_TokenGovernor *TokenGovernorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _TokenGovernor.Contract.AcceptOwnership(&_TokenGovernor.TransactOpts)
}

func (_TokenGovernor *TokenGovernorTransactor) BatchDrainFrozenAccounts(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "batchDrainFrozenAccounts", accounts)
}

func (_TokenGovernor *TokenGovernorSession) BatchDrainFrozenAccounts(accounts []common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.BatchDrainFrozenAccounts(&_TokenGovernor.TransactOpts, accounts)
}

func (_TokenGovernor *TokenGovernorTransactorSession) BatchDrainFrozenAccounts(accounts []common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.BatchDrainFrozenAccounts(&_TokenGovernor.TransactOpts, accounts)
}

func (_TokenGovernor *TokenGovernorTransactor) BatchFreeze(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "batchFreeze", accounts)
}

func (_TokenGovernor *TokenGovernorSession) BatchFreeze(accounts []common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.BatchFreeze(&_TokenGovernor.TransactOpts, accounts)
}

func (_TokenGovernor *TokenGovernorTransactorSession) BatchFreeze(accounts []common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.BatchFreeze(&_TokenGovernor.TransactOpts, accounts)
}

func (_TokenGovernor *TokenGovernorTransactor) BatchUnfreeze(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "batchUnfreeze", accounts)
}

func (_TokenGovernor *TokenGovernorSession) BatchUnfreeze(accounts []common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.BatchUnfreeze(&_TokenGovernor.TransactOpts, accounts)
}

func (_TokenGovernor *TokenGovernorTransactorSession) BatchUnfreeze(accounts []common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.BatchUnfreeze(&_TokenGovernor.TransactOpts, accounts)
}

func (_TokenGovernor *TokenGovernorTransactor) BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "beginDefaultAdminTransfer", newAdmin)
}

func (_TokenGovernor *TokenGovernorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.BeginDefaultAdminTransfer(&_TokenGovernor.TransactOpts, newAdmin)
}

func (_TokenGovernor *TokenGovernorTransactorSession) BeginDefaultAdminTransfer(newAdmin common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.BeginDefaultAdminTransfer(&_TokenGovernor.TransactOpts, newAdmin)
}

func (_TokenGovernor *TokenGovernorTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "burn", amount)
}

func (_TokenGovernor *TokenGovernorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.Contract.Burn(&_TokenGovernor.TransactOpts, amount)
}

func (_TokenGovernor *TokenGovernorTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.Contract.Burn(&_TokenGovernor.TransactOpts, amount)
}

func (_TokenGovernor *TokenGovernorTransactor) BurnFrom(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "burnFrom", from, amount)
}

func (_TokenGovernor *TokenGovernorSession) BurnFrom(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.Contract.BurnFrom(&_TokenGovernor.TransactOpts, from, amount)
}

func (_TokenGovernor *TokenGovernorTransactorSession) BurnFrom(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.Contract.BurnFrom(&_TokenGovernor.TransactOpts, from, amount)
}

func (_TokenGovernor *TokenGovernorTransactor) CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "cancelDefaultAdminTransfer")
}

func (_TokenGovernor *TokenGovernorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _TokenGovernor.Contract.CancelDefaultAdminTransfer(&_TokenGovernor.TransactOpts)
}

func (_TokenGovernor *TokenGovernorTransactorSession) CancelDefaultAdminTransfer() (*types.Transaction, error) {
	return _TokenGovernor.Contract.CancelDefaultAdminTransfer(&_TokenGovernor.TransactOpts)
}

func (_TokenGovernor *TokenGovernorTransactor) ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "changeDefaultAdminDelay", newDelay)
}

func (_TokenGovernor *TokenGovernorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.Contract.ChangeDefaultAdminDelay(&_TokenGovernor.TransactOpts, newDelay)
}

func (_TokenGovernor *TokenGovernorTransactorSession) ChangeDefaultAdminDelay(newDelay *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.Contract.ChangeDefaultAdminDelay(&_TokenGovernor.TransactOpts, newDelay)
}

func (_TokenGovernor *TokenGovernorTransactor) DrainFrozenAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "drainFrozenAccount", account)
}

func (_TokenGovernor *TokenGovernorSession) DrainFrozenAccount(account common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.DrainFrozenAccount(&_TokenGovernor.TransactOpts, account)
}

func (_TokenGovernor *TokenGovernorTransactorSession) DrainFrozenAccount(account common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.DrainFrozenAccount(&_TokenGovernor.TransactOpts, account)
}

func (_TokenGovernor *TokenGovernorTransactor) ExecuteTokenFunction(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "executeTokenFunction", data)
}

func (_TokenGovernor *TokenGovernorSession) ExecuteTokenFunction(data []byte) (*types.Transaction, error) {
	return _TokenGovernor.Contract.ExecuteTokenFunction(&_TokenGovernor.TransactOpts, data)
}

func (_TokenGovernor *TokenGovernorTransactorSession) ExecuteTokenFunction(data []byte) (*types.Transaction, error) {
	return _TokenGovernor.Contract.ExecuteTokenFunction(&_TokenGovernor.TransactOpts, data)
}

func (_TokenGovernor *TokenGovernorTransactor) Freeze(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "freeze", account)
}

func (_TokenGovernor *TokenGovernorSession) Freeze(account common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.Freeze(&_TokenGovernor.TransactOpts, account)
}

func (_TokenGovernor *TokenGovernorTransactorSession) Freeze(account common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.Freeze(&_TokenGovernor.TransactOpts, account)
}

func (_TokenGovernor *TokenGovernorTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "grantRole", role, account)
}

func (_TokenGovernor *TokenGovernorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.GrantRole(&_TokenGovernor.TransactOpts, role, account)
}

func (_TokenGovernor *TokenGovernorTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.GrantRole(&_TokenGovernor.TransactOpts, role, account)
}

func (_TokenGovernor *TokenGovernorTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "mint", recipient, amount)
}

func (_TokenGovernor *TokenGovernorSession) Mint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.Contract.Mint(&_TokenGovernor.TransactOpts, recipient, amount)
}

func (_TokenGovernor *TokenGovernorTransactorSession) Mint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.Contract.Mint(&_TokenGovernor.TransactOpts, recipient, amount)
}

func (_TokenGovernor *TokenGovernorTransactor) Mint0(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "mint0", amount)
}

func (_TokenGovernor *TokenGovernorSession) Mint0(amount *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.Contract.Mint0(&_TokenGovernor.TransactOpts, amount)
}

func (_TokenGovernor *TokenGovernorTransactorSession) Mint0(amount *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.Contract.Mint0(&_TokenGovernor.TransactOpts, amount)
}

func (_TokenGovernor *TokenGovernorTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "pause")
}

func (_TokenGovernor *TokenGovernorSession) Pause() (*types.Transaction, error) {
	return _TokenGovernor.Contract.Pause(&_TokenGovernor.TransactOpts)
}

func (_TokenGovernor *TokenGovernorTransactorSession) Pause() (*types.Transaction, error) {
	return _TokenGovernor.Contract.Pause(&_TokenGovernor.TransactOpts)
}

func (_TokenGovernor *TokenGovernorTransactor) RecoverERC20(opts *bind.TransactOpts, token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "recoverERC20", token, recipient, amount)
}

func (_TokenGovernor *TokenGovernorSession) RecoverERC20(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.Contract.RecoverERC20(&_TokenGovernor.TransactOpts, token, recipient, amount)
}

func (_TokenGovernor *TokenGovernorTransactorSession) RecoverERC20(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.Contract.RecoverERC20(&_TokenGovernor.TransactOpts, token, recipient, amount)
}

func (_TokenGovernor *TokenGovernorTransactor) RecoverGovernedTokenERC20(opts *bind.TransactOpts, token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "recoverGovernedTokenERC20", token, recipient, amount)
}

func (_TokenGovernor *TokenGovernorSession) RecoverGovernedTokenERC20(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.Contract.RecoverGovernedTokenERC20(&_TokenGovernor.TransactOpts, token, recipient, amount)
}

func (_TokenGovernor *TokenGovernorTransactorSession) RecoverGovernedTokenERC20(token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenGovernor.Contract.RecoverGovernedTokenERC20(&_TokenGovernor.TransactOpts, token, recipient, amount)
}

func (_TokenGovernor *TokenGovernorTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "renounceRole", role, account)
}

func (_TokenGovernor *TokenGovernorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.RenounceRole(&_TokenGovernor.TransactOpts, role, account)
}

func (_TokenGovernor *TokenGovernorTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.RenounceRole(&_TokenGovernor.TransactOpts, role, account)
}

func (_TokenGovernor *TokenGovernorTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "revokeRole", role, account)
}

func (_TokenGovernor *TokenGovernorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.RevokeRole(&_TokenGovernor.TransactOpts, role, account)
}

func (_TokenGovernor *TokenGovernorTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.RevokeRole(&_TokenGovernor.TransactOpts, role, account)
}

func (_TokenGovernor *TokenGovernorTransactor) RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "rollbackDefaultAdminDelay")
}

func (_TokenGovernor *TokenGovernorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _TokenGovernor.Contract.RollbackDefaultAdminDelay(&_TokenGovernor.TransactOpts)
}

func (_TokenGovernor *TokenGovernorTransactorSession) RollbackDefaultAdminDelay() (*types.Transaction, error) {
	return _TokenGovernor.Contract.RollbackDefaultAdminDelay(&_TokenGovernor.TransactOpts)
}

func (_TokenGovernor *TokenGovernorTransactor) SetChecker(opts *bind.TransactOpts, newChecker common.Address) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "setChecker", newChecker)
}

func (_TokenGovernor *TokenGovernorSession) SetChecker(newChecker common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.SetChecker(&_TokenGovernor.TransactOpts, newChecker)
}

func (_TokenGovernor *TokenGovernorTransactorSession) SetChecker(newChecker common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.SetChecker(&_TokenGovernor.TransactOpts, newChecker)
}

func (_TokenGovernor *TokenGovernorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "transferOwnership", newOwner)
}

func (_TokenGovernor *TokenGovernorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.TransferOwnership(&_TokenGovernor.TransactOpts, newOwner)
}

func (_TokenGovernor *TokenGovernorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.TransferOwnership(&_TokenGovernor.TransactOpts, newOwner)
}

func (_TokenGovernor *TokenGovernorTransactor) Unfreeze(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "unfreeze", account)
}

func (_TokenGovernor *TokenGovernorSession) Unfreeze(account common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.Unfreeze(&_TokenGovernor.TransactOpts, account)
}

func (_TokenGovernor *TokenGovernorTransactorSession) Unfreeze(account common.Address) (*types.Transaction, error) {
	return _TokenGovernor.Contract.Unfreeze(&_TokenGovernor.TransactOpts, account)
}

func (_TokenGovernor *TokenGovernorTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenGovernor.contract.Transact(opts, "unpause")
}

func (_TokenGovernor *TokenGovernorSession) Unpause() (*types.Transaction, error) {
	return _TokenGovernor.Contract.Unpause(&_TokenGovernor.TransactOpts)
}

func (_TokenGovernor *TokenGovernorTransactorSession) Unpause() (*types.Transaction, error) {
	return _TokenGovernor.Contract.Unpause(&_TokenGovernor.TransactOpts)
}

type TokenGovernorAccountFrozenIterator struct {
	Event *TokenGovernorAccountFrozen

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorAccountFrozenIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorAccountFrozen)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorAccountFrozen)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorAccountFrozenIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorAccountFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorAccountFrozen struct {
	Caller  common.Address
	Account common.Address
	Raw     types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterAccountFrozen(opts *bind.FilterOpts, caller []common.Address, account []common.Address) (*TokenGovernorAccountFrozenIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "AccountFrozen", callerRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorAccountFrozenIterator{contract: _TokenGovernor.contract, event: "AccountFrozen", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchAccountFrozen(opts *bind.WatchOpts, sink chan<- *TokenGovernorAccountFrozen, caller []common.Address, account []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "AccountFrozen", callerRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorAccountFrozen)
				if err := _TokenGovernor.contract.UnpackLog(event, "AccountFrozen", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseAccountFrozen(log types.Log) (*TokenGovernorAccountFrozen, error) {
	event := new(TokenGovernorAccountFrozen)
	if err := _TokenGovernor.contract.UnpackLog(event, "AccountFrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorAccountUnfrozenIterator struct {
	Event *TokenGovernorAccountUnfrozen

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorAccountUnfrozenIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorAccountUnfrozen)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorAccountUnfrozen)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorAccountUnfrozenIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorAccountUnfrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorAccountUnfrozen struct {
	Caller  common.Address
	Account common.Address
	Raw     types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterAccountUnfrozen(opts *bind.FilterOpts, caller []common.Address, account []common.Address) (*TokenGovernorAccountUnfrozenIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "AccountUnfrozen", callerRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorAccountUnfrozenIterator{contract: _TokenGovernor.contract, event: "AccountUnfrozen", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchAccountUnfrozen(opts *bind.WatchOpts, sink chan<- *TokenGovernorAccountUnfrozen, caller []common.Address, account []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "AccountUnfrozen", callerRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorAccountUnfrozen)
				if err := _TokenGovernor.contract.UnpackLog(event, "AccountUnfrozen", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseAccountUnfrozen(log types.Log) (*TokenGovernorAccountUnfrozen, error) {
	event := new(TokenGovernorAccountUnfrozen)
	if err := _TokenGovernor.contract.UnpackLog(event, "AccountUnfrozen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorBridgeBurnIterator struct {
	Event *TokenGovernorBridgeBurn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorBridgeBurnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorBridgeBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorBridgeBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorBridgeBurnIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorBridgeBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorBridgeBurn struct {
	Caller common.Address
	From   common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterBridgeBurn(opts *bind.FilterOpts, caller []common.Address, from []common.Address) (*TokenGovernorBridgeBurnIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "BridgeBurn", callerRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorBridgeBurnIterator{contract: _TokenGovernor.contract, event: "BridgeBurn", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchBridgeBurn(opts *bind.WatchOpts, sink chan<- *TokenGovernorBridgeBurn, caller []common.Address, from []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "BridgeBurn", callerRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorBridgeBurn)
				if err := _TokenGovernor.contract.UnpackLog(event, "BridgeBurn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseBridgeBurn(log types.Log) (*TokenGovernorBridgeBurn, error) {
	event := new(TokenGovernorBridgeBurn)
	if err := _TokenGovernor.contract.UnpackLog(event, "BridgeBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorBridgeMintIterator struct {
	Event *TokenGovernorBridgeMint

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorBridgeMintIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorBridgeMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorBridgeMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorBridgeMintIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorBridgeMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorBridgeMint struct {
	Caller    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterBridgeMint(opts *bind.FilterOpts, caller []common.Address, recipient []common.Address) (*TokenGovernorBridgeMintIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "BridgeMint", callerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorBridgeMintIterator{contract: _TokenGovernor.contract, event: "BridgeMint", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchBridgeMint(opts *bind.WatchOpts, sink chan<- *TokenGovernorBridgeMint, caller []common.Address, recipient []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "BridgeMint", callerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorBridgeMint)
				if err := _TokenGovernor.contract.UnpackLog(event, "BridgeMint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseBridgeMint(log types.Log) (*TokenGovernorBridgeMint, error) {
	event := new(TokenGovernorBridgeMint)
	if err := _TokenGovernor.contract.UnpackLog(event, "BridgeMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorCheckerUpdatedIterator struct {
	Event *TokenGovernorCheckerUpdated

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorCheckerUpdatedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorCheckerUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorCheckerUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorCheckerUpdatedIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorCheckerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorCheckerUpdated struct {
	PreviousChecker common.Address
	NewChecker      common.Address
	Raw             types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterCheckerUpdated(opts *bind.FilterOpts, previousChecker []common.Address, newChecker []common.Address) (*TokenGovernorCheckerUpdatedIterator, error) {

	var previousCheckerRule []interface{}
	for _, previousCheckerItem := range previousChecker {
		previousCheckerRule = append(previousCheckerRule, previousCheckerItem)
	}
	var newCheckerRule []interface{}
	for _, newCheckerItem := range newChecker {
		newCheckerRule = append(newCheckerRule, newCheckerItem)
	}

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "CheckerUpdated", previousCheckerRule, newCheckerRule)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorCheckerUpdatedIterator{contract: _TokenGovernor.contract, event: "CheckerUpdated", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchCheckerUpdated(opts *bind.WatchOpts, sink chan<- *TokenGovernorCheckerUpdated, previousChecker []common.Address, newChecker []common.Address) (event.Subscription, error) {

	var previousCheckerRule []interface{}
	for _, previousCheckerItem := range previousChecker {
		previousCheckerRule = append(previousCheckerRule, previousCheckerItem)
	}
	var newCheckerRule []interface{}
	for _, newCheckerItem := range newChecker {
		newCheckerRule = append(newCheckerRule, newCheckerItem)
	}

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "CheckerUpdated", previousCheckerRule, newCheckerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorCheckerUpdated)
				if err := _TokenGovernor.contract.UnpackLog(event, "CheckerUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseCheckerUpdated(log types.Log) (*TokenGovernorCheckerUpdated, error) {
	event := new(TokenGovernorCheckerUpdated)
	if err := _TokenGovernor.contract.UnpackLog(event, "CheckerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorContractPausedIterator struct {
	Event *TokenGovernorContractPaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorContractPausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorContractPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorContractPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorContractPausedIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorContractPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorContractPaused struct {
	Caller common.Address
	Raw    types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterContractPaused(opts *bind.FilterOpts, caller []common.Address) (*TokenGovernorContractPausedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "ContractPaused", callerRule)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorContractPausedIterator{contract: _TokenGovernor.contract, event: "ContractPaused", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchContractPaused(opts *bind.WatchOpts, sink chan<- *TokenGovernorContractPaused, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "ContractPaused", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorContractPaused)
				if err := _TokenGovernor.contract.UnpackLog(event, "ContractPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseContractPaused(log types.Log) (*TokenGovernorContractPaused, error) {
	event := new(TokenGovernorContractPaused)
	if err := _TokenGovernor.contract.UnpackLog(event, "ContractPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorContractUnpausedIterator struct {
	Event *TokenGovernorContractUnpaused

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorContractUnpausedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorContractUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorContractUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorContractUnpausedIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorContractUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorContractUnpaused struct {
	Caller common.Address
	Raw    types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterContractUnpaused(opts *bind.FilterOpts, caller []common.Address) (*TokenGovernorContractUnpausedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "ContractUnpaused", callerRule)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorContractUnpausedIterator{contract: _TokenGovernor.contract, event: "ContractUnpaused", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchContractUnpaused(opts *bind.WatchOpts, sink chan<- *TokenGovernorContractUnpaused, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "ContractUnpaused", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorContractUnpaused)
				if err := _TokenGovernor.contract.UnpackLog(event, "ContractUnpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseContractUnpaused(log types.Log) (*TokenGovernorContractUnpaused, error) {
	event := new(TokenGovernorContractUnpaused)
	if err := _TokenGovernor.contract.UnpackLog(event, "ContractUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorDefaultAdminDelayChangeCanceledIterator struct {
	Event *TokenGovernorDefaultAdminDelayChangeCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorDefaultAdminDelayChangeCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorDefaultAdminDelayChangeCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorDefaultAdminDelayChangeCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorDefaultAdminDelayChangeCanceledIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorDefaultAdminDelayChangeCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorDefaultAdminDelayChangeCanceled struct {
	Raw types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*TokenGovernorDefaultAdminDelayChangeCanceledIterator, error) {

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return &TokenGovernorDefaultAdminDelayChangeCanceledIterator{contract: _TokenGovernor.contract, event: "DefaultAdminDelayChangeCanceled", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *TokenGovernorDefaultAdminDelayChangeCanceled) (event.Subscription, error) {

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "DefaultAdminDelayChangeCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorDefaultAdminDelayChangeCanceled)
				if err := _TokenGovernor.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseDefaultAdminDelayChangeCanceled(log types.Log) (*TokenGovernorDefaultAdminDelayChangeCanceled, error) {
	event := new(TokenGovernorDefaultAdminDelayChangeCanceled)
	if err := _TokenGovernor.contract.UnpackLog(event, "DefaultAdminDelayChangeCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorDefaultAdminDelayChangeScheduledIterator struct {
	Event *TokenGovernorDefaultAdminDelayChangeScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorDefaultAdminDelayChangeScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorDefaultAdminDelayChangeScheduled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorDefaultAdminDelayChangeScheduled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorDefaultAdminDelayChangeScheduledIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorDefaultAdminDelayChangeScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorDefaultAdminDelayChangeScheduled struct {
	NewDelay       *big.Int
	EffectSchedule *big.Int
	Raw            types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*TokenGovernorDefaultAdminDelayChangeScheduledIterator, error) {

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return &TokenGovernorDefaultAdminDelayChangeScheduledIterator{contract: _TokenGovernor.contract, event: "DefaultAdminDelayChangeScheduled", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *TokenGovernorDefaultAdminDelayChangeScheduled) (event.Subscription, error) {

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "DefaultAdminDelayChangeScheduled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorDefaultAdminDelayChangeScheduled)
				if err := _TokenGovernor.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseDefaultAdminDelayChangeScheduled(log types.Log) (*TokenGovernorDefaultAdminDelayChangeScheduled, error) {
	event := new(TokenGovernorDefaultAdminDelayChangeScheduled)
	if err := _TokenGovernor.contract.UnpackLog(event, "DefaultAdminDelayChangeScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorDefaultAdminTransferCanceledIterator struct {
	Event *TokenGovernorDefaultAdminTransferCanceled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorDefaultAdminTransferCanceledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorDefaultAdminTransferCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorDefaultAdminTransferCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorDefaultAdminTransferCanceledIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorDefaultAdminTransferCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorDefaultAdminTransferCanceled struct {
	Raw types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*TokenGovernorDefaultAdminTransferCanceledIterator, error) {

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return &TokenGovernorDefaultAdminTransferCanceledIterator{contract: _TokenGovernor.contract, event: "DefaultAdminTransferCanceled", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *TokenGovernorDefaultAdminTransferCanceled) (event.Subscription, error) {

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "DefaultAdminTransferCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorDefaultAdminTransferCanceled)
				if err := _TokenGovernor.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseDefaultAdminTransferCanceled(log types.Log) (*TokenGovernorDefaultAdminTransferCanceled, error) {
	event := new(TokenGovernorDefaultAdminTransferCanceled)
	if err := _TokenGovernor.contract.UnpackLog(event, "DefaultAdminTransferCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorDefaultAdminTransferScheduledIterator struct {
	Event *TokenGovernorDefaultAdminTransferScheduled

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorDefaultAdminTransferScheduledIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorDefaultAdminTransferScheduled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorDefaultAdminTransferScheduled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorDefaultAdminTransferScheduledIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorDefaultAdminTransferScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorDefaultAdminTransferScheduled struct {
	NewAdmin       common.Address
	AcceptSchedule *big.Int
	Raw            types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*TokenGovernorDefaultAdminTransferScheduledIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorDefaultAdminTransferScheduledIterator{contract: _TokenGovernor.contract, event: "DefaultAdminTransferScheduled", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *TokenGovernorDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "DefaultAdminTransferScheduled", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorDefaultAdminTransferScheduled)
				if err := _TokenGovernor.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseDefaultAdminTransferScheduled(log types.Log) (*TokenGovernorDefaultAdminTransferScheduled, error) {
	event := new(TokenGovernorDefaultAdminTransferScheduled)
	if err := _TokenGovernor.contract.UnpackLog(event, "DefaultAdminTransferScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorFrozenAccountDrainedIterator struct {
	Event *TokenGovernorFrozenAccountDrained

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorFrozenAccountDrainedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorFrozenAccountDrained)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorFrozenAccountDrained)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorFrozenAccountDrainedIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorFrozenAccountDrainedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorFrozenAccountDrained struct {
	Caller  common.Address
	Account common.Address
	Raw     types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterFrozenAccountDrained(opts *bind.FilterOpts, caller []common.Address, account []common.Address) (*TokenGovernorFrozenAccountDrainedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "FrozenAccountDrained", callerRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorFrozenAccountDrainedIterator{contract: _TokenGovernor.contract, event: "FrozenAccountDrained", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchFrozenAccountDrained(opts *bind.WatchOpts, sink chan<- *TokenGovernorFrozenAccountDrained, caller []common.Address, account []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "FrozenAccountDrained", callerRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorFrozenAccountDrained)
				if err := _TokenGovernor.contract.UnpackLog(event, "FrozenAccountDrained", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseFrozenAccountDrained(log types.Log) (*TokenGovernorFrozenAccountDrained, error) {
	event := new(TokenGovernorFrozenAccountDrained)
	if err := _TokenGovernor.contract.UnpackLog(event, "FrozenAccountDrained", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorGovernedTokensRecoveredIterator struct {
	Event *TokenGovernorGovernedTokensRecovered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorGovernedTokensRecoveredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorGovernedTokensRecovered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorGovernedTokensRecovered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorGovernedTokensRecoveredIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorGovernedTokensRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorGovernedTokensRecovered struct {
	Caller    common.Address
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterGovernedTokensRecovered(opts *bind.FilterOpts, caller []common.Address, token []common.Address, recipient []common.Address) (*TokenGovernorGovernedTokensRecoveredIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "GovernedTokensRecovered", callerRule, tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorGovernedTokensRecoveredIterator{contract: _TokenGovernor.contract, event: "GovernedTokensRecovered", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchGovernedTokensRecovered(opts *bind.WatchOpts, sink chan<- *TokenGovernorGovernedTokensRecovered, caller []common.Address, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "GovernedTokensRecovered", callerRule, tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorGovernedTokensRecovered)
				if err := _TokenGovernor.contract.UnpackLog(event, "GovernedTokensRecovered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseGovernedTokensRecovered(log types.Log) (*TokenGovernorGovernedTokensRecovered, error) {
	event := new(TokenGovernorGovernedTokensRecovered)
	if err := _TokenGovernor.contract.UnpackLog(event, "GovernedTokensRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorNativeBurnIterator struct {
	Event *TokenGovernorNativeBurn

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorNativeBurnIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorNativeBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorNativeBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorNativeBurnIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorNativeBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorNativeBurn struct {
	Caller common.Address
	From   common.Address
	Amount *big.Int
	Raw    types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterNativeBurn(opts *bind.FilterOpts, caller []common.Address, from []common.Address) (*TokenGovernorNativeBurnIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "NativeBurn", callerRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorNativeBurnIterator{contract: _TokenGovernor.contract, event: "NativeBurn", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchNativeBurn(opts *bind.WatchOpts, sink chan<- *TokenGovernorNativeBurn, caller []common.Address, from []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "NativeBurn", callerRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorNativeBurn)
				if err := _TokenGovernor.contract.UnpackLog(event, "NativeBurn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseNativeBurn(log types.Log) (*TokenGovernorNativeBurn, error) {
	event := new(TokenGovernorNativeBurn)
	if err := _TokenGovernor.contract.UnpackLog(event, "NativeBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorNativeMintIterator struct {
	Event *TokenGovernorNativeMint

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorNativeMintIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorNativeMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorNativeMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorNativeMintIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorNativeMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorNativeMint struct {
	Caller    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterNativeMint(opts *bind.FilterOpts, caller []common.Address, recipient []common.Address) (*TokenGovernorNativeMintIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "NativeMint", callerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorNativeMintIterator{contract: _TokenGovernor.contract, event: "NativeMint", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchNativeMint(opts *bind.WatchOpts, sink chan<- *TokenGovernorNativeMint, caller []common.Address, recipient []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "NativeMint", callerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorNativeMint)
				if err := _TokenGovernor.contract.UnpackLog(event, "NativeMint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseNativeMint(log types.Log) (*TokenGovernorNativeMint, error) {
	event := new(TokenGovernorNativeMint)
	if err := _TokenGovernor.contract.UnpackLog(event, "NativeMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorOwnershipAcceptedIterator struct {
	Event *TokenGovernorOwnershipAccepted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorOwnershipAcceptedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorOwnershipAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorOwnershipAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorOwnershipAcceptedIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorOwnershipAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorOwnershipAccepted struct {
	Caller common.Address
	Raw    types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterOwnershipAccepted(opts *bind.FilterOpts, caller []common.Address) (*TokenGovernorOwnershipAcceptedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "OwnershipAccepted", callerRule)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorOwnershipAcceptedIterator{contract: _TokenGovernor.contract, event: "OwnershipAccepted", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchOwnershipAccepted(opts *bind.WatchOpts, sink chan<- *TokenGovernorOwnershipAccepted, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "OwnershipAccepted", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorOwnershipAccepted)
				if err := _TokenGovernor.contract.UnpackLog(event, "OwnershipAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseOwnershipAccepted(log types.Log) (*TokenGovernorOwnershipAccepted, error) {
	event := new(TokenGovernorOwnershipAccepted)
	if err := _TokenGovernor.contract.UnpackLog(event, "OwnershipAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorOwnershipTransferredIterator struct {
	Event *TokenGovernorOwnershipTransferred

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorOwnershipTransferredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorOwnershipTransferredIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorOwnershipTransferred struct {
	Caller   common.Address
	NewOwner common.Address
	Raw      types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, caller []common.Address, newOwner []common.Address) (*TokenGovernorOwnershipTransferredIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "OwnershipTransferred", callerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorOwnershipTransferredIterator{contract: _TokenGovernor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenGovernorOwnershipTransferred, caller []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "OwnershipTransferred", callerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorOwnershipTransferred)
				if err := _TokenGovernor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseOwnershipTransferred(log types.Log) (*TokenGovernorOwnershipTransferred, error) {
	event := new(TokenGovernorOwnershipTransferred)
	if err := _TokenGovernor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorRoleAdminChangedIterator struct {
	Event *TokenGovernorRoleAdminChanged

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorRoleAdminChangedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorRoleAdminChangedIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TokenGovernorRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorRoleAdminChangedIterator{contract: _TokenGovernor.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TokenGovernorRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorRoleAdminChanged)
				if err := _TokenGovernor.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseRoleAdminChanged(log types.Log) (*TokenGovernorRoleAdminChanged, error) {
	event := new(TokenGovernorRoleAdminChanged)
	if err := _TokenGovernor.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorRoleGrantedIterator struct {
	Event *TokenGovernorRoleGranted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorRoleGrantedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorRoleGrantedIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TokenGovernorRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorRoleGrantedIterator{contract: _TokenGovernor.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TokenGovernorRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorRoleGranted)
				if err := _TokenGovernor.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseRoleGranted(log types.Log) (*TokenGovernorRoleGranted, error) {
	event := new(TokenGovernorRoleGranted)
	if err := _TokenGovernor.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorRoleRevokedIterator struct {
	Event *TokenGovernorRoleRevoked

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorRoleRevokedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorRoleRevokedIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TokenGovernorRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorRoleRevokedIterator{contract: _TokenGovernor.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TokenGovernorRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorRoleRevoked)
				if err := _TokenGovernor.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseRoleRevoked(log types.Log) (*TokenGovernorRoleRevoked, error) {
	event := new(TokenGovernorRoleRevoked)
	if err := _TokenGovernor.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorTokenFunctionExecutedIterator struct {
	Event *TokenGovernorTokenFunctionExecuted

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorTokenFunctionExecutedIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorTokenFunctionExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorTokenFunctionExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorTokenFunctionExecutedIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorTokenFunctionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorTokenFunctionExecuted struct {
	Caller     common.Address
	Data       []byte
	ReturnData []byte
	Raw        types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterTokenFunctionExecuted(opts *bind.FilterOpts, caller []common.Address) (*TokenGovernorTokenFunctionExecutedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "TokenFunctionExecuted", callerRule)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorTokenFunctionExecutedIterator{contract: _TokenGovernor.contract, event: "TokenFunctionExecuted", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchTokenFunctionExecuted(opts *bind.WatchOpts, sink chan<- *TokenGovernorTokenFunctionExecuted, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "TokenFunctionExecuted", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorTokenFunctionExecuted)
				if err := _TokenGovernor.contract.UnpackLog(event, "TokenFunctionExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseTokenFunctionExecuted(log types.Log) (*TokenGovernorTokenFunctionExecuted, error) {
	event := new(TokenGovernorTokenFunctionExecuted)
	if err := _TokenGovernor.contract.UnpackLog(event, "TokenFunctionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type TokenGovernorTokensRecoveredIterator struct {
	Event *TokenGovernorTokensRecovered

	contract *bind.BoundContract
	event    string

	logs chan types.Log
	sub  ethereum.Subscription
	done bool
	fail error
}

func (it *TokenGovernorTokensRecoveredIterator) Next() bool {

	if it.fail != nil {
		return false
	}

	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenGovernorTokensRecovered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}

	select {
	case log := <-it.logs:
		it.Event = new(TokenGovernorTokensRecovered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

func (it *TokenGovernorTokensRecoveredIterator) Error() error {
	return it.fail
}

func (it *TokenGovernorTokensRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

type TokenGovernorTokensRecovered struct {
	Caller       common.Address
	TokenAddress common.Address
	Recipient    common.Address
	Amount       *big.Int
	Raw          types.Log
}

func (_TokenGovernor *TokenGovernorFilterer) FilterTokensRecovered(opts *bind.FilterOpts, caller []common.Address, tokenAddress []common.Address, recipient []common.Address) (*TokenGovernorTokensRecoveredIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenGovernor.contract.FilterLogs(opts, "TokensRecovered", callerRule, tokenAddressRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &TokenGovernorTokensRecoveredIterator{contract: _TokenGovernor.contract, event: "TokensRecovered", logs: logs, sub: sub}, nil
}

func (_TokenGovernor *TokenGovernorFilterer) WatchTokensRecovered(opts *bind.WatchOpts, sink chan<- *TokenGovernorTokensRecovered, caller []common.Address, tokenAddress []common.Address, recipient []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _TokenGovernor.contract.WatchLogs(opts, "TokensRecovered", callerRule, tokenAddressRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:

				event := new(TokenGovernorTokensRecovered)
				if err := _TokenGovernor.contract.UnpackLog(event, "TokensRecovered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

func (_TokenGovernor *TokenGovernorFilterer) ParseTokensRecovered(log types.Log) (*TokenGovernorTokensRecovered, error) {
	event := new(TokenGovernorTokensRecovered)
	if err := _TokenGovernor.contract.UnpackLog(event, "TokensRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

type PendingDefaultAdmin struct {
	NewAdmin common.Address
	Schedule *big.Int
}
type PendingDefaultAdminDelay struct {
	NewDelay *big.Int
	Schedule *big.Int
}

func (_TokenGovernor *TokenGovernor) ParseLog(log types.Log) (generated.AbigenLog, error) {
	switch log.Topics[0] {
	case _TokenGovernor.abi.Events["AccountFrozen"].ID:
		return _TokenGovernor.ParseAccountFrozen(log)
	case _TokenGovernor.abi.Events["AccountUnfrozen"].ID:
		return _TokenGovernor.ParseAccountUnfrozen(log)
	case _TokenGovernor.abi.Events["BridgeBurn"].ID:
		return _TokenGovernor.ParseBridgeBurn(log)
	case _TokenGovernor.abi.Events["BridgeMint"].ID:
		return _TokenGovernor.ParseBridgeMint(log)
	case _TokenGovernor.abi.Events["CheckerUpdated"].ID:
		return _TokenGovernor.ParseCheckerUpdated(log)
	case _TokenGovernor.abi.Events["ContractPaused"].ID:
		return _TokenGovernor.ParseContractPaused(log)
	case _TokenGovernor.abi.Events["ContractUnpaused"].ID:
		return _TokenGovernor.ParseContractUnpaused(log)
	case _TokenGovernor.abi.Events["DefaultAdminDelayChangeCanceled"].ID:
		return _TokenGovernor.ParseDefaultAdminDelayChangeCanceled(log)
	case _TokenGovernor.abi.Events["DefaultAdminDelayChangeScheduled"].ID:
		return _TokenGovernor.ParseDefaultAdminDelayChangeScheduled(log)
	case _TokenGovernor.abi.Events["DefaultAdminTransferCanceled"].ID:
		return _TokenGovernor.ParseDefaultAdminTransferCanceled(log)
	case _TokenGovernor.abi.Events["DefaultAdminTransferScheduled"].ID:
		return _TokenGovernor.ParseDefaultAdminTransferScheduled(log)
	case _TokenGovernor.abi.Events["FrozenAccountDrained"].ID:
		return _TokenGovernor.ParseFrozenAccountDrained(log)
	case _TokenGovernor.abi.Events["GovernedTokensRecovered"].ID:
		return _TokenGovernor.ParseGovernedTokensRecovered(log)
	case _TokenGovernor.abi.Events["NativeBurn"].ID:
		return _TokenGovernor.ParseNativeBurn(log)
	case _TokenGovernor.abi.Events["NativeMint"].ID:
		return _TokenGovernor.ParseNativeMint(log)
	case _TokenGovernor.abi.Events["OwnershipAccepted"].ID:
		return _TokenGovernor.ParseOwnershipAccepted(log)
	case _TokenGovernor.abi.Events["OwnershipTransferred"].ID:
		return _TokenGovernor.ParseOwnershipTransferred(log)
	case _TokenGovernor.abi.Events["RoleAdminChanged"].ID:
		return _TokenGovernor.ParseRoleAdminChanged(log)
	case _TokenGovernor.abi.Events["RoleGranted"].ID:
		return _TokenGovernor.ParseRoleGranted(log)
	case _TokenGovernor.abi.Events["RoleRevoked"].ID:
		return _TokenGovernor.ParseRoleRevoked(log)
	case _TokenGovernor.abi.Events["TokenFunctionExecuted"].ID:
		return _TokenGovernor.ParseTokenFunctionExecuted(log)
	case _TokenGovernor.abi.Events["TokensRecovered"].ID:
		return _TokenGovernor.ParseTokensRecovered(log)

	default:
		return nil, fmt.Errorf("abigen wrapper received unknown log topic: %v", log.Topics[0])
	}
}

func (TokenGovernorAccountFrozen) Topic() common.Hash {
	return common.HexToHash("0x2a3de20682fb291f444b5c1469d7e0950c558ce3dadf97163687873e29bcf4ae")
}

func (TokenGovernorAccountUnfrozen) Topic() common.Hash {
	return common.HexToHash("0xe19c610e04dba2019efcfb0f9455fad3af646853bb02abad2a452db1fd47c327")
}

func (TokenGovernorBridgeBurn) Topic() common.Hash {
	return common.HexToHash("0xabf8a0bc0c6341b64dfa026a551cda9d3beb0e0525758303026bacbc11ad1d8c")
}

func (TokenGovernorBridgeMint) Topic() common.Hash {
	return common.HexToHash("0xc3b8ae385c02b938fbbbd694d1da0761f755ee2c942f26bbfc6723c986d31b9b")
}

func (TokenGovernorCheckerUpdated) Topic() common.Hash {
	return common.HexToHash("0x35631e65f3bb46b57c2cfadf0d2ca3f08e6cbab52c2ff987b5328c9b37d45ea3")
}

func (TokenGovernorContractPaused) Topic() common.Hash {
	return common.HexToHash("0x81990fd9a5c552b8e3677917d8a03c07678f0d2cb68f88b634aca2022e9bd19f")
}

func (TokenGovernorContractUnpaused) Topic() common.Hash {
	return common.HexToHash("0x5b65b0c1363b3003db9bcc5e1fd8805a6d6bf5bf6dc9d3431ee4494cd7d11766")
}

func (TokenGovernorDefaultAdminDelayChangeCanceled) Topic() common.Hash {
	return common.HexToHash("0x2b1fa2edafe6f7b9e97c1a9e0c3660e645beb2dcaa2d45bdbf9beaf5472e1ec5")
}

func (TokenGovernorDefaultAdminDelayChangeScheduled) Topic() common.Hash {
	return common.HexToHash("0xf1038c18cf84a56e432fdbfaf746924b7ea511dfe03a6506a0ceba4888788d9b")
}

func (TokenGovernorDefaultAdminTransferCanceled) Topic() common.Hash {
	return common.HexToHash("0x8886ebfc4259abdbc16601dd8fb5678e54878f47b3c34836cfc51154a9605109")
}

func (TokenGovernorDefaultAdminTransferScheduled) Topic() common.Hash {
	return common.HexToHash("0x3377dc44241e779dd06afab5b788a35ca5f3b778836e2990bdb26a2a4b2e5ed6")
}

func (TokenGovernorFrozenAccountDrained) Topic() common.Hash {
	return common.HexToHash("0x8e535556a0a95ee3befe296cf986f7bf7d88881991e46f517c4b477c0ea69385")
}

func (TokenGovernorGovernedTokensRecovered) Topic() common.Hash {
	return common.HexToHash("0x42f446384ff78c62bf2caa065e680c6010e2f54857d15bee828d6c7d57353893")
}

func (TokenGovernorNativeBurn) Topic() common.Hash {
	return common.HexToHash("0x76bf0a63dad9216ecf94a5b1fdefa2c44ed7c809fe0028c4c2ce064b151c903f")
}

func (TokenGovernorNativeMint) Topic() common.Hash {
	return common.HexToHash("0xc8a07fbf83d1ff93486eea14a0676adef77c8c95fa79c85cd7bba0d10204102f")
}

func (TokenGovernorOwnershipAccepted) Topic() common.Hash {
	return common.HexToHash("0xb27970c1714b28277b78cc17ac2fe9556e7f048cd48358cffe3dc7d547608fdc")
}

func (TokenGovernorOwnershipTransferred) Topic() common.Hash {
	return common.HexToHash("0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0")
}

func (TokenGovernorRoleAdminChanged) Topic() common.Hash {
	return common.HexToHash("0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff")
}

func (TokenGovernorRoleGranted) Topic() common.Hash {
	return common.HexToHash("0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d")
}

func (TokenGovernorRoleRevoked) Topic() common.Hash {
	return common.HexToHash("0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b")
}

func (TokenGovernorTokenFunctionExecuted) Topic() common.Hash {
	return common.HexToHash("0x3f341ef37380014c90cd66348f1e20455df1b642e6a3445f39dd66f88110cc65")
}

func (TokenGovernorTokensRecovered) Topic() common.Hash {
	return common.HexToHash("0xa2231b10d9b4e4166c8a827c99f97691b05aa88fb04e009a4e499005b5c50fcc")
}

func (_TokenGovernor *TokenGovernor) Address() common.Address {
	return _TokenGovernor.address
}

type TokenGovernorInterface interface {
	BRIDGEMINTERORBURNERROLE(opts *bind.CallOpts) ([32]byte, error)

	BURNERROLE(opts *bind.CallOpts) ([32]byte, error)

	CHECKERADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error)

	FREEZERROLE(opts *bind.CallOpts) ([32]byte, error)

	MINTERROLE(opts *bind.CallOpts) ([32]byte, error)

	PAUSERROLE(opts *bind.CallOpts) ([32]byte, error)

	RECOVERYROLE(opts *bind.CallOpts) ([32]byte, error)

	UNFREEZERROLE(opts *bind.CallOpts) ([32]byte, error)

	UNPAUSERROLE(opts *bind.CallOpts) ([32]byte, error)

	DefaultAdmin(opts *bind.CallOpts) (common.Address, error)

	DefaultAdminDelay(opts *bind.CallOpts) (*big.Int, error)

	DefaultAdminDelayIncreaseWait(opts *bind.CallOpts) (*big.Int, error)

	GetAdmins(opts *bind.CallOpts) ([]common.Address, error)

	GetBridgeMintersOrBurners(opts *bind.CallOpts) ([]common.Address, error)

	GetBurners(opts *bind.CallOpts) ([]common.Address, error)

	GetChecker(opts *bind.CallOpts) (common.Address, error)

	GetCheckerAdmins(opts *bind.CallOpts) ([]common.Address, error)

	GetFreezers(opts *bind.CallOpts) ([]common.Address, error)

	GetMinters(opts *bind.CallOpts) ([]common.Address, error)

	GetPausers(opts *bind.CallOpts) ([]common.Address, error)

	GetRecoveryManagers(opts *bind.CallOpts) ([]common.Address, error)

	GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error)

	GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error)

	GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error)

	GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error)

	GetToken(opts *bind.CallOpts) (common.Address, error)

	GetTokenBalance(opts *bind.CallOpts) (*big.Int, error)

	GetUnfreezers(opts *bind.CallOpts) ([]common.Address, error)

	GetUnpausers(opts *bind.CallOpts) ([]common.Address, error)

	HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error)

	Owner(opts *bind.CallOpts) (common.Address, error)

	PendingDefaultAdmin(opts *bind.CallOpts) (PendingDefaultAdmin,

		error)

	PendingDefaultAdminDelay(opts *bind.CallOpts) (PendingDefaultAdminDelay,

		error)

	SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error)

	AcceptDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error)

	AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error)

	BatchDrainFrozenAccounts(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error)

	BatchFreeze(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error)

	BatchUnfreeze(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error)

	BeginDefaultAdminTransfer(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error)

	Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	BurnFrom(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error)

	CancelDefaultAdminTransfer(opts *bind.TransactOpts) (*types.Transaction, error)

	ChangeDefaultAdminDelay(opts *bind.TransactOpts, newDelay *big.Int) (*types.Transaction, error)

	DrainFrozenAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error)

	ExecuteTokenFunction(opts *bind.TransactOpts, data []byte) (*types.Transaction, error)

	Freeze(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error)

	GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	Mint(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	Mint0(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error)

	Pause(opts *bind.TransactOpts) (*types.Transaction, error)

	RecoverERC20(opts *bind.TransactOpts, token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	RecoverGovernedTokenERC20(opts *bind.TransactOpts, token common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error)

	RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error)

	RollbackDefaultAdminDelay(opts *bind.TransactOpts) (*types.Transaction, error)

	SetChecker(opts *bind.TransactOpts, newChecker common.Address) (*types.Transaction, error)

	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)

	Unfreeze(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error)

	Unpause(opts *bind.TransactOpts) (*types.Transaction, error)

	FilterAccountFrozen(opts *bind.FilterOpts, caller []common.Address, account []common.Address) (*TokenGovernorAccountFrozenIterator, error)

	WatchAccountFrozen(opts *bind.WatchOpts, sink chan<- *TokenGovernorAccountFrozen, caller []common.Address, account []common.Address) (event.Subscription, error)

	ParseAccountFrozen(log types.Log) (*TokenGovernorAccountFrozen, error)

	FilterAccountUnfrozen(opts *bind.FilterOpts, caller []common.Address, account []common.Address) (*TokenGovernorAccountUnfrozenIterator, error)

	WatchAccountUnfrozen(opts *bind.WatchOpts, sink chan<- *TokenGovernorAccountUnfrozen, caller []common.Address, account []common.Address) (event.Subscription, error)

	ParseAccountUnfrozen(log types.Log) (*TokenGovernorAccountUnfrozen, error)

	FilterBridgeBurn(opts *bind.FilterOpts, caller []common.Address, from []common.Address) (*TokenGovernorBridgeBurnIterator, error)

	WatchBridgeBurn(opts *bind.WatchOpts, sink chan<- *TokenGovernorBridgeBurn, caller []common.Address, from []common.Address) (event.Subscription, error)

	ParseBridgeBurn(log types.Log) (*TokenGovernorBridgeBurn, error)

	FilterBridgeMint(opts *bind.FilterOpts, caller []common.Address, recipient []common.Address) (*TokenGovernorBridgeMintIterator, error)

	WatchBridgeMint(opts *bind.WatchOpts, sink chan<- *TokenGovernorBridgeMint, caller []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseBridgeMint(log types.Log) (*TokenGovernorBridgeMint, error)

	FilterCheckerUpdated(opts *bind.FilterOpts, previousChecker []common.Address, newChecker []common.Address) (*TokenGovernorCheckerUpdatedIterator, error)

	WatchCheckerUpdated(opts *bind.WatchOpts, sink chan<- *TokenGovernorCheckerUpdated, previousChecker []common.Address, newChecker []common.Address) (event.Subscription, error)

	ParseCheckerUpdated(log types.Log) (*TokenGovernorCheckerUpdated, error)

	FilterContractPaused(opts *bind.FilterOpts, caller []common.Address) (*TokenGovernorContractPausedIterator, error)

	WatchContractPaused(opts *bind.WatchOpts, sink chan<- *TokenGovernorContractPaused, caller []common.Address) (event.Subscription, error)

	ParseContractPaused(log types.Log) (*TokenGovernorContractPaused, error)

	FilterContractUnpaused(opts *bind.FilterOpts, caller []common.Address) (*TokenGovernorContractUnpausedIterator, error)

	WatchContractUnpaused(opts *bind.WatchOpts, sink chan<- *TokenGovernorContractUnpaused, caller []common.Address) (event.Subscription, error)

	ParseContractUnpaused(log types.Log) (*TokenGovernorContractUnpaused, error)

	FilterDefaultAdminDelayChangeCanceled(opts *bind.FilterOpts) (*TokenGovernorDefaultAdminDelayChangeCanceledIterator, error)

	WatchDefaultAdminDelayChangeCanceled(opts *bind.WatchOpts, sink chan<- *TokenGovernorDefaultAdminDelayChangeCanceled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeCanceled(log types.Log) (*TokenGovernorDefaultAdminDelayChangeCanceled, error)

	FilterDefaultAdminDelayChangeScheduled(opts *bind.FilterOpts) (*TokenGovernorDefaultAdminDelayChangeScheduledIterator, error)

	WatchDefaultAdminDelayChangeScheduled(opts *bind.WatchOpts, sink chan<- *TokenGovernorDefaultAdminDelayChangeScheduled) (event.Subscription, error)

	ParseDefaultAdminDelayChangeScheduled(log types.Log) (*TokenGovernorDefaultAdminDelayChangeScheduled, error)

	FilterDefaultAdminTransferCanceled(opts *bind.FilterOpts) (*TokenGovernorDefaultAdminTransferCanceledIterator, error)

	WatchDefaultAdminTransferCanceled(opts *bind.WatchOpts, sink chan<- *TokenGovernorDefaultAdminTransferCanceled) (event.Subscription, error)

	ParseDefaultAdminTransferCanceled(log types.Log) (*TokenGovernorDefaultAdminTransferCanceled, error)

	FilterDefaultAdminTransferScheduled(opts *bind.FilterOpts, newAdmin []common.Address) (*TokenGovernorDefaultAdminTransferScheduledIterator, error)

	WatchDefaultAdminTransferScheduled(opts *bind.WatchOpts, sink chan<- *TokenGovernorDefaultAdminTransferScheduled, newAdmin []common.Address) (event.Subscription, error)

	ParseDefaultAdminTransferScheduled(log types.Log) (*TokenGovernorDefaultAdminTransferScheduled, error)

	FilterFrozenAccountDrained(opts *bind.FilterOpts, caller []common.Address, account []common.Address) (*TokenGovernorFrozenAccountDrainedIterator, error)

	WatchFrozenAccountDrained(opts *bind.WatchOpts, sink chan<- *TokenGovernorFrozenAccountDrained, caller []common.Address, account []common.Address) (event.Subscription, error)

	ParseFrozenAccountDrained(log types.Log) (*TokenGovernorFrozenAccountDrained, error)

	FilterGovernedTokensRecovered(opts *bind.FilterOpts, caller []common.Address, token []common.Address, recipient []common.Address) (*TokenGovernorGovernedTokensRecoveredIterator, error)

	WatchGovernedTokensRecovered(opts *bind.WatchOpts, sink chan<- *TokenGovernorGovernedTokensRecovered, caller []common.Address, token []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseGovernedTokensRecovered(log types.Log) (*TokenGovernorGovernedTokensRecovered, error)

	FilterNativeBurn(opts *bind.FilterOpts, caller []common.Address, from []common.Address) (*TokenGovernorNativeBurnIterator, error)

	WatchNativeBurn(opts *bind.WatchOpts, sink chan<- *TokenGovernorNativeBurn, caller []common.Address, from []common.Address) (event.Subscription, error)

	ParseNativeBurn(log types.Log) (*TokenGovernorNativeBurn, error)

	FilterNativeMint(opts *bind.FilterOpts, caller []common.Address, recipient []common.Address) (*TokenGovernorNativeMintIterator, error)

	WatchNativeMint(opts *bind.WatchOpts, sink chan<- *TokenGovernorNativeMint, caller []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseNativeMint(log types.Log) (*TokenGovernorNativeMint, error)

	FilterOwnershipAccepted(opts *bind.FilterOpts, caller []common.Address) (*TokenGovernorOwnershipAcceptedIterator, error)

	WatchOwnershipAccepted(opts *bind.WatchOpts, sink chan<- *TokenGovernorOwnershipAccepted, caller []common.Address) (event.Subscription, error)

	ParseOwnershipAccepted(log types.Log) (*TokenGovernorOwnershipAccepted, error)

	FilterOwnershipTransferred(opts *bind.FilterOpts, caller []common.Address, newOwner []common.Address) (*TokenGovernorOwnershipTransferredIterator, error)

	WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenGovernorOwnershipTransferred, caller []common.Address, newOwner []common.Address) (event.Subscription, error)

	ParseOwnershipTransferred(log types.Log) (*TokenGovernorOwnershipTransferred, error)

	FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TokenGovernorRoleAdminChangedIterator, error)

	WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TokenGovernorRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error)

	ParseRoleAdminChanged(log types.Log) (*TokenGovernorRoleAdminChanged, error)

	FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TokenGovernorRoleGrantedIterator, error)

	WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TokenGovernorRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleGranted(log types.Log) (*TokenGovernorRoleGranted, error)

	FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TokenGovernorRoleRevokedIterator, error)

	WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TokenGovernorRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error)

	ParseRoleRevoked(log types.Log) (*TokenGovernorRoleRevoked, error)

	FilterTokenFunctionExecuted(opts *bind.FilterOpts, caller []common.Address) (*TokenGovernorTokenFunctionExecutedIterator, error)

	WatchTokenFunctionExecuted(opts *bind.WatchOpts, sink chan<- *TokenGovernorTokenFunctionExecuted, caller []common.Address) (event.Subscription, error)

	ParseTokenFunctionExecuted(log types.Log) (*TokenGovernorTokenFunctionExecuted, error)

	FilterTokensRecovered(opts *bind.FilterOpts, caller []common.Address, tokenAddress []common.Address, recipient []common.Address) (*TokenGovernorTokensRecoveredIterator, error)

	WatchTokensRecovered(opts *bind.WatchOpts, sink chan<- *TokenGovernorTokensRecovered, caller []common.Address, tokenAddress []common.Address, recipient []common.Address) (event.Subscription, error)

	ParseTokensRecovered(log types.Log) (*TokenGovernorTokensRecovered, error)

	ParseLog(log types.Log) (generated.AbigenLog, error)

	Address() common.Address
}
