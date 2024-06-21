package main

import (
	"log"
    "os"
	"github.com/labstack/echo/v4"
    "net/http"

	"product/configs"
	"product/routes"
	"product/middlewares"
	"product/models"
)

func main() {
    configs.LoadEnv()
    
	e := echo.New()

    //middleware
    middlewares.SetMiddleware(e)
    
	configs.DatabaseInit()
    defer configs.GetDB().DB()

    // Perform migrations using AutoMigrate
    db := configs.GetDB()
    err := db.AutoMigrate(&models.ProductTypeEntity{}, &models.ProductEntity{})
    if err != nil {
        panic(err)
    }
    
    // Set up Routes
	productGroup := e.Group("/product")
	routes.SetupProductRoutes(productGroup)

	productTypeGroup := e.Group("/producttype")
	routes.SetupProductTypeRoutes(productTypeGroup)

    //health check
    e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "Service Product : OK"})
	})

    // Start the server
    serverPort := os.Getenv("SERVER_PORT")
    log.Printf(serverPort)

    e.Logger.Fatal(e.Start(":" + serverPort))
}