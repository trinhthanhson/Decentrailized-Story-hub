// Package http cho user trong user-service
package http

import (
	"net/http"
	"regexp"
	"strconv"
	"unicode"
	"user-service/internal/models"
	"user-service/internal/repository"
	"user-service/internal/transport/http/dto"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// UserHandler  Xử lý các request liên quan đến User
type UserHandler struct {
	repo *repository.UserRepository
}

// NewUserHandler tạo UserHandler với repo được truyền vào
func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

// CreateUser : POST /users Tạo user mới và profile trống kèm theo
func (h *UserHandler) CreateUser(c *gin.Context) {
	//khai báo biến input để nhận dữ liệu từ client
	var input dto.RegisterRequest

	// Bind JSON từ request body vào biến input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Username không được để trống
	if len(input.Username) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username không được để trống"})
		return
	}

	// Username phải có ít nhất 8 ký tự
	if len(input.Username) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username phải có ít nhất 8 ký tự"})
		return
	}

	// Username đã tồn tại
	existingUserByUsername, err := h.repo.GetUserByUsername(input.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi server"})
		return
	}
	if existingUserByUsername != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username đã tồn tại"})
		return
	}

	// Mật khẩu không được để trống
	if len(input.Password) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mật khẩu không được để trống"})
		return
	}

	// Mật khẩu phải có ít nhất 8 ký tự
	if len(input.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mật khẩu phải có ít nhất 8 ký tự"})
		return
	}

	// Mật khẩu phải có ít nhất 1 chữ hoa
	hasUpper := false
	for _, c := range input.Password {
		if unicode.IsUpper(c) {
			hasUpper = true
			break
		}
	}
	if !hasUpper {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mật khẩu phải có ít nhất 1 chữ hoa và 1 ký tự đặc biệt"})
		return
	}

	// Mật khẩu phải có ký tự đặc biệt
	specialCharRegex := regexp.MustCompile(`[^a-zA-Z0-9]`)
	if !specialCharRegex.MatchString(input.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mật khẩu phải có ít nhất 1 chữ hoa và 1 ký tự đặc biệt"})
		return
	}

	// Email phải đúng định dạng
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(input.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email không hợp lệ"})
		return
	}

	// Email phải là duy nhất
	existingUser, err := h.repo.GetUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi server"})
		return
	}
	if existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email đã được sử dụng"})
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

	data := dto.UserResponse{
		Username:      user.Username,
		Email:         user.Email,
		WalletAddress: user.WalletAddress,
		Role:          user.Role,
		Profile: dto.ProfileResponse{
			Avatar:      user.Profile.Avatar,
			Bio:         user.Profile.Bio,
			Preferences: user.Profile.Preferences,
		},
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusCreated, dto.ApiResponse{
		Success: true,
		Message: "Đăng ký thành công",
		Data:    data,
	})
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

// GetUserByEmail: GET /users/email/:email
func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := h.repo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi server"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Người dùng không tồn tại"})
		return
	}
	c.JSON(http.StatusOK, user)
}
