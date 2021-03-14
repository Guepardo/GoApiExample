package controllers

import (
	"log"
	"net/http"

	"exemple.com/api/repositories"
)

type ProductsController struct {
	repo *repositories.ProductRepository
}

func (controller *ProductsController) Store(w http.ResponseWriter, r *http.Request) {
	log.Println("Logging...")
}

func (controller *ProductsController) Show(w http.ResponseWriter, r *http.Request) {
	log.Println("Logging...")
}

func (controller *ProductsController) Delete(w http.ResponseWriter, r *http.Request) {
	log.Println("Logging...")
}

func (controller *ProductsController) Index(w http.ResponseWriter, r *http.Request) {
	log.Println("Logging...")
}

func (controller *ProductsController) Update(w http.ResponseWriter, r *http.Request) {
	log.Println("Logging...")
}
