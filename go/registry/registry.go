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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"CPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DISK\",\"type\":\"uint64\"}],\"internalType\":\"structRegistry.Resources\",\"name\":\"total\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"CPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"GPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"MEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"DISK\",\"type\":\"uint64\"}],\"internalType\":\"structRegistry.Resources\",\"name\":\"avail\",\"type\":\"tuple\"}],\"internalType\":\"structRegistry.Info\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getKeys\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"cpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"disk\",\"type\":\"uint64\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"cpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"disk\",\"type\":\"uint64\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611957806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80631bb8355e1461005c5780632150c5181461007857806326c7e32a14610096578063c2bc2efc146100b2578063cdf3bc6f146100e2575b600080fd5b61007660048036038101906100719190611139565b6100ec565b005b61008061061d565b60405161008d919061131f565b60405180910390f35b6100b060048036038101906100ab9190611341565b6106ab565b005b6100cc60048036038101906100c791906113d4565b610ad9565b6040516100d9919061157d565b60405180910390f35b6100ea610ef5565b005b336000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550866000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010190816101ba91906117b5565b50856000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600201908161020991906117b5565b50846000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600301908161025891906117b5565b50836000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550836000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050505050565b606060018054806020026020016040519081016040528092919081815260200182805480156106a157602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610657575b5050505050905090565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff168467ffffffffffffffff16111561072857610727611887565b5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160089054906101000a900467ffffffffffffffff1667ffffffffffffffff168367ffffffffffffffff1611156107a5576107a4611887565b5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff168267ffffffffffffffff16111561082257610821611887565b5b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160189054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff16111561089f5761089e611887565b5b836000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160008282829054906101000a900467ffffffffffffffff1661090691906118e5565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160088282829054906101000a900467ffffffffffffffff1661099391906118e5565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160108282829054906101000a900467ffffffffffffffff16610a2091906118e5565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160188282829054906101000a900467ffffffffffffffff16610aad91906118e5565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050505050565b610ae1610ef7565b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054610b90906115ce565b80601f0160208091040260200160405190810160405280929190818152602001828054610bbc906115ce565b8015610c095780601f10610bde57610100808354040283529160200191610c09565b820191906000526020600020905b815481529060010190602001808311610bec57829003601f168201915b50505050508152602001600282018054610c22906115ce565b80601f0160208091040260200160405190810160405280929190818152602001828054610c4e906115ce565b8015610c9b5780601f10610c7057610100808354040283529160200191610c9b565b820191906000526020600020905b815481529060010190602001808311610c7e57829003601f168201915b50505050508152602001600382018054610cb4906115ce565b80601f0160208091040260200160405190810160405280929190818152602001828054610ce0906115ce565b8015610d2d5780601f10610d0257610100808354040283529160200191610d2d565b820191906000526020600020905b815481529060010190602001808311610d1057829003601f168201915b50505050508152602001600482016040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250508152602001600582016040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050815250509050919050565b565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160608152602001610f3c610f4f565b8152602001610f49610f4f565b81525090565b6040518060800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61100682610fbd565b810181811067ffffffffffffffff8211171561102557611024610fce565b5b80604052505050565b6000611038610f9f565b90506110448282610ffd565b919050565b600067ffffffffffffffff82111561106457611063610fce565b5b61106d82610fbd565b9050602081019050919050565b82818337600083830152505050565b600061109c61109784611049565b61102e565b9050828152602081018484840111156110b8576110b7610fb8565b5b6110c384828561107a565b509392505050565b600082601f8301126110e0576110df610fb3565b5b81356110f0848260208601611089565b91505092915050565b600067ffffffffffffffff82169050919050565b611116816110f9565b811461112157600080fd5b50565b6000813590506111338161110d565b92915050565b600080600080600080600060e0888a03121561115857611157610fa9565b5b600088013567ffffffffffffffff81111561117657611175610fae565b5b6111828a828b016110cb565b975050602088013567ffffffffffffffff8111156111a3576111a2610fae565b5b6111af8a828b016110cb565b965050604088013567ffffffffffffffff8111156111d0576111cf610fae565b5b6111dc8a828b016110cb565b95505060606111ed8a828b01611124565b94505060806111fe8a828b01611124565b93505060a061120f8a828b01611124565b92505060c06112208a828b01611124565b91505092959891949750929550565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006112868261125b565b9050919050565b6112968161127b565b82525050565b60006112a8838361128d565b60208301905092915050565b6000602082019050919050565b60006112cc8261122f565b6112d6818561123a565b93506112e18361124b565b8060005b838110156113125781516112f9888261129c565b9750611304836112b4565b9250506001810190506112e5565b5085935050505092915050565b6000602082019050818103600083015261133981846112c1565b905092915050565b6000806000806080858703121561135b5761135a610fa9565b5b600061136987828801611124565b945050602061137a87828801611124565b935050604061138b87828801611124565b925050606061139c87828801611124565b91505092959194509250565b6113b18161127b565b81146113bc57600080fd5b50565b6000813590506113ce816113a8565b92915050565b6000602082840312156113ea576113e9610fa9565b5b60006113f8848285016113bf565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561143b578082015181840152602081019050611420565b60008484015250505050565b600061145282611401565b61145c818561140c565b935061146c81856020860161141d565b61147581610fbd565b840191505092915050565b611489816110f9565b82525050565b6080820160008201516114a56000850182611480565b5060208201516114b86020850182611480565b5060408201516114cb6040850182611480565b5060608201516114de6060850182611480565b50505050565b6000610180830160008301516114fd600086018261128d565b50602083015184820360208601526115158282611447565b9150506040830151848203604086015261152f8282611447565b915050606083015184820360608601526115498282611447565b915050608083015161155e608086018261148f565b5060a083015161157261010086018261148f565b508091505092915050565b6000602082019050818103600083015261159781846114e4565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806115e657607f821691505b6020821081036115f9576115f861159f565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026116617fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611624565b61166b8683611624565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006116b26116ad6116a884611683565b61168d565b611683565b9050919050565b6000819050919050565b6116cc83611697565b6116e06116d8826116b9565b848454611631565b825550505050565b600090565b6116f56116e8565b6117008184846116c3565b505050565b5b81811015611724576117196000826116ed565b600181019050611706565b5050565b601f8211156117695761173a816115ff565b61174384611614565b81016020851015611752578190505b61176661175e85611614565b830182611705565b50505b505050565b600082821c905092915050565b600061178c6000198460080261176e565b1980831691505092915050565b60006117a5838361177b565b9150826002028217905092915050565b6117be82611401565b67ffffffffffffffff8111156117d7576117d6610fce565b5b6117e182546115ce565b6117ec828285611728565b600060209050601f83116001811461181f576000841561180d578287015190505b6118178582611799565b86555061187f565b601f19841661182d866115ff565b60005b8281101561185557848901518255600182019150602085019450602081019050611830565b86831015611872578489015161186e601f89168261177b565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006118f0826110f9565b91506118fb836110f9565b9250828203905067ffffffffffffffff81111561191b5761191a6118b6565b5b9291505056fea26469706673582212208c18e5d61689a9be1dee0d39883f7cb1cc0fc57475ca80e91a1bc2e6ace9752264736f6c63430008100033",
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

// GetKeys is a free data retrieval call binding the contract method 0x2150c518.
//
// Solidity: function getKeys() view returns(address[])
func (_Registry *RegistryCaller) GetKeys(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getKeys")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetKeys is a free data retrieval call binding the contract method 0x2150c518.
//
// Solidity: function getKeys() view returns(address[])
func (_Registry *RegistrySession) GetKeys() ([]common.Address, error) {
	return _Registry.Contract.GetKeys(&_Registry.CallOpts)
}

// GetKeys is a free data retrieval call binding the contract method 0x2150c518.
//
// Solidity: function getKeys() view returns(address[])
func (_Registry *RegistryCallerSession) GetKeys() ([]common.Address, error) {
	return _Registry.Contract.GetKeys(&_Registry.CallOpts)
}

// Register is a paid mutator transaction binding the contract method 0x1bb8355e.
//
// Solidity: function register(string ip, string domain, string port, uint64 cpu, uint64 gpu, uint64 mem, uint64 disk) returns()
func (_Registry *RegistryTransactor) Register(opts *bind.TransactOpts, ip string, domain string, port string, cpu uint64, gpu uint64, mem uint64, disk uint64) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "register", ip, domain, port, cpu, gpu, mem, disk)
}

// Register is a paid mutator transaction binding the contract method 0x1bb8355e.
//
// Solidity: function register(string ip, string domain, string port, uint64 cpu, uint64 gpu, uint64 mem, uint64 disk) returns()
func (_Registry *RegistrySession) Register(ip string, domain string, port string, cpu uint64, gpu uint64, mem uint64, disk uint64) (*types.Transaction, error) {
	return _Registry.Contract.Register(&_Registry.TransactOpts, ip, domain, port, cpu, gpu, mem, disk)
}

// Register is a paid mutator transaction binding the contract method 0x1bb8355e.
//
// Solidity: function register(string ip, string domain, string port, uint64 cpu, uint64 gpu, uint64 mem, uint64 disk) returns()
func (_Registry *RegistryTransactorSession) Register(ip string, domain string, port string, cpu uint64, gpu uint64, mem uint64, disk uint64) (*types.Transaction, error) {
	return _Registry.Contract.Register(&_Registry.TransactOpts, ip, domain, port, cpu, gpu, mem, disk)
}

// Revise is a paid mutator transaction binding the contract method 0xcdf3bc6f.
//
// Solidity: function revise() returns()
func (_Registry *RegistryTransactor) Revise(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "revise")
}

// Revise is a paid mutator transaction binding the contract method 0xcdf3bc6f.
//
// Solidity: function revise() returns()
func (_Registry *RegistrySession) Revise() (*types.Transaction, error) {
	return _Registry.Contract.Revise(&_Registry.TransactOpts)
}

// Revise is a paid mutator transaction binding the contract method 0xcdf3bc6f.
//
// Solidity: function revise() returns()
func (_Registry *RegistryTransactorSession) Revise() (*types.Transaction, error) {
	return _Registry.Contract.Revise(&_Registry.TransactOpts)
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
