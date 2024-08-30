package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	rpcUrl, count := getParameters()

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// load all to-be-sent transactions
	workload := getWorkload(count)

	startBlock, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to get the start block: %v", err)
	}

	sendWorkload(client, workload)

	endBlock, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to get the end block: %v", err)
	}

	startHeight := startBlock.Number()
	startTime := startBlock.Time()

	endHeight := endBlock.Number()
	endTime := endBlock.Time()

	if endTime <= startTime {
		log.Fatal("More transactions are needed")
	}

	totalTxCount := 0

	// Iterate over the blocks from startHeight to the endHeight 
	for blockNumber := startHeight.Add(startHeight, big.NewInt(1)); blockNumber.Cmp(endHeight) <= 0; blockNumber.Add(blockNumber, big.NewInt(1)) {
		block, err := client.BlockByNumber(context.Background(), blockNumber)
		if err != nil {
			log.Fatalf("Failed to fetch block: %v", err)
		}

		totalTxCount += len(block.Transactions())
	}

	elapsedSeconds := endTime - startTime
	tps := float64(totalTxCount) / float64(elapsedSeconds)

	fmt.Printf("Total transactions counted in %d seconds is %d\n", elapsedSeconds, totalTxCount)
	fmt.Printf("The TPS of the chain is %.2f\n", tps)
}

func getParameters() (string, int) {
	// handle command line flags
	rpcUrl := flag.String("rpc-url", "http://127.0.0.1:8545", "RPC url of the chain")
        count := flag.Int("count", 10000, "The number of transactions to be sent")
        flag.Parse()

	if *count > 1000000 {
		log.Fatal("Too many transactions to be generated and sent")
	}

	return *rpcUrl, *count
}

func sendWorkload(client *ethclient.Client, workload [](*types.Transaction)) {
	for _, tx := range workload {
		err := client.SendTransaction(context.Background(), tx)
		if err != nil {
			log.Fatal("Failed to send transactions")
		}
	}
}

func getWorkload(n int) [](*types.Transaction) {
	return [](*types.Transaction){}
}
