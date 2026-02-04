package model

import (
	"time"
)

type BidRecord struct {
	ID uint `json:"id" gorm:"primaryKey"`

	AuctionId string `json:"auctionId" gorm:"not null;size:78;index"`
	Bidder    string `json:"bidder" gorm:"not null;size:42;index"`
	BidAmount string `json:"bidAmount" gorm:"not null;size:78"`
	TxHash    string `json:"txHash" gorm:"not null;uniqueIndex;size:66"`

	CreatedAt time.Time `json:"createdAt"`
}
