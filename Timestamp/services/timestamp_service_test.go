package services_test

import (
	"timestamp/tests/mocks/mock_repositories"
	"timestamp/services"
	"timestamp/models"

	"errors"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestGetTimestamps(t *testing.T) {
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

	type testCase struct {
		test_name       string
		isFail          bool
		expected1       []models.TimestampReport  
		expected2       error  
		returnRepo1     []models.TimestampReport
		returnRepo2 	error
	}
	cases := []testCase{
		{test_name: "test case : pass",   			isFail: false, 	expected1: timestampReportsRes,  expected2: nil,	  		returnRepo1: timestampReportsRes,  returnRepo2: nil},
		{test_name: "test case : fail repository",  isFail: true, 	expected1: nil,			   		 expected2: errors.New(""), returnRepo1: nil,  				   returnRepo2: errors.New("")},
	}

	for _, tc := range cases {
		t.Run(tc.test_name, func(t *testing.T) {
			mockRepo := mock_repositories.NewTimestampRepositoryMock()
			mockRepo.On("GetAll").Return(tc.returnRepo1, tc.returnRepo2)

			service := services.NewTimestampService(mockRepo)
			saleOrderReportRes, err := service.GetTimestamps()

			mockRepo.AssertExpectations(t)
			if tc.isFail {
				assert.Error(t, err)
				assert.Equal(t, tc.expected2, err)
				assert.Nil(t, saleOrderReportRes)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected1, saleOrderReportRes)
				assert.Nil(t, err)
			} 
		})
	}
}

func TestGetTimestampsByDay(t *testing.T) {
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

	type testCase struct {
		test_name       string
		isFail          bool
		expected1       []models.TimestampReport  
		expected2       error  
		returnRepo1     []models.TimestampReport
		returnRepo2 	error
	}
	cases := []testCase{
		{test_name: "test case : pass",   			isFail: false, 	expected1: timestampReportsRes,  expected2: nil,	  		returnRepo1: timestampReportsRes,  returnRepo2: nil},
		{test_name: "test case : fail repository",  isFail: true, 	expected1: nil,			   		 expected2: errors.New(""), returnRepo1: nil,  				   returnRepo2: errors.New("")},
	}

	for _, tc := range cases {
		t.Run(tc.test_name, func(t *testing.T) {
			mockRepo := mock_repositories.NewTimestampRepositoryMock()
			mockRepo.On("GetAllByDay", mockDate).Return(tc.returnRepo1, tc.returnRepo2)

			service := services.NewTimestampService(mockRepo)
			saleOrderReportRes, err := service.GetTimestampsByDay(mockDate)

			mockRepo.AssertExpectations(t)
			if tc.isFail {
				assert.Error(t, err)
				assert.Equal(t, tc.expected2, err)
				assert.Nil(t, saleOrderReportRes)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected1, saleOrderReportRes)
				assert.Nil(t, err)
			} 
		})
	}
}

func TestGetTimestampsByMonth(t *testing.T) {
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

	type testCase struct {
		test_name       string
		isFail          bool
		expected1       []models.TimestampReport  
		expected2       error  
		returnRepo1     []models.TimestampReport
		returnRepo2 	error
	}
	cases := []testCase{
		{test_name: "test case : pass",   			isFail: false, 	expected1: timestampReportsRes,  expected2: nil,	  		returnRepo1: timestampReportsRes,  returnRepo2: nil},
		{test_name: "test case : fail repository",  isFail: true, 	expected1: nil,			   		 expected2: errors.New(""), returnRepo1: nil,  				   returnRepo2: errors.New("")},
	}

	for _, tc := range cases {
		t.Run(tc.test_name, func(t *testing.T) {
			mockRepo := mock_repositories.NewTimestampRepositoryMock()
			mockRepo.On("GetAllByMonth", mockDate).Return(tc.returnRepo1, tc.returnRepo2)

			service := services.NewTimestampService(mockRepo)
			saleOrderReportRes, err := service.GetTimestampsByMonth(mockDate)

			mockRepo.AssertExpectations(t)
			if tc.isFail {
				assert.Error(t, err)
				assert.Equal(t, tc.expected2, err)
				assert.Nil(t, saleOrderReportRes)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected1, saleOrderReportRes)
				assert.Nil(t, err)
			} 
		})
	}
}

func TestGetTimestampsByYear(t *testing.T) {
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

	type testCase struct {
		test_name       string
		isFail          bool
		expected1       []models.TimestampReport  
		expected2       error  
		returnRepo1     []models.TimestampReport
		returnRepo2 	error
	}
	cases := []testCase{
		{test_name: "test case : pass",   			isFail: false, 	expected1: timestampReportsRes,  expected2: nil,	  		returnRepo1: timestampReportsRes,  returnRepo2: nil},
		{test_name: "test case : fail repository",  isFail: true, 	expected1: nil,			   		 expected2: errors.New(""), returnRepo1: nil,  				   returnRepo2: errors.New("")},
	}

	for _, tc := range cases {
		t.Run(tc.test_name, func(t *testing.T) {
			mockRepo := mock_repositories.NewTimestampRepositoryMock()
			mockRepo.On("GetAllByYear", mockDate).Return(tc.returnRepo1, tc.returnRepo2)

			service := services.NewTimestampService(mockRepo)
			saleOrderReportRes, err := service.GetTimestampsByYear(mockDate)

			mockRepo.AssertExpectations(t)
			if tc.isFail {
				assert.Error(t, err)
				assert.Equal(t, tc.expected2, err)
				assert.Nil(t, saleOrderReportRes)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected1, saleOrderReportRes)
				assert.Nil(t, err)
			} 
		})
	}
}