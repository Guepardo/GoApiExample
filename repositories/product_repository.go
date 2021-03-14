package repositories

import (
	"exemple.com/api/models"
)

type ProductRepository struct {
	data map[string]models.Product
}

func (repo *ProductRepository) initRepo() {
	if repo.data == nil {
		repo.data = make(map[string]models.Product)
	}
}

func (repo *ProductRepository) Insert(product models.Product) (models.Product, error) {
	repo.initRepo()
	repo.data[product.Token] = product

	return product, nil
}

func (repo *ProductRepository) List() ([]models.Product, error) {
	repo.initRepo()
	var values []models.Product

	for _, value := range repo.data {
		values = append(values, value)
	}

	return values, nil
}

func (repo *ProductRepository) Delete(token string) error {
	repo.initRepo()
	delete(repo.data, token)
	return nil
}

func (repo *ProductRepository) Update(token string, data models.Product) error {
	repo.initRepo()
	product := repo.data[token]

	product.Name = data.Name
	product.Price = data.Price

	return nil
}

func (repo *ProductRepository) Find(token string) models.Product {
	return repo.data[token]
}
