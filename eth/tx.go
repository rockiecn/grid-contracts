package eth

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/xerrors"
)

var (
	checkTxSleepTime = 6 // 先等待6s（出块时间加1）
	nextBlockTime    = 5 // 出块时间5s
	defaultGasPrice  = int64(2000)
)

func MakeAuth(chainID *big.Int, hexSk string) (*bind.TransactOpts, error) {
	auth := &bind.TransactOpts{}
	sk, err := crypto.HexToECDSA(hexSk)
	if err != nil {
		return auth, err
	}

	auth, err = bind.NewKeyedTransactorWithChainID(sk, chainID)
	if err != nil {
		return nil, xerrors.Errorf("new keyed transaction failed %s", err)
	}

	auth.Value = big.NewInt(0)
	auth.GasPrice = big.NewInt(defaultGasPrice)
	return auth, nil
}

// GetTransactionReceipt 通过交易hash获得交易详情
func GetTransactionReceipt(endPoint string, hash common.Hash) *types.Receipt {
	client, err := ethclient.Dial(endPoint)
	if err != nil {
		log.Fatal("rpc.Dial err", err)
	}
	defer client.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	receipt, err := client.TransactionReceipt(ctx, hash)
	if err != nil {
		log.Println("get transaction receipt: ", err)
	}
	return receipt
}

// CheckTx check whether transaction is successful through receipt
func CheckTx(endPoint string, txHash common.Hash, name string) error {
	var receipt *types.Receipt

	t := checkTxSleepTime
	for i := 0; i < 10; i++ {
		if i != 0 {
			t = nextBlockTime * i
		}
		time.Sleep(time.Duration(t) * time.Second)
		receipt = GetTransactionReceipt(endPoint, txHash)
		if receipt != nil {
			break
		}
	}

	if receipt == nil {
		return xerrors.Errorf("%s %s cann't get tx receipt, not packaged", name, txHash)
	}

	// 0 means fail
	if receipt.Status == 0 {
		if receipt.GasUsed != receipt.CumulativeGasUsed {
			return xerrors.Errorf("%s %s transaction exceed gas limit", name, txHash)
		}
		return xerrors.Errorf("%s %s transaction mined but execution failed, please check your tx input, status:%v", name, txHash, receipt.Status)
	}
	return nil
}

func PrintGasUsed(endPoint string, hash common.Hash, prefix string) {
	err := CheckTx(endPoint, hash, prefix)
	if err != nil {
		log.Println(err)
	}

	receipt := GetTransactionReceipt(endPoint, hash)
	fmt.Println(prefix, receipt.GasUsed)
}
