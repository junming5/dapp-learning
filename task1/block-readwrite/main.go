package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("读取 .env 文件失败")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	rpcURL := os.Getenv("RPC_URL")

	//连接 Sepolia 节点
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("连接失败:%v", err)
	}
	fmt.Println("成功连接到以太坊网络！")

	header, _ := client.HeaderByNumber(ctx, nil)
	fmt.Printf("最新区块号：%s\n", header.Number.String())

	blockNumber := big.NewInt(5000000)
	block, err := client.BlockByNumber(ctx, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("区块哈希：%s\n", block.Hash().Hex())
	fmt.Printf("时间戳：%d\n", block.Time())
	fmt.Printf("交易数量：%d\n", len(block.Transactions()))

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("导出公钥失败")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	toAddress := common.HexToAddress(os.Getenv("TO_ADDRESS"))

	chainID, err := client.NetworkID(ctx)
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasTipCap, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		log.Fatal("获取小费失败:", err)
	}

	baseFee := header.BaseFee
	if baseFee == nil {
		gasPrice, err := client.SuggestGasPrice(ctx)
		if err != nil {
			log.Fatalf("获取gas价格失败:%v", err)
		}
		baseFee = gasPrice
	}

	// fee cap = base fee * 2 + tip cap
	gasFeeCap := new(big.Int).Add(
		new(big.Int).Mul(baseFee, big.NewInt(2)),
		gasTipCap,
	)

	gasLimit := uint64(21000)

	amountEth := 0.0001
	amountWei := new(big.Float).Mul(
		big.NewFloat(amountEth),
		big.NewFloat(1e18),
	)
	valueWei, _ := amountWei.Int(nil)

	balance, err := client.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		log.Fatalf("获取余额失败：%v", err)
	}

	totalCost := new(big.Int).Add(
		valueWei,
		new(big.Int).Mul(gasFeeCap, big.NewInt(int64(gasLimit))),
	)

	if balance.Cmp(totalCost) < 0 {
		log.Fatalf("余额不足：%swei, 需要：%swei", balance.String(), totalCost.String())
	}

	txData := &types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: gasTipCap, // 小费上限
		GasFeeCap: gasFeeCap, // 总费用上限
		Gas:       gasLimit,  // 普通转账固定 Gas
		To:        &toAddress,
		Value:     valueWei,
		Data:      nil,
	}
	tx := types.NewTx(txData)

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		log.Fatal("签名失败：", err)
	}

	fmt.Println("📡 正在广播交易...")

	fmt.Println("======== 🚀 准备发送 EIP-1559 动态费用交易 ========")
	fmt.Printf("🌐 网络环境:    Sepolia (ChainID: %s)\n", chainID)
	fmt.Printf("⚓ 交易序号:    Nonce [%d]\n", nonce)
	fmt.Printf("📤 发送地址:    %s\n", fromAddress.Hex())
	fmt.Printf("📥 接收地址:    %s\n", toAddress.Hex())
	fmt.Println("-----------------------------------------------------")
	fmt.Printf("💰 转账金额:    %18s Wei (≈ %.6f ETH)\n", valueWei.String(), amountEth)

	fmt.Printf("🎁 优先小费:    %18s Wei \n", gasTipCap.String())
	fmt.Printf("🔝 最高单价:    %18s Wei \n", gasFeeCap.String())
	fmt.Printf("⛽ 消耗限额:    %18d Gas\n", gasLimit)

	fmt.Println("-----------------------------------------------------")
	fmt.Printf("💳 预计总扣费:  %18s Wei (金额 + 最高手续费)\n", totalCost.String())
	fmt.Printf("💎 当前余额:    %18s Wei\n", balance.String())
	fmt.Println("==================================================")

	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Fatalf("发送失败: %v", err)
	}

	hash := signedTx.Hash().Hex()
	fmt.Printf("🚀 交易已发送! Hash：%s\n", hash)
	fmt.Printf("🔍 查看进度: https://sepolia.etherscan.io/tx/%s\n", hash)
}
