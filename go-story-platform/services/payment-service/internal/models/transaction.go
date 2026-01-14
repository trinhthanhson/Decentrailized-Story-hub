package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        uint           `gorm:"primaryKey"`
	UserID    uint           `gorm:"index"`
	Amount    int            `gorm:"not null"`
	Type      string         `gorm:"type:varchar(20)"` // Deposit, Purchase
	Status    string         `gorm:"type:varchar(20);default:'pending'"`
	TxHash    string         `gorm:"index"` // Hash trên Blockchain
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
