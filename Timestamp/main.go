package main

import (
	"log"
    "os"
	"github.com/labstack/echo/v4"
    "net/http"

	"timestamp/configs"
	"timestamp/routes"
	"timestamp/middlewares"
	"timestamp/models"
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
    err := db.AutoMigrate(&models.TimestampEntity{})
    if err != nil {
        panic(err)
    }
    
    // Set up Routes
	timestampGroup := e.Group("/timestamp")
	routes.SetupTimestampRoutes(timestampGroup)

    //health check
    e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "Service Timestamp : OK"})
	})

    // Start the server
    serverPort := os.Getenv("SERVER_PORT")
    log.Printf(serverPort)

    e.Logger.Fatal(e.Start(":" + serverPort))
}