package main

import (
	"fmt"
	"hello-gin/m/v2/model"
	"hello-gin/m/v2/router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	_ = godotenv.Load()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	appName := os.Getenv("APP_NAME")
	fmt.Println("Starting " + appName + " ...")

	// port 預設為 8080，若有設定環境變數 PORT 則使用環境變數的值
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize MySQL connection
	model.InitDB()

	router.RegisterRoutes(r)
	r.Run(":" + port)
}
