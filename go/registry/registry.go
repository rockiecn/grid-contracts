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
	Port   string
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"CPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DISK\",\"type\":\"uint64\"}],\"internalType\":\"structRegistry.Resources\",\"name\":\"total\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"CPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DISK\",\"type\":\"uint64\"}],\"internalType\":\"structRegistry.Resources\",\"name\":\"avail\",\"type\":\"tuple\"}],\"internalType\":\"structRegistry.Info\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"cpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"disk\",\"type\":\"uint64\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"cpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"disk\",\"type\":\"uint64\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061163f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80631ef41bbd1461004657806326c7e32a14610062578063c2bc2efc1461007e575b600080fd5b610060600480360381019061005b9190610f15565b6100ae565b005b61007c6004803603810190610077919061100b565b6104fc565b005b610098600480360381019061009391906110d0565b61092a565b6040516100a59190611265565b60405180910390f35b866000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000190816100fc919061149d565b50856000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101908161014b919061149d565b50846000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201908161019a919061149d565b50836000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550836000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050505050505050565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff168467ffffffffffffffff1611156105795761057861156f565b5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160089054906101000a900467ffffffffffffffff1667ffffffffffffffff168367ffffffffffffffff1611156105f6576105f561156f565b5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff168267ffffffffffffffff1611156106735761067261156f565b5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160189054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff1611156106f0576106ef61156f565b5b836000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160008282829054906101000a900467ffffffffffffffff1661075791906115cd565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160088282829054906101000a900467ffffffffffffffff166107e491906115cd565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160108282829054906101000a900467ffffffffffffffff1661087191906115cd565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160188282829054906101000a900467ffffffffffffffff166108fe91906115cd565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050505050565b610932610cf0565b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060a001604052908160008201805461098b906112b6565b80601f01602080910402602001604051908101604052809291908181526020018280546109b7906112b6565b8015610a045780601f106109d957610100808354040283529160200191610a04565b820191906000526020600020905b8154815290600101906020018083116109e757829003601f168201915b50505050508152602001600182018054610a1d906112b6565b80601f0160208091040260200160405190810160405280929190818152602001828054610a49906112b6565b8015610a965780601f10610a6b57610100808354040283529160200191610a96565b820191906000526020600020905b815481529060010190602001808311610a7957829003601f168201915b50505050508152602001600282018054610aaf906112b6565b80601f0160208091040260200160405190810160405280929190818152602001828054610adb906112b6565b8015610b285780601f10610afd57610100808354040283529160200191610b28565b820191906000526020600020905b815481529060010190602001808311610b0b57829003601f168201915b50505050508152602001600382016040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250508152602001600482016040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050815250509050919050565b6040518060a00160405280606081526020016060815260200160608152602001610d18610d2b565b8152602001610d25610d2b565b81525090565b6040518060800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610de282610d99565b810181811067ffffffffffffffff82111715610e0157610e00610daa565b5b80604052505050565b6000610e14610d7b565b9050610e208282610dd9565b919050565b600067ffffffffffffffff821115610e4057610e3f610daa565b5b610e4982610d99565b9050602081019050919050565b82818337600083830152505050565b6000610e78610e7384610e25565b610e0a565b905082815260208101848484011115610e9457610e93610d94565b5b610e9f848285610e56565b509392505050565b600082601f830112610ebc57610ebb610d8f565b5b8135610ecc848260208601610e65565b91505092915050565b600067ffffffffffffffff82169050919050565b610ef281610ed5565b8114610efd57600080fd5b50565b600081359050610f0f81610ee9565b92915050565b600080600080600080600060e0888a031215610f3457610f33610d85565b5b600088013567ffffffffffffffff811115610f5257610f51610d8a565b5b610f5e8a828b01610ea7565b975050602088013567ffffffffffffffff811115610f7f57610f7e610d8a565b5b610f8b8a828b01610ea7565b965050604088013567ffffffffffffffff811115610fac57610fab610d8a565b5b610fb88a828b01610ea7565b9550506060610fc98a828b01610f00565b9450506080610fda8a828b01610f00565b93505060a0610feb8a828b01610f00565b92505060c0610ffc8a828b01610f00565b91505092959891949750929550565b6000806000806080858703121561102557611024610d85565b5b600061103387828801610f00565b945050602061104487828801610f00565b935050604061105587828801610f00565b925050606061106687828801610f00565b91505092959194509250565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061109d82611072565b9050919050565b6110ad81611092565b81146110b857600080fd5b50565b6000813590506110ca816110a4565b92915050565b6000602082840312156110e6576110e5610d85565b5b60006110f4848285016110bb565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561113757808201518184015260208101905061111c565b60008484015250505050565b600061114e826110fd565b6111588185611108565b9350611168818560208601611119565b61117181610d99565b840191505092915050565b61118581610ed5565b82525050565b6080820160008201516111a1600085018261117c565b5060208201516111b4602085018261117c565b5060408201516111c7604085018261117c565b5060608201516111da606085018261117c565b50505050565b60006101608301600083015184820360008601526111fe8282611143565b915050602083015184820360208601526112188282611143565b915050604083015184820360408601526112328282611143565b9150506060830151611247606086018261118b565b50608083015161125a60e086018261118b565b508091505092915050565b6000602082019050818103600083015261127f81846111e0565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806112ce57607f821691505b6020821081036112e1576112e0611287565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026113497fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261130c565b611353868361130c565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b600061139a6113956113908461136b565b611375565b61136b565b9050919050565b6000819050919050565b6113b48361137f565b6113c86113c0826113a1565b848454611319565b825550505050565b600090565b6113dd6113d0565b6113e88184846113ab565b505050565b5b8181101561140c576114016000826113d5565b6001810190506113ee565b5050565b601f82111561145157611422816112e7565b61142b846112fc565b8101602085101561143a578190505b61144e611446856112fc565b8301826113ed565b50505b505050565b600082821c905092915050565b600061147460001984600802611456565b1980831691505092915050565b600061148d8383611463565b9150826002028217905092915050565b6114a6826110fd565b67ffffffffffffffff8111156114bf576114be610daa565b5b6114c982546112b6565b6114d4828285611410565b600060209050601f83116001811461150757600084156114f5578287015190505b6114ff8582611481565b865550611567565b601f198416611515866112e7565b60005b8281101561153d57848901518255600182019150602085019450602081019050611518565b8683101561155a5784890151611556601f891682611463565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006115d882610ed5565b91506115e383610ed5565b9250828203905067ffffffffffffffff8111156116035761160261159e565b5b9291505056fea26469706673582212209287376e2bb309df926532f5a5bb2f1ca0fab313c1c6fd183f9af74ea1dadf0964736f6c63430008100033",
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
// Solidity: function get(address a) view returns((string,string,string,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64)))
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
// Solidity: function get(address a) view returns((string,string,string,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64)))
func (_Registry *RegistrySession) Get(a common.Address) (RegistryInfo, error) {
	return _Registry.Contract.Get(&_Registry.CallOpts, a)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address a) view returns((string,string,string,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64)))
func (_Registry *RegistryCallerSession) Get(a common.Address) (RegistryInfo, error) {
	return _Registry.Contract.Get(&_Registry.CallOpts, a)
}

// Set is a paid mutator transaction binding the contract method 0x1ef41bbd.
//
// Solidity: function set(string ip, string domain, string port, uint64 cpu, uint64 gpu, uint64 mem, uint64 disk) returns()
func (_Registry *RegistryTransactor) Set(opts *bind.TransactOpts, ip string, domain string, port string, cpu uint64, gpu uint64, mem uint64, disk uint64) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "set", ip, domain, port, cpu, gpu, mem, disk)
}

// Set is a paid mutator transaction binding the contract method 0x1ef41bbd.
//
// Solidity: function set(string ip, string domain, string port, uint64 cpu, uint64 gpu, uint64 mem, uint64 disk) returns()
func (_Registry *RegistrySession) Set(ip string, domain string, port string, cpu uint64, gpu uint64, mem uint64, disk uint64) (*types.Transaction, error) {
	return _Registry.Contract.Set(&_Registry.TransactOpts, ip, domain, port, cpu, gpu, mem, disk)
}

// Set is a paid mutator transaction binding the contract method 0x1ef41bbd.
//
// Solidity: function set(string ip, string domain, string port, uint64 cpu, uint64 gpu, uint64 mem, uint64 disk) returns()
func (_Registry *RegistryTransactorSession) Set(ip string, domain string, port string, cpu uint64, gpu uint64, mem uint64, disk uint64) (*types.Transaction, error) {
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
