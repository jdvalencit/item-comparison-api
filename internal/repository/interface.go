package repository

import "item-comparison-api/internal/models"

type ProductRepo interface {
	LoadProducts() ([]models.Product, error)
	SaveProducts([]models.Product) error
	//GetByID(int) (*models.Product, error)
	//Create(models.Product) error
	//Update(models.Product) error
	//DeleteByID(int) error
}
