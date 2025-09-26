package repository

import (
	"encoding/json"
	"fmt"
	"item-comparison-api/internal/models"
	"os"
	"path/filepath"
)

type ProductRepoJson struct {
	storagePath string
}

// NewProductRepo creates a new instance of ProductRepoJson
func NewProductRepo(path string) *ProductRepoJson {
	if path == "" {
		path = "data/"
	}
	fmt.Printf("[Repository][NewProductRepo] Initializing with storagePath: %s\n", path)

	return &ProductRepoJson{
		storagePath: path,
	}
}

// LoadProducts loads products from a JSON file into a slice of Product structs.
func (r *ProductRepoJson) LoadProducts() ([]models.Product, error) {
	fmt.Printf("[Repository][LoadProducts] Loading all products from directory: %s\n", r.storagePath)
	dir := r.storagePath
	// Ensure the directory exists before reading
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("[Repository][LoadProducts] Directory does not exist: %s, returning empty slice\n", dir)
			return []models.Product{}, nil
		}
		fmt.Printf("[Repository][LoadProducts][ERROR] Failed to read directory: %s, error: %v\n", dir, err)
		return []models.Product{}, fmt.Errorf("failed to read directory %s: %w", dir, err)
	}

	// Read all JSON files in the directory
	var products []models.Product
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		// Only process .json files
		if filepath.Ext(entry.Name()) != ".json" {
			continue
		}

		// Read the file
		filePath := fmt.Sprintf("%s/%s", dir, entry.Name())
		fmt.Printf("[Repository][LoadProducts] Reading product file: %s\n", filePath)
		bytes, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("[Repository][LoadProducts][ERROR] Failed to read file: %s, error: %v\n", filePath, err)
			return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
		}

		// Unmarshal the JSON data into a Product struct
		var product models.Product
		if err := json.Unmarshal(bytes, &product); err != nil {
			fmt.Printf("[Repository][LoadProducts][ERROR] Failed to unmarshal file: %s, error: %v\n", filePath, err)
			return nil, fmt.Errorf("failed to unmarshal file %s: %w", filePath, err)
		}

		// Append the product to the slice
		products = append(products, product)
	}

	fmt.Printf("[Repository][LoadProducts] Loaded %d products\n", len(products))
	return products, nil
}

// SaveProducts saves a slice of Product structs to a JSON file.
func (r *ProductRepoJson) SaveProducts(products []models.Product) error {
	fmt.Printf("[Repository][SaveProducts] Saving %d products\n", len(products))
	dir := r.storagePath

	for _, p := range products {
		// Create the file name based on product ID
		fileName := fmt.Sprintf("%s/%d.json", dir, p.ID)
		fmt.Printf("[Repository][SaveProducts] Saving product ID: %d to file: %s\n", p.ID, fileName)

		// Check if file already exists to prevent overwriting
		if _, err := os.Stat(fileName); err == nil {
			fmt.Printf("[Repository][SaveProducts][ERROR] Product with ID %d already exists\n", p.ID)
			return fmt.Errorf("product with ID %d already exists", p.ID)
		}

		// Marshal the product to JSON
		bytes, err := json.MarshalIndent(p, "", "  ")
		if err != nil {
			fmt.Printf("[Repository][SaveProducts][ERROR] Failed to marshal product ID: %d, error: %v\n", p.ID, err)
			return fmt.Errorf("failed to marshal product %s: %w", fileName, err)
		}

		// Write the JSON data to the file
		if err := os.WriteFile(fileName, bytes, 0644); err != nil {
			fmt.Printf("[Repository][SaveProducts][ERROR] Failed to write product ID: %d to file: %s, error: %v\n", p.ID, fileName, err)
			return fmt.Errorf("failed to write product %s: %w", fileName, err)
		}
	}

	fmt.Printf("[Repository][SaveProducts] Successfully saved all products\n")
	return nil
}

func (r *ProductRepoJson) UpdateProducts(products []models.Product) error {
	fmt.Printf("[Repository][UpdateProducts] Updating %d products\n", len(products))
	dir := r.storagePath

	for _, p := range products {
		// Create the file name based on product ID
		fileName := fmt.Sprintf("%s/%d.json", dir, p.ID)
		fmt.Printf("[Repository][UpdateProducts] Updating product ID: %d in file: %s\n", p.ID, fileName)

		// Marshal the product to JSON
		bytes, err := json.MarshalIndent(p, "", "  ")
		if err != nil {
			fmt.Printf("[Repository][UpdateProducts][ERROR] Failed to marshal product ID: %d, error: %v\n", p.ID, err)
			return fmt.Errorf("failed to marshal product %s: %w", fileName, err)
		}

		// Write the JSON data to the file (overwrite existing file)
		if err := os.WriteFile(fileName, bytes, 0644); err != nil {
			fmt.Printf("[Repository][UpdateProducts][ERROR] Failed to write product ID: %d to file: %s, error: %v\n", p.ID, fileName, err)
			return fmt.Errorf("failed to write product %s: %w", fileName, err)
		}
	}

	fmt.Printf("[Repository][UpdateProducts] Successfully updated all products\n")
	return nil
}

func (r *ProductRepoJson) CompareProducts(ids []int) ([]models.Product, error) {
	fmt.Printf("[Repository][CompareProducts] Comparing products with IDs: %v\n", ids)
	var products []models.Product

	// Retrieve each product by ID and add it to the result slice
	for _, id := range ids {
		fmt.Printf("[Repository][CompareProducts] Getting product by ID: %d\n", id)
		product, err := r.GetProductByID(id)
		if err != nil {
			fmt.Printf("[Repository][CompareProducts][ERROR] Failed to get product ID: %d, error: %v\n", id, err)
			return nil, err
		}

		// Append the product to the result slice
		products = append(products, *product)
	}

	fmt.Printf("[Repository][CompareProducts] Compared %d products\n", len(products))
	return products, nil
}

func (r *ProductRepoJson) GetProductByID(id int) (*models.Product, error) {
	fmt.Printf("[Repository][GetProductByID] Getting product by ID: %d\n", id)
	dir := r.storagePath

	// Create the file name based on product ID
	fileName := fmt.Sprintf("%s/%d.json", dir, id)

	// Read the file
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("[Repository][GetProductByID][ERROR] Product with ID %d does not exist\n", id)
			return nil, fmt.Errorf("product with ID %d does not exist", id)
		}
		fmt.Printf("[Repository][GetProductByID][ERROR] Failed to read file: %s, error: %v\n", fileName, err)
		return nil, fmt.Errorf("failed to read file %s: %w", fileName, err)
	}

	// Unmarshal the JSON data into a Product struct
	var product models.Product
	if err := json.Unmarshal(bytes, &product); err != nil {
		fmt.Printf("[Repository][GetProductByID][ERROR] Failed to unmarshal file: %s, error: %v\n", fileName, err)
		return nil, fmt.Errorf("failed to unmarshal file %s: %w", fileName, err)
	}

	fmt.Printf("[Repository][GetProductByID] Successfully loaded product ID: %d\n", id)
	return &product, nil
}

func (r *ProductRepoJson) DeleteByID(id int) error {
	fmt.Printf("[Repository][DeleteByID] Deleting product by ID: %d\n", id)
	dir := r.storagePath

	// Create the file name based on product ID
	fileName := fmt.Sprintf("%s/%d.json", dir, id)

	// Delete the file
	if err := os.Remove(fileName); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("[Repository][DeleteByID][ERROR] Product with ID %d does not exist\n", id)
			return fmt.Errorf("product with ID %d does not exist", id)
		}
		fmt.Printf("[Repository][DeleteByID][ERROR] Failed to delete file: %s, error: %v\n", fileName, err)
		return fmt.Errorf("failed to delete file %d: %w", id, err)
	}

	fmt.Printf("[Repository][DeleteByID] Successfully deleted product ID: %d\n", id)
	return nil
}
