package main

import (
	"log"
	"net/http"
	"shopApi/internal/config"
	"shopApi/internal/db"
	"shopApi/internal/handler/client"
	"shopApi/internal/handler/image"
	"shopApi/internal/handler/product"
	"shopApi/internal/handler/supplier"
	"shopApi/internal/repository"
	"shopApi/internal/router"

	"github.com/go-playground/validator/v10"
)

func main() {
	cfg := config.LoadConfig()
	connStr := cfg.PGXConnString()

	db, err := db.NewDataBase(connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	log.Println("Database connection established successfully")

	v := validator.New()

	clientRepo := repository.NewClientRepository(db)
	productRepo := repository.NewProductRepository(db)
	supplierRepo := repository.NewSupplierRepository(db)
	imageRepo := repository.NewImageRepository(db)

	clientHandler := client.NewClientHandler(clientRepo, v)
	productHandler := product.NewProductHandler(productRepo, v)
	supplierHandler := supplier.NewSupplierHandler(supplierRepo, v)
	imageHandler := image.NewImageHandler(imageRepo, v)

	h := &router.Handlers{
		ClientHandler:   clientHandler,
		ProductHandler:  productHandler,
		SupplierHandler: supplierHandler,
		ImageHandler:    imageHandler,
	}

	r := router.SetupRouter(h)

	log.Println("Сервер запущен на :8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Сервер завершился с ошибкой: %v", err)
	}
}
