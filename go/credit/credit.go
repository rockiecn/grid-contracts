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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"access\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200232738038062002327833981810160405281019062000037919062000269565b6040518060400160405280600681526020017f43726564697400000000000000000000000000000000000000000000000000008152506040518060400160405280600381526020017f43524500000000000000000000000000000000000000000000000000000000008152508160039081620000b4919062000515565b508060049081620000c6919062000515565b505050620000e9620000dd6200013160201b60201c565b6200013960201b60201c565b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050620005fc565b600033905090565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620002318262000204565b9050919050565b620002438162000224565b81146200024f57600080fd5b50565b600081519050620002638162000238565b92915050565b600060208284031215620002825762000281620001ff565b5b6000620002928482850162000252565b91505092915050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200031d57607f821691505b602082108103620003335762000332620002d5565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026200039d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200035e565b620003a986836200035e565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620003f6620003f0620003ea84620003c1565b620003cb565b620003c1565b9050919050565b6000819050919050565b6200041283620003d5565b6200042a6200042182620003fd565b8484546200036b565b825550505050565b600090565b6200044162000432565b6200044e81848462000407565b505050565b5b8181101562000476576200046a60008262000437565b60018101905062000454565b5050565b601f821115620004c5576200048f8162000339565b6200049a846200034e565b81016020851015620004aa578190505b620004c2620004b9856200034e565b83018262000453565b50505b505050565b600082821c905092915050565b6000620004ea60001984600802620004ca565b1980831691505092915050565b6000620005058383620004d7565b9150826002028217905092915050565b62000520826200029b565b67ffffffffffffffff8111156200053c576200053b620002a6565b5b62000548825462000304565b620005558282856200047a565b600060209050601f8311600181146200058d576000841562000578578287015190505b620005848582620004f7565b865550620005f4565b601f1984166200059d8662000339565b60005b82811015620005c757848901518255600182019150602085019450602081019050620005a0565b86831015620005e75784890151620005e3601f891682620004d7565b8355505b6001600288020188555050505b505050505050565b611d1b806200060c6000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c806370a08231116100a257806395d89b411161007157806395d89b41146102a8578063a457c2d7146102c6578063a9059cbb146102f6578063dd62ed3e14610326578063f2fde38b146103565761010b565b806370a0823114610232578063715018a61461026257806371907f171461026c5780638da5cb5b1461028a5761010b565b8063313ce567116100de578063313ce567146101ac57806339509351146101ca57806340c10f19146101fa57806342966c68146102165761010b565b806306fdde0314610110578063095ea7b31461012e57806318160ddd1461015e57806323b872dd1461017c575b600080fd5b610118610372565b604051610125919061124e565b60405180910390f35b61014860048036038101906101439190611309565b610404565b6040516101559190611364565b60405180910390f35b610166610427565b604051610173919061138e565b60405180910390f35b610196600480360381019061019191906113a9565b610431565b6040516101a39190611364565b60405180910390f35b6101b4610460565b6040516101c19190611418565b60405180910390f35b6101e460048036038101906101df9190611309565b610469565b6040516101f19190611364565b60405180910390f35b610214600480360381019061020f9190611309565b6104a0565b005b610230600480360381019061022b9190611433565b61058b565b005b61024c60048036038101906102479190611460565b610598565b604051610259919061138e565b60405180910390f35b61026a6105e0565b005b6102746105f4565b604051610281919061149c565b60405180910390f35b61029261061a565b60405161029f919061149c565b60405180910390f35b6102b0610644565b6040516102bd919061124e565b60405180910390f35b6102e060048036038101906102db9190611309565b6106d6565b6040516102ed9190611364565b60405180910390f35b610310600480360381019061030b9190611309565b61074d565b60405161031d9190611364565b60405180910390f35b610340600480360381019061033b91906114b7565b610770565b60405161034d919061138e565b60405180910390f35b610370600480360381019061036b9190611460565b6107f7565b005b60606003805461038190611526565b80601f01602080910402602001604051908101604052809291908181526020018280546103ad90611526565b80156103fa5780601f106103cf576101008083540402835291602001916103fa565b820191906000526020600020905b8154815290600101906020018083116103dd57829003601f168201915b5050505050905090565b60008061040f61087a565b905061041c818585610882565b600191505092915050565b6000600254905090565b60008061043c61087a565b9050610449858285610a4b565b610454858585610ad7565b60019150509392505050565b60006012905090565b60008061047461087a565b90506104958185856104868589610770565b6104909190611586565b610882565b600191505092915050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636fae3d76336040518263ffffffff1660e01b81526004016104fb919061149c565b6020604051808303816000875af115801561051a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061053e91906115e6565b61057d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105749061165f565b60405180910390fd5b6105878282610d4d565b5050565b6105953382610ea3565b50565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6105e8611070565b6105f260006110ee565b565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606004805461065390611526565b80601f016020809104026020016040519081016040528092919081815260200182805461067f90611526565b80156106cc5780601f106106a1576101008083540402835291602001916106cc565b820191906000526020600020905b8154815290600101906020018083116106af57829003601f168201915b5050505050905090565b6000806106e161087a565b905060006106ef8286610770565b905083811015610734576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161072b906116f1565b60405180910390fd5b6107418286868403610882565b60019250505092915050565b60008061075861087a565b9050610765818585610ad7565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6107ff611070565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361086e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086590611783565b60405180910390fd5b610877816110ee565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036108f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108e890611815565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610960576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610957906118a7565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610a3e919061138e565b60405180910390a3505050565b6000610a578484610770565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610ad15781811015610ac3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aba90611913565b60405180910390fd5b610ad08484848403610882565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610b46576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3d906119a5565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610bb5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bac90611a37565b60405180910390fd5b610bc08383836111b4565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610c46576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c3d90611ac9565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610d34919061138e565b60405180910390a3610d478484846111b9565b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610dbc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610db390611b35565b60405180910390fd5b610dc8600083836111b4565b8060026000828254610dda9190611586565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610e8b919061138e565b60405180910390a3610e9f600083836111b9565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610f12576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f0990611bc7565b60405180910390fd5b610f1e826000836111b4565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610fa4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f9b90611c59565b60405180910390fd5b8181036000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282540392505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051611057919061138e565b60405180910390a361106b836000846111b9565b505050565b61107861087a565b73ffffffffffffffffffffffffffffffffffffffff1661109661061a565b73ffffffffffffffffffffffffffffffffffffffff16146110ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110e390611cc5565b60405180910390fd5b565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b838110156111f85780820151818401526020810190506111dd565b60008484015250505050565b6000601f19601f8301169050919050565b6000611220826111be565b61122a81856111c9565b935061123a8185602086016111da565b61124381611204565b840191505092915050565b600060208201905081810360008301526112688184611215565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006112a082611275565b9050919050565b6112b081611295565b81146112bb57600080fd5b50565b6000813590506112cd816112a7565b92915050565b6000819050919050565b6112e6816112d3565b81146112f157600080fd5b50565b600081359050611303816112dd565b92915050565b600080604083850312156113205761131f611270565b5b600061132e858286016112be565b925050602061133f858286016112f4565b9150509250929050565b60008115159050919050565b61135e81611349565b82525050565b60006020820190506113796000830184611355565b92915050565b611388816112d3565b82525050565b60006020820190506113a3600083018461137f565b92915050565b6000806000606084860312156113c2576113c1611270565b5b60006113d0868287016112be565b93505060206113e1868287016112be565b92505060406113f2868287016112f4565b9150509250925092565b600060ff82169050919050565b611412816113fc565b82525050565b600060208201905061142d6000830184611409565b92915050565b60006020828403121561144957611448611270565b5b6000611457848285016112f4565b91505092915050565b60006020828403121561147657611475611270565b5b6000611484848285016112be565b91505092915050565b61149681611295565b82525050565b60006020820190506114b1600083018461148d565b92915050565b600080604083850312156114ce576114cd611270565b5b60006114dc858286016112be565b92505060206114ed858286016112be565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061153e57607f821691505b602082108103611551576115506114f7565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611591826112d3565b915061159c836112d3565b92508282019050808211156115b4576115b3611557565b5b92915050565b6115c381611349565b81146115ce57600080fd5b50565b6000815190506115e0816115ba565b92915050565b6000602082840312156115fc576115fb611270565b5b600061160a848285016115d1565b91505092915050565b7f63616c6c6572206e6f206d696e74000000000000000000000000000000000000600082015250565b6000611649600e836111c9565b915061165482611613565b602082019050919050565b600060208201905081810360008301526116788161163c565b9050919050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b60006116db6025836111c9565b91506116e68261167f565b604082019050919050565b6000602082019050818103600083015261170a816116ce565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b600061176d6026836111c9565b915061177882611711565b604082019050919050565b6000602082019050818103600083015261179c81611760565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b60006117ff6024836111c9565b915061180a826117a3565b604082019050919050565b6000602082019050818103600083015261182e816117f2565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b60006118916022836111c9565b915061189c82611835565b604082019050919050565b600060208201905081810360008301526118c081611884565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b60006118fd601d836111c9565b9150611908826118c7565b602082019050919050565b6000602082019050818103600083015261192c816118f0565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b600061198f6025836111c9565b915061199a82611933565b604082019050919050565b600060208201905081810360008301526119be81611982565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000611a216023836111c9565b9150611a2c826119c5565b604082019050919050565b60006020820190508181036000830152611a5081611a14565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b6000611ab36026836111c9565b9150611abe82611a57565b604082019050919050565b60006020820190508181036000830152611ae281611aa6565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000611b1f601f836111c9565b9150611b2a82611ae9565b602082019050919050565b60006020820190508181036000830152611b4e81611b12565b9050919050565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b6000611bb16021836111c9565b9150611bbc82611b55565b604082019050919050565b60006020820190508181036000830152611be081611ba4565b9050919050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b6000611c436022836111c9565b9150611c4e82611be7565b604082019050919050565b60006020820190508181036000830152611c7281611c36565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000611caf6020836111c9565b9150611cba82611c79565b602082019050919050565b60006020820190508181036000830152611cde81611ca2565b905091905056fea264697066735822122013b84a995bcba52e0e3233d8654b2ca29aa104c3f1bf347c20715ffffdd948e664736f6c63430008100033",
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
