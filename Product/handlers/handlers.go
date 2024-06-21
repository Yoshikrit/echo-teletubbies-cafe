package handlers

import (
	"product/utils/errs"
	"net/http"
	"github.com/labstack/echo/v4"
	"strconv"
	"github.com/go-playground/validator"
)

func HandleError(c echo.Context, err error) error {
	if e, ok := err.(errs.AppError); ok {
		return c.JSON(e.Code, map[string]interface{}{
			"code":    e.Code,
			"message": e.Message,
		})
	}

	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "Interval Server Error",
	})
}

func GetIntId(c echo.Context) (int, error) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, errs.NewBadRequestError("Invalid id: " + idParam + " is not integer")
	}
	return int(id), nil
}

var (
	v = validator.New()
)

func NewProductValidator() *ProductValidator {
	return &ProductValidator{validator: validator.New()}
}

//ProductValidator a product validator
type ProductValidator struct {
	validator *validator.Validate
}

//Validate validates a product
func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

func NewProductTypeValidator() *ProductTypeValidator {
	return &ProductTypeValidator{validator: validator.New()}
}

type ProductTypeValidator struct {
	validator *validator.Validate
}

func (p *ProductTypeValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}