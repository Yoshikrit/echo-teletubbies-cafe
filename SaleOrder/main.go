package main

import (
	"log"
    "os"
	"github.com/labstack/echo/v4"
    "net/http"

	"saleorder/configs"
	"saleorder/routes"
	"saleorder/middlewares"
	"saleorder/models"
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
    err := db.AutoMigrate(&models.SaleOrderEntity{},&models.PaymentMethodEntity{})
    if err != nil {
        panic(err)
    }
    
    // Set up Routes
	saleorderGroup := e.Group("/saleorder")
	routes.SetupSaleOrderRoutes(saleorderGroup)

    saleorderDetailGroup := e.Group("/saleorderdetail")
    routes.SetupSaleOrderDetailRoutes(saleorderDetailGroup)

    //health check
    e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "Service SaleOrder : OK"})
	})

    // Start the server
    serverPort := os.Getenv("SERVER_PORT")
    log.Printf(serverPort)

    e.Logger.Fatal(e.Start(":" + serverPort))
}