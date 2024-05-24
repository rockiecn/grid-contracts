// test for token contract
package main

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	comm "github.com/grid/contracts/common"
	"github.com/grid/contracts/go/gtoken"
)

var (
	// token contract info
	tokenAddr common.Address
)

func TestDeploy(t *testing.T) {
	t.Log("connecting to eth:", comm.Endpoint)

	// connect to an eth node with ep
	client, chainID := comm.ConnETH(comm.Endpoint)
	t.Log("chain id:", chainID)

	// make auth for sending transaction
	txAuth, err := comm.MakeAuth(chainID, comm.SK1)
	if err != nil {
		t.Error(err)
	}

	// deploy Token contract
	_tokenAddr, tx, _, err := gtoken.DeployGtoken(txAuth, client)
	if err != nil {
		t.Error("deploy token err:", err)
	}
	tokenAddr = _tokenAddr

	t.Log("created token address: ", tokenAddr.Hex())

	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(comm.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy token err:", err)
	}

	receipt := comm.GetTransactionReceipt(comm.Endpoint, tx.Hash())
	t.Log("gas used:", receipt.GasUsed)

}

// test mint some token to addr2 and check the balance
func TestMint(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(comm.Endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	tokenIns, err := gtoken.NewGtoken(tokenAddr, backend)
	if err != nil {
		t.Error("new token instance failed:", err)
	}

	// make auth to sign and send tx
	txAuth1, err := comm.MakeAuth(chainID, comm.SK1)
	if err != nil {
		t.Error(err)
	}

	// get balance of addr2
	bal1, err := tokenIns.BalanceOf(&bind.CallOpts{From: comm.Addr2}, comm.Addr2)
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
	tx, err := tokenIns.Mint(txAuth1, comm.Addr2, amount)
	if err != nil {
		t.Error("call mint err:", err)
	}

	t.Log("waiting for mint tx to be ok")
	// wait tx to complete
	err = comm.CheckTx(comm.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	// get balance of addr2
	bal2, err := tokenIns.BalanceOf(&bind.CallOpts{From: comm.Addr2}, comm.Addr2)
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
	backend, chainID := comm.ConnETH(comm.Endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	tokenIns, err := gtoken.NewGtoken(tokenAddr, backend)
	if err != nil {
		t.Error("new token instance failed:", err)
	}

	// make auth to sign and send tx
	txAuth1, err := comm.MakeAuth(chainID, comm.SK1)
	if err != nil {
		t.Error(err)
	}

	mintAmount, _ := new(big.Int).SetString("1000000000000000000", 10)
	// mint some token for acc1
	tx, err := tokenIns.Mint(txAuth1, comm.Addr1, mintAmount)
	if err != nil {
		t.Error("call mint err:", err)
	}
	t.Log("waiting for mint tx to be ok")
	// wait tx to complete
	err = comm.CheckTx(comm.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	// get balance of addr2
	bal1, err := tokenIns.BalanceOf(&bind.CallOpts{From: comm.Addr2}, comm.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("balanceof acc2:", bal1)

	// 0.02 eth to transfer
	amount, ok := new(big.Int).SetString("20000000000000000", 10)
	if !ok {
		t.Error("set string failed")
	}

	// call
	t.Log("calling transfer")
	tx, err = tokenIns.Transfer(txAuth1, comm.Addr2, amount)
	if err != nil {
		t.Error("call contract err:", err)
	}
	t.Log("waiting for tx to be ok")
	// wait tx to complete
	err = comm.CheckTx(comm.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	// get balance of addr2
	bal2, err := tokenIns.BalanceOf(&bind.CallOpts{From: comm.Addr2}, comm.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("new balanceof acc2:", bal2)

	if bal2.String() != "30000000000000000" {
		t.Error("balance error")
	}
}
