// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
)

// StorjPinMetaData contains all meta data concerning the StorjPin contract.
var StorjPinMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"parse\",\"type\":\"bool\"}],\"name\":\"Pinned\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"parse\",\"type\":\"bool\"}],\"name\":\"pin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StorjPinABI is the input ABI used to generate the binding from.
// Deprecated: Use StorjPinMetaData.ABI instead.
var StorjPinABI = StorjPinMetaData.ABI

// StorjPin is an auto generated Go binding around an Ethereum contract.
type StorjPin struct {
	StorjPinCaller     // Read-only binding to the contract
	StorjPinTransactor // Write-only binding to the contract
	StorjPinFilterer   // Log filterer for contract events
}

// StorjPinCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorjPinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorjPinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorjPinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorjPinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorjPinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorjPinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorjPinSession struct {
	Contract     *StorjPin         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorjPinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorjPinCallerSession struct {
	Contract *StorjPinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StorjPinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorjPinTransactorSession struct {
	Contract     *StorjPinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StorjPinRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorjPinRaw struct {
	Contract *StorjPin // Generic contract binding to access the raw methods on
}

// StorjPinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorjPinCallerRaw struct {
	Contract *StorjPinCaller // Generic read-only contract binding to access the raw methods on
}

// StorjPinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorjPinTransactorRaw struct {
	Contract *StorjPinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorjPin creates a new instance of StorjPin, bound to a specific deployed contract.
func NewStorjPin(address common.Address, backend bind.ContractBackend) (*StorjPin, error) {
	contract, err := bindStorjPin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorjPin{StorjPinCaller: StorjPinCaller{contract: contract}, StorjPinTransactor: StorjPinTransactor{contract: contract}, StorjPinFilterer: StorjPinFilterer{contract: contract}}, nil
}

// NewStorjPinCaller creates a new read-only instance of StorjPin, bound to a specific deployed contract.
func NewStorjPinCaller(address common.Address, caller bind.ContractCaller) (*StorjPinCaller, error) {
	contract, err := bindStorjPin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorjPinCaller{contract: contract}, nil
}

// NewStorjPinTransactor creates a new write-only instance of StorjPin, bound to a specific deployed contract.
func NewStorjPinTransactor(address common.Address, transactor bind.ContractTransactor) (*StorjPinTransactor, error) {
	contract, err := bindStorjPin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorjPinTransactor{contract: contract}, nil
}

// NewStorjPinFilterer creates a new log filterer instance of StorjPin, bound to a specific deployed contract.
func NewStorjPinFilterer(address common.Address, filterer bind.ContractFilterer) (*StorjPinFilterer, error) {
	contract, err := bindStorjPin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorjPinFilterer{contract: contract}, nil
}

// bindStorjPin binds a generic wrapper to an already deployed contract.
func bindStorjPin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorjPinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorjPin *StorjPinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorjPin.Contract.StorjPinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorjPin *StorjPinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorjPin.Contract.StorjPinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorjPin *StorjPinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorjPin.Contract.StorjPinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorjPin *StorjPinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorjPin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorjPin *StorjPinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorjPin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorjPin *StorjPinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorjPin.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StorjPin *StorjPinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StorjPin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StorjPin *StorjPinSession) Owner() (common.Address, error) {
	return _StorjPin.Contract.Owner(&_StorjPin.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StorjPin *StorjPinCallerSession) Owner() (common.Address, error) {
	return _StorjPin.Contract.Owner(&_StorjPin.CallOpts)
}

// Pin is a paid mutator transaction binding the contract method 0xdec07c35.
//
// Solidity: function pin(string ipfsHash, uint256 tokenAmount, bool parse) returns()
func (_StorjPin *StorjPinTransactor) Pin(opts *bind.TransactOpts, ipfsHash string, tokenAmount *big.Int, parse bool) (*types.Transaction, error) {
	return _StorjPin.contract.Transact(opts, "pin", ipfsHash, tokenAmount, parse)
}

// Pin is a paid mutator transaction binding the contract method 0xdec07c35.
//
// Solidity: function pin(string ipfsHash, uint256 tokenAmount, bool parse) returns()
func (_StorjPin *StorjPinSession) Pin(ipfsHash string, tokenAmount *big.Int, parse bool) (*types.Transaction, error) {
	return _StorjPin.Contract.Pin(&_StorjPin.TransactOpts, ipfsHash, tokenAmount, parse)
}

// Pin is a paid mutator transaction binding the contract method 0xdec07c35.
//
// Solidity: function pin(string ipfsHash, uint256 tokenAmount, bool parse) returns()
func (_StorjPin *StorjPinTransactorSession) Pin(ipfsHash string, tokenAmount *big.Int, parse bool) (*types.Transaction, error) {
	return _StorjPin.Contract.Pin(&_StorjPin.TransactOpts, ipfsHash, tokenAmount, parse)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StorjPin *StorjPinTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorjPin.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StorjPin *StorjPinSession) RenounceOwnership() (*types.Transaction, error) {
	return _StorjPin.Contract.RenounceOwnership(&_StorjPin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StorjPin *StorjPinTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StorjPin.Contract.RenounceOwnership(&_StorjPin.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StorjPin *StorjPinTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StorjPin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StorjPin *StorjPinSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StorjPin.Contract.TransferOwnership(&_StorjPin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StorjPin *StorjPinTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StorjPin.Contract.TransferOwnership(&_StorjPin.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address target) returns()
func (_StorjPin *StorjPinTransactor) Withdraw(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _StorjPin.contract.Transact(opts, "withdraw", target)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address target) returns()
func (_StorjPin *StorjPinSession) Withdraw(target common.Address) (*types.Transaction, error) {
	return _StorjPin.Contract.Withdraw(&_StorjPin.TransactOpts, target)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address target) returns()
func (_StorjPin *StorjPinTransactorSession) Withdraw(target common.Address) (*types.Transaction, error) {
	return _StorjPin.Contract.Withdraw(&_StorjPin.TransactOpts, target)
}

// StorjPinOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StorjPin contract.
type StorjPinOwnershipTransferredIterator struct {
	Event *StorjPinOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StorjPinOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorjPinOwnershipTransferred)
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
		it.Event = new(StorjPinOwnershipTransferred)
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
func (it *StorjPinOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorjPinOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorjPinOwnershipTransferred represents a OwnershipTransferred event raised by the StorjPin contract.
type StorjPinOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StorjPin *StorjPinFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StorjPinOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StorjPin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StorjPinOwnershipTransferredIterator{contract: _StorjPin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StorjPin *StorjPinFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StorjPinOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StorjPin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorjPinOwnershipTransferred)
				if err := _StorjPin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StorjPin *StorjPinFilterer) ParseOwnershipTransferred(log types.Log) (*StorjPinOwnershipTransferred, error) {
	event := new(StorjPinOwnershipTransferred)
	if err := _StorjPin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorjPinPinnedIterator is returned from FilterPinned and is used to iterate over the raw logs and unpacked data for Pinned events raised by the StorjPin contract.
type StorjPinPinnedIterator struct {
	Event *StorjPinPinned // Event containing the contract specifics and raw log

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
func (it *StorjPinPinnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorjPinPinned)
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
		it.Event = new(StorjPinPinned)
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
func (it *StorjPinPinnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorjPinPinnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorjPinPinned represents a Pinned event raised by the StorjPin contract.
type StorjPinPinned struct {
	Owner  common.Address
	Amount *big.Int
	Hash   string
	Parse  bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPinned is a free log retrieval operation binding the contract event 0x8ddc817a12aa59f9ca11b2f6d77f8e9bfedc970fe043fde2530971d66753dbd5.
//
// Solidity: event Pinned(address indexed owner, uint256 amount, string hash, bool parse)
func (_StorjPin *StorjPinFilterer) FilterPinned(opts *bind.FilterOpts, owner []common.Address) (*StorjPinPinnedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _StorjPin.contract.FilterLogs(opts, "Pinned", ownerRule)
	if err != nil {
		return nil, err
	}
	return &StorjPinPinnedIterator{contract: _StorjPin.contract, event: "Pinned", logs: logs, sub: sub}, nil
}

// WatchPinned is a free log subscription operation binding the contract event 0x8ddc817a12aa59f9ca11b2f6d77f8e9bfedc970fe043fde2530971d66753dbd5.
//
// Solidity: event Pinned(address indexed owner, uint256 amount, string hash, bool parse)
func (_StorjPin *StorjPinFilterer) WatchPinned(opts *bind.WatchOpts, sink chan<- *StorjPinPinned, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _StorjPin.contract.WatchLogs(opts, "Pinned", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorjPinPinned)
				if err := _StorjPin.contract.UnpackLog(event, "Pinned", log); err != nil {
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

// ParsePinned is a log parse operation binding the contract event 0x8ddc817a12aa59f9ca11b2f6d77f8e9bfedc970fe043fde2530971d66753dbd5.
//
// Solidity: event Pinned(address indexed owner, uint256 amount, string hash, bool parse)
func (_StorjPin *StorjPinFilterer) ParsePinned(log types.Log) (*StorjPinPinned, error) {
	event := new(StorjPinPinned)
	if err := _StorjPin.contract.UnpackLog(event, "Pinned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
