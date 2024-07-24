package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

func main() {
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545") // Connect to Ganache
	if err != nil {
		log.Printf("Failed to connect to the Ethereum client: %v", err)
		return
	}
	fmt.Println("We have a connection")

	// Working with the address
	address := common.HexToAddress("0xcc05BbBF7eb60d01803007eBAc0362447AA95199")
	fmt.Println(address.Hex())
	fmt.Println(address.Bytes())

	// Check balance
	account := common.HexToAddress("0xcc05BbBF7eb60d01803007eBAc0362447AA95199")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Printf("Error getting balance: %v", err)
		return
	}
	fmt.Printf("Current balance: %v\n", balance)

	// Check the pending balance
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	if err != nil {
		log.Printf("Error getting pending balance: %v", err)
		return
	}
	fmt.Printf("Pending balance: %v\n", pendingBalance)

	// Generating new Wallet
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Printf("Error generating key: %v", err)
		return
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Printf("PrivateKey: %v\n", hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Printf("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Printf("Public Key: %s\n", hexutil.Encode(publicKeyBytes)[4:])

	address2 := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Printf("Newly generated wallet address: %s\n", address2)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Printf("Encoded wallet address: %v\n", hexutil.Encode(hash.Sum(nil)[12:]))

	// Transferring ETH
	// Loading private key
	newPrivateKey, err := crypto.HexToECDSA("e5f9b5a2ddcec8fd238b7bb4412c6f73b0ea1ecce62207010fa66ed8dff2918f")
	if err != nil {
		log.Printf("Error loading private key: %v", err)
		return
	}

	// Generating public key from private key.
	newPublicKey := newPrivateKey.Public()
	newPublicKeyECDSA, ok := newPublicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Printf("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}
	fromAddress := crypto.PubkeyToAddress(*newPublicKeyECDSA)

	// Check balance of sending account
	senderBalance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Printf("Error getting sender's balance: %v", err)
		return
	}
	log.Printf("Balance of sending account: %v", senderBalance)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Printf("Error getting nonce: %v", err)
		return
	}
	log.Printf("Nonce: %d", nonce)

	// Converting ether to transact into Wei
	value := big.NewInt(5000000000000000000) // 5 ETH
	gasLimit := uint64(21000)                // Standard gas limit for simple transactions
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("Error getting suggested gas price: %v", err)
		return
	}
	log.Printf("Suggested gas price: %v", gasPrice)

	toAddress := common.HexToAddress("0xD7Cbc245b0f315531c5A93498B3579fa22da66D5")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	// chainID, err := client.NetworkID(context.Background())
	// if err != nil {
	// 	log.Printf("Error getting chain ID: %v", err)
	// 	return
	// }
	chainID := big.NewInt(1337) // or whatever your Ganache chain ID is
	log.Printf("Chain ID: %v", chainID)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), newPrivateKey)
	if err != nil {
		log.Printf("Error signing transaction: %v", err)
		return
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Printf("Error sending transaction: %v", err)
		return
	}

	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())

	// Wait for the transaction to be mined
	receipt, err := waitForTx(client, signedTx.Hash())
	if err != nil {
		log.Printf("Error waiting for transaction: %v", err)
		return
	}

	if receipt.Status == types.ReceiptStatusSuccessful {
		fmt.Println("Transaction successful")
	} else {
		fmt.Println("Transaction failed")
	}

	// Check balance of the sending account after the transaction 
	txAccount := common.HexToAddress("0xc6D3B650C9456bE4551b9e543485a2F04bBc5eD8")
	txBalance, err := client.BalanceAt(context.Background(), txAccount, nil)
	if err != nil {
		log.Printf("Error getting balance: %v", err)
		return
	}
	fmt.Printf("Current balance: %v\n", txBalance)

}

func waitForTx(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	for {
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			return receipt, nil
		}
		if err != nil && err != ethereum.NotFound {
			return nil, err
		}
		time.Sleep(time.Second)
	}
}
