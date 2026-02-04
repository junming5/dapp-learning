package dbConfig

import (
	"fmt"
	"log"
	"nft-project/model"
	"nft-project/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() *gorm.DB {
	dbHost := utils.GetEnv("DB_HOST", "")
	dbPORT := utils.GetEnv("DB_PORT", "")
	dbUser := utils.GetEnv("DB_USER", "")
	dbPassword := utils.GetEnv("DB_PASSWORD", "")
	dbName := utils.GetEnv("DB_NAME", "")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPORT, dbName)

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("无法连接到 MySQL 数据库: %v", err)
	}

	// 自动迁移所有模型
	err = DB.AutoMigrate(&model.AuctionRecord{}, &model.BidRecord{}, &model.WithdrawRecord{})

	if err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}
	log.Println("MySQL数据库成功连接并迁移")

	return DB
}
