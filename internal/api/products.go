package api

import (
	"encoding/json"
	"net/http"

	"item-comparison-api/internal/services"
)

func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := services.GetAllProducts()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

/*
func GetProduct(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	product := services.GetProductByID(id)
	if product == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
*/
