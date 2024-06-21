package handlers

import (
	"saleorder/models"
	"saleorder/services"
	"saleorder/utils/logs"
	"saleorder/utils/errs"

	"github.com/labstack/echo/v4"
    "net/http"
)

type saleorderDetailHandler struct {
	saleorderDetailSrv services.SaleOrderDetailService
}

func NewSaleOrderDetailHandler(saleorderDetailSrv services.SaleOrderDetailService) saleorderDetailHandler {
	return saleorderDetailHandler{saleorderDetailSrv: saleorderDetailSrv}
}

func (h saleorderDetailHandler) CreateSaleOrderDetail(c echo.Context) error {
	saleOrderDetailReq := new(models.SaleOrderDetailCreate)
	if err := c.Bind(saleOrderDetailReq); err != nil {
		logs.Error(err.Error())
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}

	c.Echo().Validator = &SaleOrderValidator{validator: v}
	if err := c.Validate(saleOrderDetailReq); err != nil {
		logs.Error(err.Error())
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}

	saleOrderDetailRes, err := h.saleorderDetailSrv.CreateSaleOrderDetail(saleOrderDetailReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Create SaleOrderDetail Successfully")
	return c.JSON(http.StatusCreated, saleOrderDetailRes)
}

func (h saleorderDetailHandler) GetSaleOrderDetailQtyRates(c echo.Context) error {
	saleorderDetailsRes, err := h.saleorderDetailSrv.GetSaleOrderDetailQtyRates()
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get SaleOrderDetail Qty Rate Successfully")
	return c.JSON(http.StatusOK, saleorderDetailsRes)
}

func (h saleorderDetailHandler) GetSaleOrderDetailQtyRatesByDay(c echo.Context) error {
	dateReq, err := GetParamDay(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	saleorderDetailsRes, err := h.saleorderDetailSrv.GetSaleOrderDetailQtyRatesByDay(dateReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get SaleOrderDetail Qty Rate By Day Successfully")
	return c.JSON(http.StatusOK, saleorderDetailsRes)
}

func (h saleorderDetailHandler) GetSaleOrderDetailQtyRatesByMonth(c echo.Context) error {
	dateReq, err := GetParamMonth(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	saleorderDetailsRes, err := h.saleorderDetailSrv.GetSaleOrderDetailQtyRatesByMonth(dateReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get SaleOrderDetail Qty Rate By Month Successfully")
	return c.JSON(http.StatusOK, saleorderDetailsRes)
}

func (h saleorderDetailHandler) GetSaleOrderDetailQtyRatesByYear(c echo.Context) error {
	dateReq, err := GetParamYear(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	saleorderDetailsRes, err := h.saleorderDetailSrv.GetSaleOrderDetailQtyRatesByYear(dateReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get SaleOrderDetail Qty Rate By Year Successfully")
	return c.JSON(http.StatusOK, saleorderDetailsRes)
}

func (h saleorderDetailHandler) GetSaleOrderDetailPriceRates(c echo.Context) error {
	saleorderDetailsRes, err := h.saleorderDetailSrv.GetSaleOrderDetailPriceRates()
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get SaleOrderDetail Price Rate Successfully")
	return c.JSON(http.StatusOK, saleorderDetailsRes)
}

func (h saleorderDetailHandler) GetSaleOrderDetailPriceRatesByDay(c echo.Context) error {
	dateReq, err := GetParamDay(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	saleorderDetailsRes, err := h.saleorderDetailSrv.GetSaleOrderDetailPriceRatesByDay(dateReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get SaleOrderDetail Price Rate By Day Successfully")
	return c.JSON(http.StatusOK, saleorderDetailsRes)
}

func (h saleorderDetailHandler) GetSaleOrderDetailPriceRatesByMonth(c echo.Context) error {
	dateReq, err := GetParamMonth(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	saleorderDetailsRes, err := h.saleorderDetailSrv.GetSaleOrderDetailPriceRatesByMonth(dateReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get SaleOrderDetail Price Rate By Month Successfully")
	return c.JSON(http.StatusOK, saleorderDetailsRes)
}

func (h saleorderDetailHandler) GetSaleOrderDetailPriceRatesByYear(c echo.Context) error {
	dateReq, err := GetParamYear(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	saleorderDetailsRes, err := h.saleorderDetailSrv.GetSaleOrderDetailPriceRatesByYear(dateReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get SaleOrderDetail Price Rate By Year Successfully")
	return c.JSON(http.StatusOK, saleorderDetailsRes)
}

