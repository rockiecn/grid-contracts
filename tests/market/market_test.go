// test for market contract
package main

import (
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	comm "github.com/grid/contracts/common"
	"github.com/grid/contracts/go/gtoken"
	"github.com/grid/contracts/go/market"
)

var (
	// endpoint for ganache
	endpoint string = "HTTP://127.0.0.1:7545"

	// acc1 of ganache
	addr1 string = "0x5F7F7e31399531F08C2b47eA1919F11346405a16"
	// acc2 of ganache
	addr2 string = "0xe2198eb2e931f9306ABcA68D4F093E0Ac4823B0d"
	Addr1        = common.HexToAddress(addr1)
	Addr2        = common.HexToAddress(addr2)

	// use acc1 as the admin to send tx
	sk1 string = "c1e763d955e6aea410e40b95702108a30efb4d25b32d419910fe2ac611c2229d"
	sk2 string = "e8cda8fe7c04afa4a0630af457972f88a645468cb90120a11911669deac5e96e"

	// erc20 token address
	tokenAddr common.Address
	// market contract addr
	marketAddr common.Address
)

func TestDeploy(t *testing.T) {
	t.Log("connecting to eth:", endpoint)

	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// make auth for sending transaction
	txAuth, err := comm.MakeAuth(chainID, sk1)
	if err != nil {
		t.Error(err)
	}

	// deploy market contract
	_marketAddr, tx, _, err := market.DeployMarket(txAuth, backend)
	if err != nil {
		t.Error("deploy registry err:", err)
	}
	marketAddr = _marketAddr
	t.Log("created market address: ", marketAddr.Hex())
	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	receipt := comm.GetTransactionReceipt(endpoint, tx.Hash())
	t.Log("gas used:", receipt.GasUsed)

}

func TestCreateOrder(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// get contract instance
	marketIns, err := market.NewMarket(marketAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// make auth for sending transaction
	txAuth, err := comm.MakeAuth(chainID, sk1)
	if err != nil {
		t.Error(err)
	}

	// deploy token contract
	t.Log("deploy token contract")
	_tokenAddr, tx, tokenIns, err := gtoken.DeployGtoken(txAuth, backend)
	if err != nil {
		t.Error("deploy token err:", err)
	}
	tokenAddr = _tokenAddr
	t.Log("token address:", tokenAddr)
	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	// generate an order with init data
	totalValue, ok := new(big.Int).SetString("10000000000000000", 10)
	if !ok {
		log.Fatal("big set string failed")
	}
	remu, ok := new(big.Int).SetString("0", 10)
	if !ok {
		log.Fatal("big set string failed")
	}
	order := market.MarketOrder{
		User:     Addr1,
		Provider: Addr2,

		P: market.MarketPricePerHour{
			PCPU:  1,
			PGPU:  10,
			PMEM:  10,
			PDISK: 1,
		},
		R: market.MarketResources{
			NCPU:  1,
			NGPU:  2,
			NMEM:  10,
			NDISK: 100,
		},
		// deposit 0.01 eth
		TotalValue:      totalValue,
		Remain:          totalValue,
		Remuneration:    remu,
		UserConfirm:     false,
		ProviderConfirm: false,
		ActivateTime:    new(big.Int).SetInt64(0),
		LastSettleTime:  new(big.Int).SetInt64(0),
		Probation:       new(big.Int).SetInt64(10),
		Duration:        new(big.Int).SetInt64(100),
		Status:          0, // unactive
	}

	// mint some token for approve
	t.Log("mint token to user")
	tx, err = tokenIns.Mint(txAuth, Addr1, order.TotalValue)
	if err != nil {
		t.Error("mint token err:", err)
	}
	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("wait tx err:", err)
	}

	// approve must be done by the user before create an order
	t.Log("user approving..")
	tx, err = tokenIns.Approve(txAuth, marketAddr, order.TotalValue)
	if err != nil {
		t.Error(err)
	}
	// wait for tx to be ok
	t.Log("waiting tx")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	// create order
	t.Log("call create order")
	tx, err = marketIns.CreateOrder(txAuth, tokenAddr, Addr2, order)
	if err != nil {
		t.Error(err)
	}
	// wait for tx to be ok
	t.Log("waiting tx")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	receipt := comm.GetTransactionReceipt(endpoint, tx.Hash())
	t.Log("gas used:", receipt.GasUsed)

	// check balance of market contract after create order
	b, err := tokenIns.BalanceOf(&bind.CallOpts{}, marketAddr)
	if err != nil {
		t.Error(err)
	}
	t.Log("market balance after create order:", b.String())
	if b.Cmp(order.Remain) != 0 {
		t.Error("the balance of market contract is error")
	}
}

func TestGetOrder(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// get market instance
	marketIns, err := market.NewMarket(marketAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// call
	orderInfo, err := marketIns.GetOrder(&bind.CallOpts{From: Addr1}, Addr2)
	if err != nil {
		t.Error(err)
	}

	t.Log("order info:", orderInfo)

	if orderInfo.P.PCPU != 1 {
		t.Error("pcpu error")
	}
	if orderInfo.R.NCPU != 1 {
		t.Error("ncpu error")
	}
	remain, _ := new(big.Int).SetString("10000000000000000", 10)
	if orderInfo.Remain.Cmp(remain) != 0 {
		t.Error("deposit erro")
	}
	if orderInfo.UserConfirm != false {
		t.Error("user confirm error")
	}
	if orderInfo.ActivateTime.Cmp(new(big.Int).SetInt64(0)) != 0 {
		t.Error("active time error")
	}
	if orderInfo.Status != 0 {
		t.Error("status error")
	}
}

// test provider activate an order
func TestActivate(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// get contract instance
	marketIns, err := market.NewMarket(marketAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// make auth for sending transaction
	txAuth, err := comm.MakeAuth(chainID, sk2)
	if err != nil {
		t.Error(err)
	}

	// provider call activate with user as param
	t.Log("provider activate an order")
	tx, err := marketIns.Activate(txAuth, Addr1)
	if err != nil {
		t.Error(err)
	}

	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	// get order
	orderInfo, err := marketIns.GetOrder(&bind.CallOpts{From: Addr1}, Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("order status:", orderInfo.Status)
	t.Log("activate timestamp:", orderInfo.ActivateTime.String())
	// check order status
	if orderInfo.Status != 1 {
		t.Error("activate failed, status not 1")
	}
}

// test user cancel order
func TestUserCancel(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// get contract instance
	marketIns, err := market.NewMarket(marketAddr, backend)
	if err != nil {
		t.Error("new market instance failed:", err)
	}

	// get token instance
	tokenIns, err := gtoken.NewGtoken(tokenAddr, backend)
	if err != nil {
		t.Error("new token instance failed:", err)
	}

	// get balance of user before cancel order
	oldBal, err := tokenIns.BalanceOf(&bind.CallOpts{From: Addr1}, Addr1)
	if err != nil {
		t.Error(err)
	}
	t.Log("old balance of user:", oldBal)

	// make auth for sending transaction
	txAuth, err := comm.MakeAuth(chainID, sk1)
	if err != nil {
		t.Error(err)
	}

	// user cancel order, remain value is refunded to user
	t.Log("user cancels an order")
	tx, err := marketIns.UserCancel(txAuth, tokenAddr, Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	// check status
	orderInfo, err := marketIns.GetOrder(&bind.CallOpts{From: Addr1}, Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("order info:", orderInfo)
	// check order status
	if orderInfo.Status != 2 {
		t.Error("cancel failed, status not 2")
	}

	t.Log("new remain value:", orderInfo.Remain)
	// check remain after order cancelled
	if orderInfo.Remain.Cmp(new(big.Int).SetUint64(0)) != 0 {
		t.Error("the remain token in this order must be 0 after cancelled")
	}

	// check user token balance
	newBal, err := tokenIns.BalanceOf(&bind.CallOpts{From: Addr1}, Addr1)
	if err != nil {
		t.Error(err)
	}
	t.Log("new balance of user:", newBal)

	// the remain value after settled
	remain, _ := new(big.Int).SetString("9000000000000000", 10)
	incre := newBal.Sub(newBal, oldBal)
	if incre.Cmp(remain) != 0 {
		t.Error("the increment of user balance is error, should equal to order remain value")
	}

}

// test provider withdraw from remuneration
func TestProWithdraw(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	tokenIns, err := gtoken.NewGtoken(tokenAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// get old provider token
	oldBal, err := tokenIns.BalanceOf(&bind.CallOpts{From: Addr1}, Addr2)
	if err != nil {
		t.Error("get provider balance failed:", err)
	}
	t.Log("provider old balance:", oldBal)

	// get contract instance
	marketIns, err := market.NewMarket(marketAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// make auth for sending transaction
	txAuth, err := comm.MakeAuth(chainID, sk2)
	if err != nil {
		t.Error(err)
	}

	// withdraw some token
	amount, ok := new(big.Int).SetString("1000000000000000", 10)
	if !ok {
		t.Error("set string error")
	}
	// provider call activate with user as param
	t.Log("provider withdraw")
	tx, err := marketIns.ProWithdraw(txAuth, tokenAddr, Addr1, amount)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("user withdraw err:", err)
	}

	// get order
	orderInfo, err := marketIns.GetOrder(&bind.CallOpts{From: Addr1}, Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("order info:", orderInfo)

	// check remuneration value
	remu, ok := new(big.Int).SetString("0", 10)
	if !ok {
		t.Error("set string error")
	}
	if orderInfo.Remuneration.Cmp(remu) != 0 {
		if err != nil {
			t.Error("remuneration value error")
		}
	}

	// check new provider token balance
	newBal, err := tokenIns.BalanceOf(&bind.CallOpts{From: Addr2}, Addr2)
	if err != nil {
		t.Error("get balance of provider error")
	}
	t.Log("new provider balance:", newBal)

	// check balance
	b, ok := new(big.Int).SetString("1000000000000000", 10)
	if !ok {
		t.Error("set string error")
	}
	if newBal.Cmp(b) != 0 {
		if err != nil {
			t.Error("provider balance error")
		}
	}
}

/*
// test settle
func TestSettle(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// get contract instance
	marketIns, err := market.NewMarket(marketAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// make auth for sending transaction
	txAuth, err := comm.MakeAuth(chainID, sk1)
	if err != nil {
		t.Error(err)
	}

	// get order
	orderInfo, err := marketIns.GetOrder(&bind.CallOpts{From: Addr1}, Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("order info:", orderInfo)

	// call settle
	t.Log("settle an order")
	tx, err := marketIns.Settle(txAuth, Addr1, Addr2, orderInfo.LastSettleTime.Add(orderInfo.LastSettleTime, new(big.Int).SetUint64(10)))
	if err != nil {
		t.Error(err)
	}

	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("call settle err:", err)
	}

}
*/
