package repositories

import (
	"exemple.com/api/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (repo *ProductRepository) Insert(product models.Product) (models.Product, error) {
	repo.DB.Create(&product)
	return product, nil
}

func (repo *ProductRepository) List() ([]models.Product, error) {
	var products []models.Product
	repo.DB.Find(&products)
	return products, nil
}

func (repo *ProductRepository) Delete(id int) error {
	repo.DB.Delete(&models.Product{}, 1)
	return nil
}

func (repo *ProductRepository) Update(id int, data models.Product) error {
	return nil
}

func (repo *ProductRepository) Find(id int) models.Product {
	return models.Product{}
}
