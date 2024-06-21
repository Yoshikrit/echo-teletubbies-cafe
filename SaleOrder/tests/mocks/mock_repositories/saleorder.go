package mock_repositories

import (
	"github.com/stretchr/testify/mock"
	"time"

	"saleorder/models"
)

type saleorderRepositoryMock struct {
	mock.Mock
}

func NewSaleOrderRepositoryMock() *saleorderRepositoryMock {
	return &saleorderRepositoryMock{}
}

func (m *saleorderRepositoryMock) Create(prodReq *models.SaleOrderCreate) (*models.SaleOrderEntity, error) {
	args := m.Called(prodReq)
	return args.Get(0).(*models.SaleOrderEntity), args.Error(1)
}

func (m *saleorderRepositoryMock) GetSaleOrders() ([]models.SaleOrderReport, error) {
	args := m.Called()
	return args.Get(0).([]models.SaleOrderReport), args.Error(1)
}

func (m *saleorderRepositoryMock) GetAll() ([]models.SaleOrderReport, error) {
	args := m.Called()
	return args.Get(0).([]models.SaleOrderReport), args.Error(1)
}

func (m *saleorderRepositoryMock) GetAllByDay(dateReq time.Time) ([]models.SaleOrderReport, error) {
	args := m.Called(dateReq)
	return args.Get(0).([]models.SaleOrderReport), args.Error(1)
}

func (m *saleorderRepositoryMock) GetAllByMonth(dateReq time.Time) ([]models.SaleOrderReport, error) {
	args := m.Called(dateReq)
	return args.Get(0).([]models.SaleOrderReport), args.Error(1)
}

func (m *saleorderRepositoryMock) GetAllByYear(dateReq time.Time) ([]models.SaleOrderReport, error) {
	args := m.Called(dateReq)
	return args.Get(0).([]models.SaleOrderReport), args.Error(1)
}

func (m *saleorderRepositoryMock) GetTotalPricePass() (float64, error) {
	args := m.Called()
	if val, ok := args.Get(0).(float64); ok {
		return val, args.Error(1)
	}

	return float64(args.Get(0).(int)), args.Error(1)
}

func (m *saleorderRepositoryMock) GetTotalPricePassByDay(dateReq time.Time) (float64, error) {
	args := m.Called(dateReq)
	if val, ok := args.Get(0).(float64); ok {
		return val, args.Error(1)
	}

	return float64(args.Get(0).(int)), args.Error(1)
}

func (m *saleorderRepositoryMock) GetTotalPricePassByMonth(dateReq time.Time) (float64, error) {
	args := m.Called(dateReq)
	if val, ok := args.Get(0).(float64); ok {
		return val, args.Error(1)
	}

	return float64(args.Get(0).(int)), args.Error(1)
}

func (m *saleorderRepositoryMock) GetTotalPricePassByYear(dateReq time.Time) (float64, error) {
	args := m.Called(dateReq)
	if val, ok := args.Get(0).(float64); ok {
		return val, args.Error(1)
	}

	return float64(args.Get(0).(int)), args.Error(1)
}