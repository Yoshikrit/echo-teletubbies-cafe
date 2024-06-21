package handlers

import (
    "net/http"
	"github.com/labstack/echo/v4"
	
	"user/models"
	"user/services"
	"user/utils/logs"
	"user/utils/errs"
)

type roleHandler struct {
	roleSrv services.RoleService
}

func NewRoleHandler(roleSrv services.RoleService) roleHandler {
	return roleHandler{roleSrv: roleSrv}
}

func (h roleHandler) CreateRole(c echo.Context) error {
	roleReq := new(models.RoleCreate)
	if err := c.Bind(roleReq); err != nil {
		logs.Error(err.Error())
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}

	c.Echo().Validator = &RoleValidator{validator: v}
	if err := c.Validate(roleReq); err != nil {
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}

	roleRes, err := h.roleSrv.CreateRole(roleReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Create Role Successfully")
	return c.JSON(http.StatusCreated, roleRes)
}

func (h roleHandler) GetAllRoles(c echo.Context) error {
	rolesRes, err := h.roleSrv.GetRoles()
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get Roles Successfully")
	return c.JSON(http.StatusOK, rolesRes)
}

func (h roleHandler) GetRoleByID(c echo.Context) error {
	id, err := GetIntId(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	roleRes, err := h.roleSrv.GetRole(id)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get Role Successfully")
	return c.JSON(http.StatusOK, roleRes)
}

func (h roleHandler) UpdateRoleByID(c echo.Context) error {
	id, err := GetIntId(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	var roleReq models.RoleUpdate
	if err := c.Bind(&roleReq); err != nil {
		logs.Error(err.Error())
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}

	c.Echo().Validator = &RoleValidator{validator: v}
	if err := c.Validate(roleReq); err != nil {
		return HandleError(c, errs.NewBadRequestError(err.Error()))
	}

	roleRes, err := h.roleSrv.UpdateRole(id, &roleReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Update Role Successfully")
	return c.JSON(http.StatusOK, roleRes)
}

func (h roleHandler) DeleteRoleByID(c echo.Context) error {
	id, err := GetIntId(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	if err := h.roleSrv.DeleteRole(id); err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Delete Role Successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
        "message": "Deleted Successfully",
    })
}

func (h roleHandler) GetRoleCount(c echo.Context) error {
    count, err := h.roleSrv.GetRoleCount()
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get Role's Count Successfully")
	return c.JSON(http.StatusOK, count)
}
