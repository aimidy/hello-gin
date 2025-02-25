package main

import (
	"fmt"
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.RegisterRoutes(r)
	r.Run(":" + port)
}
