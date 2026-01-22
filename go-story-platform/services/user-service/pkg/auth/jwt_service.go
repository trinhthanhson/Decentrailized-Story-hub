package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	SecretKey     string
	Issuer        string
	TokenDuration int64 // in minutes
}

// CustomClaims định nghĩa các thông tin bổ sung trong JWT
type CustomClaims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// NewJWTService khởi tạo JWTService với thời hạn 1 ngày
func NewJWTService(secretKey, issuer string) *JWTService {
	return &JWTService{
		SecretKey:     secretKey,
		Issuer:        issuer,
		TokenDuration: 1440, // 1 ngày
	}
}

// GenerateToken tạo JWT cho user với ID và vai trò
func (j *JWTService) GenerateToken(userID uint, role string) (string, error) {
	claims := &CustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.TokenDuration) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.SecretKey))
}

// ValidateToken kiểm tra tính hợp lệ của JWT và trả về các thông tin trong token
// ValidateToken kiểm tra tính hợp lệ và giải mã token
func (j *JWTService) ValidateToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Kiểm tra thuật toán ký để tránh lỗi bảo mật "None" attack
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Kiểm tra ép kiểu và tính hợp lệ của Issuer
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		if claims.Issuer != j.Issuer {
			return nil, fmt.Errorf("invalid issuer")
		}
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}
