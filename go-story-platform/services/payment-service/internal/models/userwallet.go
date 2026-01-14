package models

import (
	"time"

	"gorm.io/gorm"
)

type UserWallet struct {
	UserID          uint `gorm:"primaryKey"` // Lấy từ User Service
	BalanceInternal int  `gorm:"default:0"`  // Xu trong web
	LastSyncAt      time.Time
	Transactions    []Transaction  `gorm:"foreignKey:UserID;references:UserID"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
