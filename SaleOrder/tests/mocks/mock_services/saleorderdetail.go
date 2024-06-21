package mock_services

import (
	"github.com/stretchr/testify/mock"
	"time"

	"saleorder/models"
)

type saleorderDetailServiceMock struct {
	mock.Mock
}

func NewSaleOrderDetailServiceMock() *saleorderDetailServiceMock {
	return &saleorderDetailServiceMock{}
}

func (m *saleorderDetailServiceMock) CreateSaleOrderDetail(saleorderDetailReq *models.SaleOrderDetailCreate) (*models.SaleOrderDetail, error) {
	args := m.Called(saleorderDetailReq)
	return args.Get(0).(*models.SaleOrderDetail), args.Error(1)
}

func (m *saleorderDetailServiceMock) GetSaleOrderDetailQtyRates() ([]models.SaleOrderDetailRate, error) {
	args := m.Called()
	return args.Get(0).([]models.SaleOrderDetailRate), args.Error(1)
}

func (m *saleorderDetailServiceMock) GetSaleOrderDetailQtyRatesByDay(dateReq time.Time) ([]models.SaleOrderDetailRate, error) {
	args := m.Called(dateReq)
	return args.Get(0).([]models.SaleOrderDetailRate), args.Error(1)
}

func (m *saleorderDetailServiceMock) GetSaleOrderDetailQtyRatesByMonth(dateReq time.Time) ([]models.SaleOrderDetailRate, error) {
	args := m.Called(dateReq)
	return args.Get(0).([]models.SaleOrderDetailRate), args.Error(1)
}

func (m *saleorderDetailServiceMock) GetSaleOrderDetailQtyRatesByYear(dateReq time.Time) ([]models.SaleOrderDetailRate, error) {
	args := m.Called(dateReq)
	return args.Get(0).([]models.SaleOrderDetailRate), args.Error(1)
}

func (m *saleorderDetailServiceMock) GetSaleOrderDetailPriceRates() ([]models.SaleOrderDetailRate, error) {
	args := m.Called()
	return args.Get(0).([]models.SaleOrderDetailRate), args.Error(1)
}

func (m *saleorderDetailServiceMock) GetSaleOrderDetailPriceRatesByDay(dateReq time.Time) ([]models.SaleOrderDetailRate, error) {
	args := m.Called(dateReq)
	return args.Get(0).([]models.SaleOrderDetailRate), args.Error(1)
}

func (m *saleorderDetailServiceMock) GetSaleOrderDetailPriceRatesByMonth(dateReq time.Time) ([]models.SaleOrderDetailRate, error) {
	args := m.Called(dateReq)
	return args.Get(0).([]models.SaleOrderDetailRate), args.Error(1)
}

func (m *saleorderDetailServiceMock) GetSaleOrderDetailPriceRatesByYear(dateReq time.Time) ([]models.SaleOrderDetailRate, error) {
	args := m.Called(dateReq)
	return args.Get(0).([]models.SaleOrderDetailRate), args.Error(1)
}