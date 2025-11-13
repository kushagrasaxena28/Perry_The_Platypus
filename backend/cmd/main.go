package main

import (
	"Perry_the_Platypus/backend/features/auth"
	"Perry_the_Platypus/backend/features/content"
	"Perry_the_Platypus/backend/shared/config"
	"Perry_the_Platypus/backend/shared/database"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	// Load .env file
	config.LoadEnv()

	//Get the Port
	port := os.Getenv("PORT")
	fmt.Println(port)
	if port == "" {
		port = "8080"
	}
	// Initialize database connection
	dbConn, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer dbConn.Close(context.Background())

	// Initialize Gin router
	r := gin.Default()

	// Set trusted proxies
	_ = r.SetTrustedProxies([]string{"127.0.0.1"})

	// Apply logging middleware
	r.Use(gin.Logger())

	// Setup auth routes
	auth.AuthRoutes(r, dbConn)
	content.ContentRoutes(r, dbConn)

	// Start server
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
