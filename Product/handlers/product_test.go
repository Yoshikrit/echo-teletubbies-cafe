package handlers_test

import (
    "product/utils/errs"
	"testing"
	"net/http"
	"net/http/httptest"
	// "strconv"
	"strings"
	"product/tests/mocks/mock_services"
	"product/handlers"
	"encoding/json"

	"github.com/stretchr/testify/assert"

	"product/models"
    "github.com/labstack/echo/v4"
)

func TestCreateProduct(t *testing.T) {
	e := echo.New()

	prodReqMock := &models.ProductCreate{
		Id:   1,
		ProdType_Id:   1,
		Name: "A",
		Desc: "A",
		Price: 100,
		Discount: 0,
		CreatedUser: 1,
	}

	prodReqErrorMock := &models.ProductCreate{
		Id:   0,
		ProdType_Id:   0,
		Name: "",
		Desc: "",
		Price: 0,
		Discount: 0,
		CreatedUser: 0,
	}

	prodResMock := &models.Product{
		Id:   1,
		ProdType_Id:   1,
		Name: "B",
		Desc: "A",
		Price: 100,
		Discount: 0,
	}

	prodReqJSON, _ := json.Marshal(prodReqMock)
	prodReqErrorJSON, _ := json.Marshal(prodReqErrorMock)
	prodResJSON, _ := json.Marshal(prodResMock)

	type testCase struct {
		name         	string
		isValidate      bool
		body            string
		insertSrv      	models.ProductCreate
		expectedStatus  int
		expectedBody    string
		srvReturn1 		models.Product
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    			isValidate: false,  body: string(prodReqJSON), 	  	insertSrv: *prodReqMock,   	 	expectedStatus: 201,	expectedBody: string(prodResJSON), 				srvReturn1: *prodResMock,     srvReturn2: nil},				
		{name: "test case : failed bind",  		isValidate: true, 	body: "invalid json", 			insertSrv: *prodReqErrorMock,  	expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,    	srvReturn1: models.Product{}, srvReturn2: nil},   		
		{name: "test case : failed validator",  isValidate: true, 	body: string(prodReqErrorJSON),	insertSrv: *prodReqErrorMock,  	expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,    	srvReturn1: models.Product{}, srvReturn2: nil},     
		{name: "test case : failed srvsitory", 	isValidate: false,  body: string(prodReqJSON), 	  	insertSrv: *prodReqMock,		expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	  	srvReturn1: models.Product{}, srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			prodService := mock_services.NewProductServiceMock()
			if !tc.isValidate {
				prodService.On("CreateProduct", &tc.insertSrv).Return(&tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/product/", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
	
			prodHandler := handlers.NewProductHandler(prodService)
			prodHandler.CreateProduct(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isValidate {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
				prodService.AssertExpectations(t)
			}
		})
	}
}

func TestGetAllProducts(t *testing.T) {
	e := echo.New()

	prodsResMock := []models.Product {
		{
			Id:   1,
			ProdType_Id:   1,
			Name: "A",
			Desc: "A",
			Price: 100,
			Discount: 0,
		},
		{
			Id:   2,
			ProdType_Id:   1,
			Name: "B",
			Desc: "A",
			Price: 100,
			Discount: 0,
		},
	}

	prodsResJSON, _ := json.Marshal(prodsResMock)

	type testCase struct {
		name         	string
		isFail          bool
		body            string
		expectedStatus  int
		expectedBody    string
		srvReturn1 		[]models.Product
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    			isFail: false, 	expectedStatus: 200,	expectedBody: string(prodsResJSON), 		srvReturn1: prodsResMock,  	 srvReturn2: nil},				
		{name: "test case : failed service", 	isFail: true, 	expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	srvReturn1: []models.Product{},  srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			prodService := mock_services.NewProductServiceMock()
			prodService.On("GetProducts").Return(tc.srvReturn1, tc.srvReturn2)
	
			req := httptest.NewRequest(http.MethodPost, "/product/", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
	
			prodHandler := handlers.NewProductHandler(prodService)
			prodHandler.GetAllProducts(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isFail {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
			}
			prodService.AssertExpectations(t)
		})
	}
}

func TestGetProductByID(t *testing.T) {
	e := echo.New()

	prodResMock := models.Product {
		Id:   1,
		ProdType_Id:   1,
		Name: "A",
		Desc: "A",
		Price: 100,
		Discount: 0,
	}
	
	prodResJSON, _ := json.Marshal(prodResMock)

	type testCase struct {
		name         	string
		isValidate      bool
		param           string
		insertSrv      	int
		expectedStatus  int
		expectedBody    string
		srvReturn1 		models.Product
		srvReturn2 		error
	}

	cases := []testCase{
		{name: "test case : pass",    			isValidate: false,  param: "1", insertSrv: 1,   	expectedStatus: 200,	expectedBody: string(prodResJSON), 			srvReturn1: prodResMock,  	 	srvReturn2: nil},					
		{name: "test case : failed param int",  isValidate: true, 	param: "a", insertSrv: 0,  		expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,  srvReturn1: models.Product{},  srvReturn2: errs.NewBadRequestError("")},     
		{name: "test case : failed service", 	isValidate: false,  param: "1", insertSrv: 1,		expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	srvReturn1: models.Product{},  srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			prodService := mock_services.NewProductServiceMock()
			if !tc.isValidate {
				prodService.On("GetProduct", tc.insertSrv).Return(&tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/product/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tc.param)
	
			prodHandler := handlers.NewProductHandler(prodService)
			prodHandler.GetProductByID(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isValidate {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
				prodService.AssertExpectations(t)
			}
		})
	}
}

func TestUpdateProductByID(t *testing.T) {
	e := echo.New()

	prodReqMock := &models.ProductUpdate {
		ProdType_Id:   1,
		Name: "B",
		Desc: "A",
		Price: 100,
		Discount: 0,
		UpdatedUser: 1,
	}

	prodResMock := &models.Product {
		Id:   1,
		ProdType_Id:   1,
		Name: "B",
		Desc: "A",
		Price: 100,
		Discount: 0,
	}
	
	prodReqJSON, _ := json.Marshal(prodReqMock)
	prodResJSON, _ := json.Marshal(prodResMock)

	type testCase struct {
		name         	string
		isValidate      bool
		param           string
		body            string
		insertId        int
		insertSrv      	models.ProductUpdate
		expectedStatus  int
		expectedBody    string
		srvReturn1 		models.Product
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    			param: "1", isValidate: false,  body: string(prodReqJSON), 	  	insertId: 1, insertSrv: *prodReqMock,   	 	expectedStatus: 200,	expectedBody: string(prodResJSON), 				srvReturn1: *prodResMock,     srvReturn2: nil},				
		{name: "test case : failed param int",  param: "a", isValidate: true, 	body: string(prodReqJSON),		insertId: 1, insertSrv: *prodReqMock,  			expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,  	srvReturn1: models.Product{}, srvReturn2: nil},     
		{name: "test case : failed bind",  		param: "1", isValidate: true, 	body: "invalid json", 			insertId: 1, insertSrv: models.ProductUpdate{}, expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,    	srvReturn1: models.Product{}, srvReturn2: nil},   		
		{name: "test case : failed service", 	param: "1", isValidate: false,  body: string(prodReqJSON), 	  	insertId: 1, insertSrv: *prodReqMock,			expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	  	srvReturn1: models.Product{}, srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			prodService := mock_services.NewProductServiceMock()
			if !tc.isValidate {
				prodService.On("UpdateProduct", tc.insertId, &tc.insertSrv).Return(&tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPut, "/product/", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tc.param)
	
			prodHandler := handlers.NewProductHandler(prodService)
			prodHandler.UpdateProductByID(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isValidate {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
				prodService.AssertExpectations(t)
			}
		})
	}
}


func TestDeleteProductByID(t *testing.T) {
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
			prodService := mock_services.NewProductServiceMock()
			if !tc.isValidate {
				prodService.On("DeleteProduct", tc.insertId).Return(tc.srvReturn)
			}
	
			req := httptest.NewRequest(http.MethodDelete, "/product/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tc.param)
	
			prodHandler := handlers.NewProductHandler(prodService)
			prodHandler.DeleteProductByID(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isValidate {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
				prodService.AssertExpectations(t)
			}
		})
	}
}

func TestGetProductCount(t *testing.T) {
    e := echo.New()

    t.Run("test case : pass", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/count", nil)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		prodService := mock_services.NewProductServiceMock()
		prodService.On("GetProductCount").Return(int64(42), nil)

		prodHandler := handlers.NewProductHandler(prodService)
		prodHandler.GetProductCount(c)

		expectedCode := 200
		expectedBody := `42`
		assert.Equal(t, expectedCode, rec.Code)
		assert.Equal(t, expectedBody, strings.TrimSpace(rec.Body.String()))
		prodService.AssertExpectations(t)
	})

	t.Run("test case : pass", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/count", nil)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		prodService := mock_services.NewProductServiceMock()
		prodService.On("GetProductCount").Return(int64(0), errs.NewUnexpectedError(""))

		prodHandler := handlers.NewProductHandler(prodService)
		prodHandler.GetProductCount(c)

		expectedCode := 500
		expectedBody := `{"code":500,"message":""}`
		assert.Equal(t, expectedCode, rec.Code)
		assert.Equal(t, expectedBody, strings.TrimSpace(rec.Body.String()))
		prodService.AssertExpectations(t)
	})
}
