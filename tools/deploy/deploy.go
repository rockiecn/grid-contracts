package main

import (
	"fmt"
	"math/big"

	"github.com/grid/contracts/go/access"
	"github.com/grid/contracts/go/credit"
	"github.com/grid/contracts/go/market"
	"github.com/grid/contracts/go/registry"

	eth "github.com/grid/contracts/eth"
	"github.com/grid/contracts/log"
)

var logger = log.Logger("deploy")

// deploy all contracts for grid, save all addresses into a json file
func main() {
	logger.Info("connecting to eth:", eth.Endpoint)

	// connect to an eth node with ep
	backend, chainID := eth.ConnETH(eth.Endpoint)
	logger.Info("chain id:", chainID)

	// make auth for sending transaction with admin
	txAuth, err := eth.MakeAuth(chainID, eth.SK0)
	if err != nil {
		logger.Panic(err)
	}

	//
	txAuth.GasLimit = 5000000
	// 2 gwei
	txAuth.GasPrice = new(big.Int).SetUint64(2000000000)

	// deploy market contract
	// gas: 3428839
	logger.Info("deploying market..")
	_marketAddr, tx, _, err := market.DeployMarket(txAuth, backend)
	if err != nil {
		logger.Panic("deploy market err:", err)
	}

	logger.Info("tx hash:", tx.Hash())

	logger.Info("created market address: ", _marketAddr.Hex())
	logger.Info("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		logger.Panic("deploy contract err:", err)
	}

	receipt := eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	logger.Info("deploy market gas used:", receipt.GasUsed)

	// deploy access contract
	// gas:495059
	logger.Info("deploying access")
	_accessAddr, tx, accessIns, err := access.DeployAccess(txAuth, backend)
	if err != nil {
		logger.Panic("deploy access err:", err)
	}
	logger.Info("created access address: ", _accessAddr.Hex())
	logger.Info("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		logger.Panic("deploy contract err:", err)
	}

	receipt = eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	logger.Info("deploy access gas used:", receipt.GasUsed)

	// set access for admin
	// gas:65143
	logger.Info("set access for admin")
	tx, err = accessIns.Set(txAuth, eth.Addr0, true)
	if err != nil {
		logger.Panic(err)
	}
	logger.Info("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		logger.Panic(err)
	}

	receipt = eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	logger.Info("set access gas used:", receipt.GasUsed)

	// deploy credit contract
	// gas:1763026
	logger.Info("deploying credit contract")
	_creditAddr, tx, _, err := credit.DeployCredit(txAuth, backend, _accessAddr)
	if err != nil {
		logger.Panic("deploy credit err:", err)
	}
	logger.Info("credit address:", _creditAddr)
	logger.Info("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		logger.Panic("deploy contract err:", err)
	}

	receipt = eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	logger.Info("deploy credit gas used:", receipt.GasUsed)

	// deploy registry contract
	// gas:1260180
	logger.Info("deploying registry contract")
	_registryAddr, tx, _, err := registry.DeployRegistry(txAuth, backend)
	if err != nil {
		logger.Panic("deploy registry err:", err)
	}
	logger.Info("registry address:", _registryAddr)
	logger.Info("waiting for tx to be ok")
	err = eth.CheckTx(eth.Endpoint, tx.Hash(), "")
	if err != nil {
		logger.Panic("deploy contract err:", err)
	}

	receipt = eth.GetTransactionReceipt(eth.Endpoint, tx.Hash())
	logger.Info("deploy registry gas used:", receipt.GasUsed)

	// deploy token contract

	// deploy pledge contract

	// deploy swap contract

	// save addresses into json
	a := eth.Address{
		Market:   _marketAddr.Hex(),
		Access:   _accessAddr.Hex(),
		Credit:   _creditAddr.Hex(),
		Registry: _registryAddr.Hex(),
	}
	eth.Save(a, "../../eth/contracts.json")

	fmt.Printf("\nAll contract deployed, addresses are saved in eth/contracts.json\n")
}

/*
func main() {
	//endpoint := flag.String("ep", "https://testchain.metamemo.one:24180", "block chain endpoint")
	endpoint := flag.String("ep", "HTTP://127.0.0.1:7545", "block chain endpoint")
	sk := flag.String("sk", "c1e763d955e6aea410e40b95702108a30efb4d25b32d419910fe2ac611c2229d", "sk for sending transaction")

	flag.Parse()

	//owner := "0x5F7F7e31399531F08C2b47eA1919F11346405a16"
	//addrOnwer := ethon.HexToAddress(owner)

	fmt.Println("endpoint:", *endpoint)

	// connect to an eth node with ep
	client, chainID := eth.ConnETH(*endpoint)

	// make auth for sending transaction
	fmt.Println("chainid:", chainID)
	txAuth, err := eth.MakeAuth(chainID, *sk)
	if err != nil {
		log.Fatal(err)
	}

	// deploy Token contract
	tokenAddr, tx, tokenIns, err := gtoken.DeployGtoken(txAuth, client)
	if err != nil {
		log.Fatal("deploy token err:", err)
	}
	fmt.Println("tokenAddr: ", tokenAddr.Hex())
	go eth.PrintGasUsed(*endpoint, tx.Hash(), "deploy token gas:")

	_ = tokenIns

	ToAddr := ethon.HexToAddress("0xe2198eb2e931f9306ABcA68D4F093E0Ac4823B0d")
	// 0.01 eth
	Amount, ok := new(big.Int).SetString("10000000000000000", 10)
	if !ok {
		log.Fatal("new big int failed")
	}

	// test mint 0.01 eth to acc1 in ganache
	tokenIns.Mint(txAuth, ToAddr, Amount)

}
*/
