package handlers_test

import (
    "auth/utils/errs"
	"testing"
	"net/http"
	"strings"
	"net/http/httptest"
	"auth/tests/mocks/mock_services"
	"auth/handlers"
	"encoding/json"

	"github.com/stretchr/testify/assert"

	"auth/models"
    "github.com/labstack/echo/v4"
)

func TestLogin(t *testing.T) {
	userReq := &models.UserLogin{
		Email: "A",
		Password: "A",
	}

	userReqError := &models.UserLogin{
		Email: "",
		Password: "",
	}

	userRes := &models.Response{
		Id:   1,
		Name: "A",
		Role: "A",
		JWT: "A",
	}

	e := echo.New()

	userReqJSON, _ := json.Marshal(userReq)
	userReqErrorJSON, _ := json.Marshal(userReqError)
	userResJSON, _ := json.Marshal(userRes)

	type testCase struct {
		name         	string
		isValidate      bool
		body            string
		insertsrv      models.UserLogin
		expectedStatus  int
		expectedBody    string
		srvReturn1 	models.Response
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    			isValidate: false,  body: string(userReqJSON), 	  	insertsrv: *userReq,   	 	expectedStatus: 200,	expectedBody: string(userResJSON), 			  srvReturn1: *userRes,      		srvReturn2: nil},				
		{name: "test case : failed bind",  		isValidate: true, 	body: "invalid json", 			insertsrv: *userReqError,  	expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,    srvReturn1: models.Response{}, 	srvReturn2: nil},   		
		{name: "test case : failed validator",  isValidate: true, 	body: string(userReqErrorJSON),	insertsrv: *userReqError,  	expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,    srvReturn1: models.Response{}, 	srvReturn2: nil},     
		{name: "test case : failed srvsitory", 	isValidate: false,  body: string(userReqJSON), 	  	insertsrv: *userReq,			expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	  srvReturn1: models.Response{}, 	srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			authService := mock_services.NewAuthServiceMock()
			if !tc.isValidate {
				authService.On("Login", &tc.insertsrv).Return(&tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/login/", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
	
			authHandler := handlers.NewAuthHandler(authService)
			authHandler.Login(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isValidate {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
				authService.AssertExpectations(t)
			}
		})
	}
}

func TestLogout(t *testing.T) {
	e := echo.New()

	type testCase struct {
		name         	string
		isValidate      bool
		param           string
		insertsrv      	int
		expectedStatus  int
		expectedBody    string
		srvReturn 		error
	}

	cases := []testCase{
		{name: "test case : pass",    			isValidate: false,  param: "1", insertsrv: 1,   	expectedStatus: 200,	expectedBody: `"Log out successfully"`, 	srvReturn: nil},				
		{name: "test case : failed param int",  isValidate: true, 	param: "a", insertsrv: 0,  		expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,    					srvReturn: nil},     
		{name: "test case : failed service", 	isValidate: false,  param: "1", insertsrv: 1,		expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	  					srvReturn: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			authService := mock_services.NewAuthServiceMock()
			if !tc.isValidate {
				authService.On("Logout", tc.insertsrv).Return(tc.srvReturn)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/logout/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tc.param)
	
			authHandler := handlers.NewAuthHandler(authService)
			authHandler.Logout(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isValidate {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
				authService.AssertExpectations(t)
			}
		})
	}
}