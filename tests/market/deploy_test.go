// test for market contract
package main

import (
	"math/big"
	"testing"

	"github.com/grid/contracts/eth"
	"github.com/grid/contracts/go/access"
	"github.com/grid/contracts/go/credit"
	"github.com/grid/contracts/go/market"
	"github.com/grid/contracts/go/registry"
)

func TestDeploy(t *testing.T) {
	t.Log("connecting to eth:", eth.Endpoint)

	// connect to an eth node with ep
	endpoint, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// make auth for sending transaction
	txAuth, err := eth.MakeAuth(chainID, eth.SK0)
	if err != nil {
		t.Error(err)
	}

	//
	txAuth.GasLimit = 4000000
	// 2 gwei
	txAuth.GasPrice = new(big.Int).SetUint64(2000000000)

	// deploy access contract
	// gas:494259
	t.Log("deploying access")
	_accessAddr, tx, accessIns, err := access.DeployAccess(txAuth, endpoint)
	if err != nil {
		t.Error("deploy access err:", err)
	}
	t.Log("created access address: ", _accessAddr.Hex())
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy contract err:", err)
	}

	receipt := eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
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
	_creditAddr, tx, _, err := credit.DeployCredit(txAuth, endpoint, _accessAddr)
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

	// deploy registry contract
	// gas:1260180
	t.Log("deploying registry contract")
	_registryAddr, tx, _, err := registry.DeployRegistry(txAuth, endpoint)
	if err != nil {
		t.Fatal("deploy registry err:", err)
	}
	t.Log("registry address:", _registryAddr)
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Fatal("deploy contract err:", err)
	}

	receipt = eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	t.Log("deploy registry gas used:", receipt.GasUsed)

	// deploy market contract
	// gas: 3428839
	t.Log("deploying market..")
	_marketAddr, tx, _, err := market.DeployMarket(txAuth, endpoint, _registryAddr, _creditAddr)
	if err != nil {
		t.Fatal("deploy market err:", err)
	}

	t.Log("tx hash:", tx.Hash())

	t.Log("created market address: ", _marketAddr.Hex())
	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Fatal("deploy contract err:", err)
	}

	receipt = eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	t.Log("deploy market gas used:", receipt.GasUsed)

	// save addresses into json
	a := eth.Address{
		Market:   _marketAddr.Hex(),
		Access:   _accessAddr.Hex(),
		Credit:   _creditAddr.Hex(),
		Registry: _registryAddr.Hex(),
	}
	eth.Save(a, "./contracts.json")
}
