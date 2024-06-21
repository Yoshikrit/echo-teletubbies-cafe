package handlers

import (
    "net/http"
	"github.com/labstack/echo/v4"
	
	"user/models"
	"user/services"
	"user/utils/logs"
	"user/utils/errs"
)

type userHandler struct {
	userSrv services.UserService
}

func NewUserHandler(userSrv services.UserService) userHandler {
	return userHandler{userSrv: userSrv}
}

func (h userHandler) CreateUser(c echo.Context) error {
	userReq := new(models.UserCreate)
	if err := c.Bind(userReq); err != nil {
		logs.Error(err.Error())
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}

	c.Echo().Validator = &UserValidator{validator: v}
	if err := c.Validate(userReq); err != nil {
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}

	userRes, err := h.userSrv.CreateUser(userReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Create User Successfully")
	return c.JSON(http.StatusCreated, userRes)
}

func (h userHandler) GetAllUsers(c echo.Context) error {
	usersRes, err := h.userSrv.GetUsers()
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get Users Successfully")
	return c.JSON(http.StatusOK, usersRes)
}

func (h userHandler) GetUserByID(c echo.Context) error {
	id, err := GetIntId(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}
	
	userRes, err := h.userSrv.GetUser(id)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get User Successfully")
	return c.JSON(http.StatusOK, userRes)
}

func (h userHandler) UpdateUserByID(c echo.Context) error {
	id, err := GetIntId(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	var userReq models.UserUpdate
	if err := c.Bind(&userReq); err != nil {
		logs.Error(err.Error())
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}

	userRes, err := h.userSrv.UpdateUser(id, &userReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Update User Successfully")
	return c.JSON(http.StatusOK, userRes)
}

func (h userHandler) DeleteUserByID(c echo.Context) error {
	id, err := GetIntId(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	if err := h.userSrv.DeleteUser(id); err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Delete User Successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "Deleted Successfully",
    })
}

func (h userHandler) GetUserCount(c echo.Context) error {
    count, err := h.userSrv.GetUserCount()
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get User's Count Successfully")
	return c.JSON(http.StatusOK, count)
}
