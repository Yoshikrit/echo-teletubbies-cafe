package handlers_test

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"strings"
	"encoding/json"
	"github.com/stretchr/testify/assert"
    "github.com/labstack/echo/v4"

	"user/models"
	"user/tests/mocks/mock_services"
	"user/handlers"
    "user/utils/errs"
)

func TestCreateRole(t *testing.T) {
	e := echo.New()

	roleReqMock := &models.RoleCreate{
		Id:   1,
		Name: "A",
	}

	roleReqErrorMock := &models.RoleCreate{
		Id:   0,
		Name: "",
	}

	roleResMock := &models.Role{
		Id:   1,
		Name: "A",
	}

	roleReqJSON, _ := json.Marshal(roleReqMock)
	roleReqErrorJSON, _ := json.Marshal(roleReqErrorMock)
	roleResJSON, _ := json.Marshal(roleResMock)

	type testCase struct {
		name         	string
		isValidate      bool
		body            string
		insertSrv      	models.RoleCreate
		expectedStatus  int
		expectedBody    string
		srvReturn1 		models.Role
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    			isValidate: false,  body: string(roleReqJSON), 	  	insertSrv: *roleReqMock,   	 	expectedStatus: 201,	expectedBody: string(roleResJSON), 			srvReturn1: *roleResMock,     srvReturn2: nil},				
		{name: "test case : failed bind",  		isValidate: true, 	body: "invalid json", 			insertSrv: *roleReqErrorMock,  	expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,    	srvReturn1: models.Role{}, srvReturn2: nil},   		
		{name: "test case : failed validator",  isValidate: true, 	body: string(roleReqErrorJSON),	insertSrv: *roleReqErrorMock,  	expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,    	srvReturn1: models.Role{}, srvReturn2: nil},     
		{name: "test case : failed srvsitory", 	isValidate: false,  body: string(roleReqJSON), 	  	insertSrv: *roleReqMock,		expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	  	srvReturn1: models.Role{}, srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			roleService := mock_services.NewRoleServiceMock()
			if !tc.isValidate {
				roleService.On("CreateRole", &tc.insertSrv).Return(&tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/role/", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
	
			roleHandler := handlers.NewRoleHandler(roleService)
			roleHandler.CreateRole(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isValidate {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
				roleService.AssertExpectations(t)
			}
		})
	}
}

func TestGetAllRoles(t *testing.T) {
	e := echo.New()

	rolesResMock := []models.Role {
		{
			Id:   1,
			Name: "A",
		},
		{
			Id:   2,
			Name: "B",
		},
	}

	rolesResJSON, _ := json.Marshal(rolesResMock)

	type testCase struct {
		name         	string
		isFail          bool
		body            string
		expectedStatus  int
		expectedBody    string
		srvReturn1 		[]models.Role
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    			isFail: false, 	expectedStatus: 200,	expectedBody: string(rolesResJSON), 		srvReturn1: rolesResMock,  	 srvReturn2: nil},				
		{name: "test case : failed service", 	isFail: true, 	expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	srvReturn1: []models.Role{},  srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			roleService := mock_services.NewRoleServiceMock()
			roleService.On("GetRoles").Return(tc.srvReturn1, tc.srvReturn2)
	
			req := httptest.NewRequest(http.MethodPost, "/role/", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
	
			roleHandler := handlers.NewRoleHandler(roleService)
			roleHandler.GetAllRoles(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isFail {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
			}
			roleService.AssertExpectations(t)
		})
	}
}

func TestGetRoleByID(t *testing.T) {
	e := echo.New()

	roleResMock := models.Role {
		Id:   1,
		Name: "A",
	}
	
	roleResJSON, _ := json.Marshal(roleResMock)

	type testCase struct {
		name         	string
		isValidate      bool
		param           string
		insertSrv      	int
		expectedStatus  int
		expectedBody    string
		srvReturn1 		models.Role
		srvReturn2 		error
	}

	cases := []testCase{
		{name: "test case : pass",    			isValidate: false,  param: "1", insertSrv: 1,   	expectedStatus: 200,	expectedBody: string(roleResJSON), 			srvReturn1: roleResMock,  	 	srvReturn2: nil},					
		{name: "test case : failed param int",  isValidate: true, 	param: "a", insertSrv: 0,  		expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,  srvReturn1: models.Role{},  srvReturn2: errs.NewBadRequestError("")},     
		{name: "test case : failed service", 	isValidate: false,  param: "1", insertSrv: 1,		expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	srvReturn1: models.Role{},  srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			roleService := mock_services.NewRoleServiceMock()
			if !tc.isValidate {
				roleService.On("GetRole", tc.insertSrv).Return(&tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/role/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tc.param)
	
			roleHandler := handlers.NewRoleHandler(roleService)
			roleHandler.GetRoleByID(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isValidate {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
				roleService.AssertExpectations(t)
			}
		})
	}
}

func TestUpdateRoleByID(t *testing.T) {
	e := echo.New()

	roleReqMock := &models.RoleUpdate {
		Name: "B",
	}

	roleErrorReqMock := &models.RoleUpdate {
		Name: "",
	}

	roleResMock := &models.Role {
		Id:   1,
		Name: "B",
	}
	
	roleReqJSON, _ := json.Marshal(roleReqMock)
	roleErrorReqJSON, _ := json.Marshal(roleErrorReqMock)
	roleResJSON, _ := json.Marshal(roleResMock)

	type testCase struct {
		name         	string
		isValidate      bool
		param           string
		body            string
		insertId        int
		insertSrv      	models.RoleUpdate
		expectedStatus  int
		expectedBody    string
		srvReturn1 		models.Role
		srvReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    			param: "1", isValidate: false,  body: string(roleReqJSON), 	  		insertId: 1, insertSrv: *roleReqMock,   	 	expectedStatus: 200,	expectedBody: string(roleResJSON), 				srvReturn1: *roleResMock,  srvReturn2: nil},				
		{name: "test case : failed param int",  param: "a", isValidate: true, 	body: string(roleReqJSON),			insertId: 1, insertSrv: *roleReqMock,  			expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,  	srvReturn1: models.Role{}, srvReturn2: nil},     
		{name: "test case : failed bind",  		param: "1", isValidate: true, 	body: "invalid json", 				insertId: 1, insertSrv: models.RoleUpdate{}, 	expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,    	srvReturn1: models.Role{}, srvReturn2: nil},   
		{name: "test case : failed validator",  param: "1", isValidate: true, 	body: string(roleErrorReqJSON), 	insertId: 1, insertSrv: models.RoleUpdate{}, 	expectedStatus: 400,	expectedBody: `{"code":400,"message":""}`,    	srvReturn1: models.Role{}, srvReturn2: nil},   			
		{name: "test case : failed service", 	param: "1", isValidate: false,  body: string(roleReqJSON), 	  		insertId: 1, insertSrv: *roleReqMock,			expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	  	srvReturn1: models.Role{}, srvReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			roleService := mock_services.NewRoleServiceMock()
			if !tc.isValidate {
				roleService.On("UpdateRole", tc.insertId, &tc.insertSrv).Return(&tc.srvReturn1, tc.srvReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPut, "/role/", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tc.param)
	
			roleHandler := handlers.NewRoleHandler(roleService)
			roleHandler.UpdateRoleByID(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isValidate {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
				roleService.AssertExpectations(t)
			}
		})
	}
}


func TestDeleteRoleByID(t *testing.T) {
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
			roleService := mock_services.NewRoleServiceMock()
			if !tc.isValidate {
				roleService.On("DeleteRole", tc.insertId).Return(tc.srvReturn)
			}
	
			req := httptest.NewRequest(http.MethodDelete, "/role/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("id")
			c.SetParamValues(tc.param)
	
			roleHandler := handlers.NewRoleHandler(roleService)
			roleHandler.DeleteRoleByID(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isValidate {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
				roleService.AssertExpectations(t)
			}
		})
	}
}

func TestGetRoleCount(t *testing.T) {
    e := echo.New()
    
	t.Run("test case : pass", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/count", nil)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		roleService := mock_services.NewRoleServiceMock()
		roleService.On("GetRoleCount").Return(int64(42), nil)

		roleHandler := handlers.NewRoleHandler(roleService)
		roleHandler.GetRoleCount(c)

		expectedCode := 200
		expectedBody := `42`
		assert.Equal(t, expectedCode, rec.Code)
		assert.Equal(t, expectedBody, strings.TrimSpace(rec.Body.String()))
		roleService.AssertExpectations(t)
	})

	t.Run("test case : pass", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/count", nil)

		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		roleService := mock_services.NewRoleServiceMock()
		roleService.On("GetRoleCount").Return(int64(0), errs.NewUnexpectedError(""))

		roleHandler := handlers.NewRoleHandler(roleService)
		roleHandler.GetRoleCount(c)

		expectedCode := 500
		expectedBody := `{"code":500,"message":""}`
		assert.Equal(t, expectedCode, rec.Code)
		assert.Equal(t, expectedBody, strings.TrimSpace(rec.Body.String()))
		roleService.AssertExpectations(t)
	})
}