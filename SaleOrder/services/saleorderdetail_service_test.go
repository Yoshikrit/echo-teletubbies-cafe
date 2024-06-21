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

func TestCreateSaleOrderDetail(t *testing.T) {
	saleOrderDetailReq := &models.SaleOrderDetailCreate{
		SO_Id: 1,
		Prod_Id: 1,
		Quantity: 4,
		Price: 100.0,
		Discount: 10.0,
	}

	saleOrderDetailEntityRes := &models.SaleOrderDetailEntity{
		Seq: 1,
		SO_Id: 1,
		Prod_Id: 1,
		Quantity: 4,
		Price: 100.0,
		Discount: 10.0,
	}

	saleOrderDetailMockRes := &models.SaleOrderDetail{
		Seq: 1,
		SO_Id: 1,
		Prod_Id: 1,
		Quantity: 4,
		Price: 100.0,
		Discount: 10.0,
	}

	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderDetailRepositoryMock()
		mockRepo.On("Create", saleOrderDetailReq).Return(saleOrderDetailEntityRes, nil)

		service := services.NewSaleOrderDetailService(mockRepo)
		saleOrderDetailRes, err := service.CreateSaleOrderDetail(saleOrderDetailReq)

		expected := saleOrderDetailRes
		assert.NoError(t, err)
		assert.Equal(t, expected, saleOrderDetailMockRes)
		assert.Nil(t, err)
	})

	type testCase struct {
		test_name       string
		isNull          bool
		so_id    		int
		prod_id    		int
		qty          	int
		price          	int
		err 			error
	}
	cases := []testCase{
		{test_name: "test case : fail no saleorder id", isNull: true, 	so_id: 0, prod_id: 1, qty: 1, price: 1, 	err: errs.NewBadRequestError("SaleOrderDetail's SaleOrder Id is null")},
		{test_name: "test case : fail no product id",   isNull: true, 	so_id: 1, prod_id: 0, qty: 1, price: 1,     err: errs.NewBadRequestError("SaleOrderDetail's User Id is null")},
		{test_name: "test case : fail no qty",   		isNull: true, 	so_id: 1, prod_id: 1, qty: 0, price: 1,     err: errs.NewBadRequestError("SaleOrderDetail's Quantity is null")},
		{test_name: "test case : fail no price",   		isNull: true, 	so_id: 1, prod_id: 1, qty: 1, price: 0,     err: errs.NewBadRequestError("SaleOrderDetail's Price is null")},
		{test_name: "test case : fail repo",   			isNull: false, 	so_id: 1, prod_id: 1, qty: 1, price: 1,     err: errors.New("")},
	}

	for _, tc := range cases {
		saleOrderDetailReqFail := &models.SaleOrderDetailCreate{
			SO_Id: tc.so_id,
			Prod_Id: tc.prod_id,
			Quantity: tc.qty,
			Price: float64(tc.price),
			Discount: 0.0,
		}

		t.Run(tc.test_name, func(t *testing.T) {
			mockRepo := mock_repositories.NewSaleOrderDetailRepositoryMock()
			if !tc.isNull {
				mockRepo.On("Create", saleOrderDetailReqFail).Return(&models.SaleOrderDetailEntity{}, errors.New(""))
			}
			service := services.NewSaleOrderDetailService(mockRepo)

			saleOrderDetailRes, err := service.CreateSaleOrderDetail(saleOrderDetailReqFail)

			expected := tc.err
			assert.Error(t, err)
			assert.Equal(t, expected, err)
			assert.Nil(t, saleOrderDetailRes)
			if !tc.isNull {
				mockRepo.AssertExpectations(t)
			}
		})
	}
}

func TestGetSaleOrderDetailQtyRates(t *testing.T) {
	saleOrderDetailReportRes := []models.SaleOrderDetailRate{
		{
			Prod_Id:            1,
			Quantity: 		    1,
			Price:              1,
		},
	}
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderDetailRepositoryMock()
		mockRepo.On("GetAllQtyRates").Return(saleOrderDetailReportRes, nil)

		service := services.NewSaleOrderDetailService(mockRepo)
		saleOrderDetailRes, err := service.GetSaleOrderDetailQtyRates()

		expected := saleOrderDetailReportRes
		assert.NoError(t, err)
		assert.Equal(t, expected, saleOrderDetailRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : fail", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderDetailRepositoryMock()
		mockRepo.On("GetAllQtyRates").Return([]models.SaleOrderDetailRate{}, errors.New(""))

		service := services.NewSaleOrderDetailService(mockRepo)
		saleOrderDetailRes, err := service.GetSaleOrderDetailQtyRates()

		expected := errors.New("")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, saleOrderDetailRes)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetSaleOrderDetailQtyRatesByDay(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	saleOrderDetailReportRes := []models.SaleOrderDetailRate{
		{
			Prod_Id:            1,
			Quantity: 		    1,
			Price:              1,
		},
	}
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderDetailRepositoryMock()
		mockRepo.On("GetAllQtyRatesByDay", mockDate).Return(saleOrderDetailReportRes, nil)

		service := services.NewSaleOrderDetailService(mockRepo)
		saleOrderDetailRes, err := service.GetSaleOrderDetailQtyRatesByDay(mockDate)

		expected := saleOrderDetailReportRes
		assert.NoError(t, err)
		assert.Equal(t, expected, saleOrderDetailRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : fail", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderDetailRepositoryMock()
		mockRepo.On("GetAllQtyRatesByDay", mockDate).Return([]models.SaleOrderDetailRate{}, errors.New(""))

		service := services.NewSaleOrderDetailService(mockRepo)
		saleOrderDetailRes, err := service.GetSaleOrderDetailQtyRatesByDay(mockDate)

		expected := errors.New("")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, saleOrderDetailRes)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetSaleOrderDetailQtyRatesByMonth(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	saleOrderDetailReportRes := []models.SaleOrderDetailRate{
		{
			Prod_Id:            1,
			Quantity: 		    1,
			Price:              1,
		},
	}
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderDetailRepositoryMock()
		mockRepo.On("GetAllQtyRatesByMonth", mockDate).Return(saleOrderDetailReportRes, nil)

		service := services.NewSaleOrderDetailService(mockRepo)
		saleOrderDetailRes, err := service.GetSaleOrderDetailQtyRatesByMonth(mockDate)

		expected := saleOrderDetailReportRes
		assert.NoError(t, err)
		assert.Equal(t, expected, saleOrderDetailRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : fail", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderDetailRepositoryMock()
		mockRepo.On("GetAllQtyRatesByMonth", mockDate).Return([]models.SaleOrderDetailRate{}, errors.New(""))

		service := services.NewSaleOrderDetailService(mockRepo)
		saleOrderDetailRes, err := service.GetSaleOrderDetailQtyRatesByMonth(mockDate)

		expected := errors.New("")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, saleOrderDetailRes)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetSaleOrderDetailQtyRatesByYear(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	saleOrderDetailReportRes := []models.SaleOrderDetailRate{
		{
			Prod_Id:            1,
			Quantity: 		    1,
			Price:              1,
		},
	}
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderDetailRepositoryMock()
		mockRepo.On("GetAllQtyRatesByYear", mockDate).Return(saleOrderDetailReportRes, nil)

		service := services.NewSaleOrderDetailService(mockRepo)
		saleOrderDetailRes, err := service.GetSaleOrderDetailQtyRatesByYear(mockDate)

		expected := saleOrderDetailReportRes
		assert.NoError(t, err)
		assert.Equal(t, expected, saleOrderDetailRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : fail", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderDetailRepositoryMock()
		mockRepo.On("GetAllQtyRatesByYear", mockDate).Return([]models.SaleOrderDetailRate{}, errors.New(""))

		service := services.NewSaleOrderDetailService(mockRepo)
		saleOrderDetailRes, err := service.GetSaleOrderDetailQtyRatesByYear(mockDate)

		expected := errors.New("")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, saleOrderDetailRes)
	})
}

func TestGetSaleOrderDetailPriceRates(t *testing.T) {
	saleOrderDetailReportRes := []models.SaleOrderDetailRate{
		{
			Prod_Id:            1,
			Quantity: 		    1,
			Price:              1,
		},
	}
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderDetailRepositoryMock()
		mockRepo.On("GetAllPriceRates").Return(saleOrderDetailReportRes, nil)

		service := services.NewSaleOrderDetailService(mockRepo)
		saleOrderDetailRes, err := service.GetSaleOrderDetailPriceRates()

		expected := saleOrderDetailReportRes
		assert.NoError(t, err)
		assert.Equal(t, expected, saleOrderDetailRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : fail", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderDetailRepositoryMock()
		mockRepo.On("GetAllPriceRates").Return([]models.SaleOrderDetailRate{}, errors.New(""))

		service := services.NewSaleOrderDetailService(mockRepo)
		saleOrderDetailRes, err := service.GetSaleOrderDetailPriceRates()

		expected := errors.New("")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, saleOrderDetailRes)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetSaleOrderDetailPriceRatesByDay(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	saleOrderDetailReportRes := []models.SaleOrderDetailRate{
		{
			Prod_Id:            1,
			Quantity: 		    1,
			Price:              1,
		},
	}
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderDetailRepositoryMock()
		mockRepo.On("GetAllPriceRatesByDay", mockDate).Return(saleOrderDetailReportRes, nil)

		service := services.NewSaleOrderDetailService(mockRepo)
		saleOrderDetailRes, err := service.GetSaleOrderDetailPriceRatesByDay(mockDate)

		expected := saleOrderDetailReportRes
		assert.NoError(t, err)
		assert.Equal(t, expected, saleOrderDetailRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : fail", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderDetailRepositoryMock()
		mockRepo.On("GetAllPriceRatesByDay", mockDate).Return([]models.SaleOrderDetailRate{}, errors.New(""))

		service := services.NewSaleOrderDetailService(mockRepo)
		saleOrderDetailRes, err := service.GetSaleOrderDetailPriceRatesByDay(mockDate)

		expected := errors.New("")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, saleOrderDetailRes)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetSaleOrderDetailPriceRatesByMonth(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	saleOrderDetailReportRes := []models.SaleOrderDetailRate{
		{
			Prod_Id:            1,
			Quantity: 		    1,
			Price:              1,
		},
	}
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderDetailRepositoryMock()
		mockRepo.On("GetAllPriceRatesByMonth", mockDate).Return(saleOrderDetailReportRes, nil)

		service := services.NewSaleOrderDetailService(mockRepo)
		saleOrderDetailRes, err := service.GetSaleOrderDetailPriceRatesByMonth(mockDate)

		expected := saleOrderDetailReportRes
		assert.NoError(t, err)
		assert.Equal(t, expected, saleOrderDetailRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : fail", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderDetailRepositoryMock()
		mockRepo.On("GetAllPriceRatesByMonth", mockDate).Return([]models.SaleOrderDetailRate{}, errors.New(""))

		service := services.NewSaleOrderDetailService(mockRepo)
		saleOrderDetailRes, err := service.GetSaleOrderDetailPriceRatesByMonth(mockDate)

		expected := errors.New("")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, saleOrderDetailRes)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetSaleOrderDetailPriceRatesByYear(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	saleOrderDetailReportRes := []models.SaleOrderDetailRate{
		{
			Prod_Id:            1,
			Quantity: 		    1,
			Price:              1,
		},
	}
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderDetailRepositoryMock()
		mockRepo.On("GetAllPriceRatesByYear", mockDate).Return(saleOrderDetailReportRes, nil)

		service := services.NewSaleOrderDetailService(mockRepo)
		saleOrderDetailRes, err := service.GetSaleOrderDetailPriceRatesByYear(mockDate)

		expected := saleOrderDetailReportRes
		assert.NoError(t, err)
		assert.Equal(t, expected, saleOrderDetailRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : fail", func(t *testing.T) {
		mockRepo := mock_repositories.NewSaleOrderDetailRepositoryMock()
		mockRepo.On("GetAllPriceRatesByYear", mockDate).Return([]models.SaleOrderDetailRate{}, errors.New(""))

		service := services.NewSaleOrderDetailService(mockRepo)
		saleOrderDetailRes, err := service.GetSaleOrderDetailPriceRatesByYear(mockDate)

		expected := errors.New("")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, saleOrderDetailRes)
		mockRepo.AssertExpectations(t)
	})
}