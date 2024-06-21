package routes

import (
	"product/repositories"
	"product/services"
	"product/handlers"
	"product/configs"
	"github.com/labstack/echo/v4"
)

func SetupProductRoutes(g *echo.Group) {
	productRepository := repositories.NewProductRepositoryDB(configs.GetDB())
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	g.GET("/", productHandler.GetAllProducts)
	g.GET("/:id", productHandler.GetProductByID)
	g.POST("/", productHandler.CreateProduct)
	g.PUT("/:id", productHandler.UpdateProductByID)
	g.DELETE("/:id", productHandler.DeleteProductByID)
	g.GET("/count", productHandler.GetProductCount)
}

func SetupProductTypeRoutes(g *echo.Group) {
	productTypeRepository := repositories.NewProductTypeRepositoryDB(configs.GetDB())
	productTypeService := services.NewProductTypeService(productTypeRepository)
	productTypeHandler := handlers.NewProductTypeHandler(productTypeService)

	g.GET("/", productTypeHandler.GetAllProductTypes)
	g.GET("/:id", productTypeHandler.GetProductTypeByID)
	g.POST("/", productTypeHandler.CreateProductType)
	g.PUT("/:id", productTypeHandler.UpdateProductTypeByID)
	g.DELETE("/:id", productTypeHandler.DeleteProductTypeByID)
	g.GET("/count", productTypeHandler.GetProductTypeCount)
}