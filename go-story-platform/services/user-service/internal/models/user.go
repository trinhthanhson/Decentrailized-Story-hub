package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Username      string         `gorm:"unique;not null" json:"username"`
	Email         string         `gorm:"unique;not null" json:"email"`
	PasswordHash  string         `gorm:"not null" json:"-"`
	WalletAddress string         `gorm:"index;not null" json:"wallet_address"` //Liên kết Blockchain
	Role          string         `gorm:"not null;default:reader" json:"role"`
	Profile       Profile        `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"profile"` //Quan hệ 1-1 với Profile
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
