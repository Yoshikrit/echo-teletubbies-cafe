package handlers

import (
    "net/http"
	"product/models"
	"product/services"
	"product/utils/logs"
	"product/utils/errs"
	"github.com/labstack/echo/v4"
)

type productHandler struct {
	productSrv services.ProductService
}

func NewProductHandler(productSrv services.ProductService) productHandler {
	return productHandler{productSrv: productSrv}
}

func (h productHandler) CreateProduct(c echo.Context) error {
	prodReq := new(models.ProductCreate)
	if err := c.Bind(prodReq); err != nil {
		logs.Error(err.Error())
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}
	
	c.Echo().Validator = &ProductValidator{validator: v}
	if err := c.Validate(prodReq); err != nil {
		logs.Error(err.Error())
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}

	prodRes, err := h.productSrv.CreateProduct(prodReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Create Product Successfully")
	return c.JSON(http.StatusCreated, prodRes)
}

func (h productHandler) GetAllProducts(c echo.Context) error {
	prodsRes, err := h.productSrv.GetProducts()
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get Products Successfully")
	return c.JSON(http.StatusOK, prodsRes)
}

func (h productHandler) GetProductByID(c echo.Context) error {
	id, err := GetIntId(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	prodRes, err := h.productSrv.GetProduct(id)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get Product Successfully")
	return c.JSON(http.StatusOK, prodRes)
}

func (h productHandler) UpdateProductByID(c echo.Context) error {
	id, err := GetIntId(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	var prodReq models.ProductUpdate
	if err := c.Bind(&prodReq); err != nil {
		logs.Error(err.Error())
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}

	prodRes, err := h.productSrv.UpdateProduct(id, &prodReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Update Product Successfully")
	return c.JSON(http.StatusOK, prodRes)
}

func (h productHandler) DeleteProductByID(c echo.Context) error {
	id, err := GetIntId(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	if err := h.productSrv.DeleteProduct(id); err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Delete Product Successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "Deleted Successfully",
    })
}

func (h productHandler) GetProductCount(c echo.Context) error {
    count, err := h.productSrv.GetProductCount()
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get Product's Count Successfully")
	return c.JSON(http.StatusOK, count)
}
