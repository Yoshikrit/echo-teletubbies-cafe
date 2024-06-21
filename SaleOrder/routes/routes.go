package routes

import (
	"saleorder/configs"
	"saleorder/handlers"
	"saleorder/repositories"
	"saleorder/services"

	"github.com/labstack/echo/v4"
)

func SetupSaleOrderRoutes(g *echo.Group) {
	saleorderRepository := repositories.NewSaleOrderRepositoryDB(configs.GetDB())
	saleorderService := services.NewSaleOrderService(saleorderRepository)
	saleorderHandler := handlers.NewSaleOrderHandler(saleorderService)

	g.POST("/", saleorderHandler.CreateSaleOrder)
	g.GET("/", saleorderHandler.GetAllSaleOrders)
	g.GET("/day/:date", saleorderHandler.GetAllSaleOrdersByDay)
	g.GET("/month/:date", saleorderHandler.GetAllSaleOrdersByMonth)
	g.GET("/year/:date", saleorderHandler.GetAllSaleOrdersByYear)
	g.GET("/totalprice", saleorderHandler.GetTotalPricePass)
	g.GET("/totalprice/day/:date", saleorderHandler.GetTotalPricePassByDay)
	g.GET("/totalprice/month/:date", saleorderHandler.GetTotalPricePassByMonth)
	g.GET("/totalprice/year/:date", saleorderHandler.GetTotalPricePassByYear)
}

func SetupSaleOrderDetailRoutes(g *echo.Group) {
	saleorderDetailRepository := repositories.NewSaleOrderDetailRepositoryDB(configs.GetDB())
	saleorderDetailService := services.NewSaleOrderDetailService(saleorderDetailRepository)
	saleorderDetailHandler := handlers.NewSaleOrderDetailHandler(saleorderDetailService)

	g.POST("/", saleorderDetailHandler.CreateSaleOrderDetail)
	g.GET("/qtyrate", saleorderDetailHandler.GetSaleOrderDetailQtyRates)
	g.GET("/qtyrate/day/:date", saleorderDetailHandler.GetSaleOrderDetailQtyRatesByDay)
	g.GET("/qtyrate/month/:date", saleorderDetailHandler.GetSaleOrderDetailQtyRatesByMonth)
	g.GET("/qtyrate/year/:date", saleorderDetailHandler.GetSaleOrderDetailQtyRatesByYear)
	g.GET("/pricerate", saleorderDetailHandler.GetSaleOrderDetailPriceRates)
	g.GET("/pricerate/day/:date", saleorderDetailHandler.GetSaleOrderDetailPriceRatesByDay)
	g.GET("/pricerate/month/:date", saleorderDetailHandler.GetSaleOrderDetailPriceRatesByMonth)
	g.GET("/pricerate/year/:date", saleorderDetailHandler.GetSaleOrderDetailPriceRatesByYear)
}
