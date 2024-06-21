package handlers

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"github.com/go-playground/validator"

	"auth/utils/errs"
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

func NewAuthValidator() *AuthValidator {
	return &AuthValidator{validator: validator.New()}
}

type AuthValidator struct {
	validator *validator.Validate
}

func (p *AuthValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}