package main

import (
	"go-starter/config"
	"go-starter/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin
	router := gin.Default()

	// Initialize MongoDB
	db, err := config.ConnectMongoDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Set the database instance to the router context
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// Set up routes
	routes.SetupRoutes(router)

	// Run the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
