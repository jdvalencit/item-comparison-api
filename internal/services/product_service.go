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

// NewProductService creates a new instance of ProductService with the given repository.
func NewProductService(r repository.ProductRepo) *ProductService {
	fmt.Printf("[ProductService][NewProductService] Initializing ProductService\n")
	return &ProductService{repo: r}
}

// LoadProducts loads all products from the repository.
func (s *ProductService) LoadProducts() ([]dto.ProductResponse, error) {
	fmt.Printf("[ProductService][LoadProducts] Called\n")

	fmt.Printf("[ProductService][LoadProducts] Loading all products from repository\n")
	products, err := s.repo.LoadProducts()
	if err != nil {
		fmt.Printf("[ProductService][LoadProducts][ERROR] %v\n", err)
		return nil, err
	}

	fmt.Printf("[ProductService][LoadProducts] Loaded %d products\n", len(products))
	responseProducts := ProductsToResponses(products)
	return responseProducts, nil
}

// SaveProducts saves all products to the repository.
func (s *ProductService) SaveProducts(req []dto.ProductRequest, sellerID string) error {
	fmt.Printf("[ProductService][SaveProducts] Called with %d products, sellerID: %s\n", len(req), sellerID)

	// Define a slice to hold products to save and set CreatedAt and SellerID for each product
	productsToSave := make([]models.Product, 0, len(req))
	createdAt := time.Now().Format(time.RFC3339)
	for _, r := range req {
		product := ProductFromRequest(r)
		product.CreatedAt = createdAt
		product.SellerID = sellerID
		productsToSave = append(productsToSave, product)
	}

	fmt.Printf("[ProductService][SaveProducts] Saving products to repository\n")
	if err := s.repo.SaveProducts(productsToSave); err != nil {
		fmt.Printf("[ProductService][SaveProducts][ERROR] %v\n", err)
		return err
	}

	fmt.Printf("[ProductService][SaveProducts] Successfully saved %d products\n", len(productsToSave))
	return nil
}

func (s *ProductService) UpdateProducts(req []dto.ProductRequest, sellerID string) error {
	fmt.Printf("[ProductService][UpdateProducts] Called with %d products, sellerID: %s\n", len(req), sellerID)

	// Define a slice to hold the products to update
	productsToUpdate := make([]models.Product, 0, len(req))
	// Convert requests to models
	productsRequest := ProductsFromRequests(req)

	for _, product := range productsRequest {
		fmt.Printf("[ProductService][UpdateProducts] Getting product by ID: %d from repository\n", product.ID)

		// Retrieve the existing product from the repository
		existingProduct, err := s.repo.GetProductByID(product.ID)
		if err != nil {
			// If the product does not exist, we provide CreatedAt and SellerID for the new product
			if strings.Contains(err.Error(), "does not exist") {
				createdAt := time.Now().Format(time.RFC3339)
				product.CreatedAt = createdAt
				product.SellerID = sellerID
			}

			// if the sellerID does not match, return an error
		} else if existingProduct.SellerID != sellerID {
			fmt.Printf("[ProductService][UpdateProducts][ERROR] Unauthorized update attempt for product ID: %d\n", product.ID)
			return fmt.Errorf("unauthorized: you do not own this product")
		}

		// If the product exists, retain its CreatedAt and SellerID
		product.CreatedAt = existingProduct.CreatedAt
		product.SellerID = existingProduct.SellerID
		productsToUpdate = append(productsToUpdate, product)
	}

	fmt.Printf("[ProductService][UpdateProducts] Updating products in repository\n")
	if err := s.repo.UpdateProducts(productsToUpdate); err != nil {
		fmt.Printf("[ProductService][UpdateProducts][ERROR] %v\n", err)
		return err
	}

	fmt.Printf("[ProductService][UpdateProducts] Successfully updated %d products\n", len(productsToUpdate))
	return nil
}

func (s *ProductService) CompareProducts(ids []int) ([]dto.ProductResponse, error) {
	fmt.Printf("[ProductService][CompareProducts] Called with IDs: %v\n", ids)
	fmt.Printf("[ProductService][CompareProducts] Comparing products by IDs in repository\n")
	products, err := s.repo.CompareProducts(ids)
	if err != nil {
		fmt.Printf("[ProductService][CompareProducts][ERROR] %v\n", err)
		return nil, err
	}

	fmt.Printf("[ProductService][CompareProducts] Compared %d products\n", len(products))
	responseProducts := ProductsToResponses(products)
	return responseProducts, nil
}

func (s *ProductService) GetProductByID(id int) (*dto.ProductResponse, error) {
	fmt.Printf("[ProductService][GetProductByID] Called with ID: %d\n", id)
	fmt.Printf("[ProductService][GetProductByID] Getting product by ID from repository\n")
	product, err := s.repo.GetProductByID(id)
	if err != nil {
		fmt.Printf("[ProductService][GetProductByID][ERROR] %v\n", err)
		return nil, err
	}

	response := ProductToResponse(*product)
	fmt.Printf("[ProductService][GetProductByID] Successfully loaded product ID: %d\n", id)
	return &response, nil
}

func (s *ProductService) DeleteProductByID(id int, sellerID string) error {
	fmt.Printf("[ProductService][DeleteProductByID] Called with ID: %d, sellerID: %s\n", id, sellerID)
	fmt.Printf("[ProductService][DeleteProductByID] Getting product by ID from repository\n")
	product, err := s.repo.GetProductByID(id)
	if err != nil {
		fmt.Printf("[ProductService][DeleteProductByID][ERROR] %v\n", err)
		return err
	}

	// if the sellerID does not match, return an error
	if product.SellerID != sellerID {
		fmt.Printf("[ProductService][DeleteProductByID][ERROR] Unauthorized delete attempt for product ID: %d\n", id)
		return fmt.Errorf("unauthorized: you do not own this product")
	}

	fmt.Printf("[ProductService][DeleteProductByID] Deleting product by ID in repository\n")
	if err := s.repo.DeleteByID(id); err != nil {
		fmt.Printf("[ProductService][DeleteProductByID][ERROR] %v\n", err)
		return err
	}

	fmt.Printf("[ProductService][DeleteProductByID] Successfully deleted product ID: %d\n", id)
	return nil
}
