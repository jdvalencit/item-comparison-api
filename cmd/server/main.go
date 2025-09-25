package main

import (
	"item-comparison-api/internal"
	"log"
	"net/http"
)

func main() {
	r := internal.SetupRouter()
	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
