package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/marcosduarte-dev/go-api/configs"
	"github.com/marcosduarte-dev/go-api/internal/entity"
	"github.com/marcosduarte-dev/go-api/internal/infra/database"
	"github.com/marcosduarte-dev/go-api/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	ProductHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r  := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", ProductHandler.CreateProduct)
	r.Get("/products/{id}", ProductHandler.GetProduct)
	r.Get("/products", ProductHandler.GetProducts)
	r.Put("/products/{id}", ProductHandler.UpdateProduct)
	r.Delete("/products/{id}", ProductHandler.DeleteProduct)

	r.Post("/users", userHandler.Create)
	
	http.ListenAndServe(":8000", r)
}