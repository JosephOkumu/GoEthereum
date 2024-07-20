package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Replace YOUR_PROJECT_ID with your actual Infura project ID
	conn, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatalf("Whoops, something went wrong: %v", err)
	}

	ctx := context.Background()
	txHash := common.HexToHash("0xcc9b91dee308231ff6ac0dd09dc8731682185c87e588bb6f3f1e79bf9f3f3f12")
	tx, pending, err := conn.TransactionByHash(ctx, txHash)
	if err != nil {
		log.Fatalf("Failed to get transaction by hash: %v", err)
	}

	if !pending {
		fmt.Println(tx)
	} else {
		fmt.Println("Transaction is still pending")
	}
}
