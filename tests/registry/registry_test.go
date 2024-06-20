// test for token contract
package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/grid/contracts/eth"
	"github.com/grid/contracts/go/registry"
)

var (
	// token contract addr
	contractAddr common.Address
)

// deploy registry contract
func TestDeploy(t *testing.T) {
	t.Log("connecting to eth:", eth.Endpoint)

	// connect to an eth node with ep
	client, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// make auth for sending transaction with admin
	txAuth, err := eth.MakeAuth(chainID, eth.SK0)
	if err != nil {
		t.Error(err)
	}

	// deploy registry contract
	_contractAddr, tx, _, err := registry.DeployRegistry(txAuth, client)
	if err != nil {
		t.Error("deploy registry err:", err)
	}
	contractAddr = _contractAddr

	t.Log("created registry address: ", contractAddr.Hex())

	t.Log("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy token err:", err)
	}

	receipt := eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	t.Log("gas used:", receipt.GasUsed)

}

func TestGet(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	contractIns, err := registry.NewRegistry(contractAddr, backend)

	if err != nil {
		t.Error("new token instance failed:", err)
	}

	// get balance of addr2
	regInfo, err := contractIns.Get(&bind.CallOpts{From: eth.Addr2}, eth.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("registry info:", regInfo)
}

func TestRegister(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	contractIns, err := registry.NewRegistry(contractAddr, backend)
	if err != nil {
		t.Error("new token instance failed:", err)
	}

	// register a cp
	txAuth2, err := eth.MakeAuth(chainID, eth.SK2)
	if err != nil {
		t.Error(err)
	}

	// call registry's Set method
	tx, err := contractIns.Register(txAuth2, "123.123.123.0", "test domain", "123", 11, 22, 33, 44)
	if err != nil {
		t.Error(err)
	}

	// wait tx ok
	t.Log("waiting for set to be ok")
	eth.CheckTx(eth.Endpoint, tx.Hash(), "")

	// get reg info
	regInfo, err := contractIns.Get(&bind.CallOpts{}, eth.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("regInfo:", regInfo)

	// check if set is correct
	if regInfo.Ip != "123.123.123.0" {
		t.Error("ip is error")
	}
	if regInfo.Domain != "test domain" {
		t.Error("domain is error")
	}
	if regInfo.Port != "123" {
		t.Error("port is error")
	}
	if regInfo.Total.NCPU != 11 {
		t.Error("cpu is error")
	}
	if regInfo.Total.NGPU != 22 {
		t.Error("gpu is error")
	}
	if regInfo.Total.NMEM != 33 {
		t.Error("mem is error")
	}
	if regInfo.Total.NDISK != 44 {
		t.Error("disk is error")
	}
}

func TestGetKeys(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// get registry instance
	registryIns, err := registry.NewRegistry(contractAddr, backend)
	if err != nil {
		t.Error("new contract instance failed:", err)
	}

	// call
	keys, err := registryIns.GetKeys(&bind.CallOpts{})
	if err != nil {
		t.Error(err)
	}

	t.Log("keys:", keys)
}

func TestUpdate(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := eth.ConnETH(eth.Endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	contractIns, err := registry.NewRegistry(contractAddr, backend)
	if err != nil {
		t.Error("new token instance failed:", err)
	}

	// auth for cp
	txAuth2, err := eth.MakeAuth(chainID, eth.SK2)
	if err != nil {
		t.Error(err)
	}

	// call registry's update method
	tx, err := contractIns.Update(txAuth2, eth.Addr2, 1, 1, 1, 1)
	if err != nil {
		t.Error(err)
	}

	// wait tx ok
	t.Log("waiting for tx to be ok")
	eth.CheckTx(eth.Endpoint, tx.Hash(), "")

	// get reg info of cp
	regInfo, err := contractIns.Get(&bind.CallOpts{}, eth.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("regInfo:", regInfo)

	// check if set is correct
	if regInfo.Ip != "123.123.123.0" {
		t.Error("ip is error")
	}
	if regInfo.Domain != "test domain" {
		t.Error("domain is error")
	}
	if regInfo.Port != "123" {
		t.Error("port is error")
	}
	if regInfo.Avail.NCPU != 10 {
		t.Error("cpu is error")
	}
	if regInfo.Avail.NGPU != 21 {
		t.Error("gpu is error")
	}
	if regInfo.Avail.NMEM != 32 {
		t.Error("mem is error")
	}
	if regInfo.Avail.NDISK != 43 {
		t.Error("disk is error")
	}
}
