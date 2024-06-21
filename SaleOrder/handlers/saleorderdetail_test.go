package handlers_test

import (
	"encoding/json"
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

func TestCreateSaleOrderDetail(t *testing.T) {
	saleOrderDetailReq := &models.SaleOrderDetailCreate{
		SO_Id: 1,
		Prod_Id:  200,
		Quantity:  1,
		Price:   200,
		Discount:   0,
	}
	saleOrderDetailErrorReq := &models.SaleOrderDetailCreate{
		SO_Id: 0,
		Prod_Id:  0,
		Quantity:  0,
		Price:   0,
		Discount:   0,
	}

	saleOrderDetailRes := &models.SaleOrderDetail{
		Seq: 1,
		SO_Id: 1,
		Prod_Id:  200,
		Quantity:  1,
		Price:   200,
		Discount:   0,
	}

	e := echo.New()

	saleOrderDetailReqJSON, _ := json.Marshal(saleOrderDetailReq)
	saleOrderDetailReqErrorJSON, _ := json.Marshal(saleOrderDetailErrorReq)
	saleOrderDetailResJSON, _ := json.Marshal(saleOrderDetailRes)

	type testCase struct {
		name         	string
		isValidate      bool
		body            string
		insertRepo      models.SaleOrderDetailCreate
		expectedStatus  int
		expectedBody    string
		srvReturn1 		models.SaleOrderDetail
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    			isValidate: false,  body: string(saleOrderDetailReqJSON), 	  	insertRepo: *saleOrderDetailReq,   	 	expectedStatus: 201,	expectedBody: string(saleOrderDetailResJSON), 	srvReturn1: *saleOrderDetailRes,      srvReturn2: nil},				
		{name: "test case : failed bind",  		isValidate: true, 	body: "invalid json", 			         	insertRepo: *saleOrderDetailErrorReq,  	expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,    	srvReturn1: models.SaleOrderDetail{}, srvReturn2: nil},   		
		{name: "test case : failed validator",  isValidate: true, 	body: string(saleOrderDetailReqErrorJSON),	insertRepo: *saleOrderDetailErrorReq,  	expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,    	srvReturn1: models.SaleOrderDetail{}, srvReturn2: nil},     
		{name: "test case : failed srvsitory", 	isValidate: false,  body: string(saleOrderDetailReqJSON), 	  	insertRepo: *saleOrderDetailReq,		expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	  	srvReturn1: models.SaleOrderDetail{}, srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			saleOrderDetailService := mock_services.NewSaleOrderDetailServiceMock()
			if !tc.isValidate {
				saleOrderDetailService.On("CreateSaleOrderDetail", &tc.insertRepo).Return(&tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/saleordertail/", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
	
			saleOrderDetailHandler := handlers.NewSaleOrderDetailHandler(saleOrderDetailService)
			saleOrderDetailHandler.CreateSaleOrderDetail(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isValidate {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
				saleOrderDetailService.AssertExpectations(t)
			}
		})
	}
}

func TestGetSaleOrderDetailQtyRates(t *testing.T) {
	saleOrderDetailRes := []models.SaleOrderDetailRate{
		{
			Prod_Id:  200,
			Quantity:  1,
			Price:   200,
		},
	}

	e := echo.New()

	saleOrderDetailResJSON, _ := json.Marshal(saleOrderDetailRes)

	type testCase struct {
		name         	string
		isFail          bool
		body            string
		expectedStatus  int
		expectedBody    string
		srvReturn1 	[]models.SaleOrderDetailRate
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    			isFail: false, 	expectedStatus: 200,	expectedBody: string(saleOrderDetailResJSON), srvReturn1: saleOrderDetailRes,      	 	srvReturn2: nil},				
		{name: "test case : failed service", 	isFail: true, 	expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	  srvReturn1: []models.SaleOrderDetailRate{},  srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			saleOrderDetailService := mock_services.NewSaleOrderDetailServiceMock()
			saleOrderDetailService.On("GetSaleOrderDetailQtyRates").Return(tc.srvReturn1, tc.srvReturn2)
	
			req := httptest.NewRequest(http.MethodPost, "/saleordertail/qtyrate", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
	
			saleOrderDetailHandler := handlers.NewSaleOrderDetailHandler(saleOrderDetailService)
			saleOrderDetailHandler.GetSaleOrderDetailQtyRates(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isFail {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
			}
			saleOrderDetailService.AssertExpectations(t)
		})
	}
}

func TestGetSaleOrderDetailQtyRatesByDay(t *testing.T) {
	saleOrderDetailRes := []models.SaleOrderDetailRate{
		{
			Prod_Id:  200,
			Quantity:  1,
			Price:   200,
		},
	}

	e := echo.New()
	mockDate := time.Date(2023, time.December, 28, 0, 0, 0, 0, time.UTC)

	saleOrderDetailResJSON, _ := json.Marshal(saleOrderDetailRes)

	type testCase struct {
		name         	string
		isFail          bool
		isValidate      bool
		paramName       string
		expectedStatus  int
		expectedBody    string
		srvReturn1 	[]models.SaleOrderDetailRate
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    				isFail: false, isValidate: true, 	expectedStatus: 200, paramName: "2023-12-28",	expectedBody: string(saleOrderDetailResJSON), srvReturn1: saleOrderDetailRes,      	 	srvReturn2: nil},				
		{name: "test case : failed validate date", 	isFail: true,  isValidate: false, 	expectedStatus: 400, paramName: "2",			expectedBody: `{"code":400,"message":""}`,	  srvReturn1: []models.SaleOrderDetailRate{},  srvReturn2: nil},			
		{name: "test case : failed service", 		isFail: true,  isValidate: true, 	expectedStatus: 500, paramName: "2023-12-28",	expectedBody: `{"code":500,"message":""}`,	  srvReturn1: []models.SaleOrderDetailRate{},  srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			saleOrderDetailService := mock_services.NewSaleOrderDetailServiceMock()
			if tc.isValidate{
				saleOrderDetailService.On("GetSaleOrderDetailQtyRatesByDay", mockDate).Return(tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/saleordertail/qtyrate/day", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("date")
			c.SetParamValues(tc.paramName)
	
			saleOrderDetailHandler := handlers.NewSaleOrderDetailHandler(saleOrderDetailService)
			saleOrderDetailHandler.GetSaleOrderDetailQtyRatesByDay(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isFail {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
			}
			saleOrderDetailService.AssertExpectations(t)
		})
	}
}

func TestGetSaleOrderDetailQtyRatesByMonth(t *testing.T) {
	saleOrderDetailRes := []models.SaleOrderDetailRate{
		{
			Prod_Id:  200,
			Quantity:  1,
			Price:   200,
		},
	}

	e := echo.New()
	mockDate := time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC)

	saleOrderDetailResJSON, _ := json.Marshal(saleOrderDetailRes)

	type testCase struct {
		name         	string
		isFail          bool
		isValidate      bool
		paramName       string
		expectedStatus  int
		expectedBody    string
		srvReturn1 	[]models.SaleOrderDetailRate
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    				isFail: false, isValidate: true, 	expectedStatus: 200, paramName: "2023-12",	expectedBody: string(saleOrderDetailResJSON), srvReturn1: saleOrderDetailRes,      	 	srvReturn2: nil},				
		{name: "test case : failed validate date", 	isFail: true,  isValidate: false, 	expectedStatus: 400, paramName: "2",		expectedBody: `{"code":400,"message":""}`,	  srvReturn1: []models.SaleOrderDetailRate{},  srvReturn2: nil},			
		{name: "test case : failed service", 		isFail: true,  isValidate: true, 	expectedStatus: 500, paramName: "2023-12",	expectedBody: `{"code":500,"message":""}`,	  srvReturn1: []models.SaleOrderDetailRate{},  srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			saleOrderDetailService := mock_services.NewSaleOrderDetailServiceMock()
			if tc.isValidate{
				saleOrderDetailService.On("GetSaleOrderDetailQtyRatesByMonth", mockDate).Return(tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/saleordertail/qtyrate/month", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("date")
			c.SetParamValues(tc.paramName)
	
			saleOrderDetailHandler := handlers.NewSaleOrderDetailHandler(saleOrderDetailService)
			saleOrderDetailHandler.GetSaleOrderDetailQtyRatesByMonth(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isFail {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
			}
			saleOrderDetailService.AssertExpectations(t)
		})
	}
}

func TestGetSaleOrderDetailQtyRatesByYear(t *testing.T) {
	saleOrderDetailRes := []models.SaleOrderDetailRate{
		{
			Prod_Id:  200,
			Quantity:  1,
			Price:   200,
		},
	}

	e := echo.New()
	mockDate := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)

	saleOrderDetailResJSON, _ := json.Marshal(saleOrderDetailRes)

	type testCase struct {
		name         	string
		isFail          bool
		isValidate      bool
		paramName       string
		expectedStatus  int
		expectedBody    string
		srvReturn1 	[]models.SaleOrderDetailRate
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    				isFail: false, isValidate: true, 	expectedStatus: 200, paramName: "2023",	expectedBody: string(saleOrderDetailResJSON), srvReturn1: saleOrderDetailRes,      	 	srvReturn2: nil},				
		{name: "test case : failed validate date", 	isFail: true,  isValidate: false, 	expectedStatus: 400, paramName: "2",	expectedBody: `{"code":400,"message":""}`,	  srvReturn1: []models.SaleOrderDetailRate{},  srvReturn2: nil},			
		{name: "test case : failed service", 		isFail: true,  isValidate: true, 	expectedStatus: 500, paramName: "2023",	expectedBody: `{"code":500,"message":""}`,	  srvReturn1: []models.SaleOrderDetailRate{},  srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			saleOrderDetailService := mock_services.NewSaleOrderDetailServiceMock()
			if tc.isValidate{
				saleOrderDetailService.On("GetSaleOrderDetailQtyRatesByYear", mockDate).Return(tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/saleordertail/qtyrate/year", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("date")
			c.SetParamValues(tc.paramName)
	
			saleOrderDetailHandler := handlers.NewSaleOrderDetailHandler(saleOrderDetailService)
			saleOrderDetailHandler.GetSaleOrderDetailQtyRatesByYear(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isFail {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
			}
			saleOrderDetailService.AssertExpectations(t)
		})
	}
}

func TestGetSaleOrderDetailPriceRates(t *testing.T) {
	saleOrderDetailRes := []models.SaleOrderDetailRate{
		{
			Prod_Id:  200,
			Quantity:  1,
			Price:   200,
		},
	}

	e := echo.New()

	saleOrderDetailResJSON, _ := json.Marshal(saleOrderDetailRes)

	type testCase struct {
		name         	string
		isFail          bool
		body            string
		expectedStatus  int
		expectedBody    string
		srvReturn1 	[]models.SaleOrderDetailRate
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    			isFail: false, 	expectedStatus: 200,	expectedBody: string(saleOrderDetailResJSON), srvReturn1: saleOrderDetailRes,      	 	srvReturn2: nil},				
		{name: "test case : failed service", 	isFail: true, 	expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	  srvReturn1: []models.SaleOrderDetailRate{},  srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			saleOrderDetailService := mock_services.NewSaleOrderDetailServiceMock()
			saleOrderDetailService.On("GetSaleOrderDetailPriceRates").Return(tc.srvReturn1, tc.srvReturn2)
	
			req := httptest.NewRequest(http.MethodPost, "/saleordertail/pricerate", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
	
			saleOrderDetailHandler := handlers.NewSaleOrderDetailHandler(saleOrderDetailService)
			saleOrderDetailHandler.GetSaleOrderDetailPriceRates(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isFail {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
			}
			saleOrderDetailService.AssertExpectations(t)
		})
	}
}

func TestGetSaleOrderDetailPriceRatesByDay(t *testing.T) {
	saleOrderDetailRes := []models.SaleOrderDetailRate{
		{
			Prod_Id:  200,
			Quantity:  1,
			Price:   200,
		},
	}

	e := echo.New()
	mockDate := time.Date(2023, time.December, 28, 0, 0, 0, 0, time.UTC)

	saleOrderDetailResJSON, _ := json.Marshal(saleOrderDetailRes)

	type testCase struct {
		name         	string
		isFail          bool
		isValidate      bool
		paramName       string
		expectedStatus  int
		expectedBody    string
		srvReturn1 	[]models.SaleOrderDetailRate
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    				isFail: false, isValidate: true, 	expectedStatus: 200, paramName: "2023-12-28",	expectedBody: string(saleOrderDetailResJSON), srvReturn1: saleOrderDetailRes,      	 	srvReturn2: nil},				
		{name: "test case : failed validate date", 	isFail: true,  isValidate: false, 	expectedStatus: 400, paramName: "2",			expectedBody: `{"code":400,"message":""}`,	  srvReturn1: []models.SaleOrderDetailRate{},  srvReturn2: nil},			
		{name: "test case : failed service", 		isFail: true,  isValidate: true, 	expectedStatus: 500, paramName: "2023-12-28",	expectedBody: `{"code":500,"message":""}`,	  srvReturn1: []models.SaleOrderDetailRate{},  srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			saleOrderDetailService := mock_services.NewSaleOrderDetailServiceMock()
			if tc.isValidate{
				saleOrderDetailService.On("GetSaleOrderDetailPriceRatesByDay", mockDate).Return(tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/saleordertail/pricerate/day", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("date")
			c.SetParamValues(tc.paramName)
	
			saleOrderDetailHandler := handlers.NewSaleOrderDetailHandler(saleOrderDetailService)
			saleOrderDetailHandler.GetSaleOrderDetailPriceRatesByDay(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isFail {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
			}
			saleOrderDetailService.AssertExpectations(t)
		})
	}
}

func TestGetSaleOrderDetailPriceRatesByMonth(t *testing.T) {
	saleOrderDetailRes := []models.SaleOrderDetailRate{
		{
			Prod_Id:  200,
			Quantity:  1,
			Price:   200,
		},
	}

	e := echo.New()
	mockDate := time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC)

	saleOrderDetailResJSON, _ := json.Marshal(saleOrderDetailRes)

	type testCase struct {
		name         	string
		isFail          bool
		isValidate      bool
		paramName       string
		expectedStatus  int
		expectedBody    string
		srvReturn1 	[]models.SaleOrderDetailRate
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    				isFail: false, isValidate: true, 	expectedStatus: 200, paramName: "2023-12",	expectedBody: string(saleOrderDetailResJSON), srvReturn1: saleOrderDetailRes,      	 	srvReturn2: nil},				
		{name: "test case : failed validate date", 	isFail: true,  isValidate: false, 	expectedStatus: 400, paramName: "2",		expectedBody: `{"code":400,"message":""}`,	  srvReturn1: []models.SaleOrderDetailRate{},  srvReturn2: nil},			
		{name: "test case : failed service", 		isFail: true,  isValidate: true, 	expectedStatus: 500, paramName: "2023-12",	expectedBody: `{"code":500,"message":""}`,	  srvReturn1: []models.SaleOrderDetailRate{},  srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			saleOrderDetailService := mock_services.NewSaleOrderDetailServiceMock()
			if tc.isValidate{
				saleOrderDetailService.On("GetSaleOrderDetailPriceRatesByMonth", mockDate).Return(tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/saleordertail/pricerate/month", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("date")
			c.SetParamValues(tc.paramName)
	
			saleOrderDetailHandler := handlers.NewSaleOrderDetailHandler(saleOrderDetailService)
			saleOrderDetailHandler.GetSaleOrderDetailPriceRatesByMonth(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isFail {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
			}
			saleOrderDetailService.AssertExpectations(t)
		})
	}
}

func TestGetSaleOrderDetailPriceRatesByYear(t *testing.T) {
	saleOrderDetailRes := []models.SaleOrderDetailRate{
		{
			Prod_Id:  200,
			Quantity:  1,
			Price:   200,
		},
	}

	e := echo.New()
	mockDate := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)

	saleOrderDetailResJSON, _ := json.Marshal(saleOrderDetailRes)

	type testCase struct {
		name         	string
		isFail          bool
		isValidate      bool
		paramName       string
		expectedStatus  int
		expectedBody    string
		srvReturn1 	[]models.SaleOrderDetailRate
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    				isFail: false, isValidate: true, 	expectedStatus: 200, paramName: "2023",	expectedBody: string(saleOrderDetailResJSON), srvReturn1: saleOrderDetailRes,      	 	srvReturn2: nil},				
		{name: "test case : failed validate date", 	isFail: true,  isValidate: false, 	expectedStatus: 400, paramName: "2",	expectedBody: `{"code":400,"message":""}`,	  srvReturn1: []models.SaleOrderDetailRate{},  srvReturn2: nil},			
		{name: "test case : failed service", 		isFail: true,  isValidate: true, 	expectedStatus: 500, paramName: "2023",	expectedBody: `{"code":500,"message":""}`,	  srvReturn1: []models.SaleOrderDetailRate{},  srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			saleOrderDetailService := mock_services.NewSaleOrderDetailServiceMock()
			if tc.isValidate{
				saleOrderDetailService.On("GetSaleOrderDetailPriceRatesByYear", mockDate).Return(tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/saleordertail/pricerate/year", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("date")
			c.SetParamValues(tc.paramName)
	
			saleOrderDetailHandler := handlers.NewSaleOrderDetailHandler(saleOrderDetailService)
			saleOrderDetailHandler.GetSaleOrderDetailPriceRatesByYear(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isFail {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
			}
			saleOrderDetailService.AssertExpectations(t)
		})
	}
}