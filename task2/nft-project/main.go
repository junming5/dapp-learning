package main

import (
	"log"
	"nft-project/dbConfig"
	"nft-project/ether"
	"nft-project/router"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("警告: 未能加载 .env 文件，将尝试读取系统环境变量")
	}

	// dbConfig.InitDB()
	db := dbConfig.InitDB()

	// ether.StartListenEvent(db)

	go ether.StartListenEvent(db)

	server := router.NewApiServer(db)
	server.Start()
}
