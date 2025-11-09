package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jay13jay/asstBackend/internal/api"
	"github.com/jay13jay/asstBackend/internal/config"
	"github.com/jay13jay/asstBackend/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Initialize configuration
	cfg := config.Load()

	// Initialize database
	db, err := database.NewConnection(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Run migrations
	if err := database.RunMigrations(cfg.DatabaseURL); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// Set Gin mode
	gin.SetMode(cfg.GinMode)

	// Initialize API server
	server := api.NewServer(db, cfg)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	if err := server.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
