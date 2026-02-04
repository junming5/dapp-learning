package model

import "time"

type WithdrawRecord struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	User      string    `json:"user" gorm:"size:42;index"`    // 谁提现了
	AmountUsd string    `json:"amountUsd" gorm:"size:78"`     // 提现的金额 (USD 单位)
	TxHash    string    `json:"txHash" gorm:"size:66;unique"` // 交易哈希，防止重复计入
	CreatedAt time.Time `json:"createdAt"`
}
