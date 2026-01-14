// Package models chứa các định nghĩa cấu trúc dữ liệu (struct) cho User Service,
// được sử dụng để ánh xạ (mapping) với các bảng trong cơ sở dữ liệu PostgreSQL qua GORM.
package models

import (
	"time"

	"gorm.io/gorm"
)

// Profile đại diện cho thông tin chi tiết và tùy chọn của người dùng trong hệ thống.
type Profile struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	UserID uint   `gorm:"unique;not null" json:"user_id"`
	Avatar string `json:"avatar"`
	Bio    string `gorm:"type:text" json:"bio"`
	// Preferences lưu trữ các cài đặt cá nhân dưới dạng JSON (ví dụ: ngôn ngữ, giao diện).
	Preferences string         `gorm:"type:jsonb" json:"preferences"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
