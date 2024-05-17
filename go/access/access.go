// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package access

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

// AccessMetaData contains all meta data concerning the Access contract.
var AccessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"access\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040526001600060146101000a81548161ffff021916908361ffff16021790555034801561002e57600080fd5b5061004b61004061005060201b60201c565b61005860201b60201c565b61011c565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b61077e8061012b6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063715018a61161005b578063715018a6146100ec5780638da5cb5b146100f6578063affed0e014610114578063f2fde38b146101325761007d565b806335e3b25a1461008257806354fd4d501461009e5780636fae3d76146100bc575b600080fd5b61009c600480360381019061009791906104a9565b61014e565b005b6100a66101ca565b6040516100b39190610506565b60405180910390f35b6100d660048036038101906100d19190610521565b6101de565b6040516100e3919061055d565b60405180910390f35b6100f46101fe565b005b6100fe610212565b60405161010b9190610587565b60405180910390f35b61011c61023b565b60405161012991906105bb565b60405180910390f35b61014c60048036038101906101479190610521565b610241565b005b6101566102c4565b80600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060018060008282546101bf9190610605565b925050819055505050565b600060149054906101000a900461ffff1681565b60026020528060005260406000206000915054906101000a900460ff1681565b6102066102c4565b6102106000610342565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60015481565b6102496102c4565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036102b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102af906106bc565b60405180910390fd5b6102c181610342565b50565b6102cc610406565b73ffffffffffffffffffffffffffffffffffffffff166102ea610212565b73ffffffffffffffffffffffffffffffffffffffff1614610340576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033790610728565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061043e82610413565b9050919050565b61044e81610433565b811461045957600080fd5b50565b60008135905061046b81610445565b92915050565b60008115159050919050565b61048681610471565b811461049157600080fd5b50565b6000813590506104a38161047d565b92915050565b600080604083850312156104c0576104bf61040e565b5b60006104ce8582860161045c565b92505060206104df85828601610494565b9150509250929050565b600061ffff82169050919050565b610500816104e9565b82525050565b600060208201905061051b60008301846104f7565b92915050565b6000602082840312156105375761053661040e565b5b60006105458482850161045c565b91505092915050565b61055781610471565b82525050565b6000602082019050610572600083018461054e565b92915050565b61058181610433565b82525050565b600060208201905061059c6000830184610578565b92915050565b6000819050919050565b6105b5816105a2565b82525050565b60006020820190506105d060008301846105ac565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610610826105a2565b915061061b836105a2565b9250828201905080821115610633576106326105d6565b5b92915050565b600082825260208201905092915050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006106a6602683610639565b91506106b18261064a565b604082019050919050565b600060208201905081810360008301526106d581610699565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000610712602083610639565b915061071d826106dc565b602082019050919050565b6000602082019050818103600083015261074181610705565b905091905056fea264697066735822122031f9b27b81fc67bae7dd2a6fb9201a9d579404b34fb985432c1d20c0bc34acfa64736f6c63430008100033",
}

// AccessABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessMetaData.ABI instead.
var AccessABI = AccessMetaData.ABI

// AccessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AccessMetaData.Bin instead.
var AccessBin = AccessMetaData.Bin

// DeployAccess deploys a new Ethereum contract, binding an instance of Access to it.
func DeployAccess(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Access, error) {
	parsed, err := AccessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AccessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Access{AccessCaller: AccessCaller{contract: contract}, AccessTransactor: AccessTransactor{contract: contract}, AccessFilterer: AccessFilterer{contract: contract}}, nil
}

// Access is an auto generated Go binding around an Ethereum contract.
type Access struct {
	AccessCaller     // Read-only binding to the contract
	AccessTransactor // Write-only binding to the contract
	AccessFilterer   // Log filterer for contract events
}

// AccessCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessSession struct {
	Contract     *Access           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessCallerSession struct {
	Contract *AccessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AccessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessTransactorSession struct {
	Contract     *AccessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessRaw struct {
	Contract *Access // Generic contract binding to access the raw methods on
}

// AccessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessCallerRaw struct {
	Contract *AccessCaller // Generic read-only contract binding to access the raw methods on
}

// AccessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessTransactorRaw struct {
	Contract *AccessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccess creates a new instance of Access, bound to a specific deployed contract.
func NewAccess(address common.Address, backend bind.ContractBackend) (*Access, error) {
	contract, err := bindAccess(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Access{AccessCaller: AccessCaller{contract: contract}, AccessTransactor: AccessTransactor{contract: contract}, AccessFilterer: AccessFilterer{contract: contract}}, nil
}

// NewAccessCaller creates a new read-only instance of Access, bound to a specific deployed contract.
func NewAccessCaller(address common.Address, caller bind.ContractCaller) (*AccessCaller, error) {
	contract, err := bindAccess(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessCaller{contract: contract}, nil
}

// NewAccessTransactor creates a new write-only instance of Access, bound to a specific deployed contract.
func NewAccessTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessTransactor, error) {
	contract, err := bindAccess(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessTransactor{contract: contract}, nil
}

// NewAccessFilterer creates a new log filterer instance of Access, bound to a specific deployed contract.
func NewAccessFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessFilterer, error) {
	contract, err := bindAccess(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessFilterer{contract: contract}, nil
}

// bindAccess binds a generic wrapper to an already deployed contract.
func bindAccess(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Access *AccessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Access.Contract.AccessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Access *AccessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Access.Contract.AccessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Access *AccessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Access.Contract.AccessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Access *AccessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Access.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Access *AccessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Access.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Access *AccessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Access.Contract.contract.Transact(opts, method, params...)
}

// Access is a free data retrieval call binding the contract method 0x6fae3d76.
//
// Solidity: function access(address ) view returns(bool)
func (_Access *AccessCaller) Access(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Access.contract.Call(opts, &out, "access", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Access is a free data retrieval call binding the contract method 0x6fae3d76.
//
// Solidity: function access(address ) view returns(bool)
func (_Access *AccessSession) Access(arg0 common.Address) (bool, error) {
	return _Access.Contract.Access(&_Access.CallOpts, arg0)
}

// Access is a free data retrieval call binding the contract method 0x6fae3d76.
//
// Solidity: function access(address ) view returns(bool)
func (_Access *AccessCallerSession) Access(arg0 common.Address) (bool, error) {
	return _Access.Contract.Access(&_Access.CallOpts, arg0)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Access *AccessCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Access.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Access *AccessSession) Nonce() (*big.Int, error) {
	return _Access.Contract.Nonce(&_Access.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Access *AccessCallerSession) Nonce() (*big.Int, error) {
	return _Access.Contract.Nonce(&_Access.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Access *AccessCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Access.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Access *AccessSession) Owner() (common.Address, error) {
	return _Access.Contract.Owner(&_Access.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Access *AccessCallerSession) Owner() (common.Address, error) {
	return _Access.Contract.Owner(&_Access.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Access *AccessCaller) Version(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Access.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Access *AccessSession) Version() (uint16, error) {
	return _Access.Contract.Version(&_Access.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_Access *AccessCallerSession) Version() (uint16, error) {
	return _Access.Contract.Version(&_Access.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Access *AccessTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Access.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Access *AccessSession) RenounceOwnership() (*types.Transaction, error) {
	return _Access.Contract.RenounceOwnership(&_Access.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Access *AccessTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Access.Contract.RenounceOwnership(&_Access.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x35e3b25a.
//
// Solidity: function set(address addr, bool isSet) returns()
func (_Access *AccessTransactor) Set(opts *bind.TransactOpts, addr common.Address, isSet bool) (*types.Transaction, error) {
	return _Access.contract.Transact(opts, "set", addr, isSet)
}

// Set is a paid mutator transaction binding the contract method 0x35e3b25a.
//
// Solidity: function set(address addr, bool isSet) returns()
func (_Access *AccessSession) Set(addr common.Address, isSet bool) (*types.Transaction, error) {
	return _Access.Contract.Set(&_Access.TransactOpts, addr, isSet)
}

// Set is a paid mutator transaction binding the contract method 0x35e3b25a.
//
// Solidity: function set(address addr, bool isSet) returns()
func (_Access *AccessTransactorSession) Set(addr common.Address, isSet bool) (*types.Transaction, error) {
	return _Access.Contract.Set(&_Access.TransactOpts, addr, isSet)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Access *AccessTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Access.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Access *AccessSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Access.Contract.TransferOwnership(&_Access.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Access *AccessTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Access.Contract.TransferOwnership(&_Access.TransactOpts, newOwner)
}

// AccessOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Access contract.
type AccessOwnershipTransferredIterator struct {
	Event *AccessOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccessOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessOwnershipTransferred)
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
		it.Event = new(AccessOwnershipTransferred)
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
func (it *AccessOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessOwnershipTransferred represents a OwnershipTransferred event raised by the Access contract.
type AccessOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Access *AccessFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccessOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Access.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccessOwnershipTransferredIterator{contract: _Access.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Access *AccessFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccessOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Access.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessOwnershipTransferred)
				if err := _Access.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Access *AccessFilterer) ParseOwnershipTransferred(log types.Log) (*AccessOwnershipTransferred, error) {
	event := new(AccessOwnershipTransferred)
	if err := _Access.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
