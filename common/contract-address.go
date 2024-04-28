package common

import (
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/xerrors"
)

var (
	DevChain     = "https://devchain.metamemo.one:8501"
	TestChain    = "https://testchain.metamemo.one:24180"
	ProductChain = "https://chain.metamemo.one:8501"
	GoerliChain  = "https://eth-goerli.g.alchemy.com/v2/Bn3AbuwyuTWanFLJiflS-dc23r1Re_Af"

	DefaultGasLimit = uint64(6000000) // as small as possible
	DefaultGasPrice = big.NewInt(100) // as small as possible
	PledgeAmount    = "1 memo"        // 1 memo = 1e9 nanomemo = 1e18 attomemo
	KPledgePerB     = "909495"        // 1memo/1TB
	PPledgePerB     = "9094947"       // 10memo/1TB
	nano            = big.NewInt(1e9)
	memo            = big.NewInt(1e18)

	DefaultCap = uint64(1073741824) // capacity that provider can provide, default 1TB
	DefaultLoc = "personal"

	WaitTime = 30 * time.Second

	// AdminAddr admin address
	AdminAddr = common.HexToAddress("0x1c111472F298E4119150850c198C657DA1F8a368")
	// Foundation foundation address
	Foundation = common.HexToAddress("0x98B0B2387f98206efbF6fbCe2462cE22916BAAa3")

	// Mediate Bank with 1M MEMO
	// addres: 0xcdb5e0cb700e851a79e00347fff38d6b1b77dcb4
	// sk: 0xf572673a18b0d94c1c2c11e6d39f64aff6270842d53814c361dcad0309af38e0

	Erc20Name   = "memo"
	Erc20Symbol = "M"

	TypeProxy        = uint8(100)
	TypeControl1     = uint8(101)
	TypeControl2     = uint8(102)
	TypeControl3     = uint8(103)
	TypeAuth         = uint8(2)
	TypeAccess       = uint8(3)
	TypeERC20        = uint8(4)
	TypePoolP        = uint8(5)
	TypeRole         = uint8(6)
	TypeToken        = uint8(7)
	TypePledge       = uint8(8)
	TypeIssu         = uint8(9)
	Typefs           = uint8(10)
	TypeGroup        = uint8(11)
	TypePoolF        = uint8(12)
	TypeKmanage      = uint8(13)
	TypeSwap         = uint8(21)
	TypeReward       = uint8(22)
	TypeRewardAccess = uint8(23)
	TypeGetter       = uint8(150)

	// did
	TypeAccountDid     = uint8(30)
	TypeParsePubKey    = uint8(31)
	TypeDidControl     = uint8(32)
	TypeDidProxy       = uint8(33)
	TypeFileDid        = uint8(34)
	TypeFileDidControl = uint8(35)

	// middleware
	TypeReadPay         = uint8(40)
	TypeStorePay        = uint8(41)
	TypePaymentControl  = uint8(42)
	TypeMiddlewareProxy = uint8(43)

	UserType     = uint8(1)
	ProviderType = uint8(2)
	KeeperType   = uint8(3)
)

var (
	// dev
	InstanceAddr = common.HexToAddress("0xfCf4CE56e19A1234d6b79a12A624431A84513481")

	// test old
	//TestnetInstanceAddr = common.HexToAddress("0x3F29807CbCe7EaE9cbDE54E60392B9b7ce6F6a26")

	// test
	TestnetInstanceAddr = common.HexToAddress("0x3c9ebF885d38813632FD129f892E1E0b75371cfA")

	// product
	ProductInstanceAddr = common.HexToAddress("0xbd16029A7126C91ED42E9157dc7BADD2B3a81189")

	// goerli
	GoerliInstanceAddr = common.HexToAddress("0x8282d8CD56F293ea54372Ba66baB66DB5BBBd0Ab")
)

var (
	ErrMemoAmount = xerrors.New("memoAmount's uint should be memo、nanomemo、attomemo; number and uint should be separated with a space; such as:1 memo")
)

// group's description, saved to Group-contract in json
type GDesc struct {
	Dc         uint8
	Pc         uint8
	StorePrice *big.Int
	ReadPrice  *big.Int
}

// provider's description, saved to Role-contract in json
type PDesc struct {
	Capacity uint64 // capacity that provider can provide, default 1TB
	Location string // indicate the type of the node by location(personal or cloud), default "personal"
}

// TODO: deploy contract on other chain
func GetInsEndPointByChain(chain string) (common.Address, string) {
	switch chain {
	case "dev":
		return InstanceAddr, DevChain
	case "test":
		return TestnetInstanceAddr, TestChain
	case "product":
		return ProductInstanceAddr, ProductChain
	case "goerli":
		return GoerliInstanceAddr, GoerliChain
	}
	return InstanceAddr, DevChain
}

// ParseMemo parse an amount string with units
func ParseMemo(memoAmount string) (*big.Int, error) {
	strSli := strings.Split(memoAmount, " ")
	if len(strSli) != 2 {
		return nil, ErrMemoAmount
	}

	num, err := strconv.ParseInt(strSli[0], 10, 0)
	if err != nil {
		return nil, ErrMemoAmount
	}

	value := big.NewInt(num)

	switch strSli[1] {
	case "memo":
		value.Mul(value, memo)
	case "nanomemo":
		value.Mul(value, nano)
	case "attomemo":
	default:
		return nil, ErrMemoAmount
	}
	return value, nil
}

func FormatMemo(value *big.Int) string {
	var m, n, a int64
	var mStr, nStr, aStr string

	tmp := big.NewInt(0)

	if value.Cmp(memo) >= 0 {
		m = tmp.Div(value, memo).Int64()
		mStr = strconv.Itoa(int(m)) + " memo "
		value.Mod(value, memo)
	}
	if value.Cmp(nano) >= 0 {
		n = tmp.Div(value, nano).Int64()
		nStr = strconv.Itoa(int(n)) + " nanomemo "
		value.Mod(value, nano)
	}
	a = value.Int64()
	if a != 0 {
		aStr = strconv.Itoa(int(a)) + " attomemo"
	}

	res := mStr + nStr + aStr
	if res == "" {
		return "0 attomemo"
	}
	return res
}
