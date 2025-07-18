package router

import (
	"shopApi/internal/handler/client"
	"shopApi/internal/handler/image"
	"shopApi/internal/handler/product"
	"shopApi/internal/handler/supplier"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	ClientHandler   *client.ClientHandler
	ProductHandler  *product.ProductHandler
	SupplierHandler *supplier.SupplierHandler
	ImageHandler    *image.ImageHandler
}

func SetupRouter(h *Handlers) *gin.Engine {
	r := gin.Default()

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
		productGroop.PATCH("/:id/decrease", h.ProductHandler.DecreaseStock)
		productGroop.GET("", h.ProductHandler.GetAllAvailableProducts)
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
