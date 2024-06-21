package handlers

import (
	"saleorder/models"
	"saleorder/services"
	"saleorder/utils/logs"
	"saleorder/utils/errs"

	"github.com/labstack/echo/v4"
    "net/http"
)

type saleorderHandler struct {
	saleorderSrv services.SaleOrderService
}

func NewSaleOrderHandler(saleorderSrv services.SaleOrderService) saleorderHandler {
	return saleorderHandler{saleorderSrv: saleorderSrv}
}

func (h saleorderHandler) CreateSaleOrder(c echo.Context) error {
	saleOrderReq := new(models.SaleOrderCreate)
	if err := c.Bind(saleOrderReq); err != nil {
		logs.Error(err.Error())
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}

	c.Echo().Validator = &SaleOrderValidator{validator: v}
	if err := c.Validate(saleOrderReq); err != nil {
		logs.Error(err.Error())
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}

	saleOrderRes, err := h.saleorderSrv.CreateSaleOrder(saleOrderReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Create SaleOrder Successfully")
	return c.JSON(http.StatusCreated, saleOrderRes)
}


func (h saleorderHandler) GetAllSaleOrders(c echo.Context) error {
	saleordersRes, err := h.saleorderSrv.GetSaleOrders()
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get SaleOrders Successfully")
	return c.JSON(http.StatusOK, saleordersRes)
}

func (h saleorderHandler) GetAllSaleOrdersByDay(c echo.Context) error {
	dateReq, err := GetParamDay(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	saleordersRes, err := h.saleorderSrv.GetSaleOrdersByDay(dateReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get SaleOrders By Day Successfully")
	return c.JSON(http.StatusOK, saleordersRes)
}

func (h saleorderHandler) GetAllSaleOrdersByMonth(c echo.Context) error {
	dateReq, err := GetParamMonth(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	saleordersRes, err := h.saleorderSrv.GetSaleOrdersByMonth(dateReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get SaleOrders By Month Successfully")
	return c.JSON(http.StatusOK, saleordersRes)
}

func (h saleorderHandler) GetAllSaleOrdersByYear(c echo.Context) error {
	dateReq, err := GetParamYear(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	saleordersRes, err := h.saleorderSrv.GetSaleOrdersByYear(dateReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get SaleOrders By Year Successfully")
	return c.JSON(http.StatusOK, saleordersRes)
}

func (h saleorderHandler) GetTotalPricePass(c echo.Context) error {
	totalAmount, err := h.saleorderSrv.GetSaleOrderPriceAmountPass()
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get Total Amount of Price That Pass Successfully")
	return c.JSON(http.StatusOK, totalAmount)
}

func (h saleorderHandler) GetTotalPricePassByDay(c echo.Context) error {
	dateReq, err := GetParamDay(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	totalAmount, err := h.saleorderSrv.GetSaleOrderPriceAmountPassByDay(dateReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get Total Amount of Price That Pass Successfully")
	return c.JSON(http.StatusOK, totalAmount)
}

func (h saleorderHandler) GetTotalPricePassByMonth(c echo.Context) error {
	dateReq, err := GetParamMonth(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	totalAmount, err := h.saleorderSrv.GetSaleOrderPriceAmountPassByMonth(dateReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get Total Amount of Price That Pass Successfully")
	return c.JSON(http.StatusOK, totalAmount)
}

func (h saleorderHandler) GetTotalPricePassByYear(c echo.Context) error {
	dateReq, err := GetParamYear(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}
	
	totalAmount, err := h.saleorderSrv.GetSaleOrderPriceAmountPassByYear(dateReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get Total Amount of Price That Pass Successfully")
	return c.JSON(http.StatusOK, totalAmount)
}