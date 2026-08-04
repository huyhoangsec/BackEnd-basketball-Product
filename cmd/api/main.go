package main

import (
	"log"

	"backend-ocean-basketball/config"
	"backend-ocean-basketball/internal/models"
	"backend-ocean-basketball/internal/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Set Gin mode
	gin.SetMode(cfg.GIN_MODE)

	// Initialize Gin router
	r := gin.Default()

	// Setup CORS - Allow all origins dynamically for seamless Render & Vercel deployment
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Serve static files (uploads)
	r.Static("/uploads", "./uploads")

	// Connect to DB
	models.ConnectDatabase(cfg)

	// Setup other routes
	api := r.Group("/api")
	routes.SetupRoutes(api, cfg)

	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
