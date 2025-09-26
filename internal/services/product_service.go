package services

import (
	"fmt"
	"item-comparison-api/internal/dto"
	"item-comparison-api/internal/models"
	"item-comparison-api/internal/repository"
	"strings"
	"time"
)

// ProductService provides methods to interact with products using a repository interface.
type ProductService struct {
	repo repository.ProductRepo
}

func NewProductService(r repository.ProductRepo) *ProductService {
	return &ProductService{repo: r}
}

// GetAllProducts loads all products from the repository.
func (s *ProductService) LoadProducts() ([]dto.ProductResponse, error) {
	products, err := s.repo.LoadProducts()
	if err != nil {
		return nil, err
	}

	// Map product models to response DTOs
	responseProducts := ProductsToResponses(products)
	return responseProducts, nil
}

// SaveAllProducts saves all products to the repository.
func (s *ProductService) SaveProducts(req []dto.ProductRequest, sellerID string) error {
	productsToSave := make([]models.Product, 0, len(req))

	// Set createdAt timespamp
	createdAt := time.Now().Format(time.RFC3339)

	// Map request DTOs to product models and set createdAt
	for _, r := range req {
		product := ProductFromRequest(r)
		product.CreatedAt = createdAt
		product.SellerID = sellerID
		productsToSave = append(productsToSave, product)
	}

	if err := s.repo.SaveProducts(productsToSave); err != nil {
		return err
	}
	return nil
}

func (s *ProductService) UpdateProducts(req []dto.ProductRequest, sellerID string) error {
	productsToUpdate := make([]models.Product, 0, len(req))
	// Map request DTOs to product models
	productsRequest := ProductsFromRequests(req)

	for _, product := range productsRequest {
		existingProduct, err := s.repo.GetProductByID(product.ID)
		if err != nil {
			if strings.Contains(err.Error(), "does not exist") {
				createdAt := time.Now().Format(time.RFC3339)
				product.CreatedAt = createdAt
				product.SellerID = sellerID
			}
		} else if existingProduct.SellerID != sellerID {
			return fmt.Errorf("unauthorized: you do not own this product")
		}
		// Mantener datos originales
		product.CreatedAt = existingProduct.CreatedAt
		product.SellerID = existingProduct.SellerID

		productsToUpdate = append(productsToUpdate, product)
	}

	if err := s.repo.UpdateProducts(productsToUpdate); err != nil {
		return err
	}
	return nil
}

func (s *ProductService) CompareProducts(ids []int) ([]dto.ProductResponse, error) {
	products, err := s.repo.CompareProducts(ids)
	if err != nil {
		return nil, err
	}

	// Map product models to response DTOs
	responseProducts := ProductsToResponses(products)
	return responseProducts, nil
}

func (s *ProductService) GetProductByID(id int) (*dto.ProductResponse, error) {
	product, err := s.repo.GetProductByID(id)
	if err != nil {
		return nil, err
	}
	response := ProductToResponse(*product)
	return &response, nil
}

func (s *ProductService) DeleteProductByID(id int, sellerID string) error {
	product, err := s.repo.GetProductByID(id)
	if err != nil {
		return err
	}

	if product.SellerID != sellerID {
		return fmt.Errorf("unauthorized: you do not own this product")
	}

	if err := s.repo.DeleteByID(id); err != nil {
		return err
	}
	return nil
}
