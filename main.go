package main

import (
	"log"
	"os"

	"github.com/dennisferdian9/golang-sqlite/config"
	_ "github.com/joho/godotenv/autoload"

	router "github.com/dennisferdian9/golang-sqlite/routes"
)

func main() {
	port := os.Getenv("PORT")
	config.InitDB()
	defer config.DB.Close()

	if port == "" {
		port = "8080" // Default port if not set
	}

	r := router.SetupRouter()

	log.Printf("Starting server on port %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Could not start server: %v\n", err)
	}
}
