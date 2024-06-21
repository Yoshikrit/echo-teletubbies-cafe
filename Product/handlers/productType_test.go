package handlers_test

import (
    "product/utils/errs"
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
	"product/tests/mocks/mock_services"
	"product/handlers"
	"encoding/json"

	"github.com/stretchr/testify/assert"

	"product/models"
    "github.com/labstack/echo/v4"
)

func TestCreateProductType(t *testing.T) {
	e := echo.New()

	prodTypeReqMock := &models.ProductTypeCreate{
		Id:   1,
		Name: "A",
	}

	prodTypeReqErrorMock := &models.ProductTypeCreate{
		Id:   0,
		Name: "",
	}

	prodTypeResMock := &models.ProductType{
		Id:   1,
		Name: "B",
	}

	prodTypeReqJSON, _ := json.Marshal(prodTypeReqMock)
	prodTypeReqErrorJSON, _ := json.Marshal(prodTypeReqErrorMock)
	prodTypeResJSON, _ := json.Marshal(prodTypeResMock)

	type testCase struct {
		name         	string
		isValidate      bool
		body            string
		insertSrv      	models.ProductTypeCreate
		expectedStatus  int
		expectedBody    string
		srvReturn1 		models.ProductType
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    			isValidate: false,  body: string(prodTypeReqJSON), 	  	insertSrv: *prodTypeReqMock,   	 	expectedStatus: 201,	expectedBody: string(prodTypeResJSON), 			srvReturn1: *prodTypeResMock,     srvReturn2: nil},				
		{name: "test case : failed bind",  		isValidate: true, 	body: "invalid json", 				insertSrv: *prodTypeReqErrorMock,  	expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,    	srvReturn1: models.ProductType{}, srvReturn2: nil},   		
		{name: "test case : failed validator",  isValidate: true, 	body: string(prodTypeReqErrorJSON),	insertSrv: *prodTypeReqErrorMock,  	expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,    	srvReturn1: models.ProductType{}, srvReturn2: nil},     
		{name: "test case : failed srvsitory", 	isValidate: false,  body: string(prodTypeReqJSON), 	  	insertSrv: *prodTypeReqMock,		expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	  	srvReturn1: models.ProductType{}, srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			prodTypeService := mock_services.NewProductTypeServiceMock()
			if !tc.isValidate {
				prodTypeService.On("CreateProductType", &tc.insertSrv).Return(&tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/producttype/", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
	
			prodTypeHandler := handlers.NewProductTypeHandler(prodTypeService)
			prodTypeHandler.CreateProductType(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isValidate {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
				prodTypeService.AssertExpectations(t)
			}
		})
	}
}

func TestGetAllProductTypes(t *testing.T) {
	e := echo.New()

	prodTypesResMock := []models.ProductType {
		{
			Id:   1,
			Name: "A",
		},
		{
			Id:   2,
			Name: "B",
		},
	}

	prodTypesResJSON, _ := json.Marshal(prodTypesResMock)

	type testCase struct {
		name         	string
		isFail          bool
		body            string
		expectedStatus  int
		expectedBody    string
		srvReturn1 		[]models.ProductType
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    			isFail: false, 	expectedStatus: 200,	expectedBody: string(prodTypesResJSON), 		srvReturn1: prodTypesResMock,  	 srvReturn2: nil},				
		{name: "test case : failed service", 	isFail: true, 	expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	srvReturn1: []models.ProductType{},  srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			prodTypeService := mock_services.NewProductTypeServiceMock()
			prodTypeService.On("GetProductTypes").Return(tc.srvReturn1, tc.srvReturn2)
	
			req := httptest.NewRequest(http.MethodPost, "/producttype/", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
	
			prodTypeHandler := handlers.NewProductTypeHandler(prodTypeService)
			prodTypeHandler.GetAllProductTypes(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isFail {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
			}
			prodTypeService.AssertExpectations(t)
		})
	}
}

func TestGetProductTypeByID(t *testing.T) {
	e := echo.New()

	prodTypeResMock := models.ProductType {
		Id:   1,
		Name: "A",
	}
	
	prodTypeResJSON, _ := json.Marshal(prodTypeResMock)

	type testCase struct {
		name         	string
		isValidate      bool
		param           string
		insertSrv      	int
		expectedStatus  int
		expectedBody    string
		srvReturn1 		models.ProductType
		srvReturn2 		error
	}

	cases := []testCase{
		{name: "test case : pass",    			isValidate: false,  param: "1", insertSrv: 1,   	expectedStatus: 200,	expectedBody: string(prodTypeResJSON), 			srvReturn1: prodTypeResMock,  	 	srvReturn2: nil},					
		{name: "test case : failed param int",  isValidate: true, 	param: "a", insertSrv: 0,  		expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,  srvReturn1: models.ProductType{},  srvReturn2: errs.NewBadRequestError("")},     
		{name: "test case : failed service", 	isValidate: false,  param: "1", insertSrv: 1,		expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	srvReturn1: models.ProductType{},  srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			prodTypeService := mock_services.NewProductTypeServiceMock()
			if !tc.isValidate {
				prodTypeService.On("GetProductType", tc.insertSrv).Return(&tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/producttype/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tc.param)
	
			prodTypeHandler := handlers.NewProductTypeHandler(prodTypeService)
			prodTypeHandler.GetProductTypeByID(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isValidate {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
				prodTypeService.AssertExpectations(t)
			}
		})
	}
}

func TestUpdateProductTypeByID(t *testing.T) {
	e := echo.New()

	prodTypeReqMock := &models.ProductTypeUpdate {
		Name: "B",
	}

	prodTypeErrorReqMock := &models.ProductTypeUpdate {
		Name: "",
	}

	prodTypeResMock := &models.ProductType {
		Id:   1,
		Name: "B",
	}
	
	prodTypeReqJSON, _ := json.Marshal(prodTypeReqMock)
	prodTypeErrorReqJSON, _ := json.Marshal(prodTypeErrorReqMock)
	prodTypeResJSON, _ := json.Marshal(prodTypeResMock)

	type testCase struct {
		name         	string
		isValidate      bool
		param           string
		body            string
		insertId        int
		insertSrv      	models.ProductTypeUpdate
		expectedStatus  int
		expectedBody    string
		srvReturn1 		models.ProductType
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    			param: "1", isValidate: false,  body: string(prodTypeReqJSON), 	  		insertId: 1, insertSrv: *prodTypeReqMock,   	 	expectedStatus: 200,	expectedBody: string(prodTypeResJSON), 			srvReturn1: *prodTypeResMock,     srvReturn2: nil},				
		{name: "test case : failed param int",  param: "a", isValidate: true, 	body: string(prodTypeReqJSON),			insertId: 1, insertSrv: *prodTypeReqMock,  			expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,  	srvReturn1: models.ProductType{}, srvReturn2: nil},     
		{name: "test case : failed bind",  		param: "1", isValidate: true, 	body: "invalid json", 					insertId: 1, insertSrv: models.ProductTypeUpdate{}, expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,    	srvReturn1: models.ProductType{}, srvReturn2: nil},   
		{name: "test case : failed validator",  param: "1", isValidate: true, 	body: string(prodTypeErrorReqJSON), 	insertId: 1, insertSrv: models.ProductTypeUpdate{}, expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,    	srvReturn1: models.ProductType{}, srvReturn2: nil},   			
		{name: "test case : failed service", 	param: "1", isValidate: false,  body: string(prodTypeReqJSON), 	  		insertId: 1, insertSrv: *prodTypeReqMock,			expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	  	srvReturn1: models.ProductType{}, srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			prodTypeService := mock_services.NewProductTypeServiceMock()
			if !tc.isValidate {
				prodTypeService.On("UpdateProductType", tc.insertId, &tc.insertSrv).Return(&tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPut, "/producttype/", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tc.param)
	
			prodTypeHandler := handlers.NewProductTypeHandler(prodTypeService)
			prodTypeHandler.UpdateProductTypeByID(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isValidate {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
				prodTypeService.AssertExpectations(t)
			}
		})
	}
}


func TestDeleteProductTypeByID(t *testing.T) {
	e := echo.New()
	
	type testCase struct {
		name         	string
		isValidate      bool
		param           string
		insertId        int
		expectedStatus  int
		expectedBody    string
		srvReturn	  	error
	}

	cases := []testCase{
		{name: "test case : pass",    			param: "1", isValidate: false,  insertId: 1, expectedStatus: 200,	expectedBody: `{"message":"Deleted Successfully"}`,	srvReturn: nil},				
		{name: "test case : failed param int",  param: "a", isValidate: true, 	insertId: 0, expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,  		srvReturn: nil},    		
		{name: "test case : failed service", 	param: "1", isValidate: false,  insertId: 1, expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,			srvReturn: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			prodTypeService := mock_services.NewProductTypeServiceMock()
			if !tc.isValidate {
				prodTypeService.On("DeleteProductType", tc.insertId).Return(tc.srvReturn)
			}
	
			req := httptest.NewRequest(http.MethodDelete, "/prodTypeuct/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tc.param)
	
			prodTypeHandler := handlers.NewProductTypeHandler(prodTypeService)
			prodTypeHandler.DeleteProductTypeByID(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isValidate {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
				prodTypeService.AssertExpectations(t)
			}
		})
	}
}

func TestGetProductTypeCount(t *testing.T) {
    e := echo.New()

    t.Run("test case : pass", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/count", nil)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		prodTypeService := mock_services.NewProductTypeServiceMock()
		prodTypeService.On("GetProductTypeCount").Return(int64(42), nil)

		prodTypeHandler := handlers.NewProductTypeHandler(prodTypeService)
		prodTypeHandler.GetProductTypeCount(c)

		expectedCode := 200
		expectedBody := `42`
		assert.Equal(t, expectedCode, rec.Code)
		assert.Equal(t, expectedBody, strings.TrimSpace(rec.Body.String()))
		prodTypeService.AssertExpectations(t)
	})

	t.Run("test case : fail repository", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/count", nil)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		prodTypeService := mock_services.NewProductTypeServiceMock()
		prodTypeService.On("GetProductTypeCount").Return(int64(0), errs.NewUnexpectedError(""))

		prodTypeHandler := handlers.NewProductTypeHandler(prodTypeService)
		prodTypeHandler.GetProductTypeCount(c)

		expectedCode := 500
		expectedBody := `{"code":500,"message":""}`
		assert.Equal(t, expectedCode, rec.Code)
		assert.Equal(t, expectedBody, strings.TrimSpace(rec.Body.String()))
		prodTypeService.AssertExpectations(t)
	})
}
