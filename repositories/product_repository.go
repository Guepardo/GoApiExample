package repositories

import (
	"exemple.com./api/models"
)

type ProductRepository struct {
}

func (repo *ProductRepository) Insert(product models.Product) models.Product {

}

func (repo *ProductRepository) List() error {

}

func (repo *ProductRepository) Delete(token string) error {

}

func (repo *ProductRepository) Update(token string, params map[string]string) error {

}
