package services

import (
	"item-comparison-api/internal/dto"
	"item-comparison-api/internal/models"
)

// Convierte de request a modelo
func ProductFromRequest(req dto.ProductRequest) models.Product {
	return models.Product{
		ID:             req.ID,
		Name:           req.Name,
		Description:    req.Description,
		Price:          req.Price,
		Brand:          req.Brand,
		ImageUrl:       req.ImageUrl,
		Rating:         req.Rating,
		Specifications: req.Specifications,
	}
}

// Convierte de modelo a response
func ProductToResponse(p models.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:       p.ID,
		Name:     p.Name,
		Price:    p.Price,
		Brand:    p.Brand,
		ImageUrl: p.ImageUrl,
		Rating:   p.Rating,
		SellerID: p.SellerID,
	}
}

// Devueve una lista de modelos a partir de una lista de requests
func ProductsFromRequests(requests []dto.ProductRequest) []models.Product {
	products := make([]models.Product, 0, len(requests))
	for _, r := range requests {
		products = append(products, ProductFromRequest(r))
	}
	return products
}

// Convierte una lista de modelos a una lista de responses
func ProductsToResponses(products []models.Product) []dto.ProductResponse {
	responses := make([]dto.ProductResponse, 0, len(products))
	for _, p := range products {
		responses = append(responses, ProductToResponse(p))
	}
	return responses
}
