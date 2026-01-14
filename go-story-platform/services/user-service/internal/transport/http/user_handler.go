// Package http cho user trong user-service
package http

import (
	"net/http"
	"strconv"
	"user-service/internal/models"
	"user-service/internal/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// UserHandler không tự tạo DB, Nhận repo từ ngoài
type UserHandler struct {
	repo *repository.UserRepository
}

// NewUserHandler contructor
func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

// CreateUser Client → Handler → Repository → DB POST /users
func (h *UserHandler) CreateUser(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1. Mã hóa mật khẩu
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể xử lý mật khẩu"})
		return
	}

	// 2. Chuyển từ Input sang Model để lưu vào DB
	user := models.User{
		Username:      input.Username,
		Email:         input.Email,
		PasswordHash:  string(hashedPassword), // Lưu bản đã mã hóa
		WalletAddress: input.WalletAddress,
		Profile: models.Profile{
			Bio:         input.Profile.Bio,
			Preferences: "{}",
		},
	}

	if err := h.repo.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// UpdateUser : PUT /users/:id
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID không hợp lệ"})
		return
	}
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " Cập nhật thất bại" + err.Error()})
		return
	}

	user.ID = uint(id)
	if err := h.repo.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cập nhật thất bại" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cập nhật thành công", "user": user})
}

// DeleteUser : DELETE /users/:id
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.repo.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Xóa thất bại"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Đã xóa user ID " + strconv.Itoa(id)})
}

// ListUsers: GET /users
func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.repo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể lấy danh sách người dùng"})
		return
	}

	c.JSON(http.StatusOK, users)
}
