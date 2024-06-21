package services_test

import (
	"saleorder/tests/mocks/mock_repositories"
	"saleorder/services"
	"saleorder/utils/errs"
	"saleorder/models"

	"errors"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestCreateSaleOrder(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	saleOrderReq := &models.SaleOrderCreate{
		CreatedUser: 1,
		TotalPrice: 200,
		Status: "Pass",
		PayMethod: 1,
	}

	saleOrderEntityRes := &models.SaleOrderEntity{
		Id: 1,
		CreatedUser: 1,
		CreatedDate: mockDate,
		TotalPrice: 200,
		Status: "Pass",
		PayMethod: 1,
	}

	saleOrderMockRes := &models.SaleOrder{
		Id: 1,
		CreatedUser: 1,
		CreatedDate: mockDate,
		TotalPrice: 200,
		Status: "Pass",
		PayMethod: 1,
	}

	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderRepositoryMock()
		mockRepo.On("Create", saleOrderReq).Return(saleOrderEntityRes, nil)

		service := services.NewSaleOrderService(mockRepo)
		saleOrderRes, err := service.CreateSaleOrder(saleOrderReq)

		expected := saleOrderRes
		assert.NoError(t, err)
		assert.Equal(t, expected, saleOrderMockRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	type testCase struct {
		test_name       string
		isNull          bool
		createdUser     int
		status          string
		err 			error
	}
	cases := []testCase{
		{test_name: "test case : fail no user id",  isNull: true, 	createdUser: 0, status: "Pass", 	err: errs.NewBadRequestError("SaleOrder's User Id is null")},
		{test_name: "test case : fail no status",   isNull: true, 	createdUser: 1, status: "",       	err: errs.NewBadRequestError("SaleOrder's Status is null")},
		{test_name: "test case : fail repo",   		isNull: false, 	createdUser: 1, status: "Pass",     err: errors.New("")},
	}

	for _, tc := range cases {
		saleOrderReqFail := &models.SaleOrderCreate{
			CreatedUser: tc.createdUser,
			TotalPrice: 200,
			Status: tc.status,
			PayMethod: 1,
		}

		t.Run(tc.test_name, func(t *testing.T) {
			mockRepo := mock_repositories.NewSaleOrderRepositoryMock()
			if !tc.isNull {	
				mockRepo.On("Create", saleOrderReqFail).Return(&models.SaleOrderEntity{}, errors.New(""))
			}
			service := services.NewSaleOrderService(mockRepo)

			saleOrderRes, err := service.CreateSaleOrder(saleOrderReqFail)

			expected := tc.err
			assert.Error(t, err)
			assert.Equal(t, expected, err)
			assert.Nil(t, saleOrderRes)
			if !tc.isNull {	
				mockRepo.AssertExpectations(t)
			}
		})
	}
}


func TestGetSaleOrders(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	saleOrderReportRes := []models.SaleOrderReport{
		{
			Seq:            1,
			Id: 		    1,
			User:           "Gordon Freeman",
			Date:           mockDate,
			TotalPrice: 	450,
			Status: 	    "Pass",
			PayMethodName:  "None",
		},
	}
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderRepositoryMock()
		mockRepo.On("GetAll").Return(saleOrderReportRes, nil)

		service := services.NewSaleOrderService(mockRepo)
		saleOrderRes, err := service.GetSaleOrders()

		expected := saleOrderReportRes
		assert.NoError(t, err)
		assert.Equal(t, expected, saleOrderRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : fail", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderRepositoryMock()
		mockRepo.On("GetAll").Return([]models.SaleOrderReport{}, errors.New(""))

		service := services.NewSaleOrderService(mockRepo)
		saleOrderRes, err := service.GetSaleOrders()

		expected := errors.New("")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, saleOrderRes)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetSaleOrdersByDay(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	saleOrderReportRes := []models.SaleOrderReport{
		{
			Seq:            1,
			Id: 		    1,
			User:           "Gordon Freeman",
			Date:           mockDate,
			TotalPrice: 	450,
			Status: 	    "Pass",
			PayMethodName:  "None",
		},
	}
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderRepositoryMock()
		mockRepo.On("GetAllByDay", mockDate).Return(saleOrderReportRes, nil)

		service := services.NewSaleOrderService(mockRepo)
		saleOrderRes, err := service.GetSaleOrdersByDay(mockDate)

		expected := saleOrderReportRes
		assert.NoError(t, err)
		assert.Equal(t, expected, saleOrderRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : fail", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderRepositoryMock()
		mockRepo.On("GetAllByDay", mockDate).Return([]models.SaleOrderReport{}, errors.New(""))

		service := services.NewSaleOrderService(mockRepo)
		saleOrderRes, err := service.GetSaleOrdersByDay(mockDate)

		expected := errors.New("")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, saleOrderRes)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetSaleOrdersByMonth(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	saleOrderReportRes := []models.SaleOrderReport{
		{
			Seq:            1,
			Id: 		    1,
			User:           "Gordon Freeman",
			Date:           mockDate,
			TotalPrice: 	450,
			Status: 	    "Pass",
			PayMethodName:  "None",
		},
	}
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderRepositoryMock()
		mockRepo.On("GetAllByMonth", mockDate).Return(saleOrderReportRes, nil)

		service := services.NewSaleOrderService(mockRepo)
		saleOrderRes, err := service.GetSaleOrdersByMonth(mockDate)

		expected := saleOrderReportRes
		assert.NoError(t, err)
		assert.Equal(t, expected, saleOrderRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : fail", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderRepositoryMock()
		mockRepo.On("GetAllByMonth", mockDate).Return([]models.SaleOrderReport{}, errors.New(""))

		service := services.NewSaleOrderService(mockRepo)
		saleOrderRes, err := service.GetSaleOrdersByMonth(mockDate)

		expected := errors.New("")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, saleOrderRes)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetSaleOrdersByYear(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	saleOrderReportRes := []models.SaleOrderReport{
		{
			Seq:            1,
			Id: 		    1,
			User:           "Gordon Freeman",
			Date:           mockDate,
			TotalPrice: 	450,
			Status: 	    "Pass",
			PayMethodName:  "None",
		},
	}
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderRepositoryMock()
		mockRepo.On("GetAllByYear", mockDate).Return(saleOrderReportRes, nil)

		service := services.NewSaleOrderService(mockRepo)
		saleOrderRes, err := service.GetSaleOrdersByYear(mockDate)

		expected := saleOrderReportRes
		assert.NoError(t, err)
		assert.Equal(t, expected, saleOrderRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : fail", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderRepositoryMock()
		mockRepo.On("GetAllByYear", mockDate).Return([]models.SaleOrderReport{}, errors.New(""))

		service := services.NewSaleOrderService(mockRepo)
		saleOrderRes, err := service.GetSaleOrdersByYear(mockDate)

		expected := errors.New("")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, saleOrderRes)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetSaleOrderPriceAmountPass(t *testing.T) {
	price := 200.0
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderRepositoryMock()
		mockRepo.On("GetTotalPricePass").Return(price, nil)

		service := services.NewSaleOrderService(mockRepo)
		totalAmountRes, err := service.GetSaleOrderPriceAmountPass()

		expected := price
		assert.NoError(t, err)
		assert.Equal(t, expected, totalAmountRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : fail", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderRepositoryMock()
		mockRepo.On("GetTotalPricePass").Return(float64(0.0), errors.New(""))

		service := services.NewSaleOrderService(mockRepo)
		totalAmountRes, err := service.GetSaleOrderPriceAmountPass()

		expected := errors.New("")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Equal(t, float64(0.0), totalAmountRes)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetSaleOrderPriceAmountPassByDay(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	price := 200.0
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderRepositoryMock()
		mockRepo.On("GetTotalPricePassByDay", mockDate).Return(price, nil)

		service := services.NewSaleOrderService(mockRepo)
		totalAmountRes, err := service.GetSaleOrderPriceAmountPassByDay(mockDate)

		expected := price
		assert.NoError(t, err)
		assert.Equal(t, expected, totalAmountRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : fail", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderRepositoryMock()
		mockRepo.On("GetTotalPricePassByDay", mockDate).Return(float64(0.0), errors.New(""))

		service := services.NewSaleOrderService(mockRepo)
		totalAmountRes, err := service.GetSaleOrderPriceAmountPassByDay(mockDate)

		expected := errors.New("")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Equal(t, float64(0.0), totalAmountRes)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetSaleOrderPriceAmountPassByMonth(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	price := 200.0
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderRepositoryMock()
		mockRepo.On("GetTotalPricePassByMonth", mockDate).Return(price, nil)

		service := services.NewSaleOrderService(mockRepo)
		totalAmountRes, err := service.GetSaleOrderPriceAmountPassByMonth(mockDate)

		expected := price
		assert.NoError(t, err)
		assert.Equal(t, expected, totalAmountRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : fail", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderRepositoryMock()
		mockRepo.On("GetTotalPricePassByMonth", mockDate).Return(float64(0.0), errors.New(""))

		service := services.NewSaleOrderService(mockRepo)
		totalAmountRes, err := service.GetSaleOrderPriceAmountPassByMonth(mockDate)

		expected := errors.New("")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Equal(t, float64(0.0), totalAmountRes)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetSaleOrderPriceAmountPassByYear(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	price := 200.0
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderRepositoryMock()
		mockRepo.On("GetTotalPricePassByYear", mockDate).Return(price, nil)

		service := services.NewSaleOrderService(mockRepo)
		totalAmountRes, err := service.GetSaleOrderPriceAmountPassByYear(mockDate)

		expected := price
		assert.NoError(t, err)
		assert.Equal(t, expected, totalAmountRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : fail", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderRepositoryMock()
		mockRepo.On("GetTotalPricePassByYear", mockDate).Return(float64(0.0), errors.New(""))

		service := services.NewSaleOrderService(mockRepo)
		totalAmountRes, err := service.GetSaleOrderPriceAmountPassByYear(mockDate)

		expected := errors.New("")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Equal(t, float64(0.0), totalAmountRes)
		mockRepo.AssertExpectations(t)
	})
}