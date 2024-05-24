// test for market contract
package main

import (
	"log"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	comm "github.com/grid/contracts/common"
	"github.com/grid/contracts/go/access"
	"github.com/grid/contracts/go/credit"
	"github.com/grid/contracts/go/market"
)

var (
	// access contract address
	accessAddr common.Address
	// credit contract address
	creditAddr common.Address
	// market contract addr
	marketAddr common.Address
)

func TestDeploy(t *testing.T) {
	t.Log("connecting to eth:", comm.Endpoint)

	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(comm.Endpoint)
	t.Log("chain id:", chainID)

	// make auth for sending transaction
	txAuth, err := comm.MakeAuth(chainID, comm.SK1)
	if err != nil {
		t.Error(err)
	}

	// deploy market contract
	t.Log("deploying market..")
	_marketAddr, tx, _, err := market.DeployMarket(txAuth, backend)
	if err != nil {
		t.Error("deploy registry err:", err)
	}
	marketAddr = _marketAddr
	t.Log("created market address: ", marketAddr.Hex())
	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(comm.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	receipt := comm.GetTransactionReceipt(comm.Endpoint, tx.Hash())
	t.Log("gas used:", receipt.GasUsed)

	// deploy access contract
	t.Log("deploying access")
	_accessAddr, tx, accessIns, err := access.DeployAccess(txAuth, backend)
	if err != nil {
		t.Error("deploy access err:", err)
	}
	accessAddr = _accessAddr
	t.Log("created access address: ", accessAddr.Hex())
	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(comm.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	// set access for admin
	t.Log("set access for admin")
	tx, err = accessIns.Set(txAuth, comm.Addr1, true)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(comm.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	// deploy credit contract
	t.Log("deploying credit contract")
	_creditAddr, tx, _, err := credit.DeployCredit(txAuth, backend, accessAddr)
	if err != nil {
		t.Error("deploy credit err:", err)
	}
	creditAddr = _creditAddr
	t.Log("credit address:", creditAddr)
	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(comm.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

}

func TestCreateOrder(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(comm.Endpoint)
	t.Log("chain id:", chainID)

	// get contract instance
	marketIns, err := market.NewMarket(marketAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// make auth for sending transaction
	txAuth, err := comm.MakeAuth(chainID, comm.SK1)
	if err != nil {
		t.Error(err)
	}

	// generate an order with init data
	totalValue, ok := new(big.Int).SetString("2831300", 10)
	if !ok {
		log.Fatal("big set string failed")
	}
	remu, ok := new(big.Int).SetString("0", 10)
	if !ok {
		log.Fatal("big set string failed")
	}
	order := market.MarketOrder{
		User:     comm.Addr1,
		Provider: comm.Addr2,

		P: market.MarketPricePerHour{
			PCPU:  100,
			PGPU:  1000,
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
		Duration:        new(big.Int).SetInt64(1231),
		Status:          0, // unactive
	}

	creditIns, err := credit.NewCredit(creditAddr, backend)
	if err != nil {
		t.Error(err)
	}
	// mint some credit for approve
	t.Log("admin mint credit to user")
	tx, err := creditIns.Mint(txAuth, comm.Addr1, order.TotalValue)
	if err != nil {
		t.Error("mint credit err:", err)
	}
	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(comm.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("wait tx err:", err)
	}

	// approve must be done by the user before create an order
	t.Log("user approving credit to market..")
	tx, err = creditIns.Approve(txAuth, marketAddr, order.TotalValue)
	if err != nil {
		t.Error(err)
	}
	// wait for tx to be ok
	t.Log("waiting tx")
	err = comm.CheckTx(comm.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	// create order
	t.Log("call create order")
	tx, err = marketIns.CreateOrder(txAuth, creditAddr, comm.Addr2, order)
	if err != nil {
		t.Error(err)
	}
	// wait for tx to be ok
	t.Log("waiting tx")
	err = comm.CheckTx(comm.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	receipt := comm.GetTransactionReceipt(comm.Endpoint, tx.Hash())
	t.Log("gas used:", receipt.GasUsed)

	// check credit balance of market contract after create order
	b, err := creditIns.BalanceOf(&bind.CallOpts{}, marketAddr)
	if err != nil {
		t.Error(err)
	}
	t.Log("market balance after create order:", b.String())
	if b.Cmp(order.Remain) != 0 {
		t.Error("the balance of market contract is incorrect")
	}
}

func TestGetOrder(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(comm.Endpoint)
	t.Log("chain id:", chainID)

	// get market instance
	marketIns, err := market.NewMarket(marketAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// call
	orderInfo, err := marketIns.GetOrder(&bind.CallOpts{From: comm.Addr1}, comm.Addr2)
	if err != nil {
		t.Error(err)
	}

	t.Log("order info:", orderInfo)
}

// test provider activate an order
func TestActivate(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(comm.Endpoint)
	t.Log("chain id:", chainID)

	// get contract instance
	marketIns, err := market.NewMarket(marketAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// make auth for sending transaction
	txAuth, err := comm.MakeAuth(chainID, comm.SK2)
	if err != nil {
		t.Error(err)
	}

	// provider call activate with user as param
	t.Log("provider activate an order")
	tx, err := marketIns.Activate(txAuth, comm.Addr1)
	if err != nil {
		t.Error(err)
	}

	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(comm.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	// get order
	orderInfo, err := marketIns.GetOrder(&bind.CallOpts{From: comm.Addr1}, comm.Addr2)
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

// test provider settle
func TestProSettle(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(comm.Endpoint)
	t.Log("chain id:", chainID)

	// get contract instance
	marketIns, err := market.NewMarket(marketAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// make auth for provider
	txAuth, err := comm.MakeAuth(chainID, comm.SK2)
	if err != nil {
		t.Error(err)
	}

	// get order before
	orderBefore, err := marketIns.GetOrder(&bind.CallOpts{From: comm.Addr1}, comm.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("order info:", orderBefore)

	// call settle
	t.Log("provider settle an order")
	tx, err := marketIns.ProSettle(txAuth, comm.Addr1)
	if err != nil {
		t.Error(err)
	}

	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(comm.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("call settle err:", err)
	}

	// get order after
	orderAfter, err := marketIns.GetOrder(&bind.CallOpts{From: comm.Addr1}, comm.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("order info:", orderAfter)

	nowTime := new(big.Int).Set(orderAfter.LastSettleTime)
	probationTime := new(big.Int).Add(orderAfter.ActivateTime, orderAfter.Probation)
	lastSettleTime := new(big.Int).Set(orderBefore.LastSettleTime)
	payTime := new(big.Int)
	t.Logf("nowTime: %v, probationTime: %v, lastSettle: %v, payTime: %v", nowTime, probationTime, lastSettleTime, payTime)

	// calc the time payTime(the same as market contract)
	// lastSettleTime <= probation time
	if lastSettleTime.Cmp(probationTime) <= 0 {
		if nowTime.Cmp(probationTime) > 0 {
			payTime = payTime.Sub(nowTime, probationTime)
		}
	} else {
		payTime = payTime.Sub(nowTime, lastSettleTime)
	}

	t.Log("time to pay:", payTime)
	unitFee := new(big.Int).Div(orderAfter.TotalValue, orderAfter.Duration)
	// calc the pay fee and remueration
	fee := new(big.Int).Mul(unitFee, payTime)
	t.Log("the fee to pay during this settlement: ", fee)

	// new remu should be: remuBefore + fee
	remu := new(big.Int).Add(orderBefore.Remuneration, fee)

	// check new remueration after provider settle
	if orderAfter.Remuneration.Cmp(remu) != 0 {
		t.Errorf("remuneration is incorrect: %v, should be %v", orderAfter.Remuneration, remu)
	}

	// -------------------

	t.Log("waiting for the probation to be over, 10s...")
	// wait for the probation to be over
	time.Sleep(10 * time.Second)

	// get order before
	orderBefore, err = marketIns.GetOrder(&bind.CallOpts{From: comm.Addr1}, comm.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("order info:", orderBefore)

	// call settle
	t.Log("provider settle an order")
	tx, err = marketIns.ProSettle(txAuth, comm.Addr1)
	if err != nil {
		t.Error(err)
	}

	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(comm.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("call settle err:", err)
	}

	// get order after
	orderAfter, err = marketIns.GetOrder(&bind.CallOpts{From: comm.Addr1}, comm.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("order info:", orderAfter)

	nowTime = new(big.Int).Set(orderAfter.LastSettleTime)
	probationTime = new(big.Int).Add(orderAfter.ActivateTime, orderAfter.Probation)
	lastSettleTime = new(big.Int).Set(orderBefore.LastSettleTime)
	payTime = new(big.Int)
	t.Logf("nowTime: %v, probationTime: %v, lastSettle: %v, payTime: %v", nowTime, probationTime, lastSettleTime, payTime)

	// calc the time payTime(the same as market contract)
	// lastSettleTime <= probation time
	if lastSettleTime.Cmp(probationTime) <= 0 {
		if nowTime.Cmp(probationTime) > 0 {
			payTime = payTime.Sub(nowTime, probationTime)
		}
	} else {
		payTime = payTime.Sub(nowTime, lastSettleTime)
	}

	t.Log("time to pay:", payTime)
	unitFee = new(big.Int).Div(orderAfter.TotalValue, orderAfter.Duration)
	// calc the pay fee and remueration
	fee = new(big.Int).Mul(unitFee, payTime)
	t.Log("the fee to pay during this settlement: ", fee)

	// new remu should be: remuBefore + fee
	remu = new(big.Int).Add(orderBefore.Remuneration, fee)

	// check new remueration after provider settle
	if orderAfter.Remuneration.Cmp(remu) != 0 {
		t.Errorf("remuneration is incorrect: %v, should be %v", orderAfter.Remuneration, remu)
	}

}

// test provider withdraw from remuneration
func TestProWithdraw(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(comm.Endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	creditIns, err := credit.NewCredit(creditAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// get old provider credit
	oldBal, err := creditIns.BalanceOf(&bind.CallOpts{}, comm.Addr2)
	if err != nil {
		t.Error("get provider balance failed:", err)
	}
	t.Log("provider old balance:", oldBal)

	// get contract instance
	marketIns, err := market.NewMarket(marketAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// get order info before call
	orderBefore, err := marketIns.GetOrder(&bind.CallOpts{From: comm.Addr1}, comm.Addr2)
	if err != nil {
		t.Error(err)
	}

	// make auth for sending transaction
	txAuth, err := comm.MakeAuth(chainID, comm.SK2)
	if err != nil {
		t.Error(err)
	}

	// provider withdraw some token from remuneration
	amount, ok := new(big.Int).SetString("2124", 10)
	if !ok {
		t.Error("set string error")
	}
	// provider call withdraw with user as param
	t.Log("provider withdraw")
	tx, err := marketIns.ProWithdraw(txAuth, creditAddr, comm.Addr1, amount)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(comm.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("user withdraw err:", err)
	}

	// get order info after proWithdraw
	orderInfo, err := marketIns.GetOrder(&bind.CallOpts{From: comm.Addr1}, comm.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("order info:", orderInfo)

	// remueration decreasement
	remuDecre := new(big.Int).Sub(orderBefore.Remuneration, orderInfo.Remuneration)

	// check new provider credit balance
	newBal, err := creditIns.BalanceOf(&bind.CallOpts{From: comm.Addr2}, comm.Addr2)
	if err != nil {
		t.Error("get credit balance of provider error")
	}
	t.Log("new provider balance:", newBal)

	incre := new(big.Int).Sub(newBal, oldBal)
	t.Log("balance increased: ", incre)
	t.Log("withdraw amount: ", amount)
	t.Log("remueration decreased: ", remuDecre)

	// check remueration decreasement
	if remuDecre.Cmp(amount) != 0 {
		t.Errorf("remueration decreasement is incorrect: %v, should be: %v", remuDecre, amount)
	}
	// check balance increasement
	if incre.Cmp(amount) != 0 {
		if err != nil {
			t.Errorf("provider balance incorrect: %v, should be: %v", newBal, amount)
		}
	}
}

// test user cancel order
func TestUserCancel(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(comm.Endpoint)
	t.Log("chain id:", chainID)

	// get contract instance
	marketIns, err := market.NewMarket(marketAddr, backend)
	if err != nil {
		t.Error("new market instance failed:", err)
	}

	// get credit instance
	creditIns, err := credit.NewCredit(creditAddr, backend)
	if err != nil {
		t.Error("new token instance failed:", err)
	}

	// get balance of user before cancel order
	oldBal, err := creditIns.BalanceOf(&bind.CallOpts{From: comm.Addr1}, comm.Addr1)
	if err != nil {
		t.Error(err)
	}
	t.Log("old balance of user:", oldBal)

	// get order info before call user cancel
	orderBefore, err := marketIns.GetOrder(&bind.CallOpts{From: comm.Addr1}, comm.Addr2)
	if err != nil {
		t.Error(err)
	}

	// make auth for sending transaction
	txAuth, err := comm.MakeAuth(chainID, comm.SK1)
	if err != nil {
		t.Error(err)
	}

	// user cancel order, remain value is refunded to user
	t.Log("user cancels an order")
	tx, err := marketIns.UserCancel(txAuth, creditAddr, comm.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(comm.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	// check status
	orderInfo, err := marketIns.GetOrder(&bind.CallOpts{From: comm.Addr1}, comm.Addr2)
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
	newBal, err := creditIns.BalanceOf(&bind.CallOpts{From: comm.Addr1}, comm.Addr1)
	if err != nil {
		t.Error(err)
	}
	t.Log("new balance of user:", newBal)

	// calc the time elapsed
	elapsed := new(big.Int).Sub(orderInfo.LastSettleTime, orderBefore.LastSettleTime)
	t.Log("time elapsed:", elapsed)
	unitFee := new(big.Int).Div(orderInfo.TotalValue, orderInfo.Duration)
	// calc the fee and remain
	fee := new(big.Int).Mul(unitFee, elapsed)
	remain := new(big.Int).Sub(orderBefore.Remain, fee)

	// the user refund after settled
	//remain, _ := new(big.Int).SetString("9000000000000000", 10)
	incre := newBal.Sub(newBal, oldBal)
	if incre.Cmp(remain) != 0 {
		t.Errorf("the increment of user balance is error: %v, should be %v", incre, remain)
	}

}
