package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	setupRoutes(router)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func setupRoutes(router *gin.Engine) {
	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
			"service": "go-rest-api",
			"version": "1.0.0",
		})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", registerHandler)
			auth.POST("/login", loginHandler)
		}

		// Protected routes
		protected := v1.Group("/")
		protected.Use(authMiddleware())
		{
			// Users
			protected.GET("/users", getUsersHandler)
			protected.GET("/users/:id", getUserHandler)
			protected.PUT("/users/:id", updateUserHandler)

			// Products
			protected.GET("/products", getProductsHandler)
			protected.POST("/products", createProductHandler)
			protected.GET("/products/:id", getProductHandler)
			protected.PUT("/products/:id", updateProductHandler)
			protected.DELETE("/products/:id", deleteProductHandler)
		}
	}
}

func registerHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Register endpoint"})
}

func loginHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Login endpoint"})
}

func getUsersHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get users"})
}

func getUserHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get user"})
}

func updateUserHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Update user"})
}

func getProductsHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get products"})
}

func createProductHandler(c *gin.Context) {
	c.JSON(201, gin.H{"message": "Product created"})
}

func getProductHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Get product"})
}

func updateProductHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Update product"})
}

func deleteProductHandler(c *gin.Context) {
	c.JSON(204, nil)
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
