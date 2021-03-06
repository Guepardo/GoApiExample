package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"exemple.com/api/models"
	"exemple.com/api/repositories"
	"github.com/gorilla/mux"
)

type ProductsController struct {
	BaseController
	Repo repositories.ProductRepository
}

func (con *ProductsController) Store(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&product); err != nil {
		con.renderMessage(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	defer r.Body.Close()

	persistedProduct, err := con.Repo.Insert(product)

	if err != nil {
		con.renderMessage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	con.renderJson(w, http.StatusOK, persistedProduct)
}

func (con *ProductsController) Show(w http.ResponseWriter, r *http.Request) {
	con.renderMessage(w, http.StatusInternalServerError, "Internal Server Error")
}

func (con *ProductsController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	if err := con.Repo.Delete(id); err == nil {
		con.renderJson(w, http.StatusNoContent, nil)
		return
	}

	con.renderMessage(w, http.StatusInternalServerError, "Internal Server Error")
}

func (con *ProductsController) Index(w http.ResponseWriter, r *http.Request) {
	products, _ := con.Repo.List()

	con.renderJson(w, http.StatusOK, products)
}

func (con *ProductsController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var product models.Product

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&product); err != nil {
		con.renderMessage(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	defer r.Body.Close()

	id, _ := strconv.Atoi(vars["id"])
	con.Repo.Update(id, product)
}
