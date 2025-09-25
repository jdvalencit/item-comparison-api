package repository

import (
	"encoding/json"
	"io"
	"item-comparison-api/internal/models"
	"os"
)

// TODO: externalizar el path a un archivo de configuración
const productsFile = "data/products.json"

type ProductRepoJson struct{}

// LoadProducts loads products from a JSON file into a slice of Product structs.
func (r *ProductRepoJson) LoadProducts() ([]models.Product, error) {
	file, err := os.Open(productsFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var products []models.Product
	if err := json.Unmarshal(bytes, &products); err != nil {
		return nil, err
	}
	return products, nil
}

// SaveProducts saves a slice of Product structs to a JSON file.
func (r *ProductRepoJson) SaveProducts(products []models.Product) error {
	bytes, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(productsFile, bytes, 0644); err != nil {
		return err
	}
	return nil
}
