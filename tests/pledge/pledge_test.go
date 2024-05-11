package main

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	comm "github.com/grid/contracts/common"
	"github.com/grid/contracts/go/gtoken"
	"github.com/grid/contracts/go/pledge"
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

	// the erc20 token address
	tokenAddr common.Address
	// the pledge contract addr
	pledgeAddr common.Address
)

// deploy
// pledge
// unpledge
// withdrawInterest
// setRate

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

	// deploy pledge contract, acc1 as the owner
	_pledgeAddr, tx, _, err := pledge.DeployPledge(txAuth, backend)
	if err != nil {
		t.Error("deploy registry err:", err)
	}
	pledgeAddr = _pledgeAddr
	t.Log("created pledge address: ", pledgeAddr.Hex())
	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	receipt := comm.GetTransactionReceipt(endpoint, tx.Hash())
	t.Log("gas used:", receipt.GasUsed)

	// deploy token contract
	t.Log("deploy token contract")
	_tokenAddr, tx, _, err := gtoken.DeployGtoken(txAuth, backend)
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
}

func TestPledge(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// get contract instance
	pledgeIns, err := pledge.NewPledge(pledgeAddr, backend)
	if err != nil {
		t.Error("new pledge contract instance failed:", err)
	}
	tokenIns, err := gtoken.NewGtoken(tokenAddr, backend)
	if err != nil {
		t.Error("new token contract instance failed:", err)
	}

	// make auth for sending transaction
	txAuth1, err := comm.MakeAuth(chainID, sk1)
	if err != nil {
		t.Error(err)
	}

	txAuth2, err := comm.MakeAuth(chainID, sk2)
	if err != nil {
		t.Error(err)
	}

	// 0.01 eth to pledge
	amount, ok := new(big.Int).SetString("10000000000000000", 10)
	if !ok {
		t.Error("set string failed")
	}

	// mint some token for approve
	t.Log("mint token to user")
	tx, err := tokenIns.Mint(txAuth1, Addr2, amount)
	if err != nil {
		t.Error("mint token err:", err)
	}
	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("wait tx err:", err)
	}

	// acc2 approve to pledge contract before pledge
	t.Log("user approving..")
	tx, err = tokenIns.Approve(txAuth2, pledgeAddr, amount)
	if err != nil {
		t.Error(err)
	}
	// wait tx
	t.Log("waiting tx..")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("approve err:", err)
	}

	// for address2, check old value
	opts2 := &bind.CallOpts{From: Addr2}

	oldPledge, err := pledgeIns.GetPledge(opts2)
	if err != nil {
		t.Error(err)
	}
	oldBalance, err := tokenIns.BalanceOf(opts2, Addr2)
	if err != nil {
		t.Error(err)
	}

	t.Log("old pledge:", oldPledge)
	t.Log("old balance:", oldBalance)

	// acc2 pledge amount
	t.Log("user pledging..")
	tx, err = pledgeIns.Pledge(txAuth2, tokenAddr, amount)
	if err != nil {
		t.Error(err)
	}
	// wait tx
	t.Log("waiting tx..")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("pledge err:", err)
	}

	newPledge, err := pledgeIns.GetPledge(opts2)
	if err != nil {
		t.Error(err)
	}
	newBalance, err := tokenIns.BalanceOf(opts2, Addr2)
	if err != nil {
		t.Error(err)
	}

	t.Log("new pledge:", newPledge)
	t.Log("new balance:", newBalance)

	// calc pledge increament
	delta := newPledge.Sub(newPledge, oldPledge)
	// check the pledge amount
	if delta.Cmp(amount) != 0 {
		t.Error("pledge amount is incorrect")
	}

	// calc balance decreament
	deltaBal := newBalance.Sub(oldBalance, newBalance)
	// check the balance
	if deltaBal.Cmp(amount) != 0 {
		t.Error("pledge balance is incorrect")
	}
}

func TestUnPledge(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// get contract instance
	pledgeIns, err := pledge.NewPledge(pledgeAddr, backend)
	if err != nil {
		t.Error("new pledge contract instance failed:", err)
	}
	tokenIns, err := gtoken.NewGtoken(tokenAddr, backend)
	if err != nil {
		t.Error("new token contract instance failed:", err)
	}

	// make auth for sending transaction
	txAuth2, err := comm.MakeAuth(chainID, sk2)
	if err != nil {
		t.Error(err)
	}

	// for address2
	opts2 := &bind.CallOpts{From: Addr2}

	oldPledge, err := pledgeIns.GetPledge(opts2)
	if err != nil {
		t.Error(err)
	}
	oldBalance, err := tokenIns.BalanceOf(opts2, Addr2)
	if err != nil {
		t.Error(err)
	}

	t.Log("old pledge:", oldPledge)
	t.Log("old balance:", oldBalance)

	// unpledge 0.002 eth
	amount, ok := new(big.Int).SetString("2000000000000000", 10)
	if !ok {
		t.Error("set string failed")
	}

	// acc2 pledge amount
	t.Log("user unpledging..")
	tx, err := pledgeIns.Unpledge(txAuth2, tokenAddr, amount)
	if err != nil {
		t.Error(err)
	}
	// wait tx
	t.Log("waiting tx..")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("pledge err:", err)
	}

	newPledge, err := pledgeIns.GetPledge(opts2)
	if err != nil {
		t.Error(err)
	}
	newBalance, err := tokenIns.BalanceOf(opts2, Addr2)
	if err != nil {
		t.Error(err)
	}

	t.Log("new pledge:", newPledge)
	t.Log("new balance:", newBalance)

	// calc decreament
	delta := newPledge.Sub(oldPledge, newPledge)
	// check the pledge amount
	if delta.Cmp(amount) != 0 {
		t.Error("unpledge amount is incorrect")
	}

	// calc balance increament
	deltaBal := newBalance.Sub(newBalance, oldBalance)
	// check the balance
	if deltaBal.Cmp(amount) != 0 {
		t.Error("unpledge balance is incorrect")
	}
}

func TestWithdrawInterest(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// get contract instance
	pledgeIns, err := pledge.NewPledge(pledgeAddr, backend)
	if err != nil {
		t.Error("new pledge contract instance failed:", err)
	}
	tokenIns, err := gtoken.NewGtoken(tokenAddr, backend)
	if err != nil {
		t.Error("new token contract instance failed:", err)
	}

	// make auth for sending transaction
	txAuth2, err := comm.MakeAuth(chainID, sk2)
	if err != nil {
		t.Error(err)
	}

	// for address2
	opts2 := &bind.CallOpts{From: Addr2}

	oldInterest, err := pledgeIns.GetInterest(opts2)
	if err != nil {
		t.Error(err)
	}
	oldBalance, err := tokenIns.BalanceOf(opts2, Addr2)
	if err != nil {
		t.Error(err)
	}

	t.Log("old interest:", oldInterest)
	t.Log("old balance:", oldBalance)

	// update interest
	tx, err := pledgeIns.UpdateInterest(txAuth2)
	if err != nil {
		t.Error(err)
	}
	// wait tx
	t.Log("waiting tx..")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("update err:", err)
	}

	// withdraw amount
	amount, ok := new(big.Int).SetString("500", 10)
	if !ok {
		t.Error("set string failed")
	}

	// acc2 withdraw interest
	t.Log("user withdrawing..")
	tx, err = pledgeIns.WithdrawInterest(txAuth2, tokenAddr, amount)
	if err != nil {
		t.Error(err)
	}
	// wait tx
	t.Log("waiting tx..")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("withdraw err:", err)
	}

	newInterest, err := pledgeIns.GetInterest(opts2)
	if err != nil {
		t.Error(err)
	}
	newBalance, err := tokenIns.BalanceOf(opts2, Addr2)
	if err != nil {
		t.Error(err)
	}

	t.Log("new interest:", newInterest)
	t.Log("new balance:", newBalance)

	// calc balance increament
	deltaBal := newBalance.Sub(newBalance, oldBalance)
	// check the balance
	if deltaBal.Cmp(amount) != 0 {
		t.Error("withdraw balance is incorrect")
	}
}

func TestSetRate(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// get contract instance
	pledgeIns, err := pledge.NewPledge(pledgeAddr, backend)
	if err != nil {
		t.Error("new pledge contract instance failed:", err)
	}

	// make auth for sending transaction
	txAuth1, err := comm.MakeAuth(chainID, sk1)
	if err != nil {
		t.Error(err)
	}

	oldRate, err := pledgeIns.GetRate(&bind.CallOpts{From: Addr1})
	if err != nil {
		t.Error(err)
	}
	t.Log("old rate:", oldRate)

	rate, ok := new(big.Int).SetString("800000000", 10)
	if !ok {
		t.Error("set string error")
	}
	// set reate
	tx, err := pledgeIns.SetRate(txAuth1, rate)
	if err != nil {
		t.Error(err)
	}
	// wait tx
	t.Log("waiting tx..")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("set rate err:", err)
	}

	// get new rate
	newRate, err := pledgeIns.GetRate(&bind.CallOpts{From: Addr1})
	if err != nil {
		t.Error(err)
	}
	t.Log("new rate:", newRate)

	if newRate.Cmp(rate) != 0 {
		t.Error("new rate is incorrect")
	}
}
