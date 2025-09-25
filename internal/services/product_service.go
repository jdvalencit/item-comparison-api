package services

import (
	"item-comparison-api/internal/models"
	"item-comparison-api/internal/repository"
)

// ProductService provides methods to interact with products using a repository interface.
type ProductService struct {
	Repo repository.ProductRepo
}

// GetAllProducts loads all products from the repository.
func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	products, err := s.Repo.LoadProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

// SaveAllProducts saves all products to the repository.
func (s *ProductService) SaveAllProducts(products []models.Product) error {
	if err := s.Repo.SaveProducts(products); err != nil {
		return err
	}
	return nil
}
