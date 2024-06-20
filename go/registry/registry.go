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
	NCPU  uint64
	NGPU  uint64
	NMEM  uint64
	NDISK uint64
}

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nCPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nGPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nMEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nDISK\",\"type\":\"uint64\"}],\"internalType\":\"structRegistry.Resources\",\"name\":\"total\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nCPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nGPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nMEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nDISK\",\"type\":\"uint64\"}],\"internalType\":\"structRegistry.Resources\",\"name\":\"avail\",\"type\":\"tuple\"}],\"internalType\":\"structRegistry.Info\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getKeys\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"cpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"disk\",\"type\":\"uint64\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nCPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nGPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nMEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nDISK\",\"type\":\"uint64\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611bc6806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80631bb8355e1461005c5780632150c51814610078578063c2bc2efc14610096578063cdf3bc6f146100c6578063f0f20461146100d0575b600080fd5b61007660048036038101906100719190611202565b6100ec565b005b61008061061d565b60405161008d91906113e8565b60405180910390f35b6100b060048036038101906100ab9190611436565b6106ab565b6040516100bd91906115df565b60405180910390f35b6100ce610ac7565b005b6100ea60048036038101906100e59190611601565b610ac9565b005b336000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550866000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010190816101ba9190611892565b50856000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020190816102099190611892565b50846000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030190816102589190611892565b50836000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550836000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050505050565b606060018054806020026020016040519081016040528092919081815260200182805480156106a157602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610657575b5050505050905090565b6106b3610fc0565b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060c00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054610762906116ab565b80601f016020809104026020016040519081016040528092919081815260200182805461078e906116ab565b80156107db5780601f106107b0576101008083540402835291602001916107db565b820191906000526020600020905b8154815290600101906020018083116107be57829003601f168201915b505050505081526020016002820180546107f4906116ab565b80601f0160208091040260200160405190810160405280929190818152602001828054610820906116ab565b801561086d5780601f106108425761010080835404028352916020019161086d565b820191906000526020600020905b81548152906001019060200180831161085057829003601f168201915b50505050508152602001600382018054610886906116ab565b80601f01602080910402602001604051908101604052809291908181526020018280546108b2906116ab565b80156108ff5780601f106108d4576101008083540402835291602001916108ff565b820191906000526020600020905b8154815290600101906020018083116108e257829003601f168201915b50505050508152602001600482016040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250508152602001600582016040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050815250509050919050565b565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff168467ffffffffffffffff161115610b78576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b6f906119c1565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160089054906101000a900467ffffffffffffffff1667ffffffffffffffff168367ffffffffffffffff161115610c27576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1e90611a2d565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff168267ffffffffffffffff161115610cd6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ccd90611a99565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160189054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff161115610d85576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d7c90611b05565b60405180910390fd5b836000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160008282829054906101000a900467ffffffffffffffff16610dec9190611b54565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160088282829054906101000a900467ffffffffffffffff16610e799190611b54565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160108282829054906101000a900467ffffffffffffffff16610f069190611b54565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160188282829054906101000a900467ffffffffffffffff16610f939190611b54565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050505050565b6040518060c00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160608152602001611005611018565b8152602001611012611018565b81525090565b6040518060800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6110cf82611086565b810181811067ffffffffffffffff821117156110ee576110ed611097565b5b80604052505050565b6000611101611068565b905061110d82826110c6565b919050565b600067ffffffffffffffff82111561112d5761112c611097565b5b61113682611086565b9050602081019050919050565b82818337600083830152505050565b600061116561116084611112565b6110f7565b90508281526020810184848401111561118157611180611081565b5b61118c848285611143565b509392505050565b600082601f8301126111a9576111a861107c565b5b81356111b9848260208601611152565b91505092915050565b600067ffffffffffffffff82169050919050565b6111df816111c2565b81146111ea57600080fd5b50565b6000813590506111fc816111d6565b92915050565b600080600080600080600060e0888a03121561122157611220611072565b5b600088013567ffffffffffffffff81111561123f5761123e611077565b5b61124b8a828b01611194565b975050602088013567ffffffffffffffff81111561126c5761126b611077565b5b6112788a828b01611194565b965050604088013567ffffffffffffffff81111561129957611298611077565b5b6112a58a828b01611194565b95505060606112b68a828b016111ed565b94505060806112c78a828b016111ed565b93505060a06112d88a828b016111ed565b92505060c06112e98a828b016111ed565b91505092959891949750929550565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061134f82611324565b9050919050565b61135f81611344565b82525050565b60006113718383611356565b60208301905092915050565b6000602082019050919050565b6000611395826112f8565b61139f8185611303565b93506113aa83611314565b8060005b838110156113db5781516113c28882611365565b97506113cd8361137d565b9250506001810190506113ae565b5085935050505092915050565b60006020820190508181036000830152611402818461138a565b905092915050565b61141381611344565b811461141e57600080fd5b50565b6000813590506114308161140a565b92915050565b60006020828403121561144c5761144b611072565b5b600061145a84828501611421565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561149d578082015181840152602081019050611482565b60008484015250505050565b60006114b482611463565b6114be818561146e565b93506114ce81856020860161147f565b6114d781611086565b840191505092915050565b6114eb816111c2565b82525050565b60808201600082015161150760008501826114e2565b50602082015161151a60208501826114e2565b50604082015161152d60408501826114e2565b50606082015161154060608501826114e2565b50505050565b60006101808301600083015161155f6000860182611356565b506020830151848203602086015261157782826114a9565b9150506040830151848203604086015261159182826114a9565b915050606083015184820360608601526115ab82826114a9565b91505060808301516115c060808601826114f1565b5060a08301516115d46101008601826114f1565b508091505092915050565b600060208201905081810360008301526115f98184611546565b905092915050565b600080600080600060a0868803121561161d5761161c611072565b5b600061162b88828901611421565b955050602061163c888289016111ed565b945050604061164d888289016111ed565b935050606061165e888289016111ed565b925050608061166f888289016111ed565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806116c357607f821691505b6020821081036116d6576116d561167c565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261173e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611701565b6117488683611701565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b600061178f61178a61178584611760565b61176a565b611760565b9050919050565b6000819050919050565b6117a983611774565b6117bd6117b582611796565b84845461170e565b825550505050565b600090565b6117d26117c5565b6117dd8184846117a0565b505050565b5b81811015611801576117f66000826117ca565b6001810190506117e3565b5050565b601f82111561184657611817816116dc565b611820846116f1565b8101602085101561182f578190505b61184361183b856116f1565b8301826117e2565b50505b505050565b600082821c905092915050565b60006118696000198460080261184b565b1980831691505092915050565b60006118828383611858565b9150826002028217905092915050565b61189b82611463565b67ffffffffffffffff8111156118b4576118b3611097565b5b6118be82546116ab565b6118c9828285611805565b600060209050601f8311600181146118fc57600084156118ea578287015190505b6118f48582611876565b86555061195c565b601f19841661190a866116dc565b60005b828110156119325784890151825560018201915060208501945060208101905061190d565b8683101561194f578489015161194b601f891682611858565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b7f6e6f7420656e6f75676820637075000000000000000000000000000000000000600082015250565b60006119ab600e83611964565b91506119b682611975565b602082019050919050565b600060208201905081810360008301526119da8161199e565b9050919050565b7f6e6f7420656e6f75676820677075000000000000000000000000000000000000600082015250565b6000611a17600e83611964565b9150611a22826119e1565b602082019050919050565b60006020820190508181036000830152611a4681611a0a565b9050919050565b7f6e6f7420656e6f756768206d656d000000000000000000000000000000000000600082015250565b6000611a83600e83611964565b9150611a8e82611a4d565b602082019050919050565b60006020820190508181036000830152611ab281611a76565b9050919050565b7f6e6f7420656e6f756768206469736b0000000000000000000000000000000000600082015250565b6000611aef600f83611964565b9150611afa82611ab9565b602082019050919050565b60006020820190508181036000830152611b1e81611ae2565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611b5f826111c2565b9150611b6a836111c2565b9250828203905067ffffffffffffffff811115611b8a57611b89611b25565b5b9291505056fea2646970667358221220f839a99a871691d1f1a616b491104f8137c47ade9b8ff02ee838d654f1343a6564736f6c63430008100033",
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

// Update is a paid mutator transaction binding the contract method 0xf0f20461.
//
// Solidity: function update(address provider, uint64 nCPU, uint64 nGPU, uint64 nMEM, uint64 nDISK) returns()
func (_Registry *RegistryTransactor) Update(opts *bind.TransactOpts, provider common.Address, nCPU uint64, nGPU uint64, nMEM uint64, nDISK uint64) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "update", provider, nCPU, nGPU, nMEM, nDISK)
}

// Update is a paid mutator transaction binding the contract method 0xf0f20461.
//
// Solidity: function update(address provider, uint64 nCPU, uint64 nGPU, uint64 nMEM, uint64 nDISK) returns()
func (_Registry *RegistrySession) Update(provider common.Address, nCPU uint64, nGPU uint64, nMEM uint64, nDISK uint64) (*types.Transaction, error) {
	return _Registry.Contract.Update(&_Registry.TransactOpts, provider, nCPU, nGPU, nMEM, nDISK)
}

// Update is a paid mutator transaction binding the contract method 0xf0f20461.
//
// Solidity: function update(address provider, uint64 nCPU, uint64 nGPU, uint64 nMEM, uint64 nDISK) returns()
func (_Registry *RegistryTransactorSession) Update(provider common.Address, nCPU uint64, nGPU uint64, nMEM uint64, nDISK uint64) (*types.Transaction, error) {
	return _Registry.Contract.Update(&_Registry.TransactOpts, provider, nCPU, nGPU, nMEM, nDISK)
}
