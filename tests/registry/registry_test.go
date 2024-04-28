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

	// token contract addr
	contractAddr common.Address
)

func TestDeploy(t *testing.T) {
	t.Log("connecting to eth:", endpoint)

	// connect to an eth node with ep
	client, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// make auth for sending transaction
	txAuth, err := comm.MakeAuth(chainID, sk1)
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
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy token err:", err)
	}

	receipt := comm.GetTransactionReceipt(endpoint, tx.Hash())
	t.Log("gas used:", receipt.GasUsed)

}

func TestGet(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	contractIns, err := registry.NewRegistry(contractAddr, backend)
	if err != nil {
		t.Error("new token instance failed:", err)
	}

	// get balance of addr2
	regInfo, err := contractIns.Get(&bind.CallOpts{From: Addr2}, Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("registry info:", regInfo)
}

func TestSet(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	contractIns, err := registry.NewRegistry(contractAddr, backend)
	if err != nil {
		t.Error("new token instance failed:", err)
	}

	// register for acc2
	txAuth2, err := comm.MakeAuth(chainID, sk2)
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
	comm.CheckTx(endpoint, tx.Hash(), "")

	// get reg info
	regInfo, err := contractIns.Get(&bind.CallOpts{From: Addr2}, Addr2)
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
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	contractIns, err := registry.NewRegistry(contractAddr, backend)
	if err != nil {
		t.Error("new token instance failed:", err)
	}

	// register for acc2
	txAuth2, err := comm.MakeAuth(chainID, sk2)
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
	comm.CheckTx(endpoint, tx.Hash(), "")

	// get reg info
	regInfo, err := contractIns.Get(&bind.CallOpts{From: Addr2}, Addr2)
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
