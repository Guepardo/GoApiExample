package repositories

import (
	"exemple.com/api/models"
)

type ProductRepository struct {
}

func (repo *ProductRepository) Insert(product models.Product) models.Product {
	return models.Product{}
}

func (repo *ProductRepository) List() ([]models.Product, error) {
	return []models.Product{
		models.Product{},
	}, nil
}

func (repo *ProductRepository) Delete(token string) error {
	return nil
}

func (repo *ProductRepository) Update(token string, params map[string]string) error {
	return nil
}
