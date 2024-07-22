package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545") // Connect to Ganache
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")
	_ = client

	// Working with the address
	address := common.HexToAddress("0xdAe6cC6fB09b0be1d6D6817D0E8Af85104a2Dfe7")
	fmt.Println(address.Hex())
	// fmt.Println(address.Hash().Hex())
	fmt.Println(address.Bytes())

	// Checking balance
	account := common.HexToAddress("0xdAe6cC6fB09b0be1d6D6817D0E8Af85104a2Dfe7")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("currrent balance :%v\n", balance)

	// Checking the pending balance
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)

	if err != nil {
		fmt.Println("Error:%V", err)
	}
	fmt.Printf("pending balance: %v\n", pendingBalance)
}
