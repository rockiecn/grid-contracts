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

	// sepolia chain
	//Endpoint string = "https://rpc.sepolia.ethpandaops.io"

	// accounts info in ganache
	// 0 as admin, 1 as user, 2 as provider, 3 as backup admin
	//HexAddr0 string = "0x5F7F7e31399531F08C2b47eA1919F11346405a16"
	HexAddr0 string = "0x5Dd05D18Dc788e47b7206f3FD1B0ea6a546A124B"
	HexAddr1 string = "0xe2198eb2e931f9306ABcA68D4F093E0Ac4823B0d"
	HexAddr2 string = "0xC4EAc9E1012DFCB4833165F5d35E027EBfE1f640"

	Addr0 = common.HexToAddress(HexAddr0)

	Addr1 = common.HexToAddress(HexAddr1)
	Addr2 = common.HexToAddress(HexAddr2)
	//SK0      string = "c1e763d955e6aea410e40b95702108a30efb4d25b32d419910fe2ac611c2229d"
	SK0 string = "6ab894b68e56444ae9fb7d80fc07c0d169d1c5b62e65193ad55bd915b7fb4529"
	SK1 string = "e8cda8fe7c04afa4a0630af457972f88a645468cb90120a11911669deac5e96e"
	SK2 string = "ae313c1dc715026cf629641e3ae2dab06f95c7300d97b3121310d375b979f19b"
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
