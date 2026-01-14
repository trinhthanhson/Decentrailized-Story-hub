package main

import (
	"content-service/internal/database"
	"fmt"
)

func main() {
	// 1. Kết nối DB
	db := database.InitDB()

	// Thử kiểm tra kết nối bằng cách in ra thông tin DB
	sqlDB, _ := db.DB()
	err := sqlDB.Ping()
	if err == nil {
		fmt.Println("🚀 User Service đang chạy và sẵn sàng!")
	}

	// Sau này bạn sẽ khởi chạy Server HTTP (Gin) hoặc gRPC ở đây
	select {} // Giữ cho main không bị thoát
}
