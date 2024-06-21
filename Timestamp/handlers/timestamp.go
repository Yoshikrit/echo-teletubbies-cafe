package handlers

import (
	"timestamp/services"
	"timestamp/utils/logs"

	"github.com/labstack/echo/v4"
    "net/http"
)

type timestampHandler struct {
	timestampSrv services.TimestampService
}

func NewTimestampHandler(timestampSrv services.TimestampService) timestampHandler {
	return timestampHandler{timestampSrv: timestampSrv}
}

func (h timestampHandler) GetAllTimestamps(c echo.Context) error {
	timestampsRes, err := h.timestampSrv.GetTimestamps()
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get Timestamps Successfully")
	return c.JSON(http.StatusOK, timestampsRes)
}

func (h timestampHandler) GetAllTimestampsByDay(c echo.Context) error {
	dateReq, err := GetParamDay(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	timestampsRes, err := h.timestampSrv.GetTimestampsByDay(dateReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get Timestamps By Day Successfully")
	return c.JSON(http.StatusOK, timestampsRes)
}

func (h timestampHandler) GetAllTimestampsByMonth(c echo.Context) error {
	dateReq, err := GetParamMonth(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	timestampsRes, err := h.timestampSrv.GetTimestampsByMonth(dateReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get Timestamps By Month Successfully")
	return c.JSON(http.StatusOK, timestampsRes)
}

func (h timestampHandler) GetAllTimestampsByYear(c echo.Context) error {
	dateReq, err := GetParamYear(c)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	timestampsRes, err := h.timestampSrv.GetTimestampsByYear(dateReq)
	if err != nil {
		logs.Error(err.Error())
		return HandleError(c, err)
	}

	logs.Info("Handler: Get Timestamps By Year Successfully")
	return c.JSON(http.StatusOK, timestampsRes)
}