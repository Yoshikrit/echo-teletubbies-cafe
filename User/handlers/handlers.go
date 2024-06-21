package handlers

import (
	"user/utils/errs"
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

func NewUserValidator() *UserValidator {
	return &UserValidator{validator: validator.New()}
}

type UserValidator struct {
	validator *validator.Validate
}

func (p *UserValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

func NewRoleValidator() *RoleValidator {
	return &RoleValidator{validator: validator.New()}
}

type RoleValidator struct {
	validator *validator.Validate
}

func (p *RoleValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}