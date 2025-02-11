package events

import (
	"context"
	"finalyearproject/Backend/services"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockchainEventListener struct {
	Client            *ethclient.Client
	BlockchainService *services.BlockchainService
}

func NewBlockchainEventListener(client *ethclient.Client, blockchainService *services.BlockchainService) *BlockchainEventListener {
	return &BlockchainEventListener{
		Client:            client,
		BlockchainService: blockchainService,
	}
}

func (listener *BlockchainEventListener) StartListening() {
	log.Println("Starting Blockchain Event Listener...")
	headChan := make(chan *types.Header)
	sub, err := listener.Client.SubscribeNewHead(context.Background(), headChan)
	if err != nil {
		log.Fatal("Failed to subscribe to new block headers:", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Println("Subscription error:", err)
		case header := <-headChan:
			log.Println("New block detected:", header.Number.String())
			listener.ProcessBlock(header.Number)
		}
	}
}

func (listener *BlockchainEventListener) ProcessBlock(blockNumber *big.Int) {
	block, err := listener.Client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Println("Failed to retrieve block:", err)
		return
	}

	for _, tx := range block.Transactions() {
		log.Println("Processing transaction:", tx.Hash().Hex())
		receipt, err := listener.Client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Println("Failed to retrieve transaction receipt:", err)
			continue
		}

		for _, logEntry := range receipt.Logs {
			listener.HandleEvent(logEntry)
		}
	}
}

func (listener *BlockchainEventListener) HandleEvent(logEntry *types.Log) {
	log.Println("New event received at address:", logEntry.Address.Hex())
	if logEntry.Address == common.HexToAddress("0xYourContractAddress") {
		log.Println("Event data:", logEntry.Data)
		// Add logic to parse event data and take necessary action
	}
}
