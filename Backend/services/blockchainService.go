package services

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

type BlockchainService struct {
	client  *rpc.Client
	address common.Address
}

func NewBlockchainService(rpcURL string, contractAddress string) (*BlockchainService, error) {
	client, err := rpc.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to blockchain node: %v", err)
	}
	return &BlockchainService{
		client:  client,
		address: common.HexToAddress(contractAddress),
	}, nil
}

func (b *BlockchainService) GetLatestBlockNumber() (*big.Int, error) {
	var latestBlock *big.Int
	err := b.client.CallContext(context.Background(), &latestBlock, "eth_blockNumber")
	if err != nil {
		return nil, fmt.Errorf("failed to get latest block number: %v", err)
	}
	return latestBlock, nil
}

func (b *BlockchainService) Close() {
	b.client.Close()
}

func main() {
	service, err := NewBlockchainService("https://rpc-url", "0xYourContractAddress")
	if err != nil {
		log.Fatalf("Error initializing blockchain service: %v", err)
	}
	defer service.Close()

	blockNumber, err := service.GetLatestBlockNumber()
	if err != nil {
		log.Fatalf("Error retrieving block number: %v", err)
	}
	fmt.Println("Latest Block Number:", blockNumber)
}
