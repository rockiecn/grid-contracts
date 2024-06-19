// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package credit

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

// CreditMetaData contains all meta data concerning the Credit contract.
var CreditMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"access\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002647380380620026478339818101604052810190620000379190620003fd565b6040518060400160405280600681526020017f43726564697400000000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f43524500000000000000000000000000000000000000000000000000000000008152508160039081620000b49190620006a9565b508060049081620000c69190620006a9565b505050620000e9620000dd6200014e60201b60201c565b6200015660201b60201c565b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062000147336af8277896582678ac0000006200021c60201b60201c565b50620008ab565b600033905090565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036200028e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200028590620007f1565b60405180910390fd5b620002a2600083836200038960201b60201c565b8060026000828254620002b6919062000842565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516200036991906200088e565b60405180910390a362000385600083836200038e60201b60201c565b5050565b505050565b505050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620003c58262000398565b9050919050565b620003d781620003b8565b8114620003e357600080fd5b50565b600081519050620003f781620003cc565b92915050565b60006020828403121562000416576200041562000393565b5b60006200042684828501620003e6565b91505092915050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620004b157607f821691505b602082108103620004c757620004c662000469565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620005317fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620004f2565b6200053d8683620004f2565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200058a620005846200057e8462000555565b6200055f565b62000555565b9050919050565b6000819050919050565b620005a68362000569565b620005be620005b58262000591565b848454620004ff565b825550505050565b600090565b620005d5620005c6565b620005e28184846200059b565b505050565b5b818110156200060a57620005fe600082620005cb565b600181019050620005e8565b5050565b601f82111562000659576200062381620004cd565b6200062e84620004e2565b810160208510156200063e578190505b620006566200064d85620004e2565b830182620005e7565b50505b505050565b600082821c905092915050565b60006200067e600019846008026200065e565b1980831691505092915050565b60006200069983836200066b565b9150826002028217905092915050565b620006b4826200042f565b67ffffffffffffffff811115620006d057620006cf6200043a565b5b620006dc825462000498565b620006e98282856200060e565b600060209050601f8311600181146200072157600084156200070c578287015190505b6200071885826200068b565b86555062000788565b601f1984166200073186620004cd565b60005b828110156200075b5784890151825560018201915060208501945060208101905062000734565b868310156200077b578489015162000777601f8916826200066b565b8355505b6001600288020188555050505b505050505050565b600082825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000620007d9601f8362000790565b9150620007e682620007a1565b602082019050919050565b600060208201905081810360008301526200080c81620007ca565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006200084f8262000555565b91506200085c8362000555565b925082820190508082111562000877576200087662000813565b5b92915050565b620008888162000555565b82525050565b6000602082019050620008a560008301846200087d565b92915050565b611d8c80620008bb6000396000f3fe608060405234801561001057600080fd5b50600436106101215760003560e01c806370a08231116100ad578063a457c2d711610071578063a457c2d7146102fa578063a9059cbb1461032a578063d5abeb011461035a578063dd62ed3e14610378578063f2fde38b146103a857610121565b806370a0823114610266578063715018a61461029657806371907f17146102a05780638da5cb5b146102be57806395d89b41146102dc57610121565b8063313ce567116100f4578063313ce567146101c2578063378dc3dc146101e057806339509351146101fe57806340c10f191461022e57806342966c681461024a57610121565b806306fdde0314610126578063095ea7b31461014457806318160ddd1461017457806323b872dd14610192575b600080fd5b61012e6103c4565b60405161013b91906112bf565b60405180910390f35b61015e6004803603810190610159919061137a565b610456565b60405161016b91906113d5565b60405180910390f35b61017c610479565b60405161018991906113ff565b60405180910390f35b6101ac60048036038101906101a7919061141a565b610483565b6040516101b991906113d5565b60405180910390f35b6101ca6104b2565b6040516101d79190611489565b60405180910390f35b6101e86104bb565b6040516101f591906113ff565b60405180910390f35b6102186004803603810190610213919061137a565b6104ca565b60405161022591906113d5565b60405180910390f35b6102486004803603810190610243919061137a565b610501565b005b610264600480360381019061025f91906114a4565b6105ec565b005b610280600480360381019061027b91906114d1565b6105f9565b60405161028d91906113ff565b60405180910390f35b61029e610641565b005b6102a8610655565b6040516102b5919061150d565b60405180910390f35b6102c661067b565b6040516102d3919061150d565b60405180910390f35b6102e46106a5565b6040516102f191906112bf565b60405180910390f35b610314600480360381019061030f919061137a565b610737565b60405161032191906113d5565b60405180910390f35b610344600480360381019061033f919061137a565b6107ae565b60405161035191906113d5565b60405180910390f35b6103626107d1565b60405161036f91906113ff565b60405180910390f35b610392600480360381019061038d9190611528565b6107e1565b60405161039f91906113ff565b60405180910390f35b6103c260048036038101906103bd91906114d1565b610868565b005b6060600380546103d390611597565b80601f01602080910402602001604051908101604052809291908181526020018280546103ff90611597565b801561044c5780601f106104215761010080835404028352916020019161044c565b820191906000526020600020905b81548152906001019060200180831161042f57829003601f168201915b5050505050905090565b6000806104616108eb565b905061046e8185856108f3565b600191505092915050565b6000600254905090565b60008061048e6108eb565b905061049b858285610abc565b6104a6858585610b48565b60019150509392505050565b60006012905090565b6af8277896582678ac00000081565b6000806104d56108eb565b90506104f68185856104e785896107e1565b6104f191906115f7565b6108f3565b600191505092915050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636fae3d76336040518263ffffffff1660e01b815260040161055c919061150d565b6020604051808303816000875af115801561057b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061059f9190611657565b6105de576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105d5906116d0565b60405180910390fd5b6105e88282610dbe565b5050565b6105f63382610f14565b50565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6106496110e1565b610653600061115f565b565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600480546106b490611597565b80601f01602080910402602001604051908101604052809291908181526020018280546106e090611597565b801561072d5780601f106107025761010080835404028352916020019161072d565b820191906000526020600020905b81548152906001019060200180831161071057829003601f168201915b5050505050905090565b6000806107426108eb565b9050600061075082866107e1565b905083811015610795576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078c90611762565b60405180910390fd5b6107a282868684036108f3565b60019250505092915050565b6000806107b96108eb565b90506107c6818585610b48565b600191505092915050565b6b01f04ef12cb04cf15800000081565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6108706110e1565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036108df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d6906117f4565b60405180910390fd5b6108e88161115f565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610962576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161095990611886565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036109d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c890611918565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610aaf91906113ff565b60405180910390a3505050565b6000610ac884846107e1565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610b425781811015610b34576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b2b90611984565b60405180910390fd5b610b4184848484036108f3565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610bb7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bae90611a16565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610c26576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1d90611aa8565b60405180910390fd5b610c31838383611225565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610cb7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cae90611b3a565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610da591906113ff565b60405180910390a3610db884848461122a565b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610e2d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e2490611ba6565b60405180910390fd5b610e3960008383611225565b8060026000828254610e4b91906115f7565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610efc91906113ff565b60405180910390a3610f106000838361122a565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610f83576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f7a90611c38565b60405180910390fd5b610f8f82600083611225565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015611015576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161100c90611cca565b60405180910390fd5b8181036000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282540392505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516110c891906113ff565b60405180910390a36110dc8360008461122a565b505050565b6110e96108eb565b73ffffffffffffffffffffffffffffffffffffffff1661110761067b565b73ffffffffffffffffffffffffffffffffffffffff161461115d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161115490611d36565b60405180910390fd5b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561126957808201518184015260208101905061124e565b60008484015250505050565b6000601f19601f8301169050919050565b60006112918261122f565b61129b818561123a565b93506112ab81856020860161124b565b6112b481611275565b840191505092915050565b600060208201905081810360008301526112d98184611286565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000611311826112e6565b9050919050565b61132181611306565b811461132c57600080fd5b50565b60008135905061133e81611318565b92915050565b6000819050919050565b61135781611344565b811461136257600080fd5b50565b6000813590506113748161134e565b92915050565b60008060408385031215611391576113906112e1565b5b600061139f8582860161132f565b92505060206113b085828601611365565b9150509250929050565b60008115159050919050565b6113cf816113ba565b82525050565b60006020820190506113ea60008301846113c6565b92915050565b6113f981611344565b82525050565b600060208201905061141460008301846113f0565b92915050565b600080600060608486031215611433576114326112e1565b5b60006114418682870161132f565b93505060206114528682870161132f565b925050604061146386828701611365565b9150509250925092565b600060ff82169050919050565b6114838161146d565b82525050565b600060208201905061149e600083018461147a565b92915050565b6000602082840312156114ba576114b96112e1565b5b60006114c884828501611365565b91505092915050565b6000602082840312156114e7576114e66112e1565b5b60006114f58482850161132f565b91505092915050565b61150781611306565b82525050565b600060208201905061152260008301846114fe565b92915050565b6000806040838503121561153f5761153e6112e1565b5b600061154d8582860161132f565b925050602061155e8582860161132f565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806115af57607f821691505b6020821081036115c2576115c1611568565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061160282611344565b915061160d83611344565b9250828201905080821115611625576116246115c8565b5b92915050565b611634816113ba565b811461163f57600080fd5b50565b6000815190506116518161162b565b92915050565b60006020828403121561166d5761166c6112e1565b5b600061167b84828501611642565b91505092915050565b7f63616c6c6572206e6f206d696e74000000000000000000000000000000000000600082015250565b60006116ba600e8361123a565b91506116c582611684565b602082019050919050565b600060208201905081810360008301526116e9816116ad565b9050919050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b600061174c60258361123a565b9150611757826116f0565b604082019050919050565b6000602082019050818103600083015261177b8161173f565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006117de60268361123a565b91506117e982611782565b604082019050919050565b6000602082019050818103600083015261180d816117d1565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b600061187060248361123a565b915061187b82611814565b604082019050919050565b6000602082019050818103600083015261189f81611863565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b600061190260228361123a565b915061190d826118a6565b604082019050919050565b60006020820190508181036000830152611931816118f5565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b600061196e601d8361123a565b915061197982611938565b602082019050919050565b6000602082019050818103600083015261199d81611961565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b6000611a0060258361123a565b9150611a0b826119a4565b604082019050919050565b60006020820190508181036000830152611a2f816119f3565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000611a9260238361123a565b9150611a9d82611a36565b604082019050919050565b60006020820190508181036000830152611ac181611a85565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b6000611b2460268361123a565b9150611b2f82611ac8565b604082019050919050565b60006020820190508181036000830152611b5381611b17565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000611b90601f8361123a565b9150611b9b82611b5a565b602082019050919050565b60006020820190508181036000830152611bbf81611b83565b9050919050565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b6000611c2260218361123a565b9150611c2d82611bc6565b604082019050919050565b60006020820190508181036000830152611c5181611c15565b9050919050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b6000611cb460228361123a565b9150611cbf82611c58565b604082019050919050565b60006020820190508181036000830152611ce381611ca7565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611d2060208361123a565b9150611d2b82611cea565b602082019050919050565b60006020820190508181036000830152611d4f81611d13565b905091905056fea2646970667358221220464ea9e94b744409fb6e113d60a6ecee45ad89727cf0d424f029fdc1223840b664736f6c63430008100033",
}

// CreditABI is the input ABI used to generate the binding from.
// Deprecated: Use CreditMetaData.ABI instead.
var CreditABI = CreditMetaData.ABI

// CreditBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CreditMetaData.Bin instead.
var CreditBin = CreditMetaData.Bin

// DeployCredit deploys a new Ethereum contract, binding an instance of Credit to it.
func DeployCredit(auth *bind.TransactOpts, backend bind.ContractBackend, _a common.Address) (common.Address, *types.Transaction, *Credit, error) {
	parsed, err := CreditMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CreditBin), backend, _a)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Credit{CreditCaller: CreditCaller{contract: contract}, CreditTransactor: CreditTransactor{contract: contract}, CreditFilterer: CreditFilterer{contract: contract}}, nil
}

// Credit is an auto generated Go binding around an Ethereum contract.
type Credit struct {
	CreditCaller     // Read-only binding to the contract
	CreditTransactor // Write-only binding to the contract
	CreditFilterer   // Log filterer for contract events
}

// CreditCaller is an auto generated read-only Go binding around an Ethereum contract.
type CreditCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CreditTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CreditFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CreditSession struct {
	Contract     *Credit           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreditCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CreditCallerSession struct {
	Contract *CreditCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CreditTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CreditTransactorSession struct {
	Contract     *CreditTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreditRaw is an auto generated low-level Go binding around an Ethereum contract.
type CreditRaw struct {
	Contract *Credit // Generic contract binding to access the raw methods on
}

// CreditCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CreditCallerRaw struct {
	Contract *CreditCaller // Generic read-only contract binding to access the raw methods on
}

// CreditTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CreditTransactorRaw struct {
	Contract *CreditTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCredit creates a new instance of Credit, bound to a specific deployed contract.
func NewCredit(address common.Address, backend bind.ContractBackend) (*Credit, error) {
	contract, err := bindCredit(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Credit{CreditCaller: CreditCaller{contract: contract}, CreditTransactor: CreditTransactor{contract: contract}, CreditFilterer: CreditFilterer{contract: contract}}, nil
}

// NewCreditCaller creates a new read-only instance of Credit, bound to a specific deployed contract.
func NewCreditCaller(address common.Address, caller bind.ContractCaller) (*CreditCaller, error) {
	contract, err := bindCredit(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditCaller{contract: contract}, nil
}

// NewCreditTransactor creates a new write-only instance of Credit, bound to a specific deployed contract.
func NewCreditTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditTransactor, error) {
	contract, err := bindCredit(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditTransactor{contract: contract}, nil
}

// NewCreditFilterer creates a new log filterer instance of Credit, bound to a specific deployed contract.
func NewCreditFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditFilterer, error) {
	contract, err := bindCredit(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditFilterer{contract: contract}, nil
}

// bindCredit binds a generic wrapper to an already deployed contract.
func bindCredit(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Credit *CreditRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Credit.Contract.CreditCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Credit *CreditRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Credit.Contract.CreditTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Credit *CreditRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Credit.Contract.CreditTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Credit *CreditCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Credit.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Credit *CreditTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Credit.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Credit *CreditTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Credit.Contract.contract.Transact(opts, method, params...)
}

// Access is a free data retrieval call binding the contract method 0x71907f17.
//
// Solidity: function access() view returns(address)
func (_Credit *CreditCaller) Access(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "access")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Access is a free data retrieval call binding the contract method 0x71907f17.
//
// Solidity: function access() view returns(address)
func (_Credit *CreditSession) Access() (common.Address, error) {
	return _Credit.Contract.Access(&_Credit.CallOpts)
}

// Access is a free data retrieval call binding the contract method 0x71907f17.
//
// Solidity: function access() view returns(address)
func (_Credit *CreditCallerSession) Access() (common.Address, error) {
	return _Credit.Contract.Access(&_Credit.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Credit *CreditCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Credit *CreditSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Credit.Contract.Allowance(&_Credit.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Credit *CreditCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Credit.Contract.Allowance(&_Credit.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Credit *CreditCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Credit *CreditSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Credit.Contract.BalanceOf(&_Credit.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Credit *CreditCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Credit.Contract.BalanceOf(&_Credit.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Credit *CreditCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Credit *CreditSession) Decimals() (uint8, error) {
	return _Credit.Contract.Decimals(&_Credit.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Credit *CreditCallerSession) Decimals() (uint8, error) {
	return _Credit.Contract.Decimals(&_Credit.CallOpts)
}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() view returns(uint256)
func (_Credit *CreditCaller) InitialSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "initialSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() view returns(uint256)
func (_Credit *CreditSession) InitialSupply() (*big.Int, error) {
	return _Credit.Contract.InitialSupply(&_Credit.CallOpts)
}

// InitialSupply is a free data retrieval call binding the contract method 0x378dc3dc.
//
// Solidity: function initialSupply() view returns(uint256)
func (_Credit *CreditCallerSession) InitialSupply() (*big.Int, error) {
	return _Credit.Contract.InitialSupply(&_Credit.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Credit *CreditCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Credit *CreditSession) MaxSupply() (*big.Int, error) {
	return _Credit.Contract.MaxSupply(&_Credit.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_Credit *CreditCallerSession) MaxSupply() (*big.Int, error) {
	return _Credit.Contract.MaxSupply(&_Credit.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Credit *CreditCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Credit *CreditSession) Name() (string, error) {
	return _Credit.Contract.Name(&_Credit.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Credit *CreditCallerSession) Name() (string, error) {
	return _Credit.Contract.Name(&_Credit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Credit *CreditCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Credit *CreditSession) Owner() (common.Address, error) {
	return _Credit.Contract.Owner(&_Credit.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Credit *CreditCallerSession) Owner() (common.Address, error) {
	return _Credit.Contract.Owner(&_Credit.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Credit *CreditCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Credit *CreditSession) Symbol() (string, error) {
	return _Credit.Contract.Symbol(&_Credit.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Credit *CreditCallerSession) Symbol() (string, error) {
	return _Credit.Contract.Symbol(&_Credit.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Credit *CreditCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Credit.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Credit *CreditSession) TotalSupply() (*big.Int, error) {
	return _Credit.Contract.TotalSupply(&_Credit.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Credit *CreditCallerSession) TotalSupply() (*big.Int, error) {
	return _Credit.Contract.TotalSupply(&_Credit.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Credit *CreditTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Credit *CreditSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.Approve(&_Credit.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Credit *CreditTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.Approve(&_Credit.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns()
func (_Credit *CreditTransactor) Burn(opts *bind.TransactOpts, burnAmount *big.Int) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "burn", burnAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns()
func (_Credit *CreditSession) Burn(burnAmount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.Burn(&_Credit.TransactOpts, burnAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 burnAmount) returns()
func (_Credit *CreditTransactorSession) Burn(burnAmount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.Burn(&_Credit.TransactOpts, burnAmount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Credit *CreditTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Credit *CreditSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.DecreaseAllowance(&_Credit.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Credit *CreditTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.DecreaseAllowance(&_Credit.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Credit *CreditTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Credit *CreditSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.IncreaseAllowance(&_Credit.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Credit *CreditTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.IncreaseAllowance(&_Credit.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Credit *CreditTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Credit *CreditSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.Mint(&_Credit.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Credit *CreditTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.Mint(&_Credit.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Credit *CreditTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Credit *CreditSession) RenounceOwnership() (*types.Transaction, error) {
	return _Credit.Contract.RenounceOwnership(&_Credit.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Credit *CreditTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Credit.Contract.RenounceOwnership(&_Credit.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Credit *CreditTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Credit *CreditSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.Transfer(&_Credit.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Credit *CreditTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.Transfer(&_Credit.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Credit *CreditTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Credit *CreditSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.TransferFrom(&_Credit.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Credit *CreditTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Credit.Contract.TransferFrom(&_Credit.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Credit *CreditTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Credit.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Credit *CreditSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Credit.Contract.TransferOwnership(&_Credit.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Credit *CreditTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Credit.Contract.TransferOwnership(&_Credit.TransactOpts, newOwner)
}

// CreditApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Credit contract.
type CreditApprovalIterator struct {
	Event *CreditApproval // Event containing the contract specifics and raw log

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
func (it *CreditApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditApproval)
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
		it.Event = new(CreditApproval)
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
func (it *CreditApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditApproval represents a Approval event raised by the Credit contract.
type CreditApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Credit *CreditFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CreditApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Credit.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CreditApprovalIterator{contract: _Credit.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Credit *CreditFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CreditApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Credit.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditApproval)
				if err := _Credit.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Credit *CreditFilterer) ParseApproval(log types.Log) (*CreditApproval, error) {
	event := new(CreditApproval)
	if err := _Credit.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Credit contract.
type CreditOwnershipTransferredIterator struct {
	Event *CreditOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CreditOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditOwnershipTransferred)
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
		it.Event = new(CreditOwnershipTransferred)
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
func (it *CreditOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditOwnershipTransferred represents a OwnershipTransferred event raised by the Credit contract.
type CreditOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Credit *CreditFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CreditOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Credit.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CreditOwnershipTransferredIterator{contract: _Credit.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Credit *CreditFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CreditOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Credit.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditOwnershipTransferred)
				if err := _Credit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Credit *CreditFilterer) ParseOwnershipTransferred(log types.Log) (*CreditOwnershipTransferred, error) {
	event := new(CreditOwnershipTransferred)
	if err := _Credit.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CreditTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Credit contract.
type CreditTransferIterator struct {
	Event *CreditTransfer // Event containing the contract specifics and raw log

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
func (it *CreditTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditTransfer)
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
		it.Event = new(CreditTransfer)
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
func (it *CreditTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditTransfer represents a Transfer event raised by the Credit contract.
type CreditTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Credit *CreditFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CreditTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Credit.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CreditTransferIterator{contract: _Credit.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Credit *CreditFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CreditTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Credit.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditTransfer)
				if err := _Credit.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Credit *CreditFilterer) ParseTransfer(log types.Log) (*CreditTransfer, error) {
	event := new(CreditTransfer)
	if err := _Credit.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
