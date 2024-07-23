package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
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

	// Check balance
	account := common.HexToAddress("0xdAe6cC6fB09b0be1d6D6817D0E8Af85104a2Dfe7")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("currrent balance :%v\n", balance)

	// Check the pending balance
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)

	if err != nil {
		fmt.Println("Error:%V", err)
	}
	fmt.Printf("pending balance: %v\n", pendingBalance)

	// Generating new Wallet
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Printf("PrivateKey: %v\n", hexutil.Encode(privateKeyBytes)[2:])
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Printf("Public Key: %s\n", hexutil.Encode(publicKeyBytes)[4:])
	address2 := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Printf("Newly generated wallet address: %s\n", address2)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Printf("Encoded wallet address: %v\n", hexutil.Encode(hash.Sum(nil)[12:]))
	
}
