package repository

import (
	"encoding/json"
	"fmt"
	"item-comparison-api/internal/models"
	"os"
	"path/filepath"
)

type ProductRepoJson struct{}

func NewProductRepo() *ProductRepoJson {
	return &ProductRepoJson{}
}

// LoadProducts loads products from a JSON file into a slice of Product structs.
func (r *ProductRepoJson) LoadProducts() ([]models.Product, error) {
	// TODO: archivo de configuración
	dir := "data"
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			// Directory does not exist, return empty slice
			return []models.Product{}, nil
		}
		return []models.Product{}, fmt.Errorf("failed to read directory %s: %w", dir, err)
	}

	var products []models.Product
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		// Only process .json files
		if filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		filePath := fmt.Sprintf("%s/%s", dir, entry.Name())
		bytes, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
		}
		var product models.Product
		if err := json.Unmarshal(bytes, &product); err != nil {
			return nil, fmt.Errorf("failed to unmarshal file %s: %w", filePath, err)
		}
		products = append(products, product)
	}
	return products, nil
}

// SaveProducts saves a slice of Product structs to a JSON file.
func (r *ProductRepoJson) SaveProducts(products []models.Product) error {
	for _, p := range products {
		fileName := fmt.Sprintf("data/%d.json", p.ID)
		if _, err := os.Stat(fileName); err == nil {
			return fmt.Errorf("product with ID %d already exists", p.ID)
		}

		bytes, err := json.MarshalIndent(p, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal product %s: %w", fileName, err)
		}

		if err := os.WriteFile(fileName, bytes, 0644); err != nil {
			return fmt.Errorf("failed to write product %s: %w", fileName, err)
		}
	}
	return nil
}

func (r *ProductRepoJson) UpdateProducts(products []models.Product) error {
	for _, p := range products {
		fileName := fmt.Sprintf("data/%d.json", p.ID)

		bytes, err := json.MarshalIndent(p, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal product %s: %w", fileName, err)
		}

		if err := os.WriteFile(fileName, bytes, 0644); err != nil {
			return fmt.Errorf("failed to write product %s: %w", fileName, err)
		}
	}
	return nil
}

func (r *ProductRepoJson) CompareProducts(ids []int) ([]models.Product, error) {
	var products []models.Product
	for _, id := range ids {
		product, err := r.GetProductByID(id)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}
	return products, nil
}

func (r *ProductRepoJson) GetProductByID(id int) (*models.Product, error) {
	fileName := fmt.Sprintf("data/%d.json", id)
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("product with ID %d does not exist", id)
		}
		return nil, fmt.Errorf("failed to read file %s: %w", fileName, err)
	}

	var product models.Product
	if err := json.Unmarshal(bytes, &product); err != nil {
		return nil, fmt.Errorf("failed to unmarshal file %s: %w", fileName, err)
	}
	return &product, nil
}

func (r *ProductRepoJson) DeleteByID(id int) error {
	fileName := fmt.Sprintf("data/%d.json", id)
	if err := os.Remove(fileName); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("product with ID %d does not exist", id)
		}
		return fmt.Errorf("failed to delete file %d: %w", id, err)
	}
	return nil
}
