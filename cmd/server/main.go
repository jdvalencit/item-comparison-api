package main

import (
	"item-comparison-api/internal"
	"item-comparison-api/internal/api"
	"item-comparison-api/internal/repository"
	"item-comparison-api/internal/services"
	"log"
	"net/http"
)

func main() {
	repo := repository.NewProductRepo()
	service := services.NewProductService(repo)
	handler := api.NewProductHandler(service)

	r := internal.SetupRouter(handler)
	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
