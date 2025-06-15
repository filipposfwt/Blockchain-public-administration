// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staterecords

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/crypto"
	"backend/internal/models"
)

// Reference imports to suppress errors if they are not otherwise used.
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

// StaterecordsMetaData contains all meta data concerning the Staterecords contract.
var StaterecordsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"recordId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"RecordIssued\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"recordId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"issueRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"recordId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"verifyRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StaterecordsABI is the input ABI used to generate the binding from.
// Deprecated: Use StaterecordsMetaData.ABI instead.
var StaterecordsABI = StaterecordsMetaData.ABI

// Staterecords is an auto generated Go binding around an Ethereum contract.
type Staterecords struct {
	StaterecordsCaller     // Read-only binding to the contract
	StaterecordsTransactor // Write-only binding to the contract
	StaterecordsFilterer   // Log filterer for contract events
}

// StaterecordsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StaterecordsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaterecordsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StaterecordsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaterecordsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StaterecordsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StaterecordsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StaterecordsSession struct {
	Contract     *Staterecords     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StaterecordsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StaterecordsCallerSession struct {
	Contract *StaterecordsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StaterecordsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StaterecordsTransactorSession struct {
	Contract     *StaterecordsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StaterecordsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StaterecordsRaw struct {
	Contract *Staterecords // Generic contract binding to access the raw methods on
}

// StaterecordsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StaterecordsCallerRaw struct {
	Contract *StaterecordsCaller // Generic read-only contract binding to access the raw methods on
}

// StaterecordsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StaterecordsTransactorRaw struct {
	Contract *StaterecordsTransactor // Generic write-only contract binding to access the raw methods on
}

// StateRecordsManager handles state record operations
type StateRecordsManager struct {
	// Add any necessary dependencies here
}

// NewStateRecordsManager creates a new StateRecordsManager instance
func NewStateRecordsManager() *StateRecordsManager {
	return &StateRecordsManager{}
}

// CreateStateRecord creates a new state record
func (m *StateRecordsManager) CreateStateRecord(stateInfo models.StateInfo, privateKey string) (*models.State, error) {
	// Validate the state info
	if err := m.validateStateInfo(stateInfo); err != nil {
		return nil, fmt.Errorf("invalid state info: %v", err)
	}

	// Create the appropriate certificate based on type
	var certificateData interface{}
	switch stateInfo.Type {
	case "birthCertificate":
		if stateInfo.BirthCertificate == nil {
			return nil, fmt.Errorf("birth certificate data is required")
		}
		certificateData = stateInfo.BirthCertificate
	case "deathCertificate":
		if stateInfo.DeathCertificate == nil {
			return nil, fmt.Errorf("death certificate data is required")
		}
		certificateData = stateInfo.DeathCertificate
	case "marriageCertificate":
		if stateInfo.MarriageCertificate == nil {
			return nil, fmt.Errorf("marriage certificate data is required")
		}
		certificateData = stateInfo.MarriageCertificate
	default:
		return nil, fmt.Errorf("unsupported certificate type: %s", stateInfo.Type)
	}

	// Marshal the certificate data to JSON
	certificateJSON, err := json.Marshal(certificateData)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal certificate data: %v", err)
	}

	// Create the state hash
	stateHash := crypto.Keccak256Hash(certificateJSON).Hex()

	// Generate a unique record ID if not provided
	if stateInfo.ID == "" {
		stateInfo.ID = generateRecordID(stateInfo.Type)
	}

	// Create the state record
	state := &models.State{
		StateHash: stateHash,
		StateInfo: stateInfo,
	}

	return state, nil
}

// VerifyStateRecord verifies a state record
func (m *StateRecordsManager) VerifyStateRecord(recordID string, hash string) (bool, error) {
	// Get the state record from storage (implement your storage logic here)
	state, err := m.getStateRecord(recordID)
	if err != nil {
		return false, fmt.Errorf("failed to get state record: %v", err)
	}

	// Verify the hash matches
	return state.StateHash == hash, nil
}

// getStateRecord retrieves a state record from storage
func (m *StateRecordsManager) getStateRecord(recordID string) (*models.State, error) {
	// Implement your storage retrieval logic here
	// This is a placeholder implementation
	return nil, fmt.Errorf("not implemented")
}

// validateStateInfo validates the state info structure
func (m *StateRecordsManager) validateStateInfo(stateInfo models.StateInfo) error {
	if stateInfo.Type == "" {
		return fmt.Errorf("certificate type is required")
	}

	switch stateInfo.Type {
	case "birthCertificate":
		if stateInfo.BirthCertificate == nil {
			return fmt.Errorf("birth certificate data is required")
		}
		if stateInfo.BirthCertificate.FullName == "" {
			return fmt.Errorf("full name is required for birth certificate")
		}
		if stateInfo.BirthCertificate.BirthDate == "" {
			return fmt.Errorf("birth date is required for birth certificate")
		}
	case "deathCertificate":
		if stateInfo.DeathCertificate == nil {
			return fmt.Errorf("death certificate data is required")
		}
		if stateInfo.DeathCertificate.FullName == "" {
			return fmt.Errorf("full name is required for death certificate")
		}
		if stateInfo.DeathCertificate.DeathDate == "" {
			return fmt.Errorf("death date is required for death certificate")
		}
	case "marriageCertificate":
		if stateInfo.MarriageCertificate == nil {
			return fmt.Errorf("marriage certificate data is required")
		}
		if stateInfo.MarriageCertificate.FullName == "" {
			return fmt.Errorf("full name is required for marriage certificate")
		}
		if stateInfo.MarriageCertificate.SpouseName == "" {
			return fmt.Errorf("spouse name is required for marriage certificate")
		}
		if stateInfo.MarriageCertificate.MarriageDate == "" {
			return fmt.Errorf("marriage date is required for marriage certificate")
		}
	default:
		return fmt.Errorf("unsupported certificate type: %s", stateInfo.Type)
	}

	return nil
}

// Helper function to generate a unique record ID
func generateRecordID(certificateType string) string {
	timestamp := time.Now().UnixNano()
	return fmt.Sprintf("%s-%d", certificateType, timestamp)
}

// NewStaterecords creates a new instance of Staterecords, bound to a specific deployed contract.
func NewStaterecords(address common.Address, backend bind.ContractBackend) (*Staterecords, error) {
	contract, err := bindStaterecords(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staterecords{StaterecordsCaller: StaterecordsCaller{contract: contract}, StaterecordsTransactor: StaterecordsTransactor{contract: contract}, StaterecordsFilterer: StaterecordsFilterer{contract: contract}}, nil
}

// NewStaterecordsCaller creates a new read-only instance of Staterecords, bound to a specific deployed contract.
func NewStaterecordsCaller(address common.Address, caller bind.ContractCaller) (*StaterecordsCaller, error) {
	contract, err := bindStaterecords(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StaterecordsCaller{contract: contract}, nil
}

// NewStaterecordsTransactor creates a new write-only instance of Staterecords, bound to a specific deployed contract.
func NewStaterecordsTransactor(address common.Address, transactor bind.ContractTransactor) (*StaterecordsTransactor, error) {
	contract, err := bindStaterecords(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StaterecordsTransactor{contract: contract}, nil
}

// NewStaterecordsFilterer creates a new log filterer instance of Staterecords, bound to a specific deployed contract.
func NewStaterecordsFilterer(address common.Address, filterer bind.ContractFilterer) (*StaterecordsFilterer, error) {
	contract, err := bindStaterecords(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StaterecordsFilterer{contract: contract}, nil
}

// bindStaterecords binds a generic wrapper to an already deployed contract.
func bindStaterecords(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StaterecordsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staterecords *StaterecordsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staterecords.Contract.StaterecordsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staterecords *StaterecordsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staterecords.Contract.StaterecordsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staterecords *StaterecordsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staterecords.Contract.StaterecordsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staterecords *StaterecordsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staterecords.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staterecords *StaterecordsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staterecords.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staterecords *StaterecordsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staterecords.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staterecords *StaterecordsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staterecords.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staterecords *StaterecordsSession) Owner() (common.Address, error) {
	return _Staterecords.Contract.Owner(&_Staterecords.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staterecords *StaterecordsCallerSession) Owner() (common.Address, error) {
	return _Staterecords.Contract.Owner(&_Staterecords.CallOpts)
}

// IssueRecord is a paid mutator transaction binding the contract method 0x12345678.
//
// Solidity: function issueRecord(bytes32 recordId, bytes32 hash) returns()
func (_Staterecords *StaterecordsTransactor) IssueRecord(opts *bind.TransactOpts, recordId [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _Staterecords.contract.Transact(opts, "issueRecord", recordId, hash)
}

// IssueRecord is a paid mutator transaction binding the contract method 0x12345678.
//
// Solidity: function issueRecord(bytes32 recordId, bytes32 hash) returns()
func (_Staterecords *StaterecordsSession) IssueRecord(recordId [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _Staterecords.Contract.IssueRecord(&_Staterecords.TransactOpts, recordId, hash)
}

// IssueRecord is a paid mutator transaction binding the contract method 0x12345678.
//
// Solidity: function issueRecord(bytes32 recordId, bytes32 hash) returns()
func (_Staterecords *StaterecordsTransactorSession) IssueRecord(recordId [32]byte, hash [32]byte) (*types.Transaction, error) {
	return _Staterecords.Contract.IssueRecord(&_Staterecords.TransactOpts, recordId, hash)
}

// VerifyRecord is a free data retrieval call binding the contract method 0x87654321.
//
// Solidity: function verifyRecord(bytes32 recordId, bytes32 hash) view returns(bool)
func (_Staterecords *StaterecordsCaller) VerifyRecord(opts *bind.CallOpts, recordId [32]byte, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _Staterecords.contract.Call(opts, &out, "verifyRecord", recordId, hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

// VerifyRecord is a free data retrieval call binding the contract method 0x87654321.
//
// Solidity: function verifyRecord(bytes32 recordId, bytes32 hash) view returns(bool)
func (_Staterecords *StaterecordsSession) VerifyRecord(recordId [32]byte, hash [32]byte) (bool, error) {
	return _Staterecords.Contract.VerifyRecord(&_Staterecords.CallOpts, recordId, hash)
}

// VerifyRecord is a free data retrieval call binding the contract method 0x87654321.
//
// Solidity: function verifyRecord(bytes32 recordId, bytes32 hash) view returns(bool)
func (_Staterecords *StaterecordsCallerSession) VerifyRecord(recordId [32]byte, hash [32]byte) (bool, error) {
	return _Staterecords.Contract.VerifyRecord(&_Staterecords.CallOpts, recordId, hash)
} 