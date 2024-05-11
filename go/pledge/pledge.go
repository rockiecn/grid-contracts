// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pledge

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

// PledgeMetaData contains all meta data concerning the Pledge contract.
var PledgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPledge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"pledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"setRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unpledge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateInterest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawInterest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061002d61002261003e60201b60201c565b61004660201b60201c565b640165a0bc0060018190555061010a565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6110e4806101196000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80638da5cb5b116100715780638da5cb5b1461012a578063c483997514610148578063c8f5acb614610164578063d148279114610182578063e5a4bed31461018c578063f2fde38b146101aa576100a9565b806334fcf437146100ae5780635b730e69146100ca578063679aefce146100e6578063715018a614610104578063743b34521461010e575b600080fd5b6100c860048036038101906100c39190610a7b565b6101c6565b005b6100e460048036038101906100df9190610b06565b6101f5565b005b6100ee61035f565b6040516100fb9190610b55565b60405180910390f35b61010c610369565b005b61012860048036038101906101239190610b06565b61037d565b005b6101326104a7565b60405161013f9190610b7f565b60405180910390f35b610162600480360381019061015d9190610b06565b6104d0565b005b61016c61067d565b6040516101799190610b55565b60405180910390f35b61018a6106c7565b005b6101946106d2565b6040516101a19190610b55565b60405180910390f35b6101c460048036038101906101bf9190610b9a565b61071c565b005b6101ce61079f565b6000811180156101e25750648bb2c9700081105b6101eb57600080fd5b8060018190555050565b6101fe3361081d565b80600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101541015610283576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027a90610c24565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b81526004016102be929190610c44565b6020604051808303816000875af11580156102dd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103019190610ca5565b5080600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160008282546103549190610d01565b925050819055505050565b6000600154905090565b61037161079f565b61037b6000610974565b565b600081116103c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b790610da7565b60405180910390fd5b6103c93361081d565b8173ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b815260040161040693929190610dc7565b6020604051808303816000875af1158015610425573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104499190610ca5565b5080600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001600082825461049c9190610dfe565b925050819055505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008111610513576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050a90610da7565b60405180910390fd5b80600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600001541015610598576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058f90610ea4565b60405180910390fd5b6105a13361081d565b8173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b81526004016105dc929190610c44565b6020604051808303816000875af11580156105fb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061061f9190610ca5565b5080600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160008282546106729190610d01565b925050819055505050565b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000154905090565b6106d03361081d565b565b6000600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010154905090565b61072461079f565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610793576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078a90610f36565b60405180910390fd5b61079c81610974565b50565b6107a7610a38565b73ffffffffffffffffffffffffffffffffffffffff166107c56104a7565b73ffffffffffffffffffffffffffffffffffffffff161461081b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081290610fa2565b60405180910390fd5b565b60004290506000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002015490508082116108ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108a290611034565b60405180910390fd5b600081836108b99190610d01565b90506000600154826108cb9190611054565b905080600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101600082825461091f9190610dfe565b9250508190555083600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201819055505050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b600080fd5b6000819050919050565b610a5881610a45565b8114610a6357600080fd5b50565b600081359050610a7581610a4f565b92915050565b600060208284031215610a9157610a90610a40565b5b6000610a9f84828501610a66565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ad382610aa8565b9050919050565b610ae381610ac8565b8114610aee57600080fd5b50565b600081359050610b0081610ada565b92915050565b60008060408385031215610b1d57610b1c610a40565b5b6000610b2b85828601610af1565b9250506020610b3c85828601610a66565b9150509250929050565b610b4f81610a45565b82525050565b6000602082019050610b6a6000830184610b46565b92915050565b610b7981610ac8565b82525050565b6000602082019050610b946000830184610b70565b92915050565b600060208284031215610bb057610baf610a40565b5b6000610bbe84828501610af1565b91505092915050565b600082825260208201905092915050565b7f6e6f7420656e6f75676820696e7465726573742072656d61696e656400000000600082015250565b6000610c0e601c83610bc7565b9150610c1982610bd8565b602082019050919050565b60006020820190508181036000830152610c3d81610c01565b9050919050565b6000604082019050610c596000830185610b70565b610c666020830184610b46565b9392505050565b60008115159050919050565b610c8281610c6d565b8114610c8d57600080fd5b50565b600081519050610c9f81610c79565b92915050565b600060208284031215610cbb57610cba610a40565b5b6000610cc984828501610c90565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610d0c82610a45565b9150610d1783610a45565b9250828203905081811115610d2f57610d2e610cd2565b5b92915050565b7f616d6f756e74206f6620706c656467652073686f756c64206c6172676572207460008201527f68616e2030000000000000000000000000000000000000000000000000000000602082015250565b6000610d91602583610bc7565b9150610d9c82610d35565b604082019050919050565b60006020820190508181036000830152610dc081610d84565b9050919050565b6000606082019050610ddc6000830186610b70565b610de96020830185610b70565b610df66040830184610b46565b949350505050565b6000610e0982610a45565b9150610e1483610a45565b9250828201905080821115610e2c57610e2b610cd2565b5b92915050565b7f7468652072656d61696e656420706c65646765206e6f7420656e6f756768206660008201527f6f7220756e706c65646765000000000000000000000000000000000000000000602082015250565b6000610e8e602b83610bc7565b9150610e9982610e32565b604082019050919050565b60006020820190508181036000830152610ebd81610e81565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000610f20602683610bc7565b9150610f2b82610ec4565b604082019050919050565b60006020820190508181036000830152610f4f81610f13565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000610f8c602083610bc7565b9150610f9782610f56565b602082019050919050565b60006020820190508181036000830152610fbb81610f7f565b9050919050565b7f6e6f772074696d65206e6f74206c61746572207468616e206c6173742074696d60008201527f6500000000000000000000000000000000000000000000000000000000000000602082015250565b600061101e602183610bc7565b915061102982610fc2565b604082019050919050565b6000602082019050818103600083015261104d81611011565b9050919050565b600061105f82610a45565b915061106a83610a45565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156110a3576110a2610cd2565b5b82820290509291505056fea26469706673582212205ab70d9d916d1b26db6e025e16021b1c3e765416c312959f73db994aa65b82e264736f6c63430008100033",
}

// PledgeABI is the input ABI used to generate the binding from.
// Deprecated: Use PledgeMetaData.ABI instead.
var PledgeABI = PledgeMetaData.ABI

// PledgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PledgeMetaData.Bin instead.
var PledgeBin = PledgeMetaData.Bin

// DeployPledge deploys a new Ethereum contract, binding an instance of Pledge to it.
func DeployPledge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pledge, error) {
	parsed, err := PledgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PledgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pledge{PledgeCaller: PledgeCaller{contract: contract}, PledgeTransactor: PledgeTransactor{contract: contract}, PledgeFilterer: PledgeFilterer{contract: contract}}, nil
}

// Pledge is an auto generated Go binding around an Ethereum contract.
type Pledge struct {
	PledgeCaller     // Read-only binding to the contract
	PledgeTransactor // Write-only binding to the contract
	PledgeFilterer   // Log filterer for contract events
}

// PledgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PledgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PledgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PledgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PledgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PledgeSession struct {
	Contract     *Pledge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PledgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PledgeCallerSession struct {
	Contract *PledgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PledgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PledgeTransactorSession struct {
	Contract     *PledgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PledgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PledgeRaw struct {
	Contract *Pledge // Generic contract binding to access the raw methods on
}

// PledgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PledgeCallerRaw struct {
	Contract *PledgeCaller // Generic read-only contract binding to access the raw methods on
}

// PledgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PledgeTransactorRaw struct {
	Contract *PledgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPledge creates a new instance of Pledge, bound to a specific deployed contract.
func NewPledge(address common.Address, backend bind.ContractBackend) (*Pledge, error) {
	contract, err := bindPledge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pledge{PledgeCaller: PledgeCaller{contract: contract}, PledgeTransactor: PledgeTransactor{contract: contract}, PledgeFilterer: PledgeFilterer{contract: contract}}, nil
}

// NewPledgeCaller creates a new read-only instance of Pledge, bound to a specific deployed contract.
func NewPledgeCaller(address common.Address, caller bind.ContractCaller) (*PledgeCaller, error) {
	contract, err := bindPledge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PledgeCaller{contract: contract}, nil
}

// NewPledgeTransactor creates a new write-only instance of Pledge, bound to a specific deployed contract.
func NewPledgeTransactor(address common.Address, transactor bind.ContractTransactor) (*PledgeTransactor, error) {
	contract, err := bindPledge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PledgeTransactor{contract: contract}, nil
}

// NewPledgeFilterer creates a new log filterer instance of Pledge, bound to a specific deployed contract.
func NewPledgeFilterer(address common.Address, filterer bind.ContractFilterer) (*PledgeFilterer, error) {
	contract, err := bindPledge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PledgeFilterer{contract: contract}, nil
}

// bindPledge binds a generic wrapper to an already deployed contract.
func bindPledge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PledgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pledge *PledgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pledge.Contract.PledgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pledge *PledgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge.Contract.PledgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pledge *PledgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pledge.Contract.PledgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pledge *PledgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pledge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pledge *PledgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pledge *PledgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pledge.Contract.contract.Transact(opts, method, params...)
}

// GetInterest is a free data retrieval call binding the contract method 0xe5a4bed3.
//
// Solidity: function getInterest() view returns(uint256)
func (_Pledge *PledgeCaller) GetInterest(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "getInterest")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInterest is a free data retrieval call binding the contract method 0xe5a4bed3.
//
// Solidity: function getInterest() view returns(uint256)
func (_Pledge *PledgeSession) GetInterest() (*big.Int, error) {
	return _Pledge.Contract.GetInterest(&_Pledge.CallOpts)
}

// GetInterest is a free data retrieval call binding the contract method 0xe5a4bed3.
//
// Solidity: function getInterest() view returns(uint256)
func (_Pledge *PledgeCallerSession) GetInterest() (*big.Int, error) {
	return _Pledge.Contract.GetInterest(&_Pledge.CallOpts)
}

// GetPledge is a free data retrieval call binding the contract method 0xc8f5acb6.
//
// Solidity: function getPledge() view returns(uint256)
func (_Pledge *PledgeCaller) GetPledge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "getPledge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPledge is a free data retrieval call binding the contract method 0xc8f5acb6.
//
// Solidity: function getPledge() view returns(uint256)
func (_Pledge *PledgeSession) GetPledge() (*big.Int, error) {
	return _Pledge.Contract.GetPledge(&_Pledge.CallOpts)
}

// GetPledge is a free data retrieval call binding the contract method 0xc8f5acb6.
//
// Solidity: function getPledge() view returns(uint256)
func (_Pledge *PledgeCallerSession) GetPledge() (*big.Int, error) {
	return _Pledge.Contract.GetPledge(&_Pledge.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_Pledge *PledgeCaller) GetRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "getRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_Pledge *PledgeSession) GetRate() (*big.Int, error) {
	return _Pledge.Contract.GetRate(&_Pledge.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0x679aefce.
//
// Solidity: function getRate() view returns(uint256)
func (_Pledge *PledgeCallerSession) GetRate() (*big.Int, error) {
	return _Pledge.Contract.GetRate(&_Pledge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pledge *PledgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pledge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pledge *PledgeSession) Owner() (common.Address, error) {
	return _Pledge.Contract.Owner(&_Pledge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pledge *PledgeCallerSession) Owner() (common.Address, error) {
	return _Pledge.Contract.Owner(&_Pledge.CallOpts)
}

// Pledge is a paid mutator transaction binding the contract method 0x743b3452.
//
// Solidity: function pledge(address tokenAddr, uint256 amount) returns()
func (_Pledge *PledgeTransactor) Pledge(opts *bind.TransactOpts, tokenAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "pledge", tokenAddr, amount)
}

// Pledge is a paid mutator transaction binding the contract method 0x743b3452.
//
// Solidity: function pledge(address tokenAddr, uint256 amount) returns()
func (_Pledge *PledgeSession) Pledge(tokenAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.Pledge(&_Pledge.TransactOpts, tokenAddr, amount)
}

// Pledge is a paid mutator transaction binding the contract method 0x743b3452.
//
// Solidity: function pledge(address tokenAddr, uint256 amount) returns()
func (_Pledge *PledgeTransactorSession) Pledge(tokenAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.Pledge(&_Pledge.TransactOpts, tokenAddr, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pledge *PledgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pledge *PledgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pledge.Contract.RenounceOwnership(&_Pledge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pledge *PledgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pledge.Contract.RenounceOwnership(&_Pledge.TransactOpts)
}

// SetRate is a paid mutator transaction binding the contract method 0x34fcf437.
//
// Solidity: function setRate(uint256 _rate) returns()
func (_Pledge *PledgeTransactor) SetRate(opts *bind.TransactOpts, _rate *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "setRate", _rate)
}

// SetRate is a paid mutator transaction binding the contract method 0x34fcf437.
//
// Solidity: function setRate(uint256 _rate) returns()
func (_Pledge *PledgeSession) SetRate(_rate *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.SetRate(&_Pledge.TransactOpts, _rate)
}

// SetRate is a paid mutator transaction binding the contract method 0x34fcf437.
//
// Solidity: function setRate(uint256 _rate) returns()
func (_Pledge *PledgeTransactorSession) SetRate(_rate *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.SetRate(&_Pledge.TransactOpts, _rate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pledge *PledgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pledge *PledgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pledge.Contract.TransferOwnership(&_Pledge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pledge *PledgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pledge.Contract.TransferOwnership(&_Pledge.TransactOpts, newOwner)
}

// Unpledge is a paid mutator transaction binding the contract method 0xc4839975.
//
// Solidity: function unpledge(address tokenAddr, uint256 amount) returns()
func (_Pledge *PledgeTransactor) Unpledge(opts *bind.TransactOpts, tokenAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "unpledge", tokenAddr, amount)
}

// Unpledge is a paid mutator transaction binding the contract method 0xc4839975.
//
// Solidity: function unpledge(address tokenAddr, uint256 amount) returns()
func (_Pledge *PledgeSession) Unpledge(tokenAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.Unpledge(&_Pledge.TransactOpts, tokenAddr, amount)
}

// Unpledge is a paid mutator transaction binding the contract method 0xc4839975.
//
// Solidity: function unpledge(address tokenAddr, uint256 amount) returns()
func (_Pledge *PledgeTransactorSession) Unpledge(tokenAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.Unpledge(&_Pledge.TransactOpts, tokenAddr, amount)
}

// UpdateInterest is a paid mutator transaction binding the contract method 0xd1482791.
//
// Solidity: function updateInterest() returns()
func (_Pledge *PledgeTransactor) UpdateInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "updateInterest")
}

// UpdateInterest is a paid mutator transaction binding the contract method 0xd1482791.
//
// Solidity: function updateInterest() returns()
func (_Pledge *PledgeSession) UpdateInterest() (*types.Transaction, error) {
	return _Pledge.Contract.UpdateInterest(&_Pledge.TransactOpts)
}

// UpdateInterest is a paid mutator transaction binding the contract method 0xd1482791.
//
// Solidity: function updateInterest() returns()
func (_Pledge *PledgeTransactorSession) UpdateInterest() (*types.Transaction, error) {
	return _Pledge.Contract.UpdateInterest(&_Pledge.TransactOpts)
}

// WithdrawInterest is a paid mutator transaction binding the contract method 0x5b730e69.
//
// Solidity: function withdrawInterest(address tokenAddr, uint256 amount) returns()
func (_Pledge *PledgeTransactor) WithdrawInterest(opts *bind.TransactOpts, tokenAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pledge.contract.Transact(opts, "withdrawInterest", tokenAddr, amount)
}

// WithdrawInterest is a paid mutator transaction binding the contract method 0x5b730e69.
//
// Solidity: function withdrawInterest(address tokenAddr, uint256 amount) returns()
func (_Pledge *PledgeSession) WithdrawInterest(tokenAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.WithdrawInterest(&_Pledge.TransactOpts, tokenAddr, amount)
}

// WithdrawInterest is a paid mutator transaction binding the contract method 0x5b730e69.
//
// Solidity: function withdrawInterest(address tokenAddr, uint256 amount) returns()
func (_Pledge *PledgeTransactorSession) WithdrawInterest(tokenAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Pledge.Contract.WithdrawInterest(&_Pledge.TransactOpts, tokenAddr, amount)
}

// PledgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pledge contract.
type PledgeOwnershipTransferredIterator struct {
	Event *PledgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PledgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PledgeOwnershipTransferred)
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
		it.Event = new(PledgeOwnershipTransferred)
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
func (it *PledgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PledgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PledgeOwnershipTransferred represents a OwnershipTransferred event raised by the Pledge contract.
type PledgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pledge *PledgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PledgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pledge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PledgeOwnershipTransferredIterator{contract: _Pledge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pledge *PledgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PledgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pledge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PledgeOwnershipTransferred)
				if err := _Pledge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pledge *PledgeFilterer) ParseOwnershipTransferred(log types.Log) (*PledgeOwnershipTransferred, error) {
	event := new(PledgeOwnershipTransferred)
	if err := _Pledge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
