package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yourorg/doc-agent-demo/internal/handlers"
	"github.com/yourorg/doc-agent-demo/internal/models"
)

// RequestLoggerMiddleware logs API requests with timing
func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Process request
		c.Next()

		// Log after request
		duration := time.Since(start)
		status := c.Writer.Status()
		log.Printf("[%s] %s - Status: %d - Duration: %v", method, path, status, duration)
	}
}

func main() {
	// Create Gin router
	r := gin.Default()

	// Add custom middleware
	r.Use(RequestLoggerMiddleware())

	// API v1 group
	v1 := r.Group("/api/v1")
	{
		// Health check endpoints
		v1.GET("/health", handlers.HealthCheck)
		v1.GET("/health/details", handlers.HealthDetails) // NEW: Detailed system health

		// User endpoints
		v1.GET("/users", handlers.ListUsers)
		v1.GET("/users/:id", handlers.GetUser)
		v1.GET("/users/:id/profile", handlers.GetUserProfile) // NEW: Enhanced profile endpoint
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
