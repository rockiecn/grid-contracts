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
	Addr   common.Address
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"CPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DISK\",\"type\":\"uint64\"}],\"internalType\":\"structRegistry.Resources\",\"name\":\"total\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"CPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DISK\",\"type\":\"uint64\"}],\"internalType\":\"structRegistry.Resources\",\"name\":\"avail\",\"type\":\"tuple\"}],\"internalType\":\"structRegistry.Info\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"cpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"disk\",\"type\":\"uint64\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"cpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"disk\",\"type\":\"uint64\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611755806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80631ef41bbd1461004657806326c7e32a14610062578063c2bc2efc1461007e575b600080fd5b610060600480360381019061005b9190611008565b6100ae565b005b61007c600480360381019061007791906110fe565b61057c565b005b610098600480360381019061009391906111c3565b6109aa565b6040516100a5919061137b565b60405180910390f35b336000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550866000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600101908161017c91906115b3565b50856000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020190816101cb91906115b3565b50846000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301908161021a91906115b3565b50836000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550836000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050505050505050565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff168467ffffffffffffffff1611156105f9576105f8611685565b5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160089054906101000a900467ffffffffffffffff1667ffffffffffffffff168367ffffffffffffffff16111561067657610675611685565b5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff168267ffffffffffffffff1611156106f3576106f2611685565b5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160189054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff1611156107705761076f611685565b5b836000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160008282829054906101000a900467ffffffffffffffff166107d791906116e3565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160088282829054906101000a900467ffffffffffffffff1661086491906116e3565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160108282829054906101000a900467ffffffffffffffff166108f191906116e3565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160188282829054906101000a900467ffffffffffffffff1661097e91906116e3565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050505050565b6109b2610dc6565b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054610a61906113cc565b80601f0160208091040260200160405190810160405280929190818152602001828054610a8d906113cc565b8015610ada5780601f10610aaf57610100808354040283529160200191610ada565b820191906000526020600020905b815481529060010190602001808311610abd57829003601f168201915b50505050508152602001600282018054610af3906113cc565b80601f0160208091040260200160405190810160405280929190818152602001828054610b1f906113cc565b8015610b6c5780601f10610b4157610100808354040283529160200191610b6c565b820191906000526020600020905b815481529060010190602001808311610b4f57829003601f168201915b50505050508152602001600382018054610b85906113cc565b80601f0160208091040260200160405190810160405280929190818152602001828054610bb1906113cc565b8015610bfe5780601f10610bd357610100808354040283529160200191610bfe565b820191906000526020600020905b815481529060010190602001808311610be157829003601f168201915b50505050508152602001600482016040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250508152602001600582016040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050815250509050919050565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160608152602001610e0b610e1e565b8152602001610e18610e1e565b81525090565b6040518060800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610ed582610e8c565b810181811067ffffffffffffffff82111715610ef457610ef3610e9d565b5b80604052505050565b6000610f07610e6e565b9050610f138282610ecc565b919050565b600067ffffffffffffffff821115610f3357610f32610e9d565b5b610f3c82610e8c565b9050602081019050919050565b82818337600083830152505050565b6000610f6b610f6684610f18565b610efd565b905082815260208101848484011115610f8757610f86610e87565b5b610f92848285610f49565b509392505050565b600082601f830112610faf57610fae610e82565b5b8135610fbf848260208601610f58565b91505092915050565b600067ffffffffffffffff82169050919050565b610fe581610fc8565b8114610ff057600080fd5b50565b60008135905061100281610fdc565b92915050565b600080600080600080600060e0888a03121561102757611026610e78565b5b600088013567ffffffffffffffff81111561104557611044610e7d565b5b6110518a828b01610f9a565b975050602088013567ffffffffffffffff81111561107257611071610e7d565b5b61107e8a828b01610f9a565b965050604088013567ffffffffffffffff81111561109f5761109e610e7d565b5b6110ab8a828b01610f9a565b95505060606110bc8a828b01610ff3565b94505060806110cd8a828b01610ff3565b93505060a06110de8a828b01610ff3565b92505060c06110ef8a828b01610ff3565b91505092959891949750929550565b6000806000806080858703121561111857611117610e78565b5b600061112687828801610ff3565b945050602061113787828801610ff3565b935050604061114887828801610ff3565b925050606061115987828801610ff3565b91505092959194509250565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061119082611165565b9050919050565b6111a081611185565b81146111ab57600080fd5b50565b6000813590506111bd81611197565b92915050565b6000602082840312156111d9576111d8610e78565b5b60006111e7848285016111ae565b91505092915050565b6111f981611185565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561123957808201518184015260208101905061121e565b60008484015250505050565b6000611250826111ff565b61125a818561120a565b935061126a81856020860161121b565b61127381610e8c565b840191505092915050565b61128781610fc8565b82525050565b6080820160008201516112a3600085018261127e565b5060208201516112b6602085018261127e565b5060408201516112c9604085018261127e565b5060608201516112dc606085018261127e565b50505050565b6000610180830160008301516112fb60008601826111f0565b50602083015184820360208601526113138282611245565b9150506040830151848203604086015261132d8282611245565b915050606083015184820360608601526113478282611245565b915050608083015161135c608086018261128d565b5060a083015161137061010086018261128d565b508091505092915050565b6000602082019050818103600083015261139581846112e2565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806113e457607f821691505b6020821081036113f7576113f661139d565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261145f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611422565b6114698683611422565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006114b06114ab6114a684611481565b61148b565b611481565b9050919050565b6000819050919050565b6114ca83611495565b6114de6114d6826114b7565b84845461142f565b825550505050565b600090565b6114f36114e6565b6114fe8184846114c1565b505050565b5b81811015611522576115176000826114eb565b600181019050611504565b5050565b601f82111561156757611538816113fd565b61154184611412565b81016020851015611550578190505b61156461155c85611412565b830182611503565b50505b505050565b600082821c905092915050565b600061158a6000198460080261156c565b1980831691505092915050565b60006115a38383611579565b9150826002028217905092915050565b6115bc826111ff565b67ffffffffffffffff8111156115d5576115d4610e9d565b5b6115df82546113cc565b6115ea828285611526565b600060209050601f83116001811461161d576000841561160b578287015190505b6116158582611597565b86555061167d565b601f19841661162b866113fd565b60005b828110156116535784890151825560018201915060208501945060208101905061162e565b86831015611670578489015161166c601f891682611579565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006116ee82610fc8565b91506116f983610fc8565b9250828203905067ffffffffffffffff811115611719576117186116b4565b5b9291505056fea264697066735822122025b83460ed29ed91c1514dd3cdc78060321bd368460e104df1095a0488300a7f64736f6c63430008100033",
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
// Solidity: function get(address a) view returns((address,string,string,string,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64)))
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
// Solidity: function get(address a) view returns((address,string,string,string,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64)))
func (_Registry *RegistrySession) Get(a common.Address) (RegistryInfo, error) {
	return _Registry.Contract.Get(&_Registry.CallOpts, a)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address a) view returns((address,string,string,string,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64)))
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
