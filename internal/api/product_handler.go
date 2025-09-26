package api

import (
	"encoding/json"
	"item-comparison-api/internal/dto"
	"item-comparison-api/internal/services"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(s *services.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) LoadProducts(w http.ResponseWriter, r *http.Request) {
	responseProducts, err := h.service.LoadProducts()
	if err != nil {
		http.Error(w, "Failed to load products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the list of products in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseProducts)
}

func (h *ProductHandler) SaveProducts(w http.ResponseWriter, r *http.Request) {
	// Validate if the x-seller-id header is present
	sellerID := r.Header.Get("x-seller-id")
	if sellerID == "" {
		http.Error(w, "Missing x-seller-id header", http.StatusBadRequest)
		return
	}

	// Decode new products from request body
	var newProductsRequest []dto.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&newProductsRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Save all products
	if err := h.service.SaveProducts(newProductsRequest, sellerID); err != nil {
		http.Error(w, "Failed to save products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Products saved successfully"))
}

func (h *ProductHandler) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	// Validate if the x-seller-id header is present
	sellerID := r.Header.Get("x-seller-id")
	if sellerID == "" {
		http.Error(w, "Missing x-seller-id header", http.StatusBadRequest)
		return
	}

	// Decode updated products from request body
	var updatedProductsRequest []dto.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&updatedProductsRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update products
	if err := h.service.UpdateProducts(updatedProductsRequest, sellerID); err != nil {
		http.Error(w, "Failed to update products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Products updated successfully"))
}

func (h *ProductHandler) CompareProducts(w http.ResponseWriter, r *http.Request) {
	// Decode product IDs from request body
	var req struct {
		IDs []int `json:"ids"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if len(req.IDs) == 0 {
		http.Error(w, "No product IDs provided", http.StatusBadRequest)
		return
	}

	// Compare products by IDs
	responseProducts, err := h.service.CompareProducts(req.IDs)
	if err != nil {
		http.Error(w, "Failed to compare products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the compared products in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseProducts)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	// Extract product ID from URL parameters
	idParam := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	if idParam == "" {
		http.Error(w, "Missing product ID", http.StatusBadRequest)
		return
	}

	// Convert ID to integer
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Get product by ID
	responseProduct, err := h.service.GetProductByID(id)
	if err != nil {
		http.Error(w, "Failed to get product: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if responseProduct == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Respond with the product in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseProduct)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Validate if the x-seller-id header is present
	sellerID := r.Header.Get("x-seller-id")
	if sellerID == "" {
		http.Error(w, "Missing x-seller-id header", http.StatusBadRequest)
		return
	}

	// Extract product ID from URL parameters
	idParam := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	if idParam == "" {
		http.Error(w, "Missing product ID", http.StatusBadRequest)
		return
	}

	// Convert ID to integer
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Delete product by ID
	if err := h.service.DeleteProductByID(id, sellerID); err != nil {
		http.Error(w, "Failed to delete product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product deleted successfully"))
}
