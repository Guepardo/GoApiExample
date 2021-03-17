package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"exemple.com/api/controllers"
	"exemple.com/api/models"
	"exemple.com/api/repositories"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const DATABASE_NAME = "test.db"

type App struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (app *App) Initialize() {
	app.Router = mux.NewRouter()

	app.initializeDatabase()
	app.initializeRouters()
}

func (app *App) Start(port string) {
	log.Printf("Server Started on port %s", port)
	log.Fatal(http.ListenAndServe(port, app.Router))
}

func (app *App) initializeMiddlewares() {
	app.Router.Use()
}

func (app *App) initializeDatabase() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open(DATABASE_NAME), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Fatalln("Error on database initialization")
	}

	db.AutoMigrate(&models.Product{})

	log.Println("Database loaded.")
	app.DB = db
}

func (app *App) initializeRouters() {
	productsController := controllers.ProductsController{
		Repo: repositories.ProductRepository{
			DB: app.DB,
		},
	}

	app.Router.HandleFunc("/products", productsController.Index).Methods("GET")
	app.Router.HandleFunc("/products", productsController.Store).Methods("POST")
	app.Router.HandleFunc("/products/{id}", productsController.Show).Methods("GET")
	app.Router.HandleFunc("/products/{id}", productsController.Update).Methods("PUT")
	app.Router.HandleFunc("/products/{id}", productsController.Delete).Methods("DELETE")
}
