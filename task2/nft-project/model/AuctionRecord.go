package model

import (
	"time"

	"gorm.io/gorm"
)

type AuctionRecord struct {
	ID uint `json:"id" gorm:"primaryKey"`

	AuctionId        string    `json:"auctionId" gorm:"not null;uniqueIndex;size:78"`
	SellerAddress    string    `json:"sellerAddress" gorm:"not null;size:42;index"`
	NftAddress       string    `json:"nftAddress" gorm:"not null;size:42;index"`
	TokenId          string    `json:"tokenId" gorm:"not null;size:78;index"`
	EndTime          time.Time `json:"endTime" gorm:"not null;index"`
	Status           int       `json:"status" gorm:"default:0"` // 0:进行中, 1:已结束, 2:流拍
	HighestBidAmount string    `json:"highestBidAmount" gorm:"size:78;default:'0'"`
	WinnerAddress    string    `json:"winnerAddress" gorm:"size:42;default:null"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	Bids []BidRecord `json:"bids" gorm:"foreignKey:AuctionId;references:AuctionId"`
}
