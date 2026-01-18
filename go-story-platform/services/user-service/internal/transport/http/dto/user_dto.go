package dto

import "time"

type RegisterRequest struct {
	Username      string         `json:"username"`
	Email         string         `json:"email"`
	Password      string         `json:"password"`
	WalletAddress string         `string:"wallet_address"`
	Profile       ProfileRequest `string:"profile"`
}

type UserResponse struct {
	Username      string          `json:"username"`
	Email         string          `json:"email"`
	WalletAddress string          `json:"wallet_address"`
	Role          string          `json:"role"`
	Profile       ProfileResponse `json:"profile_response"`
	CreatedAt     time.Time       `json:"created_at"`
}
