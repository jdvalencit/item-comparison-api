package main

import (
	"item-comparison-api/internal"
	"item-comparison-api/internal/api"
	"item-comparison-api/internal/repository"
	"item-comparison-api/internal/services"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load("config.env"); err != nil {
		log.Println("No .env file found")
	}

	storagePath := os.Getenv("STORAGE_PATH")

	// Initialize repository, service and handler
	repo := repository.NewProductRepo(storagePath)
	service := services.NewProductService(repo)
	handler := api.NewProductHandler(service)

	// Setup router and start server
	r := internal.SetupRouter(handler)
	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
