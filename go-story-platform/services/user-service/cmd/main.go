package main

import (
	"user-service/internal/database"
	"user-service/internal/repository"
	"user-service/internal/transport/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Kết nối DB
	db := database.InitDB()

	// 2. Khởi tạo Repository & Handler
	userRepo := repository.NewUserRepository(db)
	userHandler := http.NewUserHandler(userRepo)

	// 3. Khởi tạo Gin
	r := gin.Default()

	// 4. Định nghĩa Routes
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", userHandler.ListUsers)        // Thêm dòng này để lấy danh sách
		userRoutes.POST("/", userHandler.CreateUser)      // Tạo mới
		userRoutes.PUT("/:id", userHandler.UpdateUser)    // Cập nhật
		userRoutes.DELETE("/:id", userHandler.DeleteUser) // Xóa
	}

	// 5. Chạy Server
	r.Run(":8080")
}
