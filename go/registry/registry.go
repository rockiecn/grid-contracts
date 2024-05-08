// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

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

// RegistryInfo is an auto generated low-level Go binding around an user-defined struct.
type RegistryInfo struct {
	Ip     string
	Domain string
	Port   uint64
	Total  RegistryResources
	Avail  RegistryResources
}

// RegistryResources is an auto generated low-level Go binding around an user-defined struct.
type RegistryResources struct {
	CPU  uint64
	GPU  uint64
	MEM  uint64
	DISK uint64
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"port\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"CPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DISK\",\"type\":\"uint64\"}],\"internalType\":\"structRegistry.Resources\",\"name\":\"total\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"CPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DISK\",\"type\":\"uint64\"}],\"internalType\":\"structRegistry.Resources\",\"name\":\"avail\",\"type\":\"tuple\"}],\"internalType\":\"structRegistry.Info\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"port\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"disk\",\"type\":\"uint64\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"cpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"disk\",\"type\":\"uint64\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506115df806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063262ae1721461004657806326c7e32a14610062578063c2bc2efc1461007e575b600080fd5b610060600480360381019061005b9190610ed8565b6100ae565b005b61007c60048036038101906100779190610fb2565b610515565b005b61009860048036038101906100939190611077565b610943565b6040516100a59190611205565b60405180910390f35b866000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000190816100fc919061143d565b50856000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101908161014b919061143d565b50846000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550836000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550836000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050505050505050565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff168467ffffffffffffffff1611156105925761059161150f565b5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160089054906101000a900467ffffffffffffffff1667ffffffffffffffff168367ffffffffffffffff16111561060f5761060e61150f565b5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff168267ffffffffffffffff16111561068c5761068b61150f565b5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160189054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff1611156107095761070861150f565b5b836000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160008282829054906101000a900467ffffffffffffffff16610770919061156d565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160088282829054906101000a900467ffffffffffffffff166107fd919061156d565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160108282829054906101000a900467ffffffffffffffff1661088a919061156d565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160188282829054906101000a900467ffffffffffffffff16610917919061156d565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050505050565b61094b610ca9565b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820180546109a490611256565b80601f01602080910402602001604051908101604052809291908181526020018280546109d090611256565b8015610a1d5780601f106109f257610100808354040283529160200191610a1d565b820191906000526020600020905b815481529060010190602001808311610a0057829003601f168201915b50505050508152602001600182018054610a3690611256565b80601f0160208091040260200160405190810160405280929190818152602001828054610a6290611256565b8015610aaf5780601f10610a8457610100808354040283529160200191610aaf565b820191906000526020600020905b815481529060010190602001808311610a9257829003601f168201915b505050505081526020016002820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff168152602001600382016040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250508152602001600482016040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050815250509050919050565b6040518060a001604052806060815260200160608152602001600067ffffffffffffffff168152602001610cdb610cee565b8152602001610ce8610cee565b81525090565b6040518060800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610da582610d5c565b810181811067ffffffffffffffff82111715610dc457610dc3610d6d565b5b80604052505050565b6000610dd7610d3e565b9050610de38282610d9c565b919050565b600067ffffffffffffffff821115610e0357610e02610d6d565b5b610e0c82610d5c565b9050602081019050919050565b82818337600083830152505050565b6000610e3b610e3684610de8565b610dcd565b905082815260208101848484011115610e5757610e56610d57565b5b610e62848285610e19565b509392505050565b600082601f830112610e7f57610e7e610d52565b5b8135610e8f848260208601610e28565b91505092915050565b600067ffffffffffffffff82169050919050565b610eb581610e98565b8114610ec057600080fd5b50565b600081359050610ed281610eac565b92915050565b600080600080600080600060e0888a031215610ef757610ef6610d48565b5b600088013567ffffffffffffffff811115610f1557610f14610d4d565b5b610f218a828b01610e6a565b975050602088013567ffffffffffffffff811115610f4257610f41610d4d565b5b610f4e8a828b01610e6a565b9650506040610f5f8a828b01610ec3565b9550506060610f708a828b01610ec3565b9450506080610f818a828b01610ec3565b93505060a0610f928a828b01610ec3565b92505060c0610fa38a828b01610ec3565b91505092959891949750929550565b60008060008060808587031215610fcc57610fcb610d48565b5b6000610fda87828801610ec3565b9450506020610feb87828801610ec3565b9350506040610ffc87828801610ec3565b925050606061100d87828801610ec3565b91505092959194509250565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061104482611019565b9050919050565b61105481611039565b811461105f57600080fd5b50565b6000813590506110718161104b565b92915050565b60006020828403121561108d5761108c610d48565b5b600061109b84828501611062565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156110de5780820151818401526020810190506110c3565b60008484015250505050565b60006110f5826110a4565b6110ff81856110af565b935061110f8185602086016110c0565b61111881610d5c565b840191505092915050565b61112c81610e98565b82525050565b6080820160008201516111486000850182611123565b50602082015161115b6020850182611123565b50604082015161116e6040850182611123565b5060608201516111816060850182611123565b50505050565b60006101608301600083015184820360008601526111a582826110ea565b915050602083015184820360208601526111bf82826110ea565b91505060408301516111d46040860182611123565b5060608301516111e76060860182611132565b5060808301516111fa60e0860182611132565b508091505092915050565b6000602082019050818103600083015261121f8184611187565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061126e57607f821691505b60208210810361128157611280611227565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026112e97fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826112ac565b6112f386836112ac565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b600061133a6113356113308461130b565b611315565b61130b565b9050919050565b6000819050919050565b6113548361131f565b61136861136082611341565b8484546112b9565b825550505050565b600090565b61137d611370565b61138881848461134b565b505050565b5b818110156113ac576113a1600082611375565b60018101905061138e565b5050565b601f8211156113f1576113c281611287565b6113cb8461129c565b810160208510156113da578190505b6113ee6113e68561129c565b83018261138d565b50505b505050565b600082821c905092915050565b6000611414600019846008026113f6565b1980831691505092915050565b600061142d8383611403565b9150826002028217905092915050565b611446826110a4565b67ffffffffffffffff81111561145f5761145e610d6d565b5b6114698254611256565b6114748282856113b0565b600060209050601f8311600181146114a75760008415611495578287015190505b61149f8582611421565b865550611507565b601f1984166114b586611287565b60005b828110156114dd578489015182556001820191506020850194506020810190506114b8565b868310156114fa57848901516114f6601f891682611403565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061157882610e98565b915061158383610e98565b9250828203905067ffffffffffffffff8111156115a3576115a261153e565b5b9291505056fea2646970667358221220cde3dc18076b3b5a46a9ede2284cf10e7b039df8e76bdd6f5b6b41f962e37a2964736f6c63430008100033",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// RegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RegistryMetaData.Bin instead.
var RegistryBin = RegistryMetaData.Bin

// DeployRegistry deploys a new Ethereum contract, binding an instance of Registry to it.
func DeployRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Registry, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address a) view returns((string,string,uint64,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64)))
func (_Registry *RegistryCaller) Get(opts *bind.CallOpts, a common.Address) (RegistryInfo, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "get", a)

	if err != nil {
		return *new(RegistryInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(RegistryInfo)).(*RegistryInfo)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address a) view returns((string,string,uint64,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64)))
func (_Registry *RegistrySession) Get(a common.Address) (RegistryInfo, error) {
	return _Registry.Contract.Get(&_Registry.CallOpts, a)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address a) view returns((string,string,uint64,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64)))
func (_Registry *RegistryCallerSession) Get(a common.Address) (RegistryInfo, error) {
	return _Registry.Contract.Get(&_Registry.CallOpts, a)
}

// Set is a paid mutator transaction binding the contract method 0x262ae172.
//
// Solidity: function set(string ip, string domain, uint64 port, uint64 cpu, uint64 gpu, uint64 mem, uint64 disk) returns()
func (_Registry *RegistryTransactor) Set(opts *bind.TransactOpts, ip string, domain string, port uint64, cpu uint64, gpu uint64, mem uint64, disk uint64) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "set", ip, domain, port, cpu, gpu, mem, disk)
}

// Set is a paid mutator transaction binding the contract method 0x262ae172.
//
// Solidity: function set(string ip, string domain, uint64 port, uint64 cpu, uint64 gpu, uint64 mem, uint64 disk) returns()
func (_Registry *RegistrySession) Set(ip string, domain string, port uint64, cpu uint64, gpu uint64, mem uint64, disk uint64) (*types.Transaction, error) {
	return _Registry.Contract.Set(&_Registry.TransactOpts, ip, domain, port, cpu, gpu, mem, disk)
}

// Set is a paid mutator transaction binding the contract method 0x262ae172.
//
// Solidity: function set(string ip, string domain, uint64 port, uint64 cpu, uint64 gpu, uint64 mem, uint64 disk) returns()
func (_Registry *RegistryTransactorSession) Set(ip string, domain string, port uint64, cpu uint64, gpu uint64, mem uint64, disk uint64) (*types.Transaction, error) {
	return _Registry.Contract.Set(&_Registry.TransactOpts, ip, domain, port, cpu, gpu, mem, disk)
}

// Update is a paid mutator transaction binding the contract method 0x26c7e32a.
//
// Solidity: function update(uint64 cpu, uint64 gpu, uint64 mem, uint64 disk) returns()
func (_Registry *RegistryTransactor) Update(opts *bind.TransactOpts, cpu uint64, gpu uint64, mem uint64, disk uint64) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "update", cpu, gpu, mem, disk)
}

// Update is a paid mutator transaction binding the contract method 0x26c7e32a.
//
// Solidity: function update(uint64 cpu, uint64 gpu, uint64 mem, uint64 disk) returns()
func (_Registry *RegistrySession) Update(cpu uint64, gpu uint64, mem uint64, disk uint64) (*types.Transaction, error) {
	return _Registry.Contract.Update(&_Registry.TransactOpts, cpu, gpu, mem, disk)
}

// Update is a paid mutator transaction binding the contract method 0x26c7e32a.
//
// Solidity: function update(uint64 cpu, uint64 gpu, uint64 mem, uint64 disk) returns()
func (_Registry *RegistryTransactorSession) Update(cpu uint64, gpu uint64, mem uint64, disk uint64) (*types.Transaction, error) {
	return _Registry.Contract.Update(&_Registry.TransactOpts, cpu, gpu, mem, disk)
}
