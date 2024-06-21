package routes

import (
	"auth/repositories"
	"auth/services"
	"auth/handlers"
	"auth/configs"
	
	"github.com/labstack/echo/v4"
)

func SetupAuthRoutes(g *echo.Group) {
	authRepository := repositories.NewAuthRepositoryDB(configs.GetDB())
	authService := services.NewAuthService(authRepository)
	authHandler := handlers.NewAuthHandler(authService)

	g.POST("/login", authHandler.Login)
	g.GET("/logout/:id", authHandler.Logout)
}
