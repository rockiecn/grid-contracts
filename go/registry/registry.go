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
	Price  RegistryPricePerHour
	Total  RegistryResources
	Avail  RegistryResources
}

// RegistryPricePerHour is an auto generated low-level Go binding around an user-defined struct.
type RegistryPricePerHour struct {
	PCPU  uint64
	PGPU  uint64
	PMEM  uint64
	PDISK uint64
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"pCPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pGPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pMEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pDISK\",\"type\":\"uint64\"}],\"internalType\":\"structRegistry.PricePerHour\",\"name\":\"price\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nCPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nGPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nMEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nDISK\",\"type\":\"uint64\"}],\"internalType\":\"structRegistry.Resources\",\"name\":\"total\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"nCPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nGPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nMEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nDISK\",\"type\":\"uint64\"}],\"internalType\":\"structRegistry.Resources\",\"name\":\"avail\",\"type\":\"tuple\"}],\"internalType\":\"structRegistry.Info\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getKeys\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"cpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"disk\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pcpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pgpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pmem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pdisk\",\"type\":\"uint64\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ip\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"cpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"gpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"mem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"disk\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pcpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pgpu\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pmem\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"pdisk\",\"type\":\"uint64\"}],\"name\":\"revise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"nCPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nGPU\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nMEM\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"nDISK\",\"type\":\"uint64\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061325e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80632150c5181461005c5780637d7282541461007a578063c2bc2efc14610096578063f0f20461146100c6578063fcf3ee06146100e2575b600080fd5b6100646100fe565b60405161007191906126f9565b60405180910390f35b610094600480360381019061008f91906128b5565b61018c565b005b6100b060048036038101906100ab9190612a29565b611434565b6040516100bd9190612c3b565b60405180910390f35b6100e060048036038101906100db9190612c5d565b61192c565b005b6100fc60048036038101906100f791906128b5565b611e23565b005b6060600180548060200260200160405190810160405280929190818152602001828054801561018257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610138575b5050505050905090565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461023d90612d07565b80601f016020809104026020016040519081016040528092919081815260200182805461026990612d07565b80156102b65780601f1061028b576101008083540402835291602001916102b6565b820191906000526020600020905b81548152906001019060200180831161029957829003601f168201915b505050505081526020016002820180546102cf90612d07565b80601f01602080910402602001604051908101604052809291908181526020018280546102fb90612d07565b80156103485780601f1061031d57610100808354040283529160200191610348565b820191906000526020600020905b81548152906001019060200180831161032b57829003601f168201915b5050505050815260200160038201805461036190612d07565b80601f016020809104026020016040519081016040528092919081815260200182805461038d90612d07565b80156103da5780601f106103af576101008083540402835291602001916103da565b820191906000526020600020905b8154815290600101906020018083116103bd57829003601f168201915b50505050508152602001600482016040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250508152602001600582016040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250508152602001600682016040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250508152505090508b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010190816106c79190612eee565b508a6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060020190816107169190612eee565b50896000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060030190816107659190612eee565b50886000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550876000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550866000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550856000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550846000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550836000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508060a001516000015167ffffffffffffffff168967ffffffffffffffff1610610b8c5760008160a00151600001518a610af79190612fef565b9050806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160008282829054906101000a900467ffffffffffffffff16610b60919061302b565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050610d18565b6000898260a0015160000151610ba29190612fef565b90506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff1610610c885760008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550610d16565b806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160008282829054906101000a900467ffffffffffffffff16610cef9190612fef565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505b505b8060a001516020015167ffffffffffffffff168867ffffffffffffffff1610610de65760008160a001516020015189610d519190612fef565b9050806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160088282829054906101000a900467ffffffffffffffff16610dba919061302b565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050610f72565b6000888260a0015160200151610dfc9190612fef565b90506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160089054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff1610610ee25760008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550610f70565b806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160088282829054906101000a900467ffffffffffffffff16610f499190612fef565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505b505b8060a001516040015167ffffffffffffffff168767ffffffffffffffff16106110405760008160a001516040015188610fab9190612fef565b9050806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160108282829054906101000a900467ffffffffffffffff16611014919061302b565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550506111cc565b6000878260a00151604001516110569190612fef565b90506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff161061113c5760008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506111ca565b806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160108282829054906101000a900467ffffffffffffffff166111a39190612fef565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505b505b8060a001516060015167ffffffffffffffff168667ffffffffffffffff161061129a5760008160a0015160600151876112059190612fef565b9050806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160188282829054906101000a900467ffffffffffffffff1661126e919061302b565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050611426565b6000868260a00151606001516112b09190612fef565b90506000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160189054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff16106113965760008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550611424565b806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160188282829054906101000a900467ffffffffffffffff166113fd9190612fef565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505b505b505050505050505050505050565b61143c612504565b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820180546114eb90612d07565b80601f016020809104026020016040519081016040528092919081815260200182805461151790612d07565b80156115645780601f1061153957610100808354040283529160200191611564565b820191906000526020600020905b81548152906001019060200180831161154757829003601f168201915b5050505050815260200160028201805461157d90612d07565b80601f01602080910402602001604051908101604052809291908181526020018280546115a990612d07565b80156115f65780601f106115cb576101008083540402835291602001916115f6565b820191906000526020600020905b8154815290600101906020018083116115d957829003601f168201915b5050505050815260200160038201805461160f90612d07565b80601f016020809104026020016040519081016040528092919081815260200182805461163b90612d07565b80156116885780601f1061165d57610100808354040283529160200191611688565b820191906000526020600020905b81548152906001019060200180831161166b57829003601f168201915b50505050508152602001600482016040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250508152602001600582016040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250508152602001600682016040518060800160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160109054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681526020016000820160189054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050815250509050919050565b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160009054906101000a900467ffffffffffffffff1667ffffffffffffffff168467ffffffffffffffff1611156119db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119d2906130c4565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160089054906101000a900467ffffffffffffffff1667ffffffffffffffff168367ffffffffffffffff161115611a8a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a8190613130565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160109054906101000a900467ffffffffffffffff1667ffffffffffffffff168267ffffffffffffffff161115611b39576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b309061319c565b60405180910390fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160189054906101000a900467ffffffffffffffff1667ffffffffffffffff168167ffffffffffffffff161115611be8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bdf90613208565b60405180910390fd5b836000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160008282829054906101000a900467ffffffffffffffff16611c4f9190612fef565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550826000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160088282829054906101000a900467ffffffffffffffff16611cdc9190612fef565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160108282829054906101000a900467ffffffffffffffff16611d699190612fef565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160188282829054906101000a900467ffffffffffffffff16611df69190612fef565b92506101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050505050565b336000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508a6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001019081611ef19190612eee565b50896000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206002019081611f409190612eee565b50886000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206003019081611f8f9190612eee565b50876000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550866000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550856000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550846000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060050160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550876000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550866000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550856000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550846000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060060160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550836000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550816000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550806000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060040160000160186101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055506001339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050505050505050565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016060815260200160608152602001612549612569565b81526020016125566125b9565b81526020016125636125b9565b81525090565b6040518060800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b6040518060800160405280600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff168152602001600067ffffffffffffffff1681525090565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061266082612635565b9050919050565b61267081612655565b82525050565b60006126828383612667565b60208301905092915050565b6000602082019050919050565b60006126a682612609565b6126b08185612614565b93506126bb83612625565b8060005b838110156126ec5781516126d38882612676565b97506126de8361268e565b9250506001810190506126bf565b5085935050505092915050565b60006020820190508181036000830152612713818461269b565b905092915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61278282612739565b810181811067ffffffffffffffff821117156127a1576127a061274a565b5b80604052505050565b60006127b461271b565b90506127c08282612779565b919050565b600067ffffffffffffffff8211156127e0576127df61274a565b5b6127e982612739565b9050602081019050919050565b82818337600083830152505050565b6000612818612813846127c5565b6127aa565b90508281526020810184848401111561283457612833612734565b5b61283f8482856127f6565b509392505050565b600082601f83011261285c5761285b61272f565b5b813561286c848260208601612805565b91505092915050565b600067ffffffffffffffff82169050919050565b61289281612875565b811461289d57600080fd5b50565b6000813590506128af81612889565b92915050565b60008060008060008060008060008060006101608c8e0312156128db576128da612725565b5b60008c013567ffffffffffffffff8111156128f9576128f861272a565b5b6129058e828f01612847565b9b505060208c013567ffffffffffffffff8111156129265761292561272a565b5b6129328e828f01612847565b9a505060408c013567ffffffffffffffff8111156129535761295261272a565b5b61295f8e828f01612847565b99505060606129708e828f016128a0565b98505060806129818e828f016128a0565b97505060a06129928e828f016128a0565b96505060c06129a38e828f016128a0565b95505060e06129b48e828f016128a0565b9450506101006129c68e828f016128a0565b9350506101206129d88e828f016128a0565b9250506101406129ea8e828f016128a0565b9150509295989b509295989b9093969950565b612a0681612655565b8114612a1157600080fd5b50565b600081359050612a23816129fd565b92915050565b600060208284031215612a3f57612a3e612725565b5b6000612a4d84828501612a14565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015612a90578082015181840152602081019050612a75565b60008484015250505050565b6000612aa782612a56565b612ab18185612a61565b9350612ac1818560208601612a72565b612aca81612739565b840191505092915050565b612ade81612875565b82525050565b608082016000820151612afa6000850182612ad5565b506020820151612b0d6020850182612ad5565b506040820151612b206040850182612ad5565b506060820151612b336060850182612ad5565b50505050565b608082016000820151612b4f6000850182612ad5565b506020820151612b626020850182612ad5565b506040820151612b756040850182612ad5565b506060820151612b886060850182612ad5565b50505050565b600061020083016000830151612ba76000860182612667565b5060208301518482036020860152612bbf8282612a9c565b91505060408301518482036040860152612bd98282612a9c565b91505060608301518482036060860152612bf38282612a9c565b9150506080830151612c086080860182612ae4565b5060a0830151612c1c610100860182612b39565b5060c0830151612c30610180860182612b39565b508091505092915050565b60006020820190508181036000830152612c558184612b8e565b905092915050565b600080600080600060a08688031215612c7957612c78612725565b5b6000612c8788828901612a14565b9550506020612c98888289016128a0565b9450506040612ca9888289016128a0565b9350506060612cba888289016128a0565b9250506080612ccb888289016128a0565b9150509295509295909350565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680612d1f57607f821691505b602082108103612d3257612d31612cd8565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302612d9a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82612d5d565b612da48683612d5d565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000612deb612de6612de184612dbc565b612dc6565b612dbc565b9050919050565b6000819050919050565b612e0583612dd0565b612e19612e1182612df2565b848454612d6a565b825550505050565b600090565b612e2e612e21565b612e39818484612dfc565b505050565b5b81811015612e5d57612e52600082612e26565b600181019050612e3f565b5050565b601f821115612ea257612e7381612d38565b612e7c84612d4d565b81016020851015612e8b578190505b612e9f612e9785612d4d565b830182612e3e565b50505b505050565b600082821c905092915050565b6000612ec560001984600802612ea7565b1980831691505092915050565b6000612ede8383612eb4565b9150826002028217905092915050565b612ef782612a56565b67ffffffffffffffff811115612f1057612f0f61274a565b5b612f1a8254612d07565b612f25828285612e61565b600060209050601f831160018114612f585760008415612f46578287015190505b612f508582612ed2565b865550612fb8565b601f198416612f6686612d38565b60005b82811015612f8e57848901518255600182019150602085019450602081019050612f69565b86831015612fab5784890151612fa7601f891682612eb4565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000612ffa82612875565b915061300583612875565b9250828203905067ffffffffffffffff81111561302557613024612fc0565b5b92915050565b600061303682612875565b915061304183612875565b9250828201905067ffffffffffffffff81111561306157613060612fc0565b5b92915050565b600082825260208201905092915050565b7f6e6f7420656e6f75676820637075000000000000000000000000000000000000600082015250565b60006130ae600e83613067565b91506130b982613078565b602082019050919050565b600060208201905081810360008301526130dd816130a1565b9050919050565b7f6e6f7420656e6f75676820677075000000000000000000000000000000000000600082015250565b600061311a600e83613067565b9150613125826130e4565b602082019050919050565b600060208201905081810360008301526131498161310d565b9050919050565b7f6e6f7420656e6f756768206d656d000000000000000000000000000000000000600082015250565b6000613186600e83613067565b915061319182613150565b602082019050919050565b600060208201905081810360008301526131b581613179565b9050919050565b7f6e6f7420656e6f756768206469736b0000000000000000000000000000000000600082015250565b60006131f2600f83613067565b91506131fd826131bc565b602082019050919050565b60006020820190508181036000830152613221816131e5565b905091905056fea26469706673582212208c230f77c135e330ff3dc6674583296ec553f1756830d3f34d5b0d64931e033964736f6c63430008100033",
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
// Solidity: function get(address a) view returns((address,string,string,string,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64)))
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
// Solidity: function get(address a) view returns((address,string,string,string,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64)))
func (_Registry *RegistrySession) Get(a common.Address) (RegistryInfo, error) {
	return _Registry.Contract.Get(&_Registry.CallOpts, a)
}

// Get is a free data retrieval call binding the contract method 0xc2bc2efc.
//
// Solidity: function get(address a) view returns((address,string,string,string,(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64),(uint64,uint64,uint64,uint64)))
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

// Register is a paid mutator transaction binding the contract method 0xfcf3ee06.
//
// Solidity: function register(string ip, string domain, string port, uint64 cpu, uint64 gpu, uint64 mem, uint64 disk, uint64 pcpu, uint64 pgpu, uint64 pmem, uint64 pdisk) returns()
func (_Registry *RegistryTransactor) Register(opts *bind.TransactOpts, ip string, domain string, port string, cpu uint64, gpu uint64, mem uint64, disk uint64, pcpu uint64, pgpu uint64, pmem uint64, pdisk uint64) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "register", ip, domain, port, cpu, gpu, mem, disk, pcpu, pgpu, pmem, pdisk)
}

// Register is a paid mutator transaction binding the contract method 0xfcf3ee06.
//
// Solidity: function register(string ip, string domain, string port, uint64 cpu, uint64 gpu, uint64 mem, uint64 disk, uint64 pcpu, uint64 pgpu, uint64 pmem, uint64 pdisk) returns()
func (_Registry *RegistrySession) Register(ip string, domain string, port string, cpu uint64, gpu uint64, mem uint64, disk uint64, pcpu uint64, pgpu uint64, pmem uint64, pdisk uint64) (*types.Transaction, error) {
	return _Registry.Contract.Register(&_Registry.TransactOpts, ip, domain, port, cpu, gpu, mem, disk, pcpu, pgpu, pmem, pdisk)
}

// Register is a paid mutator transaction binding the contract method 0xfcf3ee06.
//
// Solidity: function register(string ip, string domain, string port, uint64 cpu, uint64 gpu, uint64 mem, uint64 disk, uint64 pcpu, uint64 pgpu, uint64 pmem, uint64 pdisk) returns()
func (_Registry *RegistryTransactorSession) Register(ip string, domain string, port string, cpu uint64, gpu uint64, mem uint64, disk uint64, pcpu uint64, pgpu uint64, pmem uint64, pdisk uint64) (*types.Transaction, error) {
	return _Registry.Contract.Register(&_Registry.TransactOpts, ip, domain, port, cpu, gpu, mem, disk, pcpu, pgpu, pmem, pdisk)
}

// Revise is a paid mutator transaction binding the contract method 0x7d728254.
//
// Solidity: function revise(string ip, string domain, string port, uint64 cpu, uint64 gpu, uint64 mem, uint64 disk, uint64 pcpu, uint64 pgpu, uint64 pmem, uint64 pdisk) returns()
func (_Registry *RegistryTransactor) Revise(opts *bind.TransactOpts, ip string, domain string, port string, cpu uint64, gpu uint64, mem uint64, disk uint64, pcpu uint64, pgpu uint64, pmem uint64, pdisk uint64) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "revise", ip, domain, port, cpu, gpu, mem, disk, pcpu, pgpu, pmem, pdisk)
}

// Revise is a paid mutator transaction binding the contract method 0x7d728254.
//
// Solidity: function revise(string ip, string domain, string port, uint64 cpu, uint64 gpu, uint64 mem, uint64 disk, uint64 pcpu, uint64 pgpu, uint64 pmem, uint64 pdisk) returns()
func (_Registry *RegistrySession) Revise(ip string, domain string, port string, cpu uint64, gpu uint64, mem uint64, disk uint64, pcpu uint64, pgpu uint64, pmem uint64, pdisk uint64) (*types.Transaction, error) {
	return _Registry.Contract.Revise(&_Registry.TransactOpts, ip, domain, port, cpu, gpu, mem, disk, pcpu, pgpu, pmem, pdisk)
}

// Revise is a paid mutator transaction binding the contract method 0x7d728254.
//
// Solidity: function revise(string ip, string domain, string port, uint64 cpu, uint64 gpu, uint64 mem, uint64 disk, uint64 pcpu, uint64 pgpu, uint64 pmem, uint64 pdisk) returns()
func (_Registry *RegistryTransactorSession) Revise(ip string, domain string, port string, cpu uint64, gpu uint64, mem uint64, disk uint64, pcpu uint64, pgpu uint64, pmem uint64, pdisk uint64) (*types.Transaction, error) {
	return _Registry.Contract.Revise(&_Registry.TransactOpts, ip, domain, port, cpu, gpu, mem, disk, pcpu, pgpu, pmem, pdisk)
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
