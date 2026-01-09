package models

import (
	"time"

	"gorm.io/gorm"
)

type Chapter struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	BookID        uint           `gorm:"index" json:"book_id"`
	ChapterNumber int            `gorm:"not null" json:"chapter_number"`
	ContenrURL    string         `gorm:"not null" json:"content_url"` // Link S3/IPFS
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
