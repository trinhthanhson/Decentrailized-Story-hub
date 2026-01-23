package main

import (
	"user-service/internal/database"
	"user-service/internal/repository"
	"user-service/internal/transport/http"
	"user-service/internal/transport/http/middleware" // Import middleware của bạn

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "user-service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title User Service API

// @version 1.0

// @description API quản lý user (Gin + Swagger)

// @host localhost:8080

// @BasePath /
func main() {
	// 1. Kết nối DB
	db := database.InitDB()

	// 2. Khởi tạo Repository & Handler
	userRepo := repository.NewUserRepository(db)
	userHandler := http.NewUserHandler(userRepo)

	// 3. Khởi tạo Gin
	r := gin.Default()

	// --- QUAN TRỌNG: Cấu hình CORS để UI có thể gọi API ---
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000"}, // URL của React/Vue
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// 4. Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 5. Định nghĩa Routes

	// Nhóm các route công khai (không cần login)
	auth := r.Group("/auth")
	{
		auth.POST("/login", userHandler.Login)
		auth.POST("/register", userHandler.CreateUser) // Thường register nằm ở auth hoặc công khai
	}

	// Nhóm các route cần bảo mật (phải có Token)
	users := r.Group("/users")
	users.Use(middleware.AuthMiddleware()) // Áp dụng bảo vệ cho cả nhóm
	{
		users.GET("/", userHandler.ListUsers)
		users.GET("/email/:email", userHandler.GetUserByEmail)
		users.GET("/username/:username", userHandler.GetUserByUsername)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}

	// 6. Chạy Server
	r.Run(":8080")
}
