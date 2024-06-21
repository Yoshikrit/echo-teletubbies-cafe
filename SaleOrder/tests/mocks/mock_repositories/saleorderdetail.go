package mock_repositories

import (
	"github.com/stretchr/testify/mock"
	"time"

	"saleorder/models"
)

type saleorderDetailRepositoryMock struct {
	mock.Mock
}

func NewSaleOrderDetailRepositoryMock() *saleorderDetailRepositoryMock {
	return &saleorderDetailRepositoryMock{}
}

func (m *saleorderDetailRepositoryMock) Create(prodReq *models.SaleOrderDetailCreate) (*models.SaleOrderDetailEntity, error) {
	args := m.Called(prodReq)
	return args.Get(0).(*models.SaleOrderDetailEntity), args.Error(1)
}

func (m *saleorderDetailRepositoryMock) GetAllQtyRates() ([]models.SaleOrderDetailRate, error) {
	args := m.Called()
	return args.Get(0).([]models.SaleOrderDetailRate), args.Error(1)
}

func (m *saleorderDetailRepositoryMock) GetAllQtyRatesByDay(dateReq time.Time) ([]models.SaleOrderDetailRate, error) {
	args := m.Called(dateReq)
	return args.Get(0).([]models.SaleOrderDetailRate), args.Error(1)
}

func (m *saleorderDetailRepositoryMock) GetAllQtyRatesByMonth(dateReq time.Time) ([]models.SaleOrderDetailRate, error) {
	args := m.Called(dateReq)
	return args.Get(0).([]models.SaleOrderDetailRate), args.Error(1)
}

func (m *saleorderDetailRepositoryMock) GetAllQtyRatesByYear(dateReq time.Time) ([]models.SaleOrderDetailRate, error) {
	args := m.Called(dateReq)
	return args.Get(0).([]models.SaleOrderDetailRate), args.Error(1)
}

func (m *saleorderDetailRepositoryMock) GetAllPriceRates() ([]models.SaleOrderDetailRate, error) {
	args := m.Called()
	return args.Get(0).([]models.SaleOrderDetailRate), args.Error(1)
}

func (m *saleorderDetailRepositoryMock) GetAllPriceRatesByDay(dateReq time.Time) ([]models.SaleOrderDetailRate, error) {
	args := m.Called(dateReq)
	return args.Get(0).([]models.SaleOrderDetailRate), args.Error(1)
}

func (m *saleorderDetailRepositoryMock) GetAllPriceRatesByMonth(dateReq time.Time) ([]models.SaleOrderDetailRate, error) {
	args := m.Called(dateReq)
	return args.Get(0).([]models.SaleOrderDetailRate), args.Error(1)
}

func (m *saleorderDetailRepositoryMock) GetAllPriceRatesByYear(dateReq time.Time) ([]models.SaleOrderDetailRate, error) {
	args := m.Called(dateReq)
	return args.Get(0).([]models.SaleOrderDetailRate), args.Error(1)
}