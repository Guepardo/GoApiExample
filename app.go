package main

import (
	"log"
	"net/http"

	"exemple.com/api/controllers"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (app *App) Initialize() {
	app.Router = mux.NewRouter()
	app.initializeRouters()
}

func (app *App) Start(port string) {
	log.Printf("Server Started on port %s", port)
	log.Fatal(http.ListenAndServe(port, app.Router))
}

func (app *App) initializeRouters() {
	productsController := controllers.ProductsController{}

	app.Router.HandleFunc("/products", productsController.Index).Methods("GET")
	app.Router.HandleFunc("/products", productsController.Store).Methods("POST")
	app.Router.HandleFunc("/products/{id}", productsController.Show).Methods("GET")
	app.Router.HandleFunc("/products/{id}", productsController.Update).Methods("PUT")
	app.Router.HandleFunc("/products/{id}", productsController.Delete).Methods("DELETE")
}
