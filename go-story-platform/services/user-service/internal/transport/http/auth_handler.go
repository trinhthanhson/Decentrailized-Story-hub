package http

import (
	"net/http"
	"os"
	"user-service/internal/transport/http/dto"
	"user-service/pkg/auth"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login User godoc
// @Summary Đăng nhập user
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body dto.LoginRequest true "Login credentials"
// @Success 200 {object} dto.ApiResponse
// @Router /auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var loginReq dto.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ: " + err.Error()})
		return
	}
	user, err := h.repo.GetUserByUsername(loginReq.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi server: " + err.Error()})
		return
	}
	if user == nil || !checkPasswordHash(loginReq.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Email hoặc mật khẩu không đúng"})
		return
	}
	jwtSvc := auth.NewJWTService(
		os.Getenv("JWT_SECRET"),
		os.Getenv("ISSUER"),
	)
	token, err := jwtSvc.GenerateToken(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Không thể tạo token: " + err.Error()})
		return
	}
	data := dto.LoginResponse{
		AccessToken: token,
		User: dto.UserResponse{
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
		},
	}
	c.JSON(http.StatusOK, dto.ApiResponse{
		Success: true,
		Message: "Đăng nhập thành công",
		Data:    data,
	})
}

// checkPasswordHash kiểm tra mật khẩu đã băm với mật khẩu gốc
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
