package router

import (
	"encoding/json"
	"net/http"
	"nft-project/model"
	"nft-project/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// API 结构体，持有数据库连接
type ApiServer struct {
	db *gorm.DB
}

func NewApiServer(db *gorm.DB) *ApiServer {
	return &ApiServer{db: db}
}

// 启动 API 服务
func (s *ApiServer) Start() {
	r := gin.Default()

	// // 允许跨域 (前端开发必备)
	// r.Use(func(c *gin.Context) {
	// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// 	c.Next()
	// })

	// 路由组
	api := r.Group("/api")
	{
		api.GET("/auctions", s.getAuctions)
		api.GET("/auctions/:id", s.getAuctionDetail)
		api.GET("/statistics", s.getStatistics)
		api.GET("/users/:address/nfts", s.getUserNFTs)
	}

	r.Run(":8080") // 在 8080 端口启动
}

func (s *ApiServer) getAuctions(c *gin.Context) {
	var auctions []model.AuctionRecord

	status := c.Query("status")
	sort := c.Query("sort")

	query := s.db.Model(&model.AuctionRecord{})

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if sort == "price" {
		query = query.Order("highest_bid_amount desc")
	} else {
		query = query.Order("created_at desc") // 默认按最新创建
	}

	if err := query.Find(&auctions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, auctions)
}

// 2. 获取拍卖详情及其出价历史
func (s *ApiServer) getAuctionDetail(c *gin.Context) {
	id := c.Param("id")
	var auction model.AuctionRecord

	// 使用 ID 查询并关联出价记录
	if err := s.db.Preload("Bids").Where("auction_id = ?", id).First(&auction).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "拍卖不存在"})
		return
	}
	c.JSON(http.StatusOK, auction)
}

func (s *ApiServer) getStatistics(c *gin.Context) {
	var auctionCount int64
	var bidCount int64

	// 分别统计表行数
	s.db.Model(&model.AuctionRecord{}).Count(&auctionCount)
	s.db.Model(&model.BidRecord{}).Count(&bidCount)

	c.JSON(http.StatusOK, gin.H{
		"total_auctions": auctionCount,
		"total_bids":     bidCount,
	})
}

func (s *ApiServer) getUserNFTs(c *gin.Context) {
	address := c.Param("address")

	apiKey := utils.GetEnv("ALCHEMY_KEY", "")
	url := "https://eth-sepolia.g.alchemy.com/nft/v3/" + apiKey + "/getNFTsForOwner?owner=" + address

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "请求 Alchemy 失败"})
		return
	}
	defer resp.Body.Close()

	// 简单透传 Alchemy 的响应给前端
	var result interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	c.JSON(http.StatusOK, result)
}
