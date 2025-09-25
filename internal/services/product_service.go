package services

import (
	"item-comparison-api/internal/models"
	"item-comparison-api/internal/repository"
)

// ProductService provides methods to interact with products using a repository interface.
type ProductService struct {
	repo repository.ProductRepo
}

func NewProductService(r repository.ProductRepo) *ProductService {
	return &ProductService{repo: r}
}

// GetAllProducts loads all products from the repository.
func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	products, err := s.repo.LoadProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

// SaveAllProducts saves all products to the repository.
func (s *ProductService) SaveAllProducts(products []models.Product) error {
	if err := s.repo.SaveProducts(products); err != nil {
		return err
	}
	return nil
}

func (s *ProductService) UpdateProducts(products []models.Product) error {
	if err := s.repo.UpdateProducts(products); err != nil {
		return err
	}
	return nil
}

func (s *ProductService) CompareProducts(ids []int) ([]models.Product, error) {
	products, err := s.repo.CompareProducts(ids)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetProductByID(id int) (*models.Product, error) {
	product, err := s.repo.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) DeleteProductByID(id int) error {
	if err := s.repo.DeleteByID(id); err != nil {
		return err
	}
	return nil
}
