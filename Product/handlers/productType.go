package handlers

import (
    "net/http"
	"product/models"
	"product/services"
	"product/utils/logs"
	"product/utils/errs"
	"github.com/labstack/echo/v4"
)

type productTypeHandler struct {
	productTypeSrv services.ProductTypeService
}

func NewProductTypeHandler(productTypeSrv services.ProductTypeService) productTypeHandler {
	return productTypeHandler{productTypeSrv: productTypeSrv}
}

func (h productTypeHandler) CreateProductType(c echo.Context) error {
	prodTypeReq := new(models.ProductTypeCreate)
	if err := c.Bind(prodTypeReq); err != nil {
		logs.Error(err.Error())
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}

	c.Echo().Validator = &ProductTypeValidator{validator: v}
	if err := c.Validate(prodTypeReq); err != nil {
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}

	prodTypeRes, err := h.productTypeSrv.CreateProductType(prodTypeReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Create ProductType Successfully")
	return c.JSON(http.StatusCreated, prodTypeRes)
}

func (h productTypeHandler) GetAllProductTypes(c echo.Context) error {
	prodTypesRes, err := h.productTypeSrv.GetProductTypes()
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get ProductTypes Successfully")
	return c.JSON(http.StatusOK, prodTypesRes)
}

func (h productTypeHandler) GetProductTypeByID(c echo.Context) error {
	id, err := GetIntId(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	prodTypeRes, err := h.productTypeSrv.GetProductType(id)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get ProductType Successfully")
	return c.JSON(http.StatusOK, prodTypeRes)
}

func (h productTypeHandler) UpdateProductTypeByID(c echo.Context) error {
	id, err := GetIntId(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	var prodTypeReq models.ProductTypeUpdate
	if err := c.Bind(&prodTypeReq); err != nil {
		logs.Error(err.Error())
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}

	c.Echo().Validator = &ProductValidator{validator: v}
	if err := c.Validate(prodTypeReq); err != nil {
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}

	prodTypeRes, err := h.productTypeSrv.UpdateProductType(id, &prodTypeReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Update ProductType Successfully")
	return c.JSON(http.StatusOK, prodTypeRes)
}

func (h productTypeHandler) DeleteProductTypeByID(c echo.Context) error {
	id, err := GetIntId(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	if err := h.productTypeSrv.DeleteProductType(id); err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Delete ProductType Successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "Deleted Successfully",
    })
}

func (h productTypeHandler) GetProductTypeCount(c echo.Context) error {
    count, err := h.productTypeSrv.GetProductTypeCount()
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get ProductType's Count Successfully")
	return c.JSON(http.StatusOK, count)
}
