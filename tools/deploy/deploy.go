package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	comm "github.com/grid/contracts/common"
	"github.com/grid/contracts/go/gtoken"
)

var (
	allGas = make([]uint64, 0)
)

func main() {

	//endpoint := flag.String("ep", "https://testchain.metamemo.one:24180", "block chain endpoint")
	endpoint := flag.String("ep", "HTTP://127.0.0.1:7545", "block chain endpoint")
	sk := flag.String("sk", "c1e763d955e6aea410e40b95702108a30efb4d25b32d419910fe2ac611c2229d", "sk for sending transaction")

	flag.Parse()

	//owner := "0x5F7F7e31399531F08C2b47eA1919F11346405a16"
	//addrOnwer := common.HexToAddress(owner)

	fmt.Println("endpoint:", *endpoint)

	// connect to an eth node with ep
	client, chainID := comm.ConnETH(*endpoint)

	// make auth for sending transaction
	fmt.Println("chainid:", chainID)
	txAuth, err := comm.MakeAuth(chainID, *sk)
	if err != nil {
		log.Fatal(err)
	}

	// deploy Token contract
	tokenAddr, tx, tokenIns, err := gtoken.DeployGToken(txAuth, client)
	if err != nil {
		log.Fatal("deploy token err:", err)
	}
	fmt.Println("tokenAddr: ", tokenAddr.Hex())
	go comm.PrintGasUsed(*endpoint, tx.Hash(), "deploy token gas:", &allGas)

	_ = tokenIns

	ToAddr := common.HexToAddress("0xe2198eb2e931f9306ABcA68D4F093E0Ac4823B0d")
	// 0.01 eth
	Amount, ok := new(big.Int).SetString("10000000000000000", 10)
	if !ok {
		log.Fatal("new big int failed")
	}

	// test mint 0.01 eth to acc1 in ganache
	tokenIns.Mint(txAuth, ToAddr, Amount)

}
