package database

import (
	"content-service/internal/models" // Thay bằng tên module của bạn
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// Load biến môi trường
	godotenv.Load()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Không thể kết nối User DB:", err)
	}

	// Auto Migrate các bảng theo đúng model bạn đã tạo
	db.AutoMigrate(&models.Book{}, &models.Category{}, &models.Chapter{})

	fmt.Println("✅ User Service: Database connected & Migrated")
	return db
}
