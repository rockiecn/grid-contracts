// test for market contract
package main

import (
	"testing"

	"github.com/grid/contracts/eth"
	"github.com/grid/contracts/go/access"
	"github.com/grid/contracts/go/credit"
	"github.com/grid/contracts/go/market"
)

func TestDeploy(t *testing.T) {
	t.Log("connecting to eth:", eth.Endpoint)

	// connect to an eth node with ep
	backend, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// make auth for sending transaction
	txAuth, err := eth.MakeAuth(chainID, eth.SK0)
	if err != nil {
		t.Error(err)
	}

	// deploy market contract
	// gas: 3428839
	t.Log("deploying market..")
	_marketAddr, tx, _, err := market.DeployMarket(txAuth, backend)
	if err != nil {
		t.Error("deploy registry err:", err)
	}

	t.Log("created market address: ", _marketAddr.Hex())
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	receipt := eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	t.Log("deploy market gas used:", receipt.GasUsed)

	// deploy access contract
	// gas:494259
	t.Log("deploying access")
	_accessAddr, tx, accessIns, err := access.DeployAccess(txAuth, backend)
	if err != nil {
		t.Error("deploy access err:", err)
	}
	t.Log("created access address: ", _accessAddr.Hex())
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	receipt = eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	t.Log("deploy access gas used:", receipt.GasUsed)

	// set access for admin
	// gas:69043
	t.Log("set access for admin")
	tx, err = accessIns.Set(txAuth, eth.Addr0, true)
	if err != nil {
		t.Error(err)
	}
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	receipt = eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	t.Log("set access gas used:", receipt.GasUsed)

	// deploy credit contract
	// gas:1767526
	t.Log("deploying credit contract")
	_creditAddr, tx, _, err := credit.DeployCredit(txAuth, backend, _accessAddr)
	if err != nil {
		t.Error("deploy credit err:", err)
	}
	t.Log("credit address:", _creditAddr)
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	receipt = eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	t.Log("deploy credit gas used:", receipt.GasUsed)

	// save addresses into json
	a := eth.Address{
		Market: _marketAddr.Hex(),
		Access: _accessAddr.Hex(),
		Credit: _creditAddr.Hex(),
	}
	eth.Save(a, "./contracts.json")
}
