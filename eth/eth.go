package eth

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	// environment for ganache
	Endpoint string = "HTTP://127.0.0.1:7545"

	// environment for product chain
	//Endpoint string = "https://chain.metamemo.one:8501"

	// acc1 of ganache
	HexAddr1 string = "0x5F7F7e31399531F08C2b47eA1919F11346405a16"
	// acc2 of ganache
	HexAddr2 string = "0xe2198eb2e931f9306ABcA68D4F093E0Ac4823B0d"
	Addr1           = common.HexToAddress(HexAddr1)
	Addr2           = common.HexToAddress(HexAddr2)
	// use acc1 as the admin to send tx
	SK1 string = "c1e763d955e6aea410e40b95702108a30efb4d25b32d419910fe2ac611c2229d"
	SK2 string = "e8cda8fe7c04afa4a0630af457972f88a645468cb90120a11911669deac5e96e"
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
