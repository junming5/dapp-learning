package ether

import (
	"context"
	"fmt"
	"log"
	"nft-project/contracts"
	"nft-project/model"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// --- 事件 Signature 定义 (从 ABI 推导的 Keccak256 哈希) ---
var (
	SigCreateAuction = crypto.Keccak256Hash([]byte("AuctionCreated(uint256,address,address,uint256,uint256)"))
	SigBidPlaced     = crypto.Keccak256Hash([]byte("BidPlaced(uint256,address,uint256)"))
	SigAuctionEnded  = crypto.Keccak256Hash([]byte("AuctionEnded(uint256,address,uint256)"))
	SigWithdraw      = crypto.Keccak256Hash([]byte("Withdraw(address,uint256)"))
)

func StartListenEvent(db *gorm.DB) {
	godotenv.Load()
	client, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		log.Fatal(err)
	}

	contractAddr := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	query := ethereum.FilterQuery{Addresses: []common.Address{contractAddr}}

	logsChan := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logsChan)
	if err != nil {
		log.Fatalf("❌ 订阅日志失败: %v", err)
	}

	filterer, _ := contracts.NewAuctionFilterer(contractAddr, client)
	fmt.Println("🛰️  NFT 拍卖后端监听已启动，正在扫描 Sepolia...")

	for {
		select {
		case err := <-sub.Err():
			log.Printf("订阅异常: %v", err)
		case vLog := <-logsChan:
			// 确保 Topics 不为空防止越界
			if len(vLog.Topics) == 0 {
				continue
			}

			switch vLog.Topics[0] {
			case SigCreateAuction:
				processCreate(db, filterer, vLog)
			case SigBidPlaced:
				processBid(db, filterer, vLog)
			case SigAuctionEnded:
				processEnd(db, filterer, vLog)
			case SigWithdraw:
				processWithdraw(db, filterer, vLog)
			}
		}
	}
}

func processCreate(db *gorm.DB, f *contracts.AuctionFilterer, vLog types.Log) {
	ev, err := f.ParseAuctionCreated(vLog)
	if err != nil {
		log.Printf("解析错误: %v", err)
		return
	}

	db.Create(&model.AuctionRecord{
		AuctionId:     ev.AuctionId.String(),
		SellerAddress: ev.Seller.Hex(),
		NftAddress:    ev.NftAddress.Hex(),
		TokenId:       ev.TokenId.String(),
		EndTime:       time.Unix(int64(ev.EndTime.Uint64()), 0),
		Status:        0,
	})
	fmt.Printf("✨ [新拍卖] ID: %s, NFT: %s\n", ev.AuctionId, ev.NftAddress.Hex())
}

func processBid(db *gorm.DB, f *contracts.AuctionFilterer, vLog types.Log) {
	ev, err := f.ParseBidPlaced(vLog)
	if err != nil {
		log.Printf("解析错误: %v", err)
		return
	}

	// 记录出价记录
	db.Create(&model.BidRecord{
		AuctionId: ev.AuctionId.String(),
		Bidder:    ev.Bidder.Hex(),
		BidAmount: ev.UsdValue.String(),
		TxHash:    vLog.TxHash.Hex(),
	})

	db.Model(&model.AuctionRecord{}).
		Where("auction_id = ?", ev.AuctionId.String()).
		Updates(map[string]interface{}{
			"highest_bid_amount": ev.UsdValue.String(),
		})

	fmt.Printf("💰 [新出价] ID: %s, 金额: %s\n", ev.AuctionId, ev.UsdValue.String())
}

func processEnd(db *gorm.DB, f *contracts.AuctionFilterer, vLog types.Log) {
	ev, err := f.ParseAuctionEnded(vLog)
	if err != nil {
		log.Printf("解析错误: %v", err)
		return
	}

	// 2. 核心逻辑判断：
	// 如果 Winner 是全 0 地址 (0x000...)，说明没人出价，状态设为 2 (流拍)
	// 否则，状态设为 1 (已结束/成交)
	status := 1
	winner := ev.Winner.Hex()
	if winner == "0x0000000000000000000000000000000000000000" {
		status = 2
	}

	// 3. 更新数据库：根据 auction_id 找到那条记录，更新状态和赢家地址
	err = db.Model(&model.AuctionRecord{}).
		Where("auction_id = ?", ev.AuctionId.String()).
		Updates(map[string]interface{}{
			"status":         status,
			"winner_address": winner,
		}).Error

	if err != nil {
		log.Printf("更新数据库结拍状态失败: %v", err)
		return
	}

	fmt.Printf("🏁 [拍卖结束] ID: %s, 赢家: %s, 最终USD金额: %s, 状态: %d\n",
		ev.AuctionId.String(), winner, ev.AmountUsd.String(), status)
}

func processWithdraw(db *gorm.DB, f *contracts.AuctionFilterer, vLog types.Log) {
	ev, err := f.ParseWithdraw(vLog)
	if err != nil {
		log.Printf("解析提现事件错误: %v", err)
		return
	}

	// 2. 写入数据库
	record := model.WithdrawRecord{
		User:      ev.User.Hex(),
		AmountUsd: ev.Amount.String(),
		TxHash:    vLog.TxHash.Hex(),
	}

	if err := db.Create(&record).Error; err != nil {
		log.Printf("保存提现记录失败: %v", err)
		return
	}

	fmt.Printf("💰 [提现成功] 用户: %s, 金额: %s USD, Hash: %s\n",
		ev.User.Hex(), ev.Amount.String(), vLog.TxHash.Hex())
}
