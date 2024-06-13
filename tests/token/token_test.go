// test for token contract
package main

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/grid/contracts/eth"
	"github.com/grid/contracts/go/gtoken"
)

var (
	// token contract info
	tokenAddr common.Address
)

func TestDeploy(t *testing.T) {
	t.Log("connecting to eth:", eth.Endpoint)

	// connect to an eth node with ep
	client, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// make auth for sending transaction
	authAdmin, err := eth.MakeAuth(chainID, eth.SK0)
	if err != nil {
		t.Error(err)
	}

	// deploy Token contract
	_tokenAddr, tx, _, err := gtoken.DeployGtoken(authAdmin, client)
	if err != nil {
		t.Error("deploy token err:", err)
	}
	tokenAddr = _tokenAddr

	t.Log("created token address: ", tokenAddr.Hex())

	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy token err:", err)
	}

	receipt := eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	t.Log("gas used:", receipt.GasUsed)

}

// test mint some token to addr2 and check the balance
func TestMint(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	tokenIns, err := gtoken.NewGtoken(tokenAddr, backend)
	if err != nil {
		t.Error("new token instance failed:", err)
	}

	// make auth to sign and send tx
	authAdmin, err := eth.MakeAuth(chainID, eth.SK0)
	if err != nil {
		t.Error(err)
	}

	// get balance of provider
	bal1, err := tokenIns.BalanceOf(&bind.CallOpts{}, eth.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("balance:", bal1)

	// 0.01 eth
	amount, ok := new(big.Int).SetString("10000000000000000", 10)
	if !ok {
		t.Error("set string failed")
	}

	// call token mint
	tx, err := tokenIns.Mint(authAdmin, eth.Addr2, amount)
	if err != nil {
		t.Error("call mint err:", err)
	}

	t.Log("waiting for mint tx to be ok")
	// wait tx to complete
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	// get balance of addr2
	bal2, err := tokenIns.BalanceOf(&bind.CallOpts{}, eth.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("new balance:", bal2)

	if bal2.String() != "10000000000000000" {
		t.Error("balance error")
	}
}

func TestTransfer(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	tokenIns, err := gtoken.NewGtoken(tokenAddr, backend)
	if err != nil {
		t.Error("new token instance failed:", err)
	}

	// make auth to sign and send tx
	authAdmin, err := eth.MakeAuth(chainID, eth.SK0)
	if err != nil {
		t.Error(err)
	}
	authUser, err := eth.MakeAuth(chainID, eth.SK1)
	if err != nil {
		t.Error(err)
	}

	mintAmount, _ := new(big.Int).SetString("1000000000000000000", 10)
	// mint some token for acc1
	tx, err := tokenIns.Mint(authAdmin, eth.Addr1, mintAmount)
	if err != nil {
		t.Error("call mint err:", err)
	}
	t.Log("waiting for mint tx to be ok")
	// wait tx to complete
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	// get balance of user
	bal1, err := tokenIns.BalanceOf(&bind.CallOpts{}, eth.Addr1)
	if err != nil {
		t.Error(err)
	}
	t.Log("balanceof user:", bal1)

	// 0.02 eth to transfer
	amount, ok := new(big.Int).SetString("20000000000000000", 10)
	if !ok {
		t.Error("set string failed")
	}

	// call
	t.Log("calling transfer")
	tx, err = tokenIns.Transfer(authUser, eth.Addr2, amount)
	if err != nil {
		t.Error("call contract err:", err)
	}
	t.Log("waiting for tx to be ok")
	// wait tx to complete
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	// get balance of addr2
	bal2, err := tokenIns.BalanceOf(&bind.CallOpts{}, eth.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("new balanceof provider:", bal2)

	if bal2.String() != "30000000000000000" {
		t.Error("balance error")
	}
}
