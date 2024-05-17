// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gtoken

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

// GtokenMetaData contains all meta data concerning the Gtoken contract.
var GtokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600681526020017f47546f6b656e00000000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f47544b000000000000000000000000000000000000000000000000000000000081525081600390816200008f9190620005a6565b508060049081620000a19190620005a6565b505050620000c4620000b8620000e760201b60201c565b620000ef60201b60201c565b620000e1336af8277896582678ac000000620001b560201b60201c565b620007a8565b600033905090565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160362000227576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200021e90620006ee565b60405180910390fd5b6200023b600083836200032260201b60201c565b80600260008282546200024f91906200073f565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516200030291906200078b565b60405180910390a36200031e600083836200032760201b60201c565b5050565b505050565b505050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620003ae57607f821691505b602082108103620003c457620003c362000366565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026200042e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620003ef565b6200043a8683620003ef565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b600062000487620004816200047b8462000452565b6200045c565b62000452565b9050919050565b6000819050919050565b620004a38362000466565b620004bb620004b2826200048e565b848454620003fc565b825550505050565b600090565b620004d2620004c3565b620004df81848462000498565b505050565b5b818110156200050757620004fb600082620004c8565b600181019050620004e5565b5050565b601f82111562000556576200052081620003ca565b6200052b84620003df565b810160208510156200053b578190505b620005536200054a85620003df565b830182620004e4565b50505b505050565b600082821c905092915050565b60006200057b600019846008026200055b565b1980831691505092915050565b600062000596838362000568565b9150826002028217905092915050565b620005b1826200032c565b67ffffffffffffffff811115620005cd57620005cc62000337565b5b620005d9825462000395565b620005e68282856200050b565b600060209050601f8311600181146200061e576000841562000609578287015190505b62000615858262000588565b86555062000685565b601f1984166200062e86620003ca565b60005b82811015620006585784890151825560018201915060208501945060208101905062000631565b8683101562000678578489015162000674601f89168262000568565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000620006d6601f836200068d565b9150620006e3826200069e565b602082019050919050565b600060208201905081810360008301526200070981620006c7565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006200074c8262000452565b9150620007598362000452565b925082820190508082111562000774576200077362000710565b5b92915050565b620007858162000452565b82525050565b6000602082019050620007a260008301846200077a565b92915050565b611c1e80620007b86000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c806370a08231116100a2578063a457c2d711610071578063a457c2d7146102d1578063a9059cbb14610301578063d5abeb0114610331578063dd62ed3e1461034f578063f2fde38b1461037f57610116565b806370a082311461025b578063715018a61461028b5780638da5cb5b1461029557806395d89b41146102b357610116565b8063313ce567116100e9578063313ce567146101b7578063378dc3dc146101d557806339509351146101f357806340c10f191461022357806342966c681461023f57610116565b806306fdde031461011b578063095ea7b31461013957806318160ddd1461016957806323b872dd14610187575b600080fd5b61012361039b565b60405161013091906111e2565b60405180910390f35b610153600480360381019061014e919061129d565b61042d565b60405161016091906112f8565b60405180910390f35b610171610450565b60405161017e9190611322565b60405180910390f35b6101a1600480360381019061019c919061133d565b61045a565b6040516101ae91906112f8565b60405180910390f35b6101bf610489565b6040516101cc91906113ac565b60405180910390f35b6101dd610492565b6040516101ea9190611322565b60405180910390f35b61020d6004803603810190610208919061129d565b6104a1565b60405161021a91906112f8565b60405180910390f35b61023d6004803603810190610238919061129d565b6104d8565b005b610259600480360381019061025491906113c7565b610535565b005b610275600480360381019061027091906113f4565b610542565b6040516102829190611322565b60405180910390f35b61029361058a565b005b61029d61059e565b6040516102aa9190611430565b60405180910390f35b6102bb6105c8565b6040516102c891906111e2565b60405180910390f35b6102eb60048036038101906102e6919061129d565b61065a565b6040516102f891906112f8565b60405180910390f35b61031b6004803603810190610316919061129d565b6106d1565b60405161032891906112f8565b60405180910390f35b6103396106f4565b6040516103469190611322565b60405180910390f35b6103696004803603810190610364919061144b565b610704565b6040516103769190611322565b60405180910390f35b610399600480360381019061039491906113f4565b61078b565b005b6060600380546103aa906114ba565b80601f01602080910402602001604051908101604052809291908181526020018280546103d6906114ba565b80156104235780601f106103f857610100808354040283529160200191610423565b820191906000526020600020905b81548152906001019060200180831161040657829003601f168201915b5050505050905090565b60008061043861080e565b9050610445818585610816565b600191505092915050565b6000600254905090565b60008061046561080e565b90506104728582856109df565b61047d858585610a6b565b60019150509392505050565b60006012905090565b6af8277896582678ac00000081565b6000806104ac61080e565b90506104cd8185856104be8589610704565b6104c8919061151a565b610816565b600191505092915050565b6104e0610ce1565b60006104ea610450565b90506b01f04ef12cb04cf1580000008282610505919061151a565b111561052657806b01f04ef12cb04cf158000000610523919061154e565b91505b6105308383610d5f565b505050565b61053f3382610eb5565b50565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b610592610ce1565b61059c6000611082565b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600480546105d7906114ba565b80601f0160208091040260200160405190810160405280929190818152602001828054610603906114ba565b80156106505780601f1061062557610100808354040283529160200191610650565b820191906000526020600020905b81548152906001019060200180831161063357829003601f168201915b5050505050905090565b60008061066561080e565b905060006106738286610704565b9050838110156106b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106af906115f4565b60405180910390fd5b6106c58286868403610816565b60019250505092915050565b6000806106dc61080e565b90506106e9818585610a6b565b600191505092915050565b6b01f04ef12cb04cf15800000081565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b610793610ce1565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610802576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f990611686565b60405180910390fd5b61080b81611082565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610885576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161087c90611718565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108eb906117aa565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516109d29190611322565b60405180910390a3505050565b60006109eb8484610704565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610a655781811015610a57576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4e90611816565b60405180910390fd5b610a648484848403610816565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610ada576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ad1906118a8565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610b49576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b409061193a565b60405180910390fd5b610b54838383611148565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610bda576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bd1906119cc565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610cc89190611322565b60405180910390a3610cdb84848461114d565b50505050565b610ce961080e565b73ffffffffffffffffffffffffffffffffffffffff16610d0761059e565b73ffffffffffffffffffffffffffffffffffffffff1614610d5d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d5490611a38565b60405180910390fd5b565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610dce576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dc590611aa4565b60405180910390fd5b610dda60008383611148565b8060026000828254610dec919061151a565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610e9d9190611322565b60405180910390a3610eb16000838361114d565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610f24576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f1b90611b36565b60405180910390fd5b610f3082600083611148565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610fb6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fad90611bc8565b60405180910390fd5b8181036000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282540392505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516110699190611322565b60405180910390a361107d8360008461114d565b505050565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561118c578082015181840152602081019050611171565b60008484015250505050565b6000601f19601f8301169050919050565b60006111b482611152565b6111be818561115d565b93506111ce81856020860161116e565b6111d781611198565b840191505092915050565b600060208201905081810360008301526111fc81846111a9565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061123482611209565b9050919050565b61124481611229565b811461124f57600080fd5b50565b6000813590506112618161123b565b92915050565b6000819050919050565b61127a81611267565b811461128557600080fd5b50565b60008135905061129781611271565b92915050565b600080604083850312156112b4576112b3611204565b5b60006112c285828601611252565b92505060206112d385828601611288565b9150509250929050565b60008115159050919050565b6112f2816112dd565b82525050565b600060208201905061130d60008301846112e9565b92915050565b61131c81611267565b82525050565b60006020820190506113376000830184611313565b92915050565b60008060006060848603121561135657611355611204565b5b600061136486828701611252565b935050602061137586828701611252565b925050604061138686828701611288565b9150509250925092565b600060ff82169050919050565b6113a681611390565b82525050565b60006020820190506113c1600083018461139d565b92915050565b6000602082840312156113dd576113dc611204565b5b60006113eb84828501611288565b91505092915050565b60006020828403121561140a57611409611204565b5b600061141884828501611252565b91505092915050565b61142a81611229565b82525050565b60006020820190506114456000830184611421565b92915050565b6000806040838503121561146257611461611204565b5b600061147085828601611252565b925050602061148185828601611252565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806114d257607f821691505b6020821081036114e5576114e461148b565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061152582611267565b915061153083611267565b9250828201905080821115611548576115476114eb565b5b92915050565b600061155982611267565b915061156483611267565b925082820390508181111561157c5761157b6114eb565b5b92915050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b60006115de60258361115d565b91506115e982611582565b604082019050919050565b6000602082019050818103600083015261160d816115d1565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b600061167060268361115d565b915061167b82611614565b604082019050919050565b6000602082019050818103600083015261169f81611663565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b600061170260248361115d565b915061170d826116a6565b604082019050919050565b60006020820190508181036000830152611731816116f5565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b600061179460228361115d565b915061179f82611738565b604082019050919050565b600060208201905081810360008301526117c381611787565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b6000611800601d8361115d565b915061180b826117ca565b602082019050919050565b6000602082019050818103600083015261182f816117f3565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b600061189260258361115d565b915061189d82611836565b604082019050919050565b600060208201905081810360008301526118c181611885565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b600061192460238361115d565b915061192f826118c8565b604082019050919050565b6000602082019050818103600083015261195381611917565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b60006119b660268361115d565b91506119c18261195a565b604082019050919050565b600060208201905081810360008301526119e5816119a9565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611a2260208361115d565b9150611a2d826119ec565b602082019050919050565b60006020820190508181036000830152611a5181611a15565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000611a8e601f8361115d565b9150611a9982611a58565b602082019050919050565b60006020820190508181036000830152611abd81611a81565b9050919050565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b6000611b2060218361115d565b9150611b2b82611ac4565b604082019050919050565b60006020820190508181036000830152611b4f81611b13565b9050919050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b6000611bb260228361115d565b9150611bbd82611b56565b604082019050919050565b60006020820190508181036000830152611be181611ba5565b905091905056fea2646970667358221220a80323df211b91377b833f5f772cfbbfbeb3d714141b362089e66b4d3c1f440d64736f6c63430008100033",
}

// GtokenABI is the input ABI used to generate the binding from.
// Deprecated: Use GtokenMetaData.ABI instead.
var GtokenABI = GtokenMetaData.ABI

// GtokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GtokenMetaData.Bin instead.
var GtokenBin = GtokenMetaData.Bin

// DeployGtoken deploys a new Ethereum contract, binding an instance of Gtoken to it.
func DeployGtoken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Gtoken, error) {
	parsed, err := GtokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GtokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Gtoken{GtokenCaller: GtokenCaller{contract: contract}, GtokenTransactor: GtokenTransactor{contract: contract}, GtokenFilterer: GtokenFilterer{contract: contract}}, nil
}

// Gtoken is an auto generated Go binding around an Ethereum contract.
type Gtoken struct {
	GtokenCaller     // Read-only binding to the contract
	GtokenTransactor // Write-only binding to the contract
	GtokenFilterer   // Log filterer for contract events
}

// GtokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type GtokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GtokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GtokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GtokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GtokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GtokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GtokenSession struct {
	Contract     *Gtoken           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GtokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GtokenCallerSession struct {
	Contract *GtokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GtokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GtokenTransactorSession struct {
	Contract     *GtokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GtokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type GtokenRaw struct {
	Contract *Gtoken // Generic contract binding to access the raw methods on
}

// GtokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GtokenCallerRaw struct {
	Contract *GtokenCaller // Generic read-only contract binding to access the raw methods on
}

// GtokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GtokenTransactorRaw struct {
	Contract *GtokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGtoken creates a new instance of Gtoken, bound to a specific deployed contract.
func NewGtoken(address common.Address, backend bind.ContractBackend) (*Gtoken, error) {
	contract, err := bindGtoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gtoken{GtokenCaller: GtokenCaller{contract: contract}, GtokenTransactor: GtokenTransactor{contract: contract}, GtokenFilterer: GtokenFilterer{contract: contract}}, nil
}

// NewGtokenCaller creates a new read-only instance of Gtoken, bound to a specific deployed contract.
func NewGtokenCaller(address common.Address, caller bind.ContractCaller) (*GtokenCaller, error) {
	contract, err := bindGtoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GtokenCaller{contract: contract}, nil
}

// NewGtokenTransactor creates a new write-only instance of Gtoken, bound to a specific deployed contract.
func NewGtokenTransactor(address common.Address, transactor bind.ContractTransactor) (*GtokenTransactor, error) {
	contract, err := bindGtoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GtokenTransactor{contract: contract}, nil
}

// NewGtokenFilterer creates a new log filterer instance of Gtoken, bound to a specific deployed contract.
func NewGtokenFilterer(address common.Address, filterer bind.ContractFilterer) (*GtokenFilterer, error) {
	contract, err := bindGtoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GtokenFilterer{contract: contract}, nil
}

// bindGtoken binds a generic wrapper to an already deployed contract.
func bindGtoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GtokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gtoken *GtokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gtoken.Contract.GtokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gtoken *GtokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gtoken.Contract.GtokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gtoken *GtokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gtoken.Contract.GtokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gtoken *GtokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Gtoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gtoken *GtokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gtoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gtoken *GtokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gtoken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Gtoken *GtokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gtoken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Gtoken *GtokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Gtoken.Contract.Allowance(&_Gtoken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Gtoken *GtokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Gtoken.Contract.Allowance(&_Gtoken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Gtoken *GtokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Gtoken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Gtoken *GtokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Gtoken.Contract.BalanceOf(&_Gtoken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Gtoken *GtokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Gtoken.Contract.BalanceOf(&_Gtoken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Gtoken *GtokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Gtoken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Gtoken *GtokenSession) Decimals() (uint8, error) {
	return _Gtoken.Contract.Decimals(&_Gtoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Gtoken *GtokenCallerSession) Decimals() (uint8, error) {
	return _Gtoken.Contract.Decimals(&_Gtoken.CallOpts)
}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() view returns(uint256)
func (_Gtoken *GtokenCaller) InitialSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gtoken.contract.Call(opts, &out, "initialSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() view returns(uint256)
func (_Gtoken *GtokenSession) InitialSupply() (*big.Int, error) {
	return _Gtoken.Contract.InitialSupply(&_Gtoken.CallOpts)
}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() view returns(uint256)
func (_Gtoken *GtokenCallerSession) InitialSupply() (*big.Int, error) {
	return _Gtoken.Contract.InitialSupply(&_Gtoken.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Gtoken *GtokenCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gtoken.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Gtoken *GtokenSession) MaxSupply() (*big.Int, error) {
	return _Gtoken.Contract.MaxSupply(&_Gtoken.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Gtoken *GtokenCallerSession) MaxSupply() (*big.Int, error) {
	return _Gtoken.Contract.MaxSupply(&_Gtoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Gtoken *GtokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Gtoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Gtoken *GtokenSession) Name() (string, error) {
	return _Gtoken.Contract.Name(&_Gtoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Gtoken *GtokenCallerSession) Name() (string, error) {
	return _Gtoken.Contract.Name(&_Gtoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gtoken *GtokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Gtoken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gtoken *GtokenSession) Owner() (common.Address, error) {
	return _Gtoken.Contract.Owner(&_Gtoken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Gtoken *GtokenCallerSession) Owner() (common.Address, error) {
	return _Gtoken.Contract.Owner(&_Gtoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Gtoken *GtokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Gtoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Gtoken *GtokenSession) Symbol() (string, error) {
	return _Gtoken.Contract.Symbol(&_Gtoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Gtoken *GtokenCallerSession) Symbol() (string, error) {
	return _Gtoken.Contract.Symbol(&_Gtoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gtoken *GtokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Gtoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gtoken *GtokenSession) TotalSupply() (*big.Int, error) {
	return _Gtoken.Contract.TotalSupply(&_Gtoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Gtoken *GtokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Gtoken.Contract.TotalSupply(&_Gtoken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Gtoken *GtokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gtoken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Gtoken *GtokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gtoken.Contract.Approve(&_Gtoken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Gtoken *GtokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gtoken.Contract.Approve(&_Gtoken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns()
func (_Gtoken *GtokenTransactor) Burn(opts *bind.TransactOpts, burnAmount *big.Int) (*types.Transaction, error) {
	return _Gtoken.contract.Transact(opts, "burn", burnAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns()
func (_Gtoken *GtokenSession) Burn(burnAmount *big.Int) (*types.Transaction, error) {
	return _Gtoken.Contract.Burn(&_Gtoken.TransactOpts, burnAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns()
func (_Gtoken *GtokenTransactorSession) Burn(burnAmount *big.Int) (*types.Transaction, error) {
	return _Gtoken.Contract.Burn(&_Gtoken.TransactOpts, burnAmount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Gtoken *GtokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Gtoken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Gtoken *GtokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Gtoken.Contract.DecreaseAllowance(&_Gtoken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Gtoken *GtokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Gtoken.Contract.DecreaseAllowance(&_Gtoken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Gtoken *GtokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Gtoken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Gtoken *GtokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Gtoken.Contract.IncreaseAllowance(&_Gtoken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Gtoken *GtokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Gtoken.Contract.IncreaseAllowance(&_Gtoken.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Gtoken *GtokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gtoken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Gtoken *GtokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gtoken.Contract.Mint(&_Gtoken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Gtoken *GtokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gtoken.Contract.Mint(&_Gtoken.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gtoken *GtokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gtoken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gtoken *GtokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _Gtoken.Contract.RenounceOwnership(&_Gtoken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Gtoken *GtokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Gtoken.Contract.RenounceOwnership(&_Gtoken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Gtoken *GtokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gtoken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Gtoken *GtokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gtoken.Contract.Transfer(&_Gtoken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Gtoken *GtokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gtoken.Contract.Transfer(&_Gtoken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Gtoken *GtokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gtoken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Gtoken *GtokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gtoken.Contract.TransferFrom(&_Gtoken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Gtoken *GtokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Gtoken.Contract.TransferFrom(&_Gtoken.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gtoken *GtokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Gtoken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gtoken *GtokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gtoken.Contract.TransferOwnership(&_Gtoken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Gtoken *GtokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Gtoken.Contract.TransferOwnership(&_Gtoken.TransactOpts, newOwner)
}

// GtokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Gtoken contract.
type GtokenApprovalIterator struct {
	Event *GtokenApproval // Event containing the contract specifics and raw log

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
func (it *GtokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GtokenApproval)
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
		it.Event = new(GtokenApproval)
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
func (it *GtokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GtokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GtokenApproval represents a Approval event raised by the Gtoken contract.
type GtokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Gtoken *GtokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GtokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Gtoken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GtokenApprovalIterator{contract: _Gtoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Gtoken *GtokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GtokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Gtoken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GtokenApproval)
				if err := _Gtoken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Gtoken *GtokenFilterer) ParseApproval(log types.Log) (*GtokenApproval, error) {
	event := new(GtokenApproval)
	if err := _Gtoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GtokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Gtoken contract.
type GtokenOwnershipTransferredIterator struct {
	Event *GtokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GtokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GtokenOwnershipTransferred)
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
		it.Event = new(GtokenOwnershipTransferred)
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
func (it *GtokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GtokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GtokenOwnershipTransferred represents a OwnershipTransferred event raised by the Gtoken contract.
type GtokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gtoken *GtokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GtokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gtoken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GtokenOwnershipTransferredIterator{contract: _Gtoken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Gtoken *GtokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GtokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Gtoken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GtokenOwnershipTransferred)
				if err := _Gtoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Gtoken *GtokenFilterer) ParseOwnershipTransferred(log types.Log) (*GtokenOwnershipTransferred, error) {
	event := new(GtokenOwnershipTransferred)
	if err := _Gtoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GtokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Gtoken contract.
type GtokenTransferIterator struct {
	Event *GtokenTransfer // Event containing the contract specifics and raw log

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
func (it *GtokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GtokenTransfer)
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
		it.Event = new(GtokenTransfer)
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
func (it *GtokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GtokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GtokenTransfer represents a Transfer event raised by the Gtoken contract.
type GtokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Gtoken *GtokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GtokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Gtoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GtokenTransferIterator{contract: _Gtoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Gtoken *GtokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GtokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Gtoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GtokenTransfer)
				if err := _Gtoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Gtoken *GtokenFilterer) ParseTransfer(log types.Log) (*GtokenTransfer, error) {
	event := new(GtokenTransfer)
	if err := _Gtoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
