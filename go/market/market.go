// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package market

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

// MarketOrder is an auto generated low-level Go binding around an user-defined struct.
type MarketOrder struct {
	User            common.Address
	Provider        common.Address
	P               MarketPricePerHour
	R               MarketResources
	TotalValue      *big.Int
	Remain          *big.Int
	Remuneration    *big.Int
	UserConfirm     bool
	ProviderConfirm bool
	ActivateTime    *big.Int
	LastSettleTime  *big.Int
	Probation       *big.Int
	Duration        *big.Int
	Status          uint8
}

// MarketPricePerHour is an auto generated low-level Go binding around an user-defined struct.
type MarketPricePerHour struct {
	PCPU  uint64
	PGPU  uint64
	PMEM  uint64
	PDISK uint64
}

// MarketResources is an auto generated low-level Go binding around an user-defined struct.
type MarketResources struct {
	NCPU  uint64
	NGPU  uint64
	NMEM  uint64
	NDISK uint64
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketMetaData contains all meta data concerning the Market contract.
var MarketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pCPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pGPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pMEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pDISK\",\"type\":\"uint64\"}],\"internalType\":\"structMarket.PricePerHour\",\"name\":\"p\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nCPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nGPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nMEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nDISK\",\"type\":\"uint64\"}],\"internalType\":\"structMarket.Resources\",\"name\":\"r\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remuneration\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"userConfirm\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"providerConfirm\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"activateTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastSettleTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"probation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structMarket.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"createOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pCPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pGPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pMEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pDISK\",\"type\":\"uint64\"}],\"internalType\":\"structMarket.PricePerHour\",\"name\":\"p\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nCPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nGPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nMEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nDISK\",\"type\":\"uint64\"}],\"internalType\":\"structMarket.Resources\",\"name\":\"r\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remuneration\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"userConfirm\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"providerConfirm\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"activateTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastSettleTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"probation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structMarket.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"feeOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"getOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pCPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pGPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pMEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pDISK\",\"type\":\"uint64\"}],\"internalType\":\"structMarket.PricePerHour\",\"name\":\"p\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nCPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nGPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nMEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nDISK\",\"type\":\"uint64\"}],\"internalType\":\"structMarket.Resources\",\"name\":\"r\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"totalValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remain\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remuneration\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"userConfirm\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"providerConfirm\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"activateTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastSettleTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"probation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structMarket.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"proWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"providerCancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"userCancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"userWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1c5a9d9c": "activate(address)",
		"ef380cbe": "createOrder(address,address,(address,address,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64),uint256,uint256,uint256,bool,bool,uint256,uint256,uint256,uint256,uint8))",
		"b565fd09": "feeOrder((address,address,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64),uint256,uint256,uint256,bool,bool,uint256,uint256,uint256,uint256,uint8))",
		"6eba2b13": "getOrder(address)",
		"8ebab3ea": "proWithdraw(address,address,uint256)",
		"504cb846": "providerCancel(address)",
		"39034c72": "userCancel(address)",
		"77ecf411": "userWithdraw(address,address,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061196e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806377ecf4111161005b57806377ecf411146100f15780638ebab3ea14610104578063b565fd0914610117578063ef380cbe1461013957600080fd5b80631c5a9d9c1461008d57806339034c72146100a2578063504cb846146100b55780636eba2b13146100c8575b600080fd5b6100a061009b3660046114a8565b61014c565b005b6100a06100b03660046114a8565b610389565b6100a06100c33660046114a8565b6105b9565b6100db6100d63660046114a8565b6107e5565b6040516100e891906114ca565b60405180910390f35b6100a06100ff366004611606565b6109e5565b6100a0610112366004611606565b610c91565b61012b610125366004611836565b50600090565b6040519081526020016100e8565b6100a0610147366004611853565b610f3a565b6001600160a01b03808216600090815260208181526040808320338085529083529281902081516101c0810183528154861681526001820154861681850190815283516080808201865260028501546001600160401b038082168452600160401b8083048216858b0152600160801b8084048316868b0152600160c01b938490048316606080880191909152888b019690965289518086018b5260038a0154808516825292830484169b81019b909b5281048216988a0198909852960490951686820152820194909452600482015492810192909252600581015460a0830152600681015460c0830152600781015460ff808216151560e085015261010091829004811615159184019190915260088201546101208401526009820154610140840152600a820154610160840152600b820154610180840152600c90910154166101a08201529051909216146102e95760405162461bcd60e51b815260206004820181905260248201527f6163746976617465206f6e6c792063616c6c65642062792070726f766964657260448201526064015b60405180910390fd5b6101a081015160ff161561034b5760405162461bcd60e51b8152602060048201526024808201527f6f6e6c7920756e616374697665206f726465722063616e206265206163746976604482015263185d195960e21b60648201526084016102e0565b506001600160a01b0316600090815260208181526040808320338452909152902042600982018190556008820155600c01805460ff19166001179055565b336000818152602081815260408083206001600160a01b0386811685529083529281902081516101c081018352815485168152600182015485168185015282516080808201855260028401546001600160401b038082168452600160401b8083048216858a0152600160801b8084048316868a0152600160c01b938490048316606080880191909152878a019690965288518086018a526003890154808516825292830484169a81019a909a528104821697890197909752950490941685820152810193909352600481015491830191909152600581015460a0830152600681015460c0830152600781015460ff808216151560e085015261010091829004811615159184019190915260088201546101208401526009820154610140840152600a820154610160840152600b820154610180840152600c90910154166101a08201528051909291161461051f5760405162461bcd60e51b815260206004820152601b60248201527f7468652063616c6c6572206e6f74206f7264657227732075736572000000000060448201526064016102e0565b816001600160a01b031681602001516001600160a01b03161461057c5760405162461bcd60e51b81526020600482015260156024820152743a343290383937bb34b232b91034b99032b93937b960591b60448201526064016102e0565b6105863383611167565b50336000908152602081815260408083206001600160a01b0394909416835292905220600c01805460ff19166002179055565b6001600160a01b0380821660008181526020818152604080832033845282529182902082516101c081018452815486168152600182015486168184015283516080808201865260028401546001600160401b038082168452600160401b808304821685890152600160801b8084048316868b0152600160c01b938490048316606080880191909152878b019690965289518086018b52600389015480851682529283048416998101999099528104821698880198909852960490951684820152810192909252600481015492820192909252600582015460a0820152600682015460c0820152600782015460ff808216151560e084015261010091829004811615159183019190915260088301546101208301526009830154610140830152600a830154610160830152600b830154610180830152600c909201549091166101a08201528051909216146107435760405162461bcd60e51b81526020600482015260116024820152703a3432903ab9b2b91034b99032b93937b960791b60448201526064016102e0565b60208101516001600160a01b031633146107aa5760405162461bcd60e51b815260206004820152602260248201527f7468652063616c6c6572206973206e6f74206f7264657227732070726f76696460448201526132b960f11b60648201526084016102e0565b6107b48233611167565b506001600160a01b03166000908152602081815260408083203384529091529020600c01805460ff19166002179055565b61089e604080516101c08101825260008082526020808301829052835160808101855282815290810182905280840182905260608101919091529091820190815260408051608081018252600080825260208281018290529282018190526060820152910190815260200160008152602001600081526020016000815260200160001515815260200160001515815260200160008152602001600081526020016000815260200160008152602001600060ff1681525090565b50336000908152602081815260408083206001600160a01b03948516845282529182902082516101c08101845281548516815260018201549094168483015282516080808201855260028301546001600160401b038082168452600160401b808304821685880152600160801b8084048316868a0152600160c01b9384900483166060808801919091528a8a019690965288518086018a52600388015480851682529283048416988101989098528104821697870197909752950490941683820152840191909152600481015491830191909152600581015460a0830152600681015460c0830152600781015460ff808216151560e085015261010091829004811615159184019190915260088201546101208401526009820154610140840152600a820154610160840152600b820154610180840152600c90910154166101a082015290565b336000818152602081815260408083206001600160a01b0387811685529083529281902081516101c081018352815485168152600182015485168185015282516080808201855260028401546001600160401b038082168452600160401b8083048216858a0152600160801b8084048316868a0152600160c01b938490048316606080880191909152878a019690965288518086018a526003890154808516825292830484169a81019a909a528104821697890197909752950490941685820152810193909352600481015491830191909152600581015460a0830152600681015460c0830152600781015460ff808216151560e085015261010091829004811615159184019190915260088201546101208401526009820154610140840152600a820154610160840152600b820154610180840152600c90910154166101a082015280519092911614610b7b5760405162461bcd60e51b815260206004820152601760248201527f63616c6c6572206d75737420626520746865207573657200000000000000000060448201526064016102e0565b8060a00151821115610be05760405162461bcd60e51b815260206004820152602860248201527f74686520616d6f756e742073686f756c64206e6f74206c61726765722074686160448201526737103932b6b0b4b760c11b60648201526084016102e0565b60405163a9059cbb60e01b8152336004820152602481018390526001600160a01b0385169063a9059cbb906044016020604051808303816000875af1158015610c2d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c519190611898565b50336000908152602081815260408083206001600160a01b038716845290915281206005018054849290610c869084906118cb565b909155505050505050565b6001600160a01b03808316600090815260208181526040808320338085529083529281902081516101c0810183528154861681526001820154861681850190815283516080808201865260028501546001600160401b038082168452600160401b8083048216858b0152600160801b8084048316868b0152600160c01b938490048316606080880191909152888b019690965289518086018b5260038a0154808516825292830484169b81019b909b5281048216988a0198909852960490951686820152820194909452600482015492810192909252600581015460a0830152600681015460c0830152600781015460ff808216151560e085015261010091829004811615159184019190915260088201546101208401526009820154610140840152600a820154610160840152600b820154610180840152600c90910154166101a0820152905190921614610e295760405162461bcd60e51b815260206004820152601b60248201527f63616c6c6572206d757374206265207468652070726f7669646572000000000060448201526064016102e0565b8060c00151821115610e945760405162461bcd60e51b815260206004820152602e60248201527f74686520616d6f756e742073686f756c64206e6f74206c61726765722074686160448201526d37103932b6bab732b930ba34b7b760911b60648201526084016102e0565b60405163a9059cbb60e01b8152336004820152602481018390526001600160a01b0385169063a9059cbb906044016020604051808303816000875af1158015610ee1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f059190611898565b506001600160a01b03831660009081526020818152604080832033845290915281206006018054849290610c869084906118cb565b60808101516040516323b872dd60e01b815233600482015230602482015260448101919091526001600160a01b038416906323b872dd906064016020604051808303816000875af1158015610f93573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fb79190611898565b50336000908152602081815260408083206001600160a01b039586168452825291829020835181549086166001600160a01b0319918216178255848301516001830180549190971691161790945582820151805160028601805483850151848701516060958601516001600160401b03908116600160c01b9081026001600160c01b03938316600160801b90810285166fffffffffffffffffffffffffffffffff968516600160401b9081026fffffffffffffffffffffffffffffffff19998a169b87169b909b179a909a1787161791909117909655878b0151805160038e0180549b8301519c83015192909a015184169092029083169096029092169881169095029690921691909316179390931716929092171790556080810151600483015560a0810151600583015560c0810151600683015560e08101516007830180546101008085015115150261ff00199315159390931661ffff199091161791909117905561012081015160088301556101408101516009830155610160810151600a830155610180810151600b8301556101a00151600c909101805460ff90921660ff1990921691909117905550565b6001600160a01b038083166000908152602081815260408083208585168452825280832081516101c08101835281548616815260018201549095168584015281516080808201845260028301546001600160401b038082168452600160401b808304821685890152600160801b808404831686890152600160c01b9384900483166060808801919091528b89019690965287518086018952600388015480851682529283048416998101999099528104821696880196909652940490931684820152850192909252600482015490840152600581015460a0840152600681015460c0840152600781015460ff808216151560e0860152610100918290048116151591850191909152600882015461012085015260098201546101408501819052600a80840154610160870152600b840154610180870152600c909301549091166101a08501526112b6916118e4565b9050816101a0015160ff166001146113235760405162461bcd60e51b815260206004820152602a60248201527f6f72646572206d75737420626520696e2061637469766174652073746174757360448201526920746f20736574746c6560b01b60648201526084016102e0565b81610140015181101561133557600080fd5b60008261014001518261134891906118cb565b905060008361012001518361135d91906118cb565b90506000846101800151856080015161137691906118f7565b905060008060009050866101800151841061139957505060a085015160016113a6565b6113a38584611919565b91505b6001600160a01b03808a16600090815260208181526040808320938c16835292905290812060050180548492906113de9084906118cb565b90915550506001600160a01b03808a16600090815260208181526040808320938c168352929052908120600601805484929061141b9084906118e4565b90915550506001600160a01b03808a16600090815260208181526040808320938c168352929052206009018690558015611481576001600160a01b03808a16600090815260208181526040808320938c16835292905220600c01805460ff191660031790555b505050505050505050565b80356001600160a01b03811681146114a357600080fd5b919050565b6000602082840312156114ba57600080fd5b6114c38261148c565b9392505050565b81516001600160a01b03168152610280810160208301516114f660208401826001600160a01b03169052565b506040830151611537604084018280516001600160401b03908116835260208083015182169084015260408083015182169084015260609182015116910152565b5060608381015180516001600160401b0390811660c08601526020820151811660e08601526040820151811661010086015291810151909116610120840152506080830151610140818185015260a08501519150610160828186015260c08601519250610180838187015260e087015193506101a06115b98188018615159052565b61010088015115156101c08801526101208801516101e088015292870151610200870152908601516102208601528501516102408501529093015160ff1661026090920191909152919050565b60008060006060848603121561161b57600080fd5b6116248461148c565b92506116326020850161148c565b9150604084013590509250925092565b6040516101c081016001600160401b038111828210171561167357634e487b7160e01b600052604160045260246000fd5b60405290565b80356001600160401b03811681146114a357600080fd5b6000608082840312156116a257600080fd5b604051608081018181106001600160401b03821117156116d257634e487b7160e01b600052604160045260246000fd5b6040529050806116e183611679565b81526116ef60208401611679565b602082015261170060408401611679565b604082015261171160608401611679565b60608201525092915050565b801515811461172b57600080fd5b50565b80356114a38161171d565b803560ff811681146114a357600080fd5b6000610280828403121561175d57600080fd5b611765611642565b90506117708261148c565b815261177e6020830161148c565b60208201526117908360408401611690565b60408201526117a28360c08401611690565b60608201526101408083013560808301526101608084013560a08401526101808085013560c08501526101a06117d981870161172e565b60e08601526117eb6101c0870161172e565b6101008601526101e08601356101208601526102008601358486015261022086013583860152610240860135828601526118286102608701611739565b818601525050505092915050565b6000610280828403121561184957600080fd5b6114c3838361174a565b60008060006102c0848603121561186957600080fd5b6118728461148c565b92506118806020850161148c565b915061188f856040860161174a565b90509250925092565b6000602082840312156118aa57600080fd5b81516114c38161171d565b634e487b7160e01b600052601160045260246000fd5b818103818111156118de576118de6118b5565b92915050565b808201808211156118de576118de6118b5565b60008261191457634e487b7160e01b600052601260045260246000fd5b500490565b6000816000190483118215151615611933576119336118b5565b50029056fea2646970667358221220f7d18919281c4cdca415d3a89b3b836a8e2fe46214497a3e93e7c53b6ac9012c64736f6c63430008100033",
}

// MarketABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketMetaData.ABI instead.
var MarketABI = MarketMetaData.ABI

// Deprecated: Use MarketMetaData.Sigs instead.
// MarketFuncSigs maps the 4-byte function signature to its string representation.
var MarketFuncSigs = MarketMetaData.Sigs

// MarketBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MarketMetaData.Bin instead.
var MarketBin = MarketMetaData.Bin

// DeployMarket deploys a new Ethereum contract, binding an instance of Market to it.
func DeployMarket(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Market, error) {
	parsed, err := MarketMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MarketBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Market{MarketCaller: MarketCaller{contract: contract}, MarketTransactor: MarketTransactor{contract: contract}, MarketFilterer: MarketFilterer{contract: contract}}, nil
}

// Market is an auto generated Go binding around an Ethereum contract.
type Market struct {
	MarketCaller     // Read-only binding to the contract
	MarketTransactor // Write-only binding to the contract
	MarketFilterer   // Log filterer for contract events
}

// MarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketSession struct {
	Contract     *Market           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketCallerSession struct {
	Contract *MarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketTransactorSession struct {
	Contract     *MarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketRaw struct {
	Contract *Market // Generic contract binding to access the raw methods on
}

// MarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketCallerRaw struct {
	Contract *MarketCaller // Generic read-only contract binding to access the raw methods on
}

// MarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketTransactorRaw struct {
	Contract *MarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarket creates a new instance of Market, bound to a specific deployed contract.
func NewMarket(address common.Address, backend bind.ContractBackend) (*Market, error) {
	contract, err := bindMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Market{MarketCaller: MarketCaller{contract: contract}, MarketTransactor: MarketTransactor{contract: contract}, MarketFilterer: MarketFilterer{contract: contract}}, nil
}

// NewMarketCaller creates a new read-only instance of Market, bound to a specific deployed contract.
func NewMarketCaller(address common.Address, caller bind.ContractCaller) (*MarketCaller, error) {
	contract, err := bindMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCaller{contract: contract}, nil
}

// NewMarketTransactor creates a new write-only instance of Market, bound to a specific deployed contract.
func NewMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketTransactor, error) {
	contract, err := bindMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketTransactor{contract: contract}, nil
}

// NewMarketFilterer creates a new log filterer instance of Market, bound to a specific deployed contract.
func NewMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketFilterer, error) {
	contract, err := bindMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketFilterer{contract: contract}, nil
}

// bindMarket binds a generic wrapper to an already deployed contract.
func bindMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Market.Contract.MarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Market.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.contract.Transact(opts, method, params...)
}

// FeeOrder is a free data retrieval call binding the contract method 0xb565fd09.
//
// Solidity: function feeOrder((address,address,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64),uint256,uint256,uint256,bool,bool,uint256,uint256,uint256,uint256,uint8) order) view returns(uint256)
func (_Market *MarketCaller) FeeOrder(opts *bind.CallOpts, order MarketOrder) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "feeOrder", order)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeOrder is a free data retrieval call binding the contract method 0xb565fd09.
//
// Solidity: function feeOrder((address,address,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64),uint256,uint256,uint256,bool,bool,uint256,uint256,uint256,uint256,uint8) order) view returns(uint256)
func (_Market *MarketSession) FeeOrder(order MarketOrder) (*big.Int, error) {
	return _Market.Contract.FeeOrder(&_Market.CallOpts, order)
}

// FeeOrder is a free data retrieval call binding the contract method 0xb565fd09.
//
// Solidity: function feeOrder((address,address,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64),uint256,uint256,uint256,bool,bool,uint256,uint256,uint256,uint256,uint8) order) view returns(uint256)
func (_Market *MarketCallerSession) FeeOrder(order MarketOrder) (*big.Int, error) {
	return _Market.Contract.FeeOrder(&_Market.CallOpts, order)
}

// GetOrder is a free data retrieval call binding the contract method 0x6eba2b13.
//
// Solidity: function getOrder(address provider) view returns((address,address,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64),uint256,uint256,uint256,bool,bool,uint256,uint256,uint256,uint256,uint8))
func (_Market *MarketCaller) GetOrder(opts *bind.CallOpts, provider common.Address) (MarketOrder, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getOrder", provider)

	if err != nil {
		return *new(MarketOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(MarketOrder)).(*MarketOrder)

	return out0, err

}

// GetOrder is a free data retrieval call binding the contract method 0x6eba2b13.
//
// Solidity: function getOrder(address provider) view returns((address,address,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64),uint256,uint256,uint256,bool,bool,uint256,uint256,uint256,uint256,uint8))
func (_Market *MarketSession) GetOrder(provider common.Address) (MarketOrder, error) {
	return _Market.Contract.GetOrder(&_Market.CallOpts, provider)
}

// GetOrder is a free data retrieval call binding the contract method 0x6eba2b13.
//
// Solidity: function getOrder(address provider) view returns((address,address,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64),uint256,uint256,uint256,bool,bool,uint256,uint256,uint256,uint256,uint8))
func (_Market *MarketCallerSession) GetOrder(provider common.Address) (MarketOrder, error) {
	return _Market.Contract.GetOrder(&_Market.CallOpts, provider)
}

// Activate is a paid mutator transaction binding the contract method 0x1c5a9d9c.
//
// Solidity: function activate(address user) returns()
func (_Market *MarketTransactor) Activate(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "activate", user)
}

// Activate is a paid mutator transaction binding the contract method 0x1c5a9d9c.
//
// Solidity: function activate(address user) returns()
func (_Market *MarketSession) Activate(user common.Address) (*types.Transaction, error) {
	return _Market.Contract.Activate(&_Market.TransactOpts, user)
}

// Activate is a paid mutator transaction binding the contract method 0x1c5a9d9c.
//
// Solidity: function activate(address user) returns()
func (_Market *MarketTransactorSession) Activate(user common.Address) (*types.Transaction, error) {
	return _Market.Contract.Activate(&_Market.TransactOpts, user)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xef380cbe.
//
// Solidity: function createOrder(address tokenAddr, address provider, (address,address,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64),uint256,uint256,uint256,bool,bool,uint256,uint256,uint256,uint256,uint8) order) returns()
func (_Market *MarketTransactor) CreateOrder(opts *bind.TransactOpts, tokenAddr common.Address, provider common.Address, order MarketOrder) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "createOrder", tokenAddr, provider, order)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xef380cbe.
//
// Solidity: function createOrder(address tokenAddr, address provider, (address,address,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64),uint256,uint256,uint256,bool,bool,uint256,uint256,uint256,uint256,uint8) order) returns()
func (_Market *MarketSession) CreateOrder(tokenAddr common.Address, provider common.Address, order MarketOrder) (*types.Transaction, error) {
	return _Market.Contract.CreateOrder(&_Market.TransactOpts, tokenAddr, provider, order)
}

// CreateOrder is a paid mutator transaction binding the contract method 0xef380cbe.
//
// Solidity: function createOrder(address tokenAddr, address provider, (address,address,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64),uint256,uint256,uint256,bool,bool,uint256,uint256,uint256,uint256,uint8) order) returns()
func (_Market *MarketTransactorSession) CreateOrder(tokenAddr common.Address, provider common.Address, order MarketOrder) (*types.Transaction, error) {
	return _Market.Contract.CreateOrder(&_Market.TransactOpts, tokenAddr, provider, order)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0x8ebab3ea.
//
// Solidity: function proWithdraw(address tokenAddr, address user, uint256 amount) returns()
func (_Market *MarketTransactor) ProWithdraw(opts *bind.TransactOpts, tokenAddr common.Address, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "proWithdraw", tokenAddr, user, amount)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0x8ebab3ea.
//
// Solidity: function proWithdraw(address tokenAddr, address user, uint256 amount) returns()
func (_Market *MarketSession) ProWithdraw(tokenAddr common.Address, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.ProWithdraw(&_Market.TransactOpts, tokenAddr, user, amount)
}

// ProWithdraw is a paid mutator transaction binding the contract method 0x8ebab3ea.
//
// Solidity: function proWithdraw(address tokenAddr, address user, uint256 amount) returns()
func (_Market *MarketTransactorSession) ProWithdraw(tokenAddr common.Address, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.ProWithdraw(&_Market.TransactOpts, tokenAddr, user, amount)
}

// ProviderCancel is a paid mutator transaction binding the contract method 0x504cb846.
//
// Solidity: function providerCancel(address user) returns()
func (_Market *MarketTransactor) ProviderCancel(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "providerCancel", user)
}

// ProviderCancel is a paid mutator transaction binding the contract method 0x504cb846.
//
// Solidity: function providerCancel(address user) returns()
func (_Market *MarketSession) ProviderCancel(user common.Address) (*types.Transaction, error) {
	return _Market.Contract.ProviderCancel(&_Market.TransactOpts, user)
}

// ProviderCancel is a paid mutator transaction binding the contract method 0x504cb846.
//
// Solidity: function providerCancel(address user) returns()
func (_Market *MarketTransactorSession) ProviderCancel(user common.Address) (*types.Transaction, error) {
	return _Market.Contract.ProviderCancel(&_Market.TransactOpts, user)
}

// UserCancel is a paid mutator transaction binding the contract method 0x39034c72.
//
// Solidity: function userCancel(address provider) returns()
func (_Market *MarketTransactor) UserCancel(opts *bind.TransactOpts, provider common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "userCancel", provider)
}

// UserCancel is a paid mutator transaction binding the contract method 0x39034c72.
//
// Solidity: function userCancel(address provider) returns()
func (_Market *MarketSession) UserCancel(provider common.Address) (*types.Transaction, error) {
	return _Market.Contract.UserCancel(&_Market.TransactOpts, provider)
}

// UserCancel is a paid mutator transaction binding the contract method 0x39034c72.
//
// Solidity: function userCancel(address provider) returns()
func (_Market *MarketTransactorSession) UserCancel(provider common.Address) (*types.Transaction, error) {
	return _Market.Contract.UserCancel(&_Market.TransactOpts, provider)
}

// UserWithdraw is a paid mutator transaction binding the contract method 0x77ecf411.
//
// Solidity: function userWithdraw(address tokenAddr, address provider, uint256 amount) returns()
func (_Market *MarketTransactor) UserWithdraw(opts *bind.TransactOpts, tokenAddr common.Address, provider common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "userWithdraw", tokenAddr, provider, amount)
}

// UserWithdraw is a paid mutator transaction binding the contract method 0x77ecf411.
//
// Solidity: function userWithdraw(address tokenAddr, address provider, uint256 amount) returns()
func (_Market *MarketSession) UserWithdraw(tokenAddr common.Address, provider common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.UserWithdraw(&_Market.TransactOpts, tokenAddr, provider, amount)
}

// UserWithdraw is a paid mutator transaction binding the contract method 0x77ecf411.
//
// Solidity: function userWithdraw(address tokenAddr, address provider, uint256 amount) returns()
func (_Market *MarketTransactorSession) UserWithdraw(tokenAddr common.Address, provider common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Market.Contract.UserWithdraw(&_Market.TransactOpts, tokenAddr, provider, amount)
}
