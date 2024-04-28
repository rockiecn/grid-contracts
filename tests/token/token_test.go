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
	// endpoint for ganache
	endpoint string = "HTTP://127.0.0.1:7545"

	addr1 string = "0x5F7F7e31399531F08C2b47eA1919F11346405a16"
	addr2 string = "0xe2198eb2e931f9306ABcA68D4F093E0Ac4823B0d"
	Addr1        = common.HexToAddress(addr1)
	Addr2        = common.HexToAddress(addr2)
	sk1   string = "c1e763d955e6aea410e40b95702108a30efb4d25b32d419910fe2ac611c2229d"
	//sk2   string = "e8cda8fe7c04afa4a0630af457972f88a645468cb90120a11911669deac5e96e"

	// token contract info
	tokenAddr common.Address
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

	// deploy Token contract
	_tokenAddr, tx, _, err := gtoken.DeployGToken(txAuth, client)
	if err != nil {
		t.Error("deploy token err:", err)
	}
	tokenAddr = _tokenAddr

	t.Log("created token address: ", tokenAddr.Hex())

	t.Log("waiting for tx to be ok")
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error("deploy token err:", err)
	}

	receipt := comm.GetTransactionReceipt(endpoint, tx.Hash())
	t.Log("gas used:", receipt.GasUsed)

}

// test mint some token to addr2 and check the balance
func TestMint(t *testing.T) {
	// connect to an eth node with ep
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	tokenIns, err := gtoken.NewGToken(tokenAddr, backend)
	if err != nil {
		t.Error("new token instance failed:", err)
	}

	// make auth to sign and send tx
	txAuth1, err := comm.MakeAuth(chainID, sk1)
	if err != nil {
		t.Error(err)
	}

	// get balance of addr2
	bal1, err := tokenIns.BalanceOf(&bind.CallOpts{From: Addr2}, Addr2)
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
	tx, err := tokenIns.Mint(txAuth1, Addr2, amount)
	if err != nil {
		t.Error("call mint err:", err)
	}

	t.Log("waiting for mint tx to be ok")
	// wait tx to complete
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	// get balance of addr2
	bal2, err := tokenIns.BalanceOf(&bind.CallOpts{From: Addr2}, Addr2)
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
	backend, chainID := comm.ConnETH(endpoint)
	t.Log("chain id:", chainID)

	// get token instance
	tokenIns, err := gtoken.NewGToken(tokenAddr, backend)
	if err != nil {
		t.Error("new token instance failed:", err)
	}

	// make auth to sign and send tx
	txAuth1, err := comm.MakeAuth(chainID, sk1)
	if err != nil {
		t.Error(err)
	}

	mintAmount, _ := new(big.Int).SetString("1000000000000000000", 10)
	// mint some token for acc1
	tx, err := tokenIns.Mint(txAuth1, Addr1, mintAmount)
	if err != nil {
		t.Error("call mint err:", err)
	}
	t.Log("waiting for mint tx to be ok")
	// wait tx to complete
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	// get balance of addr2
	bal1, err := tokenIns.BalanceOf(&bind.CallOpts{From: Addr2}, Addr2)
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
	tx, err = tokenIns.Transfer(txAuth1, Addr2, amount)
	if err != nil {
		t.Error("call contract err:", err)
	}
	t.Log("waiting for tx to be ok")
	// wait tx to complete
	err = comm.CheckTx(endpoint, tx.Hash(), "")
	if err != nil {
		t.Error(err)
	}

	// get balance of addr2
	bal2, err := tokenIns.BalanceOf(&bind.CallOpts{From: Addr2}, Addr2)
	if err != nil {
		t.Error(err)
	}
	t.Log("new balanceof acc2:", bal2)

	if bal2.String() != "30000000000000000" {
		t.Error("balance error")
	}
}
