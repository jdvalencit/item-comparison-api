package api

import (
	"encoding/json"
	"fmt"
	"item-comparison-api/internal/dto"
	"item-comparison-api/internal/services"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	service services.ProductServiceInterface
}

// NewProductHandler creates a new instance of ProductHandler
func NewProductHandler(s services.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{service: s}
}

// LoadProducts handles the loading of all products
func (h *ProductHandler) LoadProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[ProductHandler][LoadProducts] %s %s\n", r.Method, r.URL.String())
	// Load all products
	responseProducts, err := h.service.LoadProducts()
	if err != nil {
		fmt.Printf("[ProductHandler][LoadProducts][ERROR] %v\n", err)
		http.Error(w, "Failed to load products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the products in JSON format
	fmt.Printf("[ProductHandler][LoadProducts] Loaded %d products\n", len(responseProducts))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseProducts)
}

func (h *ProductHandler) SaveProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[ProductHandler][SaveProducts] %s %s\n", r.Method, r.URL.String())
	// Validate if the x-seller-id header is present
	sellerID := r.Header.Get("x-seller-id")
	if sellerID == "" {
		fmt.Printf("[ProductHandler][SaveProducts][ERROR] Missing x-seller-id header\n")
		http.Error(w, "Missing x-seller-id header", http.StatusBadRequest)
		return
	}

	// Decode new products from request body
	var newProductsRequest []dto.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&newProductsRequest); err != nil {
		fmt.Printf("[ProductHandler][SaveProducts][ERROR] %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Save all products
	if err := h.service.SaveProducts(newProductsRequest, sellerID); err != nil {
		fmt.Printf("[ProductHandler][SaveProducts][ERROR] %v\n", err)
		http.Error(w, "Failed to save products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	fmt.Printf("[ProductHandler][SaveProducts] Saved %d products for seller %s\n", len(newProductsRequest), sellerID)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Products saved successfully"))
}

func (h *ProductHandler) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[ProductHandler][UpdateProducts] %s %s\n", r.Method, r.URL.String())
	// Validate if the x-seller-id header is present
	sellerID := r.Header.Get("x-seller-id")
	if sellerID == "" {
		fmt.Printf("[ProductHandler][UpdateProducts][ERROR] Missing x-seller-id header\n")
		http.Error(w, "Missing x-seller-id header", http.StatusBadRequest)
		return
	}

	// Decode updated products from request body
	var updatedProductsRequest []dto.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&updatedProductsRequest); err != nil {
		fmt.Printf("[ProductHandler][UpdateProducts][ERROR] %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update products
	if err := h.service.UpdateProducts(updatedProductsRequest, sellerID); err != nil {
		fmt.Printf("[ProductHandler][UpdateProducts][ERROR] %v\n", err)
		http.Error(w, "Failed to update products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	fmt.Printf("[ProductHandler][UpdateProducts] Updated %d products for seller %s\n", len(updatedProductsRequest), sellerID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Products updated successfully"))
}

func (h *ProductHandler) CompareProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[ProductHandler][CompareProducts] %s %s\n", r.Method, r.URL.String())
	// Decode product IDs from request body
	var req struct {
		IDs []int `json:"ids"`
	}

	// Decode product IDs from request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Printf("[ProductHandler][CompareProducts][ERROR] %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate if at least one ID is provided
	if len(req.IDs) == 0 {
		fmt.Printf("[ProductHandler][CompareProducts][ERROR] No product IDs provided\n")
		http.Error(w, "No product IDs provided", http.StatusBadRequest)
		return
	}

	// Compare products by IDs
	responseProducts, err := h.service.CompareProducts(req.IDs)
	if err != nil {
		fmt.Printf("[ProductHandler][CompareProducts][ERROR] %v\n", err)
		http.Error(w, "Failed to compare products: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the compared products in JSON format
	fmt.Printf("[ProductHandler][CompareProducts] Compared %d products\n", len(responseProducts))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseProducts)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[ProductHandler][GetProduct] %s %s\n", r.Method, r.URL.String())
	// Extract product ID from URL parameters
	idParam := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	if idParam == "" {
		fmt.Printf("[ProductHandler][GetProduct][ERROR] Missing product ID\n")
		http.Error(w, "Missing product ID", http.StatusBadRequest)
		return
	}

	// Convert ID to integer
	id, err := strconv.Atoi(idParam)
	if err != nil {
		fmt.Printf("[ProductHandler][GetProduct][ERROR] Invalid product ID: %v\n", err)
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Get product by ID
	responseProduct, err := h.service.GetProductByID(id)
	if err != nil {
		fmt.Printf("[ProductHandler][GetProduct][ERROR] %v\n", err)
		http.Error(w, "Failed to get product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// verify if product was found
	if responseProduct == nil {
		fmt.Printf("[ProductHandler][GetProduct][ERROR] Product not found: %d\n", id)
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Respond with the product in JSON format
	fmt.Printf("[ProductHandler][GetProduct] Returned product ID %d\n", id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseProduct)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[ProductHandler][DeleteProduct] %s %s\n", r.Method, r.URL.String())
	// Validate if the x-seller-id header is present
	sellerID := r.Header.Get("x-seller-id")
	if sellerID == "" {
		fmt.Printf("[ProductHandler][DeleteProduct][ERROR] Missing x-seller-id header\n")
		http.Error(w, "Missing x-seller-id header", http.StatusBadRequest)
		return
	}

	// Extract product ID from URL parameters
	idParam := r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:]
	if idParam == "" {
		fmt.Printf("[ProductHandler][DeleteProduct][ERROR] Missing product ID\n")
		http.Error(w, "Missing product ID", http.StatusBadRequest)
		return
	}

	// Convert ID to integer
	id, err := strconv.Atoi(idParam)
	if err != nil {
		fmt.Printf("[ProductHandler][DeleteProduct][ERROR] Invalid product ID: %v\n", err)
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Delete product by ID
	if err := h.service.DeleteProductByID(id, sellerID); err != nil {
		fmt.Printf("[ProductHandler][DeleteProduct][ERROR] %v\n", err)
		http.Error(w, "Failed to delete product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	fmt.Printf("[ProductHandler][DeleteProduct] Deleted product ID %d for seller %s\n", id, sellerID)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product deleted successfully"))
}
