package routes

import (
	"timestamp/repositories"
	"timestamp/services"
	"timestamp/handlers"
	"timestamp/configs"
	
	"github.com/labstack/echo/v4"
)

func SetupTimestampRoutes(g *echo.Group) {
	timestampRepository := repositories.NewTimestampRepositoryDB(configs.GetDB())
	timestampService := services.NewTimestampService(timestampRepository)
	timestampHandler := handlers.NewTimestampHandler(timestampService)

	g.GET("/", timestampHandler.GetAllTimestamps)
	g.GET("/day/:date", timestampHandler.GetAllTimestampsByDay)
	g.GET("/month/:date", timestampHandler.GetAllTimestampsByMonth)
	g.GET("/year/:date", timestampHandler.GetAllTimestampsByYear)
}
