package model

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"size:255"`
	Account   string `gorm:"size:255;unique"`
	Password  string `gorm:"size:255"`
	Detail    string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func InitDB() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN is not set in the environment variables")
	}

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Database connected successfully")

	// Auto migrate the User model
	DB.AutoMigrate(&User{})
}
