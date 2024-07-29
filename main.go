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
	client, err := connectToEthereum()
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	checkBalance(client, "0x1547d00fd8fcf012D04163B758A3D6110bCD2C1a")
	generateNewWallet()
	transferETH(client)
}

func connectToEthereum() (*ethclient.Client, error) {
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
	if err != nil {
		return nil, err
	}
	fmt.Println("We have a connection")
	return client, nil
}

func checkBalance(client *ethclient.Client, addressStr string) {
	addressStr = "0x1547d00fd8fcf012D04163B758A3D6110bCD2C1a"
	address := common.HexToAddress(addressStr)
	fmt.Println(address.Hex())
	fmt.Println(address.Bytes())

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Printf("Error getting balance: %v", err)
		return
	}
	fmt.Printf("Current balance: %v\n", balance)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), address)
	if err != nil {
		log.Printf("Error getting pending balance: %v", err)
		return
	}
	fmt.Printf("Pending balance: %v\n", pendingBalance)
}

func generateNewWallet() {
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

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Printf("Newly generated wallet address: %s\n", address)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Printf("Encoded wallet address: %v\n", hexutil.Encode(hash.Sum(nil)[12:]))
}

func transferETH(client *ethclient.Client) {
	privateKey, err := crypto.HexToECDSA("1b3a311ed6b94f6a897aa9bc0849eec8ea9c227503de64094e6170688f83cb80")
	if err != nil {
		log.Printf("Error loading private key: %v", err)
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Printf("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Printf("Error getting nonce: %v", err)
		return
	}

	value := big.NewInt(5000000000000000000) // 5 ETH
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Printf("Error getting suggested gas price: %v", err)
		return
	}

	toAddress := common.HexToAddress("0x98Db73401ccB0621Aa104E800646447F39f64456")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID := big.NewInt(1337)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
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

	checkBalance(client, "0x98Db73401ccB0621Aa104E800646447F39f64456")
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