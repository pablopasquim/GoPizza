package main

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/pablopasquim/GoPizza/internal/routes"
)

func main() {
	fmt.Println("🚀 Starting GoPizza server...")

	// Create a Gin router
	r := gin.Default()

	// Register routes
	routes.HandleRequest(r)

	// Log before starting the server (Run is blocking)
	log.Println("✅ Server running at http://localhost:8080")

	// Start the server and handle errors
	if err := r.Run(); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}

