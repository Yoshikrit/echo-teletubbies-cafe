package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"saleorder/handlers"
	"saleorder/models"
	"saleorder/tests/mocks/mock_services"
	"saleorder/utils/errs"
)

func TestCreateSaleOrder(t *testing.T) {

	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)

	saleOrderReq := &models.SaleOrderCreate{
		CreatedUser: 1,
		TotalPrice:  200,
		Status:      "Pass",
		PayMethod:   1,
	}

	saleOrderRes := &models.SaleOrder{
		Id: 1,
		CreatedUser: 1,
		CreatedDate: mockDate,
		TotalPrice: 200,
		Status: "Pass",
		PayMethod: 1,
	}

	e := echo.New()

	saleOrderReqJSON, _ := json.Marshal(saleOrderReq)
	saleOrderResJSON, _ := json.Marshal(saleOrderRes)

	t.Run("test case : pass", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
		saleOrderService.On("CreateSaleOrder", saleOrderReq).Return(saleOrderRes, nil)
			
		req := httptest.NewRequest(http.MethodPost, "/saleorder/", strings.NewReader(string(saleOrderReqJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.CreateSaleOrder(c)

		assert.Equal(t, 201, rec.Code)
		assert.Equal(t, strings.TrimSpace(string(saleOrderResJSON)), strings.TrimSpace(rec.Body.String()))
		saleOrderService.AssertExpectations(t)
	})

	t.Run("test case : fail bind", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()

		req := httptest.NewRequest(http.MethodPost, "/saleorder/", strings.NewReader("invalid json"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.CreateSaleOrder(c)

		assert.Equal(t, 400, rec.Code)
	})

	saleOrderErrReq := &models.SaleOrderCreate{
        CreatedUser:  0,
        TotalPrice:   0,
        Status:       "",
        PayMethod:    0,
    }
	saleOrderErrJSON, _ := json.Marshal(saleOrderErrReq)

	t.Run("test case : fail validator", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
		
		req := httptest.NewRequest(http.MethodPost, "/saleorder/", strings.NewReader(string(saleOrderErrJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.CreateSaleOrder(c)

		assert.Equal(t, 400, rec.Code)
	})

	t.Run("test case : fail repository", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
		saleOrderService.On("CreateSaleOrder", saleOrderReq).Return(&models.SaleOrder{}, errs.NewUnexpectedError(""))

		req := httptest.NewRequest(http.MethodPost, "/saleorder/", strings.NewReader(string(saleOrderResJSON)))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.CreateSaleOrder(c)

		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, strings.TrimSpace(`{"code":500,"message":""}`), strings.TrimSpace(rec.Body.String()))
		saleOrderService.AssertExpectations(t)
	})
}

func TestGetAllSaleOrders(t *testing.T) {
	e := echo.New()
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	saleOrderReportRes := []models.SaleOrderReport{
		{
			Seq:  			1,
			Id: 			1,
			User: 			"Walter White",
			Date:   		mockDate,
			TotalPrice:   	0,
			Status:        	"Fail",
			PayMethodName: 	"None",
		},
    }
	saleOrderReportResJSON, _ := json.Marshal(saleOrderReportRes)

	t.Run("test case : pass", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
		saleOrderService.On("GetSaleOrders").Return(saleOrderReportRes, nil)
			
		req := httptest.NewRequest(http.MethodGet, "/saleorder/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetAllSaleOrders(c)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, strings.TrimSpace(string(saleOrderReportResJSON)), strings.TrimSpace(rec.Body.String()))
		saleOrderService.AssertExpectations(t)
	})

	t.Run("test case : fail repository", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
		saleOrderService.On("GetSaleOrders").Return([]models.SaleOrderReport{}, errs.NewUnexpectedError(""))
			
		req := httptest.NewRequest(http.MethodGet, "/saleorder/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetAllSaleOrders(c)

		assert.Equal(t, 500, rec.Code)
		saleOrderService.AssertExpectations(t)
	})
}


func TestGetAllSaleOrdersByDay(t *testing.T) {
	e := echo.New()
	mockDate := time.Date(2023, time.December, 28, 0, 0, 0, 0, time.UTC)
	saleOrderReportRes := []models.SaleOrderReport{
		{
			Seq:  			1,
			Id: 			1,
			User: 			"Walter White",
			Date:   		mockDate,
			TotalPrice:   	0,
			Status:        	"Fail",
			PayMethodName: 	"None",
		},
    }
	saleOrderReportResJSON, _ := json.Marshal(saleOrderReportRes)

	t.Run("test case : pass", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
		saleOrderService.On("GetSaleOrdersByDay", mockDate).Return(saleOrderReportRes, nil)
			
		req := httptest.NewRequest(http.MethodGet, "/saleorder/day/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2023-12-28")

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetAllSaleOrdersByDay(c)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, strings.TrimSpace(string(saleOrderReportResJSON)), strings.TrimSpace(rec.Body.String()))
		saleOrderService.AssertExpectations(t)
	})

	t.Run("test case : fail validate date", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
			
		req := httptest.NewRequest(http.MethodGet, "/saleorder/day/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2")

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetAllSaleOrdersByDay(c)

		assert.Equal(t, 400, rec.Code)
	})

	t.Run("test case : fail repository", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
		saleOrderService.On("GetSaleOrdersByDay", mockDate).Return([]models.SaleOrderReport{}, errs.NewUnexpectedError(""))
			
		req := httptest.NewRequest(http.MethodGet, "/saleorder/day/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2023-12-28")

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetAllSaleOrdersByDay(c)

		assert.Equal(t, 500, rec.Code)
		saleOrderService.AssertExpectations(t)
	})
}

func TestGetAllSaleOrdersByMonth(t *testing.T) {
	e := echo.New()
	mockDate := time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC)
	saleOrderReportRes := []models.SaleOrderReport{
		{
			Seq:  			1,
			Id: 			1,
			User: 			"Walter White",
			Date:   		mockDate,
			TotalPrice:   	0,
			Status:        	"Fail",
			PayMethodName: 	"None",
		},
    }
	saleOrderReportResJSON, _ := json.Marshal(saleOrderReportRes)

	t.Run("test case : pass", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
		saleOrderService.On("GetSaleOrdersByMonth", mockDate).Return(saleOrderReportRes, nil)
			
		req := httptest.NewRequest(http.MethodGet, "/saleorder/month/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2023-12")

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetAllSaleOrdersByMonth(c)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, strings.TrimSpace(string(saleOrderReportResJSON)), strings.TrimSpace(rec.Body.String()))
		saleOrderService.AssertExpectations(t)
	})

	t.Run("test case : fail validate date", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
			
		req := httptest.NewRequest(http.MethodGet, "/saleorder/month/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2")

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetAllSaleOrdersByMonth(c)

		assert.Equal(t, 400, rec.Code)
	})

	t.Run("test case : fail repository", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
		saleOrderService.On("GetSaleOrdersByMonth", mockDate).Return([]models.SaleOrderReport{}, errs.NewUnexpectedError(""))
			
		req := httptest.NewRequest(http.MethodGet, "/saleorder/month/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2023-12")

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetAllSaleOrdersByMonth(c)

		assert.Equal(t, 500, rec.Code)
		saleOrderService.AssertExpectations(t)
	})
}

func TestGetAllSaleOrdersByYear(t *testing.T) {
	e := echo.New()
	mockDate := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	saleOrderReportRes := []models.SaleOrderReport{
		{
			Seq:  			1,
			Id: 			1,
			User: 			"Walter White",
			Date:   		mockDate,
			TotalPrice:   	0,
			Status:        	"Fail",
			PayMethodName: 	"None",
		},
    }
	saleOrderReportResJSON, _ := json.Marshal(saleOrderReportRes)

	t.Run("test case : pass", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
		saleOrderService.On("GetSaleOrdersByYear", mockDate).Return(saleOrderReportRes, nil)
			
		req := httptest.NewRequest(http.MethodGet, "/saleorder/year/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2023")

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetAllSaleOrdersByYear(c)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, strings.TrimSpace(string(saleOrderReportResJSON)), strings.TrimSpace(rec.Body.String()))
		saleOrderService.AssertExpectations(t)
	})

	t.Run("test case : fail validate date", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
			
		req := httptest.NewRequest(http.MethodGet, "/saleorder/year/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2")

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetAllSaleOrdersByYear(c)

		assert.Equal(t, 400, rec.Code)
	})

	t.Run("test case : fail repository", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
		saleOrderService.On("GetSaleOrdersByYear", mockDate).Return([]models.SaleOrderReport{}, errs.NewUnexpectedError(""))
			
		req := httptest.NewRequest(http.MethodGet, "/saleorder/year/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2023")

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetAllSaleOrdersByYear(c)

		assert.Equal(t, 500, rec.Code)
		saleOrderService.AssertExpectations(t)
	})
}

func TestGetTotalPricePass(t *testing.T) {
	e := echo.New()
	totalPrice := 200.55
	totalPriceStr := fmt.Sprintf("%.2f", totalPrice)

	t.Run("test case : pass", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
		saleOrderService.On("GetSaleOrderPriceAmountPass").Return(totalPrice, nil)
			
		req := httptest.NewRequest(http.MethodGet, "/totalprice", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetTotalPricePass(c)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, strings.TrimSpace(totalPriceStr), strings.TrimSpace(rec.Body.String()))
		saleOrderService.AssertExpectations(t)
	})

	t.Run("test case : fail repository", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
		saleOrderService.On("GetSaleOrderPriceAmountPass").Return(0, errs.NewUnexpectedError(""))
			
		req := httptest.NewRequest(http.MethodGet, "/totalprice/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetTotalPricePass(c)

		assert.Equal(t, 500, rec.Code)
		saleOrderService.AssertExpectations(t)
	})
}

func TestGetTotalPricePassByDay(t *testing.T) {
	e := echo.New()
	mockDate := time.Date(2023, time.December, 28, 0, 0, 0, 0, time.UTC)
	totalPrice := 200.55
	totalPriceStr := fmt.Sprintf("%.2f", totalPrice)

	t.Run("test case : pass", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
		saleOrderService.On("GetSaleOrderPriceAmountPassByDay", mockDate).Return(totalPrice, nil)
			
		req := httptest.NewRequest(http.MethodGet, "/totalprice/day/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2023-12-28")

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetTotalPricePassByDay(c)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, strings.TrimSpace(totalPriceStr), strings.TrimSpace(rec.Body.String()))
		saleOrderService.AssertExpectations(t)
	})

	t.Run("test case : fail validate date", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
			
		req := httptest.NewRequest(http.MethodGet, "/totalprice/day/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2")

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetTotalPricePassByDay(c)

		assert.Equal(t, 400, rec.Code)
	})

	t.Run("test case : fail repository", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
		saleOrderService.On("GetSaleOrderPriceAmountPassByDay", mockDate).Return(0, errs.NewUnexpectedError(""))
			
		req := httptest.NewRequest(http.MethodGet, "/totalprice/day/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2023-12-28")

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetTotalPricePassByDay(c)

		assert.Equal(t, 500, rec.Code)
		saleOrderService.AssertExpectations(t)
	})
}

func TestGetTotalPricePassByMonth(t *testing.T) {
	e := echo.New()
	mockDate := time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC)
	totalPrice := 200.55
	totalPriceStr := fmt.Sprintf("%.2f", totalPrice)

	t.Run("test case : pass", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
		saleOrderService.On("GetSaleOrderPriceAmountPassByMonth", mockDate).Return(totalPrice, nil)
			
		req := httptest.NewRequest(http.MethodGet, "/totalprice/month/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2023-12")

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetTotalPricePassByMonth(c)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, strings.TrimSpace(totalPriceStr), strings.TrimSpace(rec.Body.String()))
		saleOrderService.AssertExpectations(t)
	})

	t.Run("test case : fail validate date", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
			
		req := httptest.NewRequest(http.MethodGet, "/totalprice/month/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2")

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetTotalPricePassByMonth(c)

		assert.Equal(t, 400, rec.Code)
	})

	t.Run("test case : fail repository", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
		saleOrderService.On("GetSaleOrderPriceAmountPassByMonth", mockDate).Return(0, errs.NewUnexpectedError(""))
			
		req := httptest.NewRequest(http.MethodGet, "/totalprice/month/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2023-12")

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetTotalPricePassByMonth(c)

		assert.Equal(t, 500, rec.Code)
		saleOrderService.AssertExpectations(t)
	})
}

func TestGetTotalPricePassByYear(t *testing.T) {
	e := echo.New()
	mockDate := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	totalPrice := 200.55
	totalPriceStr := fmt.Sprintf("%.2f", totalPrice)

	t.Run("test case : pass", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
		saleOrderService.On("GetSaleOrderPriceAmountPassByYear", mockDate).Return(totalPrice, nil)
			
		req := httptest.NewRequest(http.MethodGet, "/totalprice/year/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2023")

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetTotalPricePassByYear(c)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, strings.TrimSpace(totalPriceStr), strings.TrimSpace(rec.Body.String()))
		saleOrderService.AssertExpectations(t)
	})

	t.Run("test case : fail validate date", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
			
		req := httptest.NewRequest(http.MethodGet, "/totalprice/year/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2")

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetTotalPricePassByYear(c)

		assert.Equal(t, 400, rec.Code)
	})

	t.Run("test case : fail repository", func(t *testing.T) {
		saleOrderService := mock_services.NewSaleOrderServiceMock()
		saleOrderService.On("GetSaleOrderPriceAmountPassByYear", mockDate).Return(0, errs.NewUnexpectedError(""))
			
		req := httptest.NewRequest(http.MethodGet, "/totalprice/year/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("date")
		c.SetParamValues("2023")

		saleOrderHandler := handlers.NewSaleOrderHandler(saleOrderService)
		saleOrderHandler.GetTotalPricePassByYear(c)

		assert.Equal(t, 500, rec.Code)
		saleOrderService.AssertExpectations(t)
	})
}