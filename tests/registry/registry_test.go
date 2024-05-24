// test for token contract
package main

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	comm "github.com/grid/contracts/common"
	"github.com/grid/contracts/go/registry"
)

var (
	// token contract addr
	contractAddr common.Address
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

	// deploy registry contract
	_contractAddr, tx, _, err := registry.DeployRegistry(txAuth, client)
	if err != nil {
		t.Error("deploy registry err:", err)
	}
	contractAddr = _contractAddr

	t.Log("created registry address: ", contractAddr.Hex())

	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(comm.Endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy token err:", err)
	}

	receipt := comm.GetTransactionReceipt(comm.Endpoint, tx.Hash())
	t.Log("gas used:", receipt.GasUsed)

}

func TestGet(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(comm.Endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	contractIns, err := registry.NewRegistry(contractAddr, backend)

	if err != nil {
		t.Error("new token instance failed:", err)
	}

	// get balance of addr2
	regInfo, err := contractIns.Get(&bind.CallOpts{From: comm.Addr2}, comm.Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("registry info:", regInfo)
}

func TestSet(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(comm.Endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	contractIns, err := registry.NewRegistry(contractAddr, backend)
	if err != nil {
		t.Error("new token instance failed:", err)
	}

	// register for acc2
	txAuth2, err := comm.MakeAuth(chainID, comm.SK2)
	if err != nil {
		t.Error(err)
	}

	// call registry's Set method
	tx, err := contractIns.Set(txAuth2, "123.123.123.0", "test domain", 123, 11, 22, 33, 44)
	if err != nil {
		t.Error(err)
	}

	// wait tx ok
	t.Log("waiting for set to be ok")
	comm.CheckTx(comm.Endpoint, tx.Hash(), "")

	// get reg info
	regInfo, err := contractIns.Get(&bind.CallOpts{From: comm.Addr2}, comm.Addr2)
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
	if regInfo.Port != 123 {
		t.Error("port is error")
	}
	if regInfo.Total.CPU != 11 {
		t.Error("cpu is error")
	}
	if regInfo.Total.GPU != 22 {
		t.Error("gpu is error")
	}
	if regInfo.Total.MEM != 33 {
		t.Error("mem is error")
	}
	if regInfo.Total.DISK != 44 {
		t.Error("disk is error")
	}
}

func TestUpdate(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(comm.Endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	contractIns, err := registry.NewRegistry(contractAddr, backend)
	if err != nil {
		t.Error("new token instance failed:", err)
	}

	// register for acc2
	txAuth2, err := comm.MakeAuth(chainID, comm.SK2)
	if err != nil {
		t.Error(err)
	}

	// call registry's update method
	tx, err := contractIns.Update(txAuth2, 1, 1, 1, 1)
	if err != nil {
		t.Error(err)
	}

	// wait tx ok
	t.Log("waiting for tx to be ok")
	comm.CheckTx(comm.Endpoint, tx.Hash(), "")

	// get reg info
	regInfo, err := contractIns.Get(&bind.CallOpts{From: comm.Addr2}, comm.Addr2)
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
	if regInfo.Port != 123 {
		t.Error("port is error")
	}
	if regInfo.Avail.CPU != 10 {
		t.Error("cpu is error")
	}
	if regInfo.Avail.GPU != 21 {
		t.Error("gpu is error")
	}
	if regInfo.Avail.MEM != 32 {
		t.Error("mem is error")
	}
	if regInfo.Avail.DISK != 43 {
		t.Error("disk is error")
	}
}
