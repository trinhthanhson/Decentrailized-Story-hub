package models

type Profile struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `gorm:"unique;not null" json:"user_id"`
	Avatar    string `json:"avatar"`
	Bio       string `gorm:"type:text" json:"bio"`
	Prefences string `gorm:"type:jsonb" json:"preferences"` // lưu sở thích dưới dạng JSON
}
