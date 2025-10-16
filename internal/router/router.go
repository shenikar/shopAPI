package router

import (
	"github.com/shenikar/shopAPI/internal/handler/client"
	"github.com/shenikar/shopAPI/internal/handler/image"
	"github.com/shenikar/shopAPI/internal/handler/product"
	"github.com/shenikar/shopAPI/internal/handler/supplier"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handlers struct {
	ClientHandler   *client.Handler
	ProductHandler  *product.Handler
	SupplierHandler *supplier.Handler
	ImageHandler    *image.Handler
}

func SetupRouter(h *Handlers) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")

	clientGroop := api.Group("/clients")
	{
		clientGroop.POST("", h.ClientHandler.CreateClient)
		clientGroop.GET("", h.ClientHandler.GetAllClients)
		clientGroop.GET("/search", h.ClientHandler.GetClientByNameSurname)
		clientGroop.PATCH("/:id/address", h.ClientHandler.UpdateAddress)
		clientGroop.DELETE("/:id", h.ClientHandler.DeleteClient)
	}

	productGroop := api.Group("/products")
	{
		productGroop.POST("", h.ProductHandler.CreateProduct)
		productGroop.PATCH("/:id/decrease-stock", h.ProductHandler.DecreaseStock)
		productGroop.GET("/available", h.ProductHandler.GetAllAvailableProducts)
		productGroop.GET("/:id", h.ProductHandler.GetProductByID)
		productGroop.DELETE("/:id", h.ProductHandler.DeleteProduct)
	}

	supplierGroop := api.Group("/suppliers")
	{
		supplierGroop.POST("", h.SupplierHandler.CreateSupplier)
		supplierGroop.PATCH("/:id/address", h.SupplierHandler.UpdateAddressSupplier)
		supplierGroop.DELETE("/:id", h.SupplierHandler.DeleteSupplier)
		supplierGroop.GET("", h.SupplierHandler.GetAllSuppliers)
		supplierGroop.GET("/:id", h.SupplierHandler.GetSupplierByID)
	}

	imageGroop := api.Group("/images")
	{
		imageGroop.POST("", h.ImageHandler.AddImage)
		imageGroop.PATCH("/:id", h.ImageHandler.UpdateImage)
		imageGroop.DELETE("/:id", h.ImageHandler.DeleteImage)
		imageGroop.GET("/product/:product_id", h.ImageHandler.GetImageByProductID)
		imageGroop.GET("/:id", h.ImageHandler.GetImageByID)
	}

	return r
}
