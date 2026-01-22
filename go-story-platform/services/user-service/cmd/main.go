package main

import (
	"user-service/internal/database"
	"user-service/internal/repository"
	"user-service/internal/transport/http"

	"github.com/gin-gonic/gin"

	// swagger
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

	// 4. Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 5. Định nghĩa Routes
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", userHandler.ListUsers)
		userRoutes.GET("/email/:email", userHandler.GetUserByEmail)
		userRoutes.GET("/username/:username", userHandler.GetUserByUsername)
		userRoutes.POST("/create", userHandler.CreateUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)

	}
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/login", userHandler.Login)
	}

	// 6. Chạy Server
	r.Run(":8080")
}
