package services

import (
	"item-comparison-api/internal/dto"
	"item-comparison-api/internal/models"
)

// Maps from request dto to model
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

// Maps from model to response dto
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

// Maps a list of models form a list of requests
func ProductsFromRequests(requests []dto.ProductRequest) []models.Product {
	products := make([]models.Product, 0, len(requests))
	for _, r := range requests {
		products = append(products, ProductFromRequest(r))
	}
	return products
}

// Maps a list of responses from a list of models
func ProductsToResponses(products []models.Product) []dto.ProductResponse {
	responses := make([]dto.ProductResponse, 0, len(products))
	for _, p := range products {
		responses = append(responses, ProductToResponse(p))
	}
	return responses
}
