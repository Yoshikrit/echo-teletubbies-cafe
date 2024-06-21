package handlers

import (
    "net/http"
	"auth/models"
	"auth/services"
	"auth/utils/logs"
	"auth/utils/errs"
	"github.com/labstack/echo/v4"
)

type authHandler struct {
	authSrv services.AuthService
}

func NewAuthHandler(authSrv services.AuthService) authHandler {
	return authHandler{authSrv: authSrv}
}

func (h authHandler) Login(c echo.Context) error {
	userReq := new(models.UserLogin)
	if err := c.Bind(userReq); err != nil {
		logs.Error(err.Error())
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}

	c.Echo().Validator = &AuthValidator{validator: v}
	if err := c.Validate(userReq); err != nil {
		logs.Error(err.Error())
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}

	response, err := h.authSrv.Login(userReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Log in Successfully")
	return c.JSON(http.StatusOK, response)
}

func (h authHandler) Logout(c echo.Context) error {
	id, err := GetIntId(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	if err := h.authSrv.Logout(id); err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Log out Successfully")
	return c.JSON(http.StatusOK, "Log out successfully")
}