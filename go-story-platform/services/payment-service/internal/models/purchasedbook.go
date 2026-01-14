package models

import (
	"time"

	"gorm.io/gorm"
)

type PurchasedBook struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"index"`
	BookID      uint   `gorm:"index"`
	ChapterID   uint   `gorm:"index"` // Có thể mua lẻ chương
	NFTTokenID  string `gorm:"index"` // Bằng chứng sở hữu Blockchain
	PurchasedAt time.Time
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
