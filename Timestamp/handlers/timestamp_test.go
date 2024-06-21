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

	"timestamp/handlers"
	"timestamp/models"
	"timestamp/tests/mocks/mock_services"
	"timestamp/utils/errs"
)

func TestGetAllTimestamps(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	timestampReportsRes := []models.TimestampReport{
		{
			Seq:              1,
			UserName:         "Morgan Freeman",
			LoginAt:          mockDate,  
			LogoutAt:         mockDate,
			Hour:             0,  
		},      
	}

	e := echo.New()

	timestampReportsResJSON, _ := json.Marshal(timestampReportsRes)

	type testCase struct {
		name         	string
		isFail          bool
		body            string
		expectedStatus  int
		expectedBody    string
		repoReturn1 	[]models.TimestampReport
		repoReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    			isFail: false, 	expectedStatus: 200,	expectedBody: string(timestampReportsResJSON), 	repoReturn1: timestampReportsRes,      	 	repoReturn2: nil},				
		{name: "test case : failed repository", isFail: true, 	expectedStatus: 500,	expectedBody: `{"code":500,"message":""}`,	  	repoReturn1: []models.TimestampReport{},  	repoReturn2: errs.NewUnexpectedError("")},			
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			timeStampService := mock_services.NewTimestampServiceMock()
			timeStampService.On("GetTimestamps").Return(tc.repoReturn1, tc.repoReturn2)
	
			req := httptest.NewRequest(http.MethodPost, "/saleordertail/qtyrate", strings.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
	
			timeStampHandler := handlers.NewTimestampHandler(timeStampService)
			timeStampHandler.GetAllTimestamps(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isFail {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
			}
			timeStampService.AssertExpectations(t)
		})
	}
}

func TestGetAllTimestampsByDay(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 0, 0, 0, 0, time.UTC)
	timestampReportsRes := []models.TimestampReport{
		{
			Seq:              1,
			UserName:         "Morgan Freeman",
			LoginAt:          mockDate,  
			LogoutAt:         mockDate,
			Hour:             0,  
		},      
	}

	e := echo.New()

	timestampReportsResJSON, _ := json.Marshal(timestampReportsRes)

	type testCase struct {
		name         	string
		isFail          bool
		isValidate      bool
		paramName       string
		expectedStatus  int
		expectedBody    string
		repoReturn1 	[]models.TimestampReport
		repoReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    				isFail: false, isValidate: false, 	expectedStatus: 200, paramName: "2023-12-28",	expectedBody: string(timestampReportsResJSON), 	repoReturn1: timestampReportsRes,      	  repoReturn2: nil},				
		{name: "test case : failed validate date", 	isFail: true,  isValidate: true, 	expectedStatus: 400, paramName: "2",			expectedBody: `{"code":400,"message":""}`,	  	repoReturn1: []models.TimestampReport{},  repoReturn2: nil},			
		{name: "test case : failed repository", 	isFail: true,  isValidate: false, 	expectedStatus: 500, paramName: "2023-12-28",	expectedBody: `{"code":500,"message":""}`,	  	repoReturn1: []models.TimestampReport{},  repoReturn2: errs.NewUnexpectedError("")},			
	}


	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			timestampService := mock_services.NewTimestampServiceMock()
			if !tc.isValidate{
				timestampService.On("GetTimestampsByDay", mockDate).Return(tc.repoReturn1, tc.repoReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/timestamp/day/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("date")
			c.SetParamValues(tc.paramName)
	
			timestampHandler := handlers.NewTimestampHandler(timestampService)
			timestampHandler.GetAllTimestampsByDay(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isFail {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
			}
			timestampService.AssertExpectations(t)
		})
	}
}

func TestGetAllTimestampsByMonth(t *testing.T) {
	mockDate := time.Date(2023, time.December, 1, 0, 0, 0, 0, time.UTC)
	timestampReportsRes := []models.TimestampReport{
		{
			Seq:              1,
			UserName:         "Morgan Freeman",
			LoginAt:          mockDate,  
			LogoutAt:         mockDate,
			Hour:             0,  
		},      
	}

	e := echo.New()

	timestampReportsResJSON, _ := json.Marshal(timestampReportsRes)

	type testCase struct {
		name         	string
		isFail          bool
		isValidate      bool
		paramName       string
		expectedStatus  int
		expectedBody    string
		repoReturn1 	[]models.TimestampReport
		repoReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    				isFail: false, isValidate: false, 	expectedStatus: 200, paramName: "2023-12",	expectedBody: string(timestampReportsResJSON), 	repoReturn1: timestampReportsRes,      	  repoReturn2: nil},				
		{name: "test case : failed validate date", 	isFail: true,  isValidate: true, 	expectedStatus: 400, paramName: "2",		expectedBody: `{"code":400,"message":""}`,	  	repoReturn1: []models.TimestampReport{},  repoReturn2: nil},			
		{name: "test case : failed repository", 	isFail: true,  isValidate: false, 	expectedStatus: 500, paramName: "2023-12",	expectedBody: `{"code":500,"message":""}`,	  	repoReturn1: []models.TimestampReport{},  repoReturn2: errs.NewUnexpectedError("")},			
	}


	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			timestampService := mock_services.NewTimestampServiceMock()
			if !tc.isValidate{
				timestampService.On("GetTimestampsByMonth", mockDate).Return(tc.repoReturn1, tc.repoReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/timestamp/month/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("date")
			c.SetParamValues(tc.paramName)
	
			timestampHandler := handlers.NewTimestampHandler(timestampService)
			timestampHandler.GetAllTimestampsByMonth(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isFail {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
			}
			timestampService.AssertExpectations(t)
		})
	}
}

func TestGetAllTimestampsByYear(t *testing.T) {
	mockDate := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	timestampReportsRes := []models.TimestampReport{
		{
			Seq:              1,
			UserName:         "Morgan Freeman",
			LoginAt:          mockDate,  
			LogoutAt:         mockDate,
			Hour:             0,  
		},      
	}

	e := echo.New()

	timestampReportsResJSON, _ := json.Marshal(timestampReportsRes)

	type testCase struct {
		name         	string
		isFail          bool
		isValidate      bool
		paramName       string
		expectedStatus  int
		expectedBody    string
		repoReturn1 	[]models.TimestampReport
		repoReturn2   	error
	}

	cases := []testCase{
		{name: "test case : pass",    				isFail: false, isValidate: false, 	expectedStatus: 200, paramName: "2023",	expectedBody: string(timestampReportsResJSON), 	repoReturn1: timestampReportsRes,      	  repoReturn2: nil},				
		{name: "test case : failed validate date", 	isFail: true,  isValidate: true, 	expectedStatus: 400, paramName: "2",	expectedBody: `{"code":400,"message":""}`,	  	repoReturn1: []models.TimestampReport{},  repoReturn2: nil},			
		{name: "test case : failed repository", 	isFail: true,  isValidate: false, 	expectedStatus: 500, paramName: "2023",	expectedBody: `{"code":500,"message":""}`,	  	repoReturn1: []models.TimestampReport{},  repoReturn2: errs.NewUnexpectedError("")},			
	}


	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			timestampService := mock_services.NewTimestampServiceMock()
			if !tc.isValidate{
				timestampService.On("GetTimestampsByYear", mockDate).Return(tc.repoReturn1, tc.repoReturn2)
			}
	
			req := httptest.NewRequest(http.MethodPost, "/timestamp/year/", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	
			rec := httptest.NewRecorder()
	
			c := e.NewContext(req, rec)
			c.SetParamNames("date")
			c.SetParamValues(tc.paramName)
	
			timestampHandler := handlers.NewTimestampHandler(timestampService)
			timestampHandler.GetAllTimestampsByYear(c)
	
			assert.Equal(t, tc.expectedStatus, rec.Code)
			if !tc.isFail {
				assert.Equal(t, strings.TrimSpace(tc.expectedBody), strings.TrimSpace(rec.Body.String()))
			}
			timestampService.AssertExpectations(t)
		})
	}
}
