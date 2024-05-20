// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swap

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

// SwapMetaData contains all meta data concerning the Swap contract.
var SwapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_credit\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gtoken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rate_numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rate_denominator\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Bought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Sold\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_creditAmount\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creditAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gtokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate_denominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate_numerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settlementToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040523480156200001157600080fd5b50604051620018f5380380620018f5833981810160405281019062000037919062000264565b620000576200004b620000f360201b60201c565b620000fb60201b60201c565b83600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081608081815250508060a0818152505050505050620002d6565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001f182620001c4565b9050919050565b6200020381620001e4565b81146200020f57600080fd5b50565b6000815190506200022381620001f8565b92915050565b6000819050919050565b6200023e8162000229565b81146200024a57600080fd5b50565b6000815190506200025e8162000233565b92915050565b60008060008060808587031215620002815762000280620001bf565b5b6000620002918782880162000212565b9450506020620002a48782880162000212565b9350506040620002b7878288016200024d565b9250506060620002ca878288016200024d565b91505092959194509250565b60805160a0516115dd6200031860003960008181610255015281816107150152610ad5015260008181610231015281816106f40152610af601526115dd6000f3fe6080604052600436106100915760003560e01c80638da5cb5b116100595780638da5cb5b14610145578063aca1b54c14610170578063d96a094a1461019b578063e4849b32146101b7578063f2fde38b146101e057610091565b8063035a1eda146100965780631099da26146100c15780631c6e5648146100ec578063715018a6146101175780637b9e618d1461012e575b600080fd5b3480156100a257600080fd5b506100ab610209565b6040516100b89190610ef3565b60405180910390f35b3480156100cd57600080fd5b506100d661022f565b6040516100e39190610f27565b60405180910390f35b3480156100f857600080fd5b50610101610253565b60405161010e9190610f27565b60405180910390f35b34801561012357600080fd5b5061012c610277565b005b34801561013a57600080fd5b5061014361028b565b005b34801561015157600080fd5b5061015a610426565b6040516101679190610ef3565b60405180910390f35b34801561017c57600080fd5b5061018561044f565b6040516101929190610ef3565b60405180910390f35b6101b560048036038101906101b09190610f73565b610475565b005b3480156101c357600080fd5b506101de60048036038101906101d99190610f73565b6108c0565b005b3480156101ec57600080fd5b5061020760048036038101906102029190610fcc565b610c0a565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b61027f610c8d565b6102896000610d0b565b565b610293610c8d565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016102f09190610ef3565b602060405180830381865afa15801561030d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610331919061100e565b9050600081111561042357600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb610382610426565b836040518363ffffffff1660e01b81526004016103a092919061103b565b6020604051808303816000875af11580156103bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103e3919061109c565b610422576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041990611126565b60405180910390fd5b5b50565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600081116104b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104af90611192565b60405180910390fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b81526004016105179291906111b2565b602060405180830381865afa158015610534573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610558919061100e565b90508181101561056757600080fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330866040518463ffffffff1660e01b81526004016105c8939291906111db565b6020604051808303816000875af11580156105e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061060b919061109c565b90508061064d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106449061125e565b60405180910390fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016106aa9190610ef3565b602060405180830381865afa1580156106c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106eb919061100e565b905060006107447f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000087610dcf9092919063ffffffff16565b905081811115610789576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610780906112f0565b60405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb33836040518363ffffffff1660e01b81526004016107e692919061103b565b6020604051808303816000875af1158015610805573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610829919061109c565b92508261086b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086290611126565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff167fc55650ccda1011e1cdc769b1fbf546ebb8c97800b6072b49e06cd560305b1d67866040516108b19190610f27565b60405180910390a25050505050565b60008111610903576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108fa90611382565b60405180910390fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b81526004016109629291906111b2565b602060405180830381865afa15801561097f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109a3919061100e565b9050818110156109e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109df906113ee565b60405180910390fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330866040518463ffffffff1660e01b8152600401610a49939291906111db565b6020604051808303816000875af1158015610a68573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a8c919061109c565b905080610ace576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac59061145a565b60405180910390fd5b6000610b257f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000086610dcf9092919063ffffffff16565b9050600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1933836040518363ffffffff1660e01b8152600401610b8492919061103b565b600060405180830381600087803b158015610b9e57600080fd5b505af1158015610bb2573d6000803e3d6000fd5b505050503373ffffffffffffffffffffffffffffffffffffffff167fae92ab4b6f8f401ead768d3273e6bb937a13e39827d19c6376e8fd4512a05d9a85604051610bfc9190610f27565b60405180910390a250505050565b610c12610c8d565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610c81576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c78906114ec565b60405180910390fd5b610c8a81610d0b565b50565b610c95610eaa565b73ffffffffffffffffffffffffffffffffffffffff16610cb3610426565b73ffffffffffffffffffffffffffffffffffffffff1614610d09576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d0090611558565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6000806000801985870985870292508281108382030391505060008103610e0a57838281610e0057610dff611578565b5b0492505050610ea3565b808411610e1657600080fd5b60008486880990508281118203915080830392506000600186190186169050808604955080840493506001818260000304019050808302841793506000600287600302189050808702600203810290508087026002038102905080870260020381029050808702600203810290508087026002038102905080870260020381029050808502955050505050505b9392505050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610edd82610eb2565b9050919050565b610eed81610ed2565b82525050565b6000602082019050610f086000830184610ee4565b92915050565b6000819050919050565b610f2181610f0e565b82525050565b6000602082019050610f3c6000830184610f18565b92915050565b600080fd5b610f5081610f0e565b8114610f5b57600080fd5b50565b600081359050610f6d81610f47565b92915050565b600060208284031215610f8957610f88610f42565b5b6000610f9784828501610f5e565b91505092915050565b610fa981610ed2565b8114610fb457600080fd5b50565b600081359050610fc681610fa0565b92915050565b600060208284031215610fe257610fe1610f42565b5b6000610ff084828501610fb7565b91505092915050565b60008151905061100881610f47565b92915050565b60006020828403121561102457611023610f42565b5b600061103284828501610ff9565b91505092915050565b60006040820190506110506000830185610ee4565b61105d6020830184610f18565b9392505050565b60008115159050919050565b61107981611064565b811461108457600080fd5b50565b60008151905061109681611070565b92915050565b6000602082840312156110b2576110b1610f42565b5b60006110c084828501611087565b91505092915050565b600082825260208201905092915050565b7f5472616e73666572206661696c00000000000000000000000000000000000000600082015250565b6000611110600d836110c9565b915061111b826110da565b602082019050919050565b6000602082019050818103600083015261113f81611103565b9050919050565b7f63726564697420616d6f756e74206d757374206c6172676572207468616e2030600082015250565b600061117c6020836110c9565b915061118782611146565b602082019050919050565b600060208201905081810360008301526111ab8161116f565b9050919050565b60006040820190506111c76000830185610ee4565b6111d46020830184610ee4565b9392505050565b60006060820190506111f06000830186610ee4565b6111fd6020830185610ee4565b61120a6040830184610f18565b949350505050565b7f5472616e7366657246726f6d206661696c000000000000000000000000000000600082015250565b60006112486011836110c9565b915061125382611212565b602082019050919050565b600060208201905081810360008301526112778161123b565b9050919050565b7f6e6f7420656e6f75676820746f6b656e20696e207377617020636f6e7472616360008201527f7420746f20627579000000000000000000000000000000000000000000000000602082015250565b60006112da6028836110c9565b91506112e58261127e565b604082019050919050565b60006020820190508181036000830152611309816112cd565b9050919050565b7f596f75206e65656420746f2073656c6c206174206c6561737420736f6d65207460008201527f6f6b656e73000000000000000000000000000000000000000000000000000000602082015250565b600061136c6025836110c9565b915061137782611310565b604082019050919050565b6000602082019050818103600083015261139b8161135f565b9050919050565b7f436865636b2074686520746f6b656e20616c6c6f77616e636500000000000000600082015250565b60006113d86019836110c9565b91506113e3826113a2565b602082019050919050565b60006020820190508181036000830152611407816113cb565b9050919050565b7f7472616e7366657246726f6d206661696c656400000000000000000000000000600082015250565b60006114446013836110c9565b915061144f8261140e565b602082019050919050565b6000602082019050818103600083015261147381611437565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006114d66026836110c9565b91506114e18261147a565b604082019050919050565b60006020820190508181036000830152611505816114c9565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006115426020836110c9565b915061154d8261150c565b602082019050919050565b6000602082019050818103600083015261157181611535565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fdfea2646970667358221220ce6c8b0a051ee7d32b6725d2000b9e2059316dc51a5cf7b4f294eaf3f7baa04c64736f6c63430008100033",
}

// SwapABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapMetaData.ABI instead.
var SwapABI = SwapMetaData.ABI

// SwapBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SwapMetaData.Bin instead.
var SwapBin = SwapMetaData.Bin

// DeploySwap deploys a new Ethereum contract, binding an instance of Swap to it.
func DeploySwap(auth *bind.TransactOpts, backend bind.ContractBackend, _credit common.Address, _gtoken common.Address, _rate_numerator *big.Int, _rate_denominator *big.Int) (common.Address, *types.Transaction, *Swap, error) {
	parsed, err := SwapMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SwapBin), backend, _credit, _gtoken, _rate_numerator, _rate_denominator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Swap{SwapCaller: SwapCaller{contract: contract}, SwapTransactor: SwapTransactor{contract: contract}, SwapFilterer: SwapFilterer{contract: contract}}, nil
}

// Swap is an auto generated Go binding around an Ethereum contract.
type Swap struct {
	SwapCaller     // Read-only binding to the contract
	SwapTransactor // Write-only binding to the contract
	SwapFilterer   // Log filterer for contract events
}

// SwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapSession struct {
	Contract     *Swap             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapCallerSession struct {
	Contract *SwapCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapTransactorSession struct {
	Contract     *SwapTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapRaw struct {
	Contract *Swap // Generic contract binding to access the raw methods on
}

// SwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapCallerRaw struct {
	Contract *SwapCaller // Generic read-only contract binding to access the raw methods on
}

// SwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapTransactorRaw struct {
	Contract *SwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwap creates a new instance of Swap, bound to a specific deployed contract.
func NewSwap(address common.Address, backend bind.ContractBackend) (*Swap, error) {
	contract, err := bindSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swap{SwapCaller: SwapCaller{contract: contract}, SwapTransactor: SwapTransactor{contract: contract}, SwapFilterer: SwapFilterer{contract: contract}}, nil
}

// NewSwapCaller creates a new read-only instance of Swap, bound to a specific deployed contract.
func NewSwapCaller(address common.Address, caller bind.ContractCaller) (*SwapCaller, error) {
	contract, err := bindSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapCaller{contract: contract}, nil
}

// NewSwapTransactor creates a new write-only instance of Swap, bound to a specific deployed contract.
func NewSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapTransactor, error) {
	contract, err := bindSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapTransactor{contract: contract}, nil
}

// NewSwapFilterer creates a new log filterer instance of Swap, bound to a specific deployed contract.
func NewSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapFilterer, error) {
	contract, err := bindSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapFilterer{contract: contract}, nil
}

// bindSwap binds a generic wrapper to an already deployed contract.
func bindSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.SwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transact(opts, method, params...)
}

// CreditAddr is a free data retrieval call binding the contract method 0x035a1eda.
//
// Solidity: function creditAddr() view returns(address)
func (_Swap *SwapCaller) CreditAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "creditAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CreditAddr is a free data retrieval call binding the contract method 0x035a1eda.
//
// Solidity: function creditAddr() view returns(address)
func (_Swap *SwapSession) CreditAddr() (common.Address, error) {
	return _Swap.Contract.CreditAddr(&_Swap.CallOpts)
}

// CreditAddr is a free data retrieval call binding the contract method 0x035a1eda.
//
// Solidity: function creditAddr() view returns(address)
func (_Swap *SwapCallerSession) CreditAddr() (common.Address, error) {
	return _Swap.Contract.CreditAddr(&_Swap.CallOpts)
}

// GtokenAddr is a free data retrieval call binding the contract method 0xaca1b54c.
//
// Solidity: function gtokenAddr() view returns(address)
func (_Swap *SwapCaller) GtokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "gtokenAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GtokenAddr is a free data retrieval call binding the contract method 0xaca1b54c.
//
// Solidity: function gtokenAddr() view returns(address)
func (_Swap *SwapSession) GtokenAddr() (common.Address, error) {
	return _Swap.Contract.GtokenAddr(&_Swap.CallOpts)
}

// GtokenAddr is a free data retrieval call binding the contract method 0xaca1b54c.
//
// Solidity: function gtokenAddr() view returns(address)
func (_Swap *SwapCallerSession) GtokenAddr() (common.Address, error) {
	return _Swap.Contract.GtokenAddr(&_Swap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swap *SwapCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swap *SwapSession) Owner() (common.Address, error) {
	return _Swap.Contract.Owner(&_Swap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swap *SwapCallerSession) Owner() (common.Address, error) {
	return _Swap.Contract.Owner(&_Swap.CallOpts)
}

// RateDenominator is a free data retrieval call binding the contract method 0x1c6e5648.
//
// Solidity: function rate_denominator() view returns(uint256)
func (_Swap *SwapCaller) RateDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "rate_denominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RateDenominator is a free data retrieval call binding the contract method 0x1c6e5648.
//
// Solidity: function rate_denominator() view returns(uint256)
func (_Swap *SwapSession) RateDenominator() (*big.Int, error) {
	return _Swap.Contract.RateDenominator(&_Swap.CallOpts)
}

// RateDenominator is a free data retrieval call binding the contract method 0x1c6e5648.
//
// Solidity: function rate_denominator() view returns(uint256)
func (_Swap *SwapCallerSession) RateDenominator() (*big.Int, error) {
	return _Swap.Contract.RateDenominator(&_Swap.CallOpts)
}

// RateNumerator is a free data retrieval call binding the contract method 0x1099da26.
//
// Solidity: function rate_numerator() view returns(uint256)
func (_Swap *SwapCaller) RateNumerator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "rate_numerator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RateNumerator is a free data retrieval call binding the contract method 0x1099da26.
//
// Solidity: function rate_numerator() view returns(uint256)
func (_Swap *SwapSession) RateNumerator() (*big.Int, error) {
	return _Swap.Contract.RateNumerator(&_Swap.CallOpts)
}

// RateNumerator is a free data retrieval call binding the contract method 0x1099da26.
//
// Solidity: function rate_numerator() view returns(uint256)
func (_Swap *SwapCallerSession) RateNumerator() (*big.Int, error) {
	return _Swap.Contract.RateNumerator(&_Swap.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 _creditAmount) payable returns()
func (_Swap *SwapTransactor) Buy(opts *bind.TransactOpts, _creditAmount *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "buy", _creditAmount)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 _creditAmount) payable returns()
func (_Swap *SwapSession) Buy(_creditAmount *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.Buy(&_Swap.TransactOpts, _creditAmount)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 _creditAmount) payable returns()
func (_Swap *SwapTransactorSession) Buy(_creditAmount *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.Buy(&_Swap.TransactOpts, _creditAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swap *SwapTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swap *SwapSession) RenounceOwnership() (*types.Transaction, error) {
	return _Swap.Contract.RenounceOwnership(&_Swap.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swap *SwapTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Swap.Contract.RenounceOwnership(&_Swap.TransactOpts)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(uint256 _tokenAmount) returns()
func (_Swap *SwapTransactor) Sell(opts *bind.TransactOpts, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "sell", _tokenAmount)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(uint256 _tokenAmount) returns()
func (_Swap *SwapSession) Sell(_tokenAmount *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.Sell(&_Swap.TransactOpts, _tokenAmount)
}

// Sell is a paid mutator transaction binding the contract method 0xe4849b32.
//
// Solidity: function sell(uint256 _tokenAmount) returns()
func (_Swap *SwapTransactorSession) Sell(_tokenAmount *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.Sell(&_Swap.TransactOpts, _tokenAmount)
}

// SettlementToken is a paid mutator transaction binding the contract method 0x7b9e618d.
//
// Solidity: function settlementToken() returns()
func (_Swap *SwapTransactor) SettlementToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "settlementToken")
}

// SettlementToken is a paid mutator transaction binding the contract method 0x7b9e618d.
//
// Solidity: function settlementToken() returns()
func (_Swap *SwapSession) SettlementToken() (*types.Transaction, error) {
	return _Swap.Contract.SettlementToken(&_Swap.TransactOpts)
}

// SettlementToken is a paid mutator transaction binding the contract method 0x7b9e618d.
//
// Solidity: function settlementToken() returns()
func (_Swap *SwapTransactorSession) SettlementToken() (*types.Transaction, error) {
	return _Swap.Contract.SettlementToken(&_Swap.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swap *SwapTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swap *SwapSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Swap.Contract.TransferOwnership(&_Swap.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swap *SwapTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Swap.Contract.TransferOwnership(&_Swap.TransactOpts, newOwner)
}

// SwapBoughtIterator is returned from FilterBought and is used to iterate over the raw logs and unpacked data for Bought events raised by the Swap contract.
type SwapBoughtIterator struct {
	Event *SwapBought // Event containing the contract specifics and raw log

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
func (it *SwapBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapBought)
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
		it.Event = new(SwapBought)
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
func (it *SwapBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapBought represents a Bought event raised by the Swap contract.
type SwapBought struct {
	Buyer  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBought is a free log retrieval operation binding the contract event 0xc55650ccda1011e1cdc769b1fbf546ebb8c97800b6072b49e06cd560305b1d67.
//
// Solidity: event Bought(address indexed buyer, uint256 amount)
func (_Swap *SwapFilterer) FilterBought(opts *bind.FilterOpts, buyer []common.Address) (*SwapBoughtIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "Bought", buyerRule)
	if err != nil {
		return nil, err
	}
	return &SwapBoughtIterator{contract: _Swap.contract, event: "Bought", logs: logs, sub: sub}, nil
}

// WatchBought is a free log subscription operation binding the contract event 0xc55650ccda1011e1cdc769b1fbf546ebb8c97800b6072b49e06cd560305b1d67.
//
// Solidity: event Bought(address indexed buyer, uint256 amount)
func (_Swap *SwapFilterer) WatchBought(opts *bind.WatchOpts, sink chan<- *SwapBought, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "Bought", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapBought)
				if err := _Swap.contract.UnpackLog(event, "Bought", log); err != nil {
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

// ParseBought is a log parse operation binding the contract event 0xc55650ccda1011e1cdc769b1fbf546ebb8c97800b6072b49e06cd560305b1d67.
//
// Solidity: event Bought(address indexed buyer, uint256 amount)
func (_Swap *SwapFilterer) ParseBought(log types.Log) (*SwapBought, error) {
	event := new(SwapBought)
	if err := _Swap.contract.UnpackLog(event, "Bought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Swap contract.
type SwapOwnershipTransferredIterator struct {
	Event *SwapOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SwapOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapOwnershipTransferred)
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
		it.Event = new(SwapOwnershipTransferred)
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
func (it *SwapOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapOwnershipTransferred represents a OwnershipTransferred event raised by the Swap contract.
type SwapOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Swap *SwapFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SwapOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SwapOwnershipTransferredIterator{contract: _Swap.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Swap *SwapFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SwapOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapOwnershipTransferred)
				if err := _Swap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Swap *SwapFilterer) ParseOwnershipTransferred(log types.Log) (*SwapOwnershipTransferred, error) {
	event := new(SwapOwnershipTransferred)
	if err := _Swap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Swap contract.
type SwapReceivedIterator struct {
	Event *SwapReceived // Event containing the contract specifics and raw log

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
func (it *SwapReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapReceived)
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
		it.Event = new(SwapReceived)
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
func (it *SwapReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapReceived represents a Received event raised by the Swap contract.
type SwapReceived struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address sender, uint256 amount)
func (_Swap *SwapFilterer) FilterReceived(opts *bind.FilterOpts) (*SwapReceivedIterator, error) {

	logs, sub, err := _Swap.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &SwapReceivedIterator{contract: _Swap.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address sender, uint256 amount)
func (_Swap *SwapFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *SwapReceived) (event.Subscription, error) {

	logs, sub, err := _Swap.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapReceived)
				if err := _Swap.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address sender, uint256 amount)
func (_Swap *SwapFilterer) ParseReceived(log types.Log) (*SwapReceived, error) {
	event := new(SwapReceived)
	if err := _Swap.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapSoldIterator is returned from FilterSold and is used to iterate over the raw logs and unpacked data for Sold events raised by the Swap contract.
type SwapSoldIterator struct {
	Event *SwapSold // Event containing the contract specifics and raw log

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
func (it *SwapSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapSold)
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
		it.Event = new(SwapSold)
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
func (it *SwapSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapSold represents a Sold event raised by the Swap contract.
type SwapSold struct {
	Seller common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSold is a free log retrieval operation binding the contract event 0xae92ab4b6f8f401ead768d3273e6bb937a13e39827d19c6376e8fd4512a05d9a.
//
// Solidity: event Sold(address indexed seller, uint256 amount)
func (_Swap *SwapFilterer) FilterSold(opts *bind.FilterOpts, seller []common.Address) (*SwapSoldIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "Sold", sellerRule)
	if err != nil {
		return nil, err
	}
	return &SwapSoldIterator{contract: _Swap.contract, event: "Sold", logs: logs, sub: sub}, nil
}

// WatchSold is a free log subscription operation binding the contract event 0xae92ab4b6f8f401ead768d3273e6bb937a13e39827d19c6376e8fd4512a05d9a.
//
// Solidity: event Sold(address indexed seller, uint256 amount)
func (_Swap *SwapFilterer) WatchSold(opts *bind.WatchOpts, sink chan<- *SwapSold, seller []common.Address) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "Sold", sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapSold)
				if err := _Swap.contract.UnpackLog(event, "Sold", log); err != nil {
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

// ParseSold is a log parse operation binding the contract event 0xae92ab4b6f8f401ead768d3273e6bb937a13e39827d19c6376e8fd4512a05d9a.
//
// Solidity: event Sold(address indexed seller, uint256 amount)
func (_Swap *SwapFilterer) ParseSold(log types.Log) (*SwapSold, error) {
	event := new(SwapSold)
	if err := _Swap.contract.UnpackLog(event, "Sold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
