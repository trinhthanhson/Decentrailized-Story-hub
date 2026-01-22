package middleware

import (
	"net/http"
	"os"
	"strings"

	"user-service/internal/transport/http/dto"
	"user-service/pkg/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Lấy header Authorization: Bearer <token>
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ApiResponse{
				Success: false,
				Message: "Bạn cần đăng nhập để truy cập",
			})
			return
		}
		// Tách chuỗi để lấy token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ApiResponse{
				Success: false,
				Message: "Định dạng Token phải là 'Bearer <token>'",
			})
			return
		}
		//Gọi service để xác thực
		jwtSvc := auth.NewJWTService(
			os.Getenv("JWT_SECRET"),
			os.Getenv("ISSUER"),
		)
		claims, err := jwtSvc.ValidateToken(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ApiResponse{
				Success: false,
				Message: "Token không hợp lệ: " + err.Error(),
			})
			return
		}
		// Lưu thông tin user vào context để các handler khác sử dụng
		c.Set("userID", claims.UserID)
		c.Set("userRole", claims.Role)

		c.Next()
	}
}
