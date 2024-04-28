package common

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

// connect to an eth node with endpoint and return client and chain id
func ConnETH(ep string) (*ethclient.Client, *big.Int) {
	// get eth client
	client, err := ethclient.DialContext(context.Background(), ep)
	if err != nil {
		log.Fatal(err)
	}

	// make auth to send transaction
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return client, chainId
}
