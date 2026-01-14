package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	AuthorID    uint           `gorm:"index" json:"author_id"` //ID từ User Service (Soft link)
	CategoryID  uint           `gorm:"index" json:"category_id"`
	Description string         `gorm:"type:text" json:"description"`
	IsPremium   bool           `gorm:"default:false" json:"is_premium"`
	Chapters    []Chapter      `gorm:"foreignKey:BookID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"chapters"` //Quan hệ 1-n với Chapters
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	gorm.Model
}
