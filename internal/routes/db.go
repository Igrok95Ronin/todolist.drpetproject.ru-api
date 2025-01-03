package routes

import (
	"fmt"
	"github.com/Igrok95Ronin/todolist.drpetproject.ru-golang.git/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

// Получаем поля все записей
type AllNotes struct {
	ID        int64     `gorm:"primaryKey:column:id"`
	Note      string    `gorm:"column:note"`
	Completed bool      `gorm:"column:completed"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func InitDB() (db *gorm.DB) {

	cfg := config.GetConfig()

	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=%s client_encoding=UTF8", cfg.DB.User, cfg.DB.DBName, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.SslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return nil
	}

	return db
}
