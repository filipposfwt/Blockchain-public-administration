// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package diplomamanager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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

// DiplomamanagerMetaData contains all meta data concerning the Diplomamanager contract.
var DiplomamanagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"diplomaHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"issuerDID\",\"type\":\"string\"}],\"name\":\"DiplomaIssued\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"diplomaHash\",\"type\":\"bytes32\"}],\"name\":\"getIssuerDID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"diplomaHash\",\"type\":\"bytes32\"}],\"name\":\"isDiplomaStored\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"diplomaHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"issuerDID\",\"type\":\"string\"}],\"name\":\"storeDiploma\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DiplomamanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DiplomamanagerMetaData.ABI instead.
var DiplomamanagerABI = DiplomamanagerMetaData.ABI

// Diplomamanager is an auto generated Go binding around an Ethereum contract.
type Diplomamanager struct {
	DiplomamanagerCaller     // Read-only binding to the contract
	DiplomamanagerTransactor // Write-only binding to the contract
	DiplomamanagerFilterer   // Log filterer for contract events
}

// DiplomamanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiplomamanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiplomamanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiplomamanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiplomamanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiplomamanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiplomamanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiplomamanagerSession struct {
	Contract     *Diplomamanager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiplomamanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiplomamanagerCallerSession struct {
	Contract *DiplomamanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DiplomamanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiplomamanagerTransactorSession struct {
	Contract     *DiplomamanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DiplomamanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiplomamanagerRaw struct {
	Contract *Diplomamanager // Generic contract binding to access the raw methods on
}

// DiplomamanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiplomamanagerCallerRaw struct {
	Contract *DiplomamanagerCaller // Generic read-only contract binding to access the raw methods on
}

// DiplomamanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiplomamanagerTransactorRaw struct {
	Contract *DiplomamanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiplomamanager creates a new instance of Diplomamanager, bound to a specific deployed contract.
func NewDiplomamanager(address common.Address, backend bind.ContractBackend) (*Diplomamanager, error) {
	contract, err := bindDiplomamanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Diplomamanager{DiplomamanagerCaller: DiplomamanagerCaller{contract: contract}, DiplomamanagerTransactor: DiplomamanagerTransactor{contract: contract}, DiplomamanagerFilterer: DiplomamanagerFilterer{contract: contract}}, nil
}

// NewDiplomamanagerCaller creates a new read-only instance of Diplomamanager, bound to a specific deployed contract.
func NewDiplomamanagerCaller(address common.Address, caller bind.ContractCaller) (*DiplomamanagerCaller, error) {
	contract, err := bindDiplomamanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiplomamanagerCaller{contract: contract}, nil
}

// NewDiplomamanagerTransactor creates a new write-only instance of Diplomamanager, bound to a specific deployed contract.
func NewDiplomamanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*DiplomamanagerTransactor, error) {
	contract, err := bindDiplomamanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiplomamanagerTransactor{contract: contract}, nil
}

// NewDiplomamanagerFilterer creates a new log filterer instance of Diplomamanager, bound to a specific deployed contract.
func NewDiplomamanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*DiplomamanagerFilterer, error) {
	contract, err := bindDiplomamanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiplomamanagerFilterer{contract: contract}, nil
}

// bindDiplomamanager binds a generic wrapper to an already deployed contract.
func bindDiplomamanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DiplomamanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Diplomamanager *DiplomamanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Diplomamanager.Contract.DiplomamanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Diplomamanager *DiplomamanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Diplomamanager.Contract.DiplomamanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Diplomamanager *DiplomamanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Diplomamanager.Contract.DiplomamanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Diplomamanager *DiplomamanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Diplomamanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Diplomamanager *DiplomamanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Diplomamanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Diplomamanager *DiplomamanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Diplomamanager.Contract.contract.Transact(opts, method, params...)
}

// GetIssuerDID is a free data retrieval call binding the contract method 0x53fa866d.
//
// Solidity: function getIssuerDID(bytes32 diplomaHash) view returns(string)
func (_Diplomamanager *DiplomamanagerCaller) GetIssuerDID(opts *bind.CallOpts, diplomaHash [32]byte) (string, error) {
	var out []interface{}
	err := _Diplomamanager.contract.Call(opts, &out, "getIssuerDID", diplomaHash)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetIssuerDID is a free data retrieval call binding the contract method 0x53fa866d.
//
// Solidity: function getIssuerDID(bytes32 diplomaHash) view returns(string)
func (_Diplomamanager *DiplomamanagerSession) GetIssuerDID(diplomaHash [32]byte) (string, error) {
	return _Diplomamanager.Contract.GetIssuerDID(&_Diplomamanager.CallOpts, diplomaHash)
}

// GetIssuerDID is a free data retrieval call binding the contract method 0x53fa866d.
//
// Solidity: function getIssuerDID(bytes32 diplomaHash) view returns(string)
func (_Diplomamanager *DiplomamanagerCallerSession) GetIssuerDID(diplomaHash [32]byte) (string, error) {
	return _Diplomamanager.Contract.GetIssuerDID(&_Diplomamanager.CallOpts, diplomaHash)
}

// IsDiplomaStored is a free data retrieval call binding the contract method 0x29bebaca.
//
// Solidity: function isDiplomaStored(bytes32 diplomaHash) view returns(bool)
func (_Diplomamanager *DiplomamanagerCaller) IsDiplomaStored(opts *bind.CallOpts, diplomaHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Diplomamanager.contract.Call(opts, &out, "isDiplomaStored", diplomaHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDiplomaStored is a free data retrieval call binding the contract method 0x29bebaca.
//
// Solidity: function isDiplomaStored(bytes32 diplomaHash) view returns(bool)
func (_Diplomamanager *DiplomamanagerSession) IsDiplomaStored(diplomaHash [32]byte) (bool, error) {
	return _Diplomamanager.Contract.IsDiplomaStored(&_Diplomamanager.CallOpts, diplomaHash)
}

// IsDiplomaStored is a free data retrieval call binding the contract method 0x29bebaca.
//
// Solidity: function isDiplomaStored(bytes32 diplomaHash) view returns(bool)
func (_Diplomamanager *DiplomamanagerCallerSession) IsDiplomaStored(diplomaHash [32]byte) (bool, error) {
	return _Diplomamanager.Contract.IsDiplomaStored(&_Diplomamanager.CallOpts, diplomaHash)
}

// StoreDiploma is a paid mutator transaction binding the contract method 0xeb5c5bd4.
//
// Solidity: function storeDiploma(bytes32 diplomaHash, string issuerDID) returns()
func (_Diplomamanager *DiplomamanagerTransactor) StoreDiploma(opts *bind.TransactOpts, diplomaHash [32]byte, issuerDID string) (*types.Transaction, error) {
	return _Diplomamanager.contract.Transact(opts, "storeDiploma", diplomaHash, issuerDID)
}

// StoreDiploma is a paid mutator transaction binding the contract method 0xeb5c5bd4.
//
// Solidity: function storeDiploma(bytes32 diplomaHash, string issuerDID) returns()
func (_Diplomamanager *DiplomamanagerSession) StoreDiploma(diplomaHash [32]byte, issuerDID string) (*types.Transaction, error) {
	return _Diplomamanager.Contract.StoreDiploma(&_Diplomamanager.TransactOpts, diplomaHash, issuerDID)
}

// StoreDiploma is a paid mutator transaction binding the contract method 0xeb5c5bd4.
//
// Solidity: function storeDiploma(bytes32 diplomaHash, string issuerDID) returns()
func (_Diplomamanager *DiplomamanagerTransactorSession) StoreDiploma(diplomaHash [32]byte, issuerDID string) (*types.Transaction, error) {
	return _Diplomamanager.Contract.StoreDiploma(&_Diplomamanager.TransactOpts, diplomaHash, issuerDID)
}

// DiplomamanagerDiplomaIssuedIterator is returned from FilterDiplomaIssued and is used to iterate over the raw logs and unpacked data for DiplomaIssued events raised by the Diplomamanager contract.
type DiplomamanagerDiplomaIssuedIterator struct {
	Event *DiplomamanagerDiplomaIssued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DiplomamanagerDiplomaIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiplomamanagerDiplomaIssued)
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
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DiplomamanagerDiplomaIssued)
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

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DiplomamanagerDiplomaIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiplomamanagerDiplomaIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiplomamanagerDiplomaIssued represents a DiplomaIssued event raised by the Diplomamanager contract.
type DiplomamanagerDiplomaIssued struct {
	DiplomaHash [32]byte
	IssuerDID   string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDiplomaIssued is a free log retrieval operation binding the contract event 0xeb1174ba23edb80a45dbc88570362dd0353de26c1bd285c6ed99d3a97871daae.
//
// Solidity: event DiplomaIssued(bytes32 indexed diplomaHash, string issuerDID)
func (_Diplomamanager *DiplomamanagerFilterer) FilterDiplomaIssued(opts *bind.FilterOpts, diplomaHash [][32]byte) (*DiplomamanagerDiplomaIssuedIterator, error) {

	var diplomaHashRule []interface{}
	for _, diplomaHashItem := range diplomaHash {
		diplomaHashRule = append(diplomaHashRule, diplomaHashItem)
	}

	logs, sub, err := _Diplomamanager.contract.FilterLogs(opts, "DiplomaIssued", diplomaHashRule)
	if err != nil {
		return nil, err
	}
	return &DiplomamanagerDiplomaIssuedIterator{contract: _Diplomamanager.contract, event: "DiplomaIssued", logs: logs, sub: sub}, nil
}

// WatchDiplomaIssued is a free log subscription operation binding the contract event 0xeb1174ba23edb80a45dbc88570362dd0353de26c1bd285c6ed99d3a97871daae.
//
// Solidity: event DiplomaIssued(bytes32 indexed diplomaHash, string issuerDID)
func (_Diplomamanager *DiplomamanagerFilterer) WatchDiplomaIssued(opts *bind.WatchOpts, sink chan<- *DiplomamanagerDiplomaIssued, diplomaHash [][32]byte) (event.Subscription, error) {

	var diplomaHashRule []interface{}
	for _, diplomaHashItem := range diplomaHash {
		diplomaHashRule = append(diplomaHashRule, diplomaHashItem)
	}

	logs, sub, err := _Diplomamanager.contract.WatchLogs(opts, "DiplomaIssued", diplomaHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiplomamanagerDiplomaIssued)
				if err := _Diplomamanager.contract.UnpackLog(event, "DiplomaIssued", log); err != nil {
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

// ParseDiplomaIssued is a log parse operation binding the contract event 0xeb1174ba23edb80a45dbc88570362dd0353de26c1bd285c6ed99d3a97871daae.
//
// Solidity: event DiplomaIssued(bytes32 indexed diplomaHash, string issuerDID)
func (_Diplomamanager *DiplomamanagerFilterer) ParseDiplomaIssued(log types.Log) (*DiplomamanagerDiplomaIssued, error) {
	event := new(DiplomamanagerDiplomaIssued)
	if err := _Diplomamanager.contract.UnpackLog(event, "DiplomaIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
