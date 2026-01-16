// Package request cho register: danh sách thông tin để gửi request đăng ký
package dto

// RegisterRequest Định nghĩa một cấu trúc riêng để nhận dữ liệu từ Postman (Register Request)
type RegisterRequest struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password      string `json:"password"` // Nhận pass từ Postman
	WalletAddress string `json:"wallet_address"`
	Profile       struct {
		Bio string `json:"bio"`
	} `json:"profile"`
}
