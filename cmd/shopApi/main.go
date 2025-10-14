package main

import (
	"log"
	"net/http"
	_ "shopApi/docs"
	"shopApi/internal/config"
	"shopApi/internal/db"
	"shopApi/internal/handler/client"
	"shopApi/internal/handler/image"
	"shopApi/internal/handler/product"
	"shopApi/internal/handler/supplier"
	"shopApi/internal/repository"
	"shopApi/internal/router"
	service "shopApi/internal/service"

	"github.com/go-playground/validator/v10"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// @title Shop API
// @version 1.0
// @description This is a sample server for a shop API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	cfg := config.LoadConfig()
	connStr := cfg.PGXConnString()

	db, err := db.NewDataBase(connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	log.Println("Database connection established successfully")

	runMigrations(connStr, cfg.MigrationsPath)

	v := validator.New()

	addressRepo := repository.NewAddressRepository(db)
	addressService := service.NewAddressService(addressRepo)

	clientRepo := repository.NewClientRepository(db)
	clientService := service.NewClientService(clientRepo, addressService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)

	supplierRepo := repository.NewSupplierRepository(db)
	supplierService := service.NewSupplierService(supplierRepo, addressService)

	imageRepo := repository.NewImageRepository(db)
	imageService := service.NewImageService(imageRepo)

	clientHandler := client.NewClientHandler(clientService, v)
	productHandler := product.NewProductHandler(productService, v)
	supplierHandler := supplier.NewSupplierHandler(supplierService, v)
	imageHandler := image.NewImageHandler(imageService, v)

	h := &router.Handlers{
		ClientHandler:   clientHandler,
		ProductHandler:  productHandler,
		SupplierHandler: supplierHandler,
		ImageHandler:    imageHandler,
	}

	r := router.SetupRouter(h)

	serverAddr := ":" + cfg.ServerPort
	log.Printf("Сервер запущен на %s...", serverAddr)
	if err := http.ListenAndServe(serverAddr, r); err != nil {
		log.Fatalf("Сервер завершился с ошибкой: %v", err)
	}
}

func runMigrations(connStr, path string) {
	if path == "" {
		path = "migrations" // fallback for local run
	}
	m, err := migrate.New(
		"file://"+path,
		connStr)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully")
}
