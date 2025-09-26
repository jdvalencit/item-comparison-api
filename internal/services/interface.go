package services

import (
	"item-comparison-api/internal/dto"
)

// ProductServiceInterface defines the methods for product service operations
type ProductServiceInterface interface {
	LoadProducts() ([]dto.ProductResponse, error)
	SaveProducts(req []dto.ProductRequest, sellerID string) error
	UpdateProducts(req []dto.ProductRequest, sellerID string) error
	CompareProducts(ids []int) ([]dto.ProductResponse, error)
	GetProductByID(id int) (*dto.ProductResponse, error)
	DeleteProductByID(id int, sellerID string) error
}
