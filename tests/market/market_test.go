// test for market contract
package main

import (
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/grid/contracts/eth"
	"github.com/grid/contracts/go/credit"
	"github.com/grid/contracts/go/market"
	"github.com/grid/contracts/go/registry"
)

var (
	// market contract addr
	MarketAddr common.Address
	// access contract address
	AccessAddr common.Address
	// credit contract address
	CreditAddr common.Address
	// registry contract address
	RegistryAddr common.Address
)

// load all addresses from json
func TestLoad(t *testing.T) {
	t.Log("test load addresses")

	// loading
	a := eth.Load("./contracts.json")
	t.Logf("%+v\n", a)

	if a.Market == "" || a.Access == "" || a.Credit == "" || a.Registry == "" {
		t.Error("all contract addresses must exist in json file")
	}

	MarketAddr = common.HexToAddress(a.Market)
	AccessAddr = common.HexToAddress(a.Access)
	CreditAddr = common.HexToAddress(a.Credit)
	RegistryAddr = common.HexToAddress(a.Registry)

}

// test create an order with test data
// gas:256945
func TestCreateOrder(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// get contract instance
	marketIns, err := market.NewMarket(MarketAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// get contract instance
	registryIns, err := registry.NewRegistry(RegistryAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// auth for admin
	authAdmin, err := eth.MakeAuth(chainID, eth.SK0)
	if err != nil {
		t.Error(err)
	}
	// auth for user
	authUser, err := eth.MakeAuth(chainID, eth.SK1)
	if err != nil {
		t.Error(err)
	}
	// auth for cp
	authProvider, err := eth.MakeAuth(chainID, eth.SK2)
	if err != nil {
		t.Error(err)
	}

	// register a cp for test createorder
	tx, err := registryIns.Register(authProvider, "123.123.123.0", "test domain", "123", 11, 22, 33, 44, 10, 20, 10, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("wait tx err:", err)
	}
	// check cp info
	cpInfo, err := registryIns.Get(&bind.CallOpts{}, eth.Addr2)
	if err != nil {
		t.Error("get cp info err:", err)
	}
	t.Log("cp info:", cpInfo)

	// generate an order with init data
	totalValue, ok := new(big.Int).SetString("2626954", 10)
	if !ok {
		log.Fatal("big set string failed")
	}
	remu, ok := new(big.Int).SetString("0", 10)
	if !ok {
		log.Fatal("big set string failed")
	}

	// a test order
	order := market.MarketOrder{
		User:     eth.Addr1,
		Provider: eth.Addr2,

		P: market.MarketPricePerHour{
			PCPU:  100,
			PGPU:  1000,
			PMEM:  10,
			PDISK: 1,
		},
		R: market.MarketResources{
			NCPU:  1,
			NGPU:  2,
			NMEM:  3,
			NDISK: 4,
		},
		// deposit 0.01 eth
		TotalValue:      totalValue,
		Remain:          totalValue,
		Remuneration:    remu,
		UserConfirm:     false,
		ProviderConfirm: false,
		ActivateTime:    new(big.Int).SetInt64(0),
		LastSettleTime:  new(big.Int).SetInt64(0),
		Probation:       new(big.Int).SetInt64(5),
		Duration:        new(big.Int).SetInt64(1231),
		Status:          0, // unactive
	}

	// credit contract
	creditIns, err := credit.NewCredit(CreditAddr, backend)
	if err != nil {
		t.Error(err)
	}
	// transfer some credit for approve
	t.Log("admin mint some credit to user for create order")
	tx, err = creditIns.Transfer(authAdmin, eth.Addr1, order.TotalValue)
	if err != nil {
		t.Error("transfer credit err:", err)
	}
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("wait tx err:", err)
	}

	// approve must be done by the user before create an order
	t.Log("user approving credit to market.., approve value: ", order.TotalValue)
	tx, err = creditIns.Approve(authUser, MarketAddr, order.TotalValue)
	if err != nil {
		t.Error(err)
	}
	// wait for tx to be ok
	t.Log("waiting tx")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	// check credit balance of market contract before create order
	b_before, err := creditIns.BalanceOf(&bind.CallOpts{}, MarketAddr)
	if err != nil {
		t.Error(err)
	}
	t.Log("market balance before create order:", b_before.String())

	t.Log("order to create: ", order)

	// create order
	t.Log("user call create order, order totalValue: ", order.TotalValue)
	tx, err = marketIns.CreateOrder(authUser, eth.Addr2, order)
	if err != nil {
		t.Error(err)
	}
	// wait for tx to be ok
	t.Log("waiting tx")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}
	receipt := eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	t.Log("create order gas used:", receipt.GasUsed)

	// cp info
	cpInfo, err = registryIns.Get(&bind.CallOpts{}, eth.Addr2)
	if err != nil {
		t.Error("get cp info err:", err)
	}
	t.Log("cp info:", cpInfo)

	// check updated avail resources in cp after createorder
	if cpInfo.Avail.NCPU != 10 {
		t.Fatalf("the available cpu is incorrect after createorder: %v, should be 10", cpInfo.Avail.NCPU)
	}
	if cpInfo.Avail.NGPU != 20 {
		t.Fatalf("the available gpu is incorrect after createorder: %v, should be 20", cpInfo.Avail.NGPU)
	}
	if cpInfo.Avail.NMEM != 30 {
		t.Fatalf("the available mem is incorrect after createorder: %v, should be 30", cpInfo.Avail.NMEM)
	}
	if cpInfo.Avail.NDISK != 40 {
		t.Fatalf("the available disk is incorrect after createorder: %v, should be 40", cpInfo.Avail.NDISK)
	}

	// check credit balance of market contract after create order
	b_after, err := creditIns.BalanceOf(&bind.CallOpts{}, MarketAddr)
	if err != nil {
		t.Error(err)
	}
	t.Log("market balance after create order:", b_after.String())

	delta := new(big.Int).Sub(b_after, b_before)
	t.Log("market balance increase:", delta)
	t.Log("order total value:", order.TotalValue)

	if delta.Cmp(order.TotalValue) != 0 {
		t.Error("the balance increament of market contract is incorrect")
	}
}

func TestGetKeys(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// get market instance
	marketIns, err := market.NewMarket(MarketAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// call
	keys, err := marketIns.GetKeys(&bind.CallOpts{})
	if err != nil {
		t.Error(err)
	}

	t.Log("keys:", keys)
}

func TestGetOrder(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// get market instance
	marketIns, err := market.NewMarket(MarketAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// call
	orderInfo, err := marketIns.GetOrder(&bind.CallOpts{From: eth.Addr1}, eth.Addr2)
	if err != nil {
		t.Error(err)
	}

	t.Log("order info:", orderInfo)
}

// test provider activate an order
// gas:112490
func TestActivate(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// get contract instance
	marketIns, err := market.NewMarket(MarketAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// make auth for sending transaction
	txAuth, err := eth.MakeAuth(chainID, eth.SK2)
	if err != nil {
		t.Error(err)
	}

	// provider call activate with user as param
	t.Log("provider activate an order")
	tx, err := marketIns.Activate(txAuth, eth.Addr1)
	if err != nil {
		t.Error(err)
	}

	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	receipt := eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	t.Log("activate order gas used:", receipt.GasUsed)

	// get order
	orderInfo, err := marketIns.GetOrder(&bind.CallOpts{From: eth.Addr1}, eth.Addr2)
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
// gas:55216 ~ 81422
func TestProSettle(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// get contract instance
	marketIns, err := market.NewMarket(MarketAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// make auth for provider
	txAuth, err := eth.MakeAuth(chainID, eth.SK2)
	if err != nil {
		t.Error(err)
	}

	// set gas for calling settle, in case the estimated gas is incorrect sometimes
	txAuth.GasLimit = 100000

	// get order before
	orderBefore, err := marketIns.GetOrder(&bind.CallOpts{From: eth.Addr1}, eth.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("order info:", orderBefore)

	// control nowtime to cross probation or not
	//time.Sleep(5 * time.Second)

	// call settle
	t.Log("provider settle an order, 1st ====>")
	tx, err := marketIns.ProSettle(txAuth, eth.Addr1)
	if err != nil {
		t.Error(err)
	}

	t.Log("waiting for tx to be ok, txHash:", tx.Hash())
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("call settle err:", err)
	}

	receipt := eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	t.Log("proSettle gas used:", receipt.GasUsed)

	// get order after
	orderAfter, err := marketIns.GetOrder(&bind.CallOpts{From: eth.Addr1}, eth.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("order info(after):", orderAfter)

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

	t.Log("--------------- waiting for the probation to be over...")
	// wait for the probation to be over
	//time.Sleep(5 * time.Second)

	// get order before
	orderBefore, err = marketIns.GetOrder(&bind.CallOpts{From: eth.Addr1}, eth.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("order info:", orderBefore)

	// call settle
	t.Log("provider settle an order, 2nd ====>")
	tx, err = marketIns.ProSettle(txAuth, eth.Addr1)
	if err != nil {
		t.Error(err)
	}

	t.Log("waiting for tx to be ok, tx hash:", tx.Hash())
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("call settle err:", err)
	}

	receipt = eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	t.Log("proSettle gas used:", receipt.GasUsed)

	// get order after
	orderAfter, err = marketIns.GetOrder(&bind.CallOpts{From: eth.Addr1}, eth.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("order info(after):", orderAfter)

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
// gas:89931
func TestProWithdraw(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	creditIns, err := credit.NewCredit(CreditAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// get old provider credit
	oldBal, err := creditIns.BalanceOf(&bind.CallOpts{}, eth.Addr2)
	if err != nil {
		t.Error("get provider balance failed:", err)
	}
	t.Log("provider old balance:", oldBal)

	// get contract instance
	marketIns, err := market.NewMarket(MarketAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// get order info before call
	orderBefore, err := marketIns.GetOrder(&bind.CallOpts{From: eth.Addr1}, eth.Addr2)
	if err != nil {
		t.Error(err)
	}

	// make auth for sending transaction
	txAuth, err := eth.MakeAuth(chainID, eth.SK2)
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
	tx, err := marketIns.ProWithdraw(txAuth, CreditAddr, eth.Addr1, amount)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("provider withdraw err:", err)
	}

	receipt := eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	t.Log("proWithdraw gas used:", receipt.GasUsed)

	// get order info after proWithdraw
	orderInfo, err := marketIns.GetOrder(&bind.CallOpts{From: eth.Addr1}, eth.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("order info:", orderInfo)

	// remueration decreasement
	remuDecre := new(big.Int).Sub(orderBefore.Remuneration, orderInfo.Remuneration)

	// check new provider credit balance
	newBal, err := creditIns.BalanceOf(&bind.CallOpts{From: eth.Addr2}, eth.Addr2)
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
// gas:105098
func TestUserCancel(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// get contract instance
	marketIns, err := market.NewMarket(MarketAddr, backend)
	if err != nil {
		t.Error("new market instance failed:", err)
	}

	// get credit instance
	creditIns, err := credit.NewCredit(CreditAddr, backend)
	if err != nil {
		t.Error("new token instance failed:", err)
	}

	// get balance of user before cancel order
	oldBal, err := creditIns.BalanceOf(&bind.CallOpts{From: eth.Addr1}, eth.Addr1)
	if err != nil {
		t.Error(err)
	}
	t.Log("old balance of user:", oldBal)

	// get order info before call user cancel
	orderBefore, err := marketIns.GetOrder(&bind.CallOpts{From: eth.Addr1}, eth.Addr2)
	if err != nil {
		t.Error(err)
	}

	// make auth for sending transaction
	txAuth, err := eth.MakeAuth(chainID, eth.SK1)
	if err != nil {
		t.Error(err)
	}

	// user cancel order, remain value is refunded to user
	t.Log("user cancels an order")
	tx, err := marketIns.UserCancel(txAuth, CreditAddr, eth.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("user cancel err:", err)
	}

	receipt := eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	t.Log("user cancel gas used:", receipt.GasUsed)

	// check status
	orderInfo, err := marketIns.GetOrder(&bind.CallOpts{From: eth.Addr1}, eth.Addr2)
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
	newBal, err := creditIns.BalanceOf(&bind.CallOpts{From: eth.Addr1}, eth.Addr1)
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
