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
	contractAddr common.Address
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
	_contractAddr, tx, _, err := market.DeployMarket(txAuth, backend)
	if err != nil {
		t.Error("deploy registry err:", err)
	}
	contractAddr = _contractAddr
	t.Log("created registry address: ", contractAddr.Hex())
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
	contractIns, err := market.NewMarket(contractAddr, backend)
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
	_tokenAddr, tx, tokenIns, err := gtoken.DeployGToken(txAuth, backend)
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
	tx, err = tokenIns.Mint(txAuth, Addr1, order.Remain)
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
	tx, err = tokenIns.Approve(txAuth, contractAddr, order.Remain)
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
	tx, err = contractIns.CreateOrder(txAuth, tokenAddr, Addr2, order)
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
	b, err := tokenIns.BalanceOf(&bind.CallOpts{}, contractAddr)
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
	contractIns, err := market.NewMarket(contractAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// call
	orderInfo, err := contractIns.GetOrder(&bind.CallOpts{From: Addr1}, Addr2)
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
	contractIns, err := market.NewMarket(contractAddr, backend)
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
	tx, err := contractIns.Activate(txAuth, Addr1)
	if err != nil {
		t.Error(err)
	}

	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	// get order
	orderInfo, err := contractIns.GetOrder(&bind.CallOpts{From: Addr1}, Addr2)
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
func TestCancel(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// get contract instance
	contractIns, err := market.NewMarket(contractAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// make auth for sending transaction
	txAuth, err := comm.MakeAuth(chainID, sk1)
	if err != nil {
		t.Error(err)
	}

	// provider call activate with user as param
	t.Log("user cancels an order")
	tx, err := contractIns.UserCancel(txAuth, Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	// get order
	orderInfo, err := contractIns.GetOrder(&bind.CallOpts{From: Addr1}, Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("order info:", orderInfo)
	// check order status
	if orderInfo.Status != 2 {
		t.Error("cancel failed, status not 2")
	}
}

// test user withdraw
func TestUserWithdraw(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	tokenIns, err := gtoken.NewGToken(tokenAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// get old user token
	oldBal, err := tokenIns.BalanceOf(&bind.CallOpts{From: Addr1}, Addr1)
	if err != nil {
		t.Error("get user balance failed:", err)
	}
	t.Log("user old balance:", oldBal)

	// get contract instance
	contractIns, err := market.NewMarket(contractAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// make auth for sending transaction
	txAuth, err := comm.MakeAuth(chainID, sk1)
	if err != nil {
		t.Error(err)
	}

	// withdraw some token
	amount, ok := new(big.Int).SetString("3000000000000000", 10)
	if !ok {
		t.Error("set string error")
	}
	// provider call activate with user as param
	t.Log("user withdraw")
	tx, err := contractIns.UserWithdraw(txAuth, tokenAddr, Addr2, amount)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("user withdraw err:", err)
	}

	// get order
	orderInfo, err := contractIns.GetOrder(&bind.CallOpts{From: Addr1}, Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("order info:", orderInfo)

	// check new user token balance
	newBal, err := tokenIns.BalanceOf(&bind.CallOpts{From: Addr1}, Addr1)
	if err != nil {
		t.Error("get balance of user error")
	}
	t.Log("new user balance:", newBal)
}

// test provider withdraw from remuneration
func TestProWithdraw(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	tokenIns, err := gtoken.NewGToken(tokenAddr, backend)
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
	contractIns, err := market.NewMarket(contractAddr, backend)
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
	tx, err := contractIns.ProWithdraw(txAuth, tokenAddr, Addr1, amount)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("user withdraw err:", err)
	}

	// get order
	orderInfo, err := contractIns.GetOrder(&bind.CallOpts{From: Addr1}, Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("order info:", orderInfo)

	// check new provider token balance
	newBal, err := tokenIns.BalanceOf(&bind.CallOpts{From: Addr2}, Addr2)
	if err != nil {
		t.Error("get balance of provider error")
	}
	t.Log("new provider balance:", newBal)
}

/*
// test settle
func TestSettle(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// get contract instance
	contractIns, err := market.NewMarket(contractAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// make auth for sending transaction
	txAuth, err := comm.MakeAuth(chainID, sk1)
	if err != nil {
		t.Error(err)
	}

	// get order
	orderInfo, err := contractIns.GetOrder(&bind.CallOpts{From: Addr1}, Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("order info:", orderInfo)

	// call settle
	t.Log("settle an order")
	tx, err := contractIns.Settle(txAuth, Addr1, Addr2, orderInfo.LastSettleTime.Add(orderInfo.LastSettleTime, new(big.Int).SetUint64(10)))
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
