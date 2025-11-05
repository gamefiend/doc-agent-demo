package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourorg/doc-agent-demo/internal/handlers"
	"github.com/yourorg/doc-agent-demo/internal/models"
)

func main() {
	// Create Gin router
	r := gin.Default()

	// API v1 group
	v1 := r.Group("/api/v1")
	{
		// Health check endpoint
		v1.GET("/health", handlers.HealthCheck)

		// User endpoints
		v1.GET("/users", handlers.ListUsers)
		v1.GET("/users/:id", handlers.GetUser)
		v1.POST("/users", handlers.CreateUser)
		v1.PUT("/users/:id", handlers.UpdateUser)
		v1.DELETE("/users/:id", handlers.DeleteUser)

		// Product endpoints
		v1.GET("/products", handlers.ListProducts)
		v1.GET("/products/:id", handlers.GetProduct)
		v1.POST("/products", handlers.CreateProduct)
	}

	// Initialize some sample data
	models.InitSampleData()

	// Start server
	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}
