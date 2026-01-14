// Package repository cho user trong user-service
package repository

import (
	"user-service/internal/models"

	"gorm.io/gorm"
)

// UserRepository Giữ kết nối với db
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository dùng để tạo UserRepository và gắn kết nối DB vào nó
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser : Thêm User mới (Tự động tạo luôn Profile trống)
func (r *UserRepository) CreateUser(user *models.User) error {
	// GORM sẽ tự động chèn dữ liệu vào cả 2 bảng users và profiles nếu user.Profile được khởi tạo
	return r.db.Create(user).Error
}

// GetUserByID Hàm này lấy User theo ID từ DB, tự động load Profile, và trả về user + lỗi
func (r *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Profile").First(&user, id).Error
	return &user, err
}

// UpdateUser Hàm này cập nhật User và toàn bộ dữ liệu liên quan, nhờ FullSaveAssociations: true
func (r *UserRepository) UpdateUser(user *models.User) error {
	return r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(user).Error
}

// DeleteUser đánh dấu User là đã xóa bằng deleted_at, không xóa khỏi DB
func (r *UserRepository) DeleteUser(id uint) error {
	// GORM sẽ đánh dấu DeletedAt thay vì xóa vĩnh viễn
	return r.db.Delete(&models.User{}, id).Error
}

// UpdateProfileByUserID Hàm này cập nhật Profile theo user_id, chỉ thay đổi các field có giá trị
func (r *UserRepository) UpdateProfileByUserID(profile *models.Profile) error {
	return r.db.Model(&models.Profile{}).Where("user_id = ?", profile.UserID).Updates(profile).Error
}

// GetAllUsers lấy danh sách toàn bộ người dùng kèm thông tin Profile
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	// Preload giúp nạp thông tin từ bảng profiles để tránh lỗi N+1 query
	err := r.db.Preload("Profile").Find(&users).Error
	return users, err
}
