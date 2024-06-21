package handlers

import (
	"saleorder/utils/errs"

	"time"
	"net/http"
	"github.com/labstack/echo/v4"
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


func GetParamDay(c echo.Context) (time.Time, error) {
	dateStr := c.Param("date")
	dateReq, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, errs.NewBadRequestError("Invalid date format")
	}
	return dateReq, nil
}

func GetParamMonth(c echo.Context) (time.Time, error) {
	dateStr := c.Param("date")
	dateReq, err := time.Parse("2006-01", dateStr)
	if err != nil {
		return time.Time{}, errs.NewBadRequestError("Invalid date format")
	}
	return dateReq, nil
}

func GetParamYear(c echo.Context) (time.Time, error) {
	dateStr := c.Param("date")
	dateReq, err := time.Parse("2006", dateStr)
	if err != nil {
		return time.Time{}, errs.NewBadRequestError("Invalid date format")
	}
	return dateReq, nil
}

var (
	v = validator.New()
)

func NewSaleOrderValidator() *SaleOrderValidator {
	return &SaleOrderValidator{validator: validator.New()}
}

type SaleOrderValidator struct {
	validator *validator.Validate
}

func (p *SaleOrderValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

func NewSaleOrderDetailValidator() *SaleOrderDetailValidator {
	return &SaleOrderDetailValidator{validator: validator.New()}
}

type SaleOrderDetailValidator struct {
	validator *validator.Validate
}

func (p *SaleOrderDetailValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}