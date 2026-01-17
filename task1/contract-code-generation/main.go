package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func DeployContract(auth *bind.TransactOpts, client *ethclient.Client) (common.Address, *Main, error) {
	fmt.Println("🚀 正在发起合约部署...")

	address, tx, instance, err := DeployMain(auth, client)
	if err != nil {
		return common.Address{}, nil, err
	}

	fmt.Printf("📦 交易已发送，等待确认... Hash: %s\n", tx.Hash().Hex())

	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return common.Address{}, nil, err
	}

	fmt.Printf("✅ 合约部署成功！地址: %s\n", address.Hex())
	return address, instance, nil
}

func InteractWithContract(auth *bind.TransactOpts, client *ethclient.Client, address common.Address) (*big.Int, error) {
	instance, err := NewMain(address, client)
	if err != nil {
		return nil, fmt.Errorf("加载合约失败: %v", err)
	}

	fmt.Println("➕ 正在执行 Increment...")
	tx, err := instance.Increment(auth)
	if err != nil {
		return nil, fmt.Errorf("Increment 失败: %v", err)
	}

	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return nil, fmt.Errorf("等待交易确认失败: %v", err)
	}

	count, err := instance.GetCount(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("获取计数失败: %v", err)
	}

	return count, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("无法加载 .env 文件，请检查该文件是否存在")
	}

	rpcURL := os.Getenv("RPC_URL")
	privKeyHex := os.Getenv("PRIVATE_KEY")

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("连接节点失败: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(privKeyHex)
	if err != nil {
		log.Fatalf("私钥无效: %v", err)
	}
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("获取 ChainID 失败: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("创建身份认证失败: %v", err)
	}

	addr, _, err := DeployContract(auth, client)
	if err != nil {
		log.Fatalf("部署失败: %v", err)
	}

	auth2, _ := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	finalCount, err := InteractWithContract(auth2, client, addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("🔢 最终结果: %d\n", finalCount)
}
