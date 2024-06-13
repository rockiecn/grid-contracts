package main

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/grid/contracts/eth"
	"github.com/grid/contracts/go/access"
	"github.com/grid/contracts/go/credit"
	"github.com/grid/contracts/go/gtoken"
	"github.com/grid/contracts/go/swap"
)

var (
	// access address
	accessAddr common.Address
	// erc20 token address
	tokenAddr common.Address
	// credit contract addr
	creditAddr common.Address
	// swap addr
	swapAddr common.Address

	rate_num = new(big.Int).SetUint64(10)
	rate_den = new(big.Int).SetUint64(1)
)

// deploy credit, gtoken, swap contracts
func TestDeploy(t *testing.T) {
	t.Log("connecting to eth:", eth.Endpoint)

	// connect to an eth node with ep
	backend, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// make auth for sending transaction
	authAdmin, err := eth.MakeAuth(chainID, eth.SK0)
	if err != nil {
		t.Error(err)
	}

	// deploy access contract
	_accessAddr, tx, accessIns, err := access.DeployAccess(authAdmin, backend)
	if err != nil {
		t.Error("deploy access err:", err)
	}
	accessAddr = _accessAddr
	t.Log("created access address: ", accessAddr.Hex())
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	// deploy credit contract, access contract is needed
	_creditAddr, tx, _, err := credit.DeployCredit(authAdmin, backend, accessAddr)
	if err != nil {
		t.Error("deploy credit err:", err)
	}
	creditAddr = _creditAddr
	t.Log("created credit address: ", creditAddr.Hex())
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	// deploy gtoken contract
	_tokenAddr, tx, tokenIns, err := gtoken.DeployGtoken(authAdmin, backend)
	if err != nil {
		t.Error("deploy gtoken err:", err)
	}
	tokenAddr = _tokenAddr
	t.Log("created gtoken address: ", tokenAddr.Hex())
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	// deploy swap contract
	_swapAddr, tx, _, err := swap.DeploySwap(authAdmin, backend, creditAddr, tokenAddr, rate_num, rate_den)
	if err != nil {
		t.Error("deploy swap err:", err)
	}
	swapAddr = _swapAddr
	t.Log("created swap address: ", swapAddr.Hex())
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	// admin transfer some token to swap contract
	amount := new(big.Int).SetUint64(10000)
	t.Log("transfering some gtoken to swap")
	tx, err = tokenIns.Transfer(authAdmin, swapAddr, amount)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	// set access for admin and swap contract
	t.Log("set access for admin")
	tx, err = accessIns.Set(authAdmin, eth.Addr0, true)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	t.Log("set access for swap contract")
	tx, err = accessIns.Set(authAdmin, swapAddr, true)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	// receipt := eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	// t.Log("gas used:", receipt.GasUsed)
}

// test credit buy gtoken
// admin mint credit for user, user buy gtoken with credit, check balance
func TestBuy(t *testing.T) {
	t.Log("connecting to eth:", eth.Endpoint)

	// connect to an eth node with ep
	backend, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// make auth for sending transaction
	authAdmin, err := eth.MakeAuth(chainID, eth.SK0)
	if err != nil {
		t.Error(err)
	}
	authUser, err := eth.MakeAuth(chainID, eth.SK1)
	if err != nil {
		t.Error(err)
	}

	creditIns, err := credit.NewCredit(creditAddr, backend)
	if err != nil {
		t.Error(err)
	}

	// mint 10 credit
	t.Log("admin mint credit to user")
	credit := new(big.Int).SetUint64(10)
	tx, err := creditIns.Mint(authAdmin, eth.Addr1, credit)
	if err != nil {
		t.Error(err)
	}

	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}
	// check credit
	creditBefore, err := creditIns.BalanceOf(&bind.CallOpts{}, eth.Addr1)
	if err != nil {
		t.Error(err)
	}
	t.Log("credit balance before buy:", creditBefore)
	if creditBefore.Cmp(credit) != 0 {
		t.Error("credit balance error")
	}

	// get token before buy
	tokenIns, err := gtoken.NewGtoken(tokenAddr, backend)
	if err != nil {
		t.Error(err)
	}
	tokenBefore, err := tokenIns.BalanceOf(&bind.CallOpts{}, eth.Addr1)
	if err != nil {
		t.Error(err)
	}
	t.Log("token before:", tokenBefore)

	// user approve to swap before buy
	buyAmount := new(big.Int).SetUint64(4)
	t.Log("user approving to swap")
	tx, err = creditIns.Approve(authUser, swapAddr, buyAmount)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	// user buy gtoken with credit
	swapIns, err := swap.NewSwap(swapAddr, backend)
	if err != nil {
		t.Error(err)
	}
	t.Log("user buy gtoken with credit amount:", buyAmount)
	tx, err = swapIns.Buy(authUser, buyAmount)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("buy gtoken err:", err)
	}

	// check balance after

	// user's credit balance
	creditAfer, err := creditIns.BalanceOf(&bind.CallOpts{}, eth.Addr1)
	if err != nil {
		t.Error(err)
	}
	t.Log("credit after buy:", creditAfer)
	should := creditBefore.Sub(creditBefore, buyAmount)
	if creditAfer.Cmp(should) != 0 {
		t.Errorf("credit balance is incorrect: %v, should be: %v", creditAfer, should)
	}

	// user's gtoken balance
	tokenAfter, err := tokenIns.BalanceOf(&bind.CallOpts{}, eth.Addr1)
	if err != nil {
		t.Error(err)
	}
	t.Log("token balance after buy:", tokenAfter)
	should = credit.Mul(buyAmount, rate_num)
	should = should.Div(should, rate_den)
	if tokenAfter.Cmp(should) != 0 {
		t.Errorf("token balance is incorrect: %v, should be: %v", tokenAfter, should)
	}

}

// sell gtoken for credit
func TestSell(t *testing.T) {
	t.Log("connecting to eth:", eth.Endpoint)

	// connect to an eth node with ep
	backend, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// make auth for sending transaction
	authUser, err := eth.MakeAuth(chainID, eth.SK1)
	if err != nil {
		t.Error(err)
	}

	creditIns, err := credit.NewCredit(creditAddr, backend)
	if err != nil {
		t.Error(err)
	}

	// check credit before sell
	creditBefore, err := creditIns.BalanceOf(&bind.CallOpts{}, eth.Addr1)
	if err != nil {
		t.Error(err)
	}
	t.Log("credit balance before sell:", creditBefore)

	// get token before sell
	tokenIns, err := gtoken.NewGtoken(tokenAddr, backend)
	if err != nil {
		t.Error(err)
	}
	tokenBefore, err := tokenIns.BalanceOf(&bind.CallOpts{}, eth.Addr1)
	if err != nil {
		t.Error(err)
	}
	t.Log("token before:", tokenBefore)

	// user approve to gtoken before sell, sell 20 token for 2 credit
	sellAmount := new(big.Int).SetUint64(20)
	t.Log("user approving some token to swap")
	tx, err := tokenIns.Approve(authUser, swapAddr, sellAmount)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	// user buy gtoken with credit
	swapIns, err := swap.NewSwap(swapAddr, backend)
	if err != nil {
		t.Error(err)
	}
	t.Log("user sell gtoken for credit:", sellAmount)
	tx, err = swapIns.Sell(authUser, sellAmount)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("buy gtoken err:", err)
	}

	// check balance after

	// user's credit balance
	creditAfer, err := creditIns.BalanceOf(&bind.CallOpts{}, eth.Addr1)
	if err != nil {
		t.Error(err)
	}
	t.Log("credit after sell:", creditAfer)
	gotCredit := new(big.Int).Set(sellAmount)
	gotCredit = gotCredit.Mul(sellAmount, rate_den)
	gotCredit = gotCredit.Div(gotCredit, rate_num)
	should := creditBefore.Add(creditBefore, gotCredit)
	if creditAfer.Cmp(should) != 0 {
		t.Errorf("credit balance is incorrect: %v, should be: %v", creditAfer, should)
	}

	// gtoken balance
	tokenAfter, err := tokenIns.BalanceOf(&bind.CallOpts{}, eth.Addr1)
	if err != nil {
		t.Error(err)
	}
	t.Log("token balance after buy:", tokenAfter)
	should = tokenBefore.Sub(tokenBefore, sellAmount)
	if tokenAfter.Cmp(should) != 0 {
		t.Errorf("token balance is incorrect: %v, should be: %v", tokenAfter, should)
	}

}

func TestSettle(t *testing.T) {
	t.Log("connecting to eth:", eth.Endpoint)

	// connect to an eth node with ep
	backend, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// make auth for sending transaction
	authAdmin, err := eth.MakeAuth(chainID, eth.SK0)
	if err != nil {
		t.Error(err)
	}

	swapIns, err := swap.NewSwap(swapAddr, backend)
	if err != nil {
		t.Error(err)
	}
	tokenIns, err := gtoken.NewGtoken(tokenAddr, backend)
	if err != nil {
		t.Error(err)
	}

	// token balance in swap
	balInSwap, err := tokenIns.BalanceOf(&bind.CallOpts{}, swapAddr)
	if err != nil {
		t.Error(err)
	}
	t.Log("balance in swap before settle:", balInSwap)

	// token balance of admin before settle
	balBefore, err := tokenIns.BalanceOf(&bind.CallOpts{}, eth.Addr0)
	if err != nil {
		t.Error(err)
	}
	t.Log("balance of admin before settle::", balBefore)

	// admin call settlement
	t.Log("admin call settlement")
	tx, err := swapIns.SettlementToken(authAdmin)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	// token balance after settle
	balAfter, err := tokenIns.BalanceOf(&bind.CallOpts{}, eth.Addr0)
	if err != nil {
		t.Error(err)
	}
	t.Log("balance of admin after settle:", balAfter)

	// balance in swap after settle
	balInSwapAfter, err := tokenIns.BalanceOf(&bind.CallOpts{}, swapAddr)
	if err != nil {
		t.Error(err)
	}
	t.Log("balance in swap after settle:", balInSwapAfter)

	// check balance
	if balInSwapAfter.Cmp(new(big.Int).SetUint64(0)) != 0 {
		t.Error("balance in swap should be 0 after settled")
	}

	// balance of admin should increased
	if balAfter.Sub(balAfter, balBefore).Cmp(balInSwap) != 0 {
		t.Error("the balance of admin should increased with balance in swap")
	}
}
