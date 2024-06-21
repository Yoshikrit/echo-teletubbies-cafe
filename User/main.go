package main

import (
	"log"
	"os"
	"github.com/labstack/echo/v4"
	"net/http"

	"user/configs"
	"user/models"
	"user/routes"
	"user/middlewares"
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
    err := db.AutoMigrate(&models.UserEntity{}, &models.RoleEntity{})
    if err != nil {
        panic(err)
    }
    
    // Set up Routes
	userGroup := e.Group("/user")
	routes.SetupUserRoutes(userGroup)

	roleGroup := e.Group("/role")
	routes.SetupRoleRoutes(roleGroup)

    //health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "Service User : OK"})
	})

    // Start the server
    serverPort := os.Getenv("SERVER_PORT")
    log.Printf(serverPort)

    e.Logger.Fatal(e.Start(":" + serverPort))
}