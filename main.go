package main

import (
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
	_ = client // we'll use this in the upcoming sections

	address := common.HexToAddress("0xdAe6cC6fB09b0be1d6D6817D0E8Af85104a2Dfe7")
	fmt.Println(address.Hex())
	// fmt.Println(address.Hash().Hex()) 
	fmt.Println(address.Bytes())     
}
