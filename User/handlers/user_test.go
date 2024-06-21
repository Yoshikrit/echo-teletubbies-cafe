package handlers_test

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
	"encoding/json"
	"github.com/stretchr/testify/assert"
    "github.com/labstack/echo/v4"
	"time"

    "user/utils/errs"
	"user/models"
	"user/tests/mocks/mock_services"
	"user/handlers"
)

func TestCreateUser(t *testing.T) {
	e := echo.New()

	mockDate := time.Date(2011, time.December, 28, 12, 30, 0, 0, time.UTC)
	userReqMock := &models.UserCreate{
		Id:   1,
		Role_Id:   1,
		FName: "A",
		LName: "A",
		Email: "aaa@example.com",
		Password: "A",
		Sex: "M",
		TelNo: "xxxxxxxxxx",
		Salary: 100,
		Address: "A",
		WorkStatus: "A",
		BirthDate: mockDate,
	}

	userReqErrorMock := &models.UserCreate{
		Id:   0,
		Role_Id:   0,
		FName: "",
		LName: "",
		Email: "",
		Password: "",
		Sex: "",
		TelNo: "",
		Salary: 0,
		Address: "",
		WorkStatus: "",
		BirthDate: time.Time{},
	}

	userResMock := &models.User{
		Id:   1,
		Role_Id:   1,
		FName: "A",
		LName: "A",
		Email: "A",
		Sex: "M",
		TelNo: "aaa@example.com",
		Salary: 100,
		Address: "A",
		WorkStatus: "A",
		BirthDate: mockDate,
	}

	userReqJSON, _ := json.Marshal(userReqMock)
	userReqErrorJSON, _ := json.Marshal(userReqErrorMock)
	userResJSON, _ := json.Marshal(userResMock)

	type testCase struct {
		name         	string
		isValidate      bool
		body            string
		insertSrv      	models.UserCreate
		expectedStatus  int
		expectedBody    string
		srvReturn1 		models.User
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    			isValidate: false,  body: string(userReqJSON), 	  	insertSrv: *userReqMock,   	 	expectedStatus: 201,	expectedBody: string(userResJSON), 				srvReturn1: *userResMock,     	srvReturn2: nil},				
		{name: "test case : failed bind",  		isValidate: true, 	body: "invalid json", 			insertSrv: *userReqErrorMock,  	expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,    	srvReturn1: models.User{}, 		srvReturn2: nil},   		
		{name: "test case : failed validator",  isValidate: true, 	body: string(userReqErrorJSON),	insertSrv: *userReqErrorMock,  	expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,    	srvReturn1: models.User{}, 		srvReturn2: nil},     
		{name: "test case : failed srvsitory", 	isValidate: false,  body: string(userReqJSON), 	  	insertSrv: *userReqMock,		expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	  	srvReturn1: models.User{}, 		srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			userService := mock_services.NewUserServiceMock()
			if !tc.isValidate {
				userService.On("CreateUser", &tc.insertSrv).Return(&tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/user/", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
	
			userHandler := handlers.NewUserHandler(userService)
			userHandler.CreateUser(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isValidate {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
				userService.AssertExpectations(t)
			}
		})
	}
}

func TestGetAllUsers(t *testing.T) {
	e := echo.New()

	mockDate := time.Date(2011, time.December, 28, 12, 30, 0, 0, time.UTC)
	usersResMock := []models.User {
		{
			Id:   1,
			Role_Id:   1,
			FName: "A",
			LName: "A",
			Email: "aaa@example.com",
			Sex: "M",
			TelNo: "xxxxxxxxxx",
			Salary: 100,
			Address: "A",
			WorkStatus: "A",
			BirthDate: mockDate,
		},
		{
			Id:   2,
			Role_Id:   1,
			FName: "A",
			LName: "A",
			Email: "aaa@example.com",
			Sex: "M",
			TelNo: "xxxxxxxxxx",
			Salary: 100,
			Address: "A",
			WorkStatus: "A",
			BirthDate: mockDate,
		},
	}

	usersResJSON, _ := json.Marshal(usersResMock)

	type testCase struct {
		name         	string
		isFail          bool
		body            string
		expectedStatus  int
		expectedBody    string
		srvReturn1 		[]models.User
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    			isFail: false, 	expectedStatus: 200,	expectedBody: string(usersResJSON), 		srvReturn1: usersResMock,  	  srvReturn2: nil},				
		{name: "test case : failed service", 	isFail: true, 	expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	srvReturn1: []models.User{},  srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			userService := mock_services.NewUserServiceMock()
			userService.On("GetUsers").Return(tc.srvReturn1, tc.srvReturn2)
	
			req := httptest.NewRequest(http.MethodPost, "/user/", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
	
			userHandler := handlers.NewUserHandler(userService)
			userHandler.GetAllUsers(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isFail {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
			}
			userService.AssertExpectations(t)
		})
	}
}

func TestGetUserByID(t *testing.T) {
	e := echo.New()

	mockDate := time.Date(2011, time.December, 28, 12, 30, 0, 0, time.UTC)
	userResMock := models.User {
		Id:   1,
		Role_Id:   1,
		FName: "A",
		LName: "A",
		Email: "aaa@example.com",
		Sex: "M",
		TelNo: "xxxxxxxxxx",
		Salary: 100,
		Address: "A",
		WorkStatus: "A",
		BirthDate: mockDate,
	}
	
	userResJSON, _ := json.Marshal(userResMock)

	type testCase struct {
		name         	string
		isValidate      bool
		param           string
		insertSrv      	int
		expectedStatus  int
		expectedBody    string
		srvReturn1 		models.User
		srvReturn2 		error
	}

	cases := []testCase{
		{name: "test case : pass",    			isValidate: false,  param: "1", insertSrv: 1,   	expectedStatus: 200,	expectedBody: string(userResJSON), 			srvReturn1: userResMock,  	 	srvReturn2: nil},					
		{name: "test case : failed param int",  isValidate: true, 	param: "a", insertSrv: 0,  		expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,  srvReturn1: models.User{},  	srvReturn2: errs.NewBadRequestError("")},     
		{name: "test case : failed service", 	isValidate: false,  param: "1", insertSrv: 1,		expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	srvReturn1: models.User{},  	srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			userService := mock_services.NewUserServiceMock()
			if !tc.isValidate {
				userService.On("GetUser", tc.insertSrv).Return(&tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/user/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tc.param)
	
			userHandler := handlers.NewUserHandler(userService)
			userHandler.GetUserByID(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isValidate {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
				userService.AssertExpectations(t)
			}
		})
	}
}

func TestUpdateUserByID(t *testing.T) {
	e := echo.New()

	mockDate := time.Date(2011, time.December, 28, 12, 30, 0, 0, time.UTC)
	userReqMock := &models.UserUpdate {
		Role_Id:   1,
		FName: "A",
		LName: "A",
		Email: "aaa@example.com",
		Password: "A",
		Sex: "M",
		TelNo: "xxxxxxxxxx",
		Salary: 100,
		Address: "A",
		WorkStatus: "A",
		BirthDate: mockDate,
	}

	userResMock := &models.User {
		Id:   1,
		Role_Id:   1,
		FName: "A",
		LName: "A",
		Email: "A",
		Sex: "M",
		TelNo: "xxxxxxxxxx",
		Salary: 100,
		Address: "A",
		WorkStatus: "A",
		BirthDate: mockDate,
	}
	
	userReqJSON, _ := json.Marshal(userReqMock)
	userResJSON, _ := json.Marshal(userResMock)

	type testCase struct {
		name         	string
		isValidate      bool
		param           string
		body            string
		insertId        int
		insertSrv      	models.UserUpdate
		expectedStatus  int
		expectedBody    string
		srvReturn1 		models.User
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    			param: "1", isValidate: false,  body: string(userReqJSON), 	  	insertId: 1, insertSrv: *userReqMock,   	 	expectedStatus: 200,	expectedBody: string(userResJSON), 				srvReturn1: *userResMock,     srvReturn2: nil},				
		{name: "test case : failed param int",  param: "a", isValidate: true, 	body: string(userReqJSON),		insertId: 1, insertSrv: *userReqMock,  			expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,  	srvReturn1: models.User{}, srvReturn2: nil},     
		{name: "test case : failed bind",  		param: "1", isValidate: true, 	body: "invalid json", 			insertId: 1, insertSrv: models.UserUpdate{}, expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,    	srvReturn1: models.User{}, srvReturn2: nil},   		
		{name: "test case : failed service", 	param: "1", isValidate: false,  body: string(userReqJSON), 	  	insertId: 1, insertSrv: *userReqMock,			expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	  	srvReturn1: models.User{}, srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			userService := mock_services.NewUserServiceMock()
			if !tc.isValidate {
				userService.On("UpdateUser", tc.insertId, &tc.insertSrv).Return(&tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPut, "/user/", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tc.param)
	
			userHandler := handlers.NewUserHandler(userService)
			userHandler.UpdateUserByID(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isValidate {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
				userService.AssertExpectations(t)
			}
		})
	}
}


func TestDeleteUserByID(t *testing.T) {
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
			userService := mock_services.NewUserServiceMock()
			if !tc.isValidate {
				userService.On("DeleteUser", tc.insertId).Return(tc.srvReturn)
			}
	
			req := httptest.NewRequest(http.MethodDelete, "/user/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tc.param)
	
			userHandler := handlers.NewUserHandler(userService)
			userHandler.DeleteUserByID(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isValidate {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
				userService.AssertExpectations(t)
			}
		})
	}
}

func TestGetUserCount(t *testing.T) {
    e := echo.New()
    
	t.Run("test case : pass", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/count", nil)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		userService := mock_services.NewUserServiceMock()
		userService.On("GetUserCount").Return(int64(42), nil)

		userHandler := handlers.NewUserHandler(userService)
		userHandler.GetUserCount(c)

		expectedCode := 200
		expectedBody := `42`
		assert.Equal(t, expectedCode, rec.Code)
		assert.Equal(t, expectedBody, strings.TrimSpace(rec.Body.String()))
		userService.AssertExpectations(t)
	})

	t.Run("test case : pass", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/count", nil)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		userService := mock_services.NewUserServiceMock()
		userService.On("GetUserCount").Return(int64(0), errs.NewUnexpectedError(""))

		userHandler := handlers.NewUserHandler(userService)
		userHandler.GetUserCount(c)

		expectedCode := 500
		expectedBody := `{"code":500,"message":""}`
		assert.Equal(t, expectedCode, rec.Code)
		assert.Equal(t, expectedBody, strings.TrimSpace(rec.Body.String()))
		userService.AssertExpectations(t)
	})
}
