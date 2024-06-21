package routes

import (
	"user/repositories"
	"user/services"
	"user/handlers"
	"user/configs"
	
	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(g *echo.Group) {
	userRepository := repositories.NewUserRepositoryDB(configs.GetDB())
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	g.GET("/", userHandler.GetAllUsers)
	g.GET("/:id", userHandler.GetUserByID)
	g.POST("/", userHandler.CreateUser)
	g.PUT("/:id", userHandler.UpdateUserByID)
	g.DELETE("/:id", userHandler.DeleteUserByID)
	g.GET("/count", userHandler.GetUserCount)
}

func SetupRoleRoutes(g *echo.Group) {
	roleRepository := repositories.NewRoleRepositoryDB(configs.GetDB())
	roleService := services.NewRoleService(roleRepository)
	roleHandler := handlers.NewRoleHandler(roleService)

	g.GET("/", roleHandler.GetAllRoles)
	g.GET("/:id", roleHandler.GetRoleByID)
	g.POST("/", roleHandler.CreateRole)
	g.PUT("/:id", roleHandler.UpdateRoleByID)
	g.DELETE("/:id", roleHandler.DeleteRoleByID)
	g.GET("/count", roleHandler.GetRoleCount)
}
