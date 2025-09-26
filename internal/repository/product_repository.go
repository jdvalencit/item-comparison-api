package repository

import "item-comparison-api/internal/models"

// ProductRepo defines the interface for product repository operations
type ProductRepo interface {
	LoadProducts() ([]models.Product, error)
	SaveProducts([]models.Product) error
	UpdateProducts([]models.Product) error
	CompareProducts([]int) ([]models.Product, error)
	GetProductByID(int) (*models.Product, error)
	DeleteByID(int) error
}
