package mock_services

import (
	"github.com/stretchr/testify/mock"
	"time"
	"fmt"

	"saleorder/models"
)

type saleorderServiceMock struct {
	mock.Mock
}

func NewSaleOrderServiceMock() *saleorderServiceMock {
	return &saleorderServiceMock{}
}

func (m *saleorderServiceMock) CreateSaleOrder(saleorderReq *models.SaleOrderCreate) (*models.SaleOrder, error) {
	args := m.Called(saleorderReq)
	fmt.Printf("Method arguments: %v\n", args)
	return args.Get(0).(*models.SaleOrder), args.Error(1)
}

func (m *saleorderServiceMock) GetSaleOrders() ([]models.SaleOrderReport, error) {
	args := m.Called()
	return args.Get(0).([]models.SaleOrderReport), args.Error(1)
}

func (m *saleorderServiceMock) GetSaleOrdersByDay(dateReq time.Time) ([]models.SaleOrderReport, error) {
	args := m.Called(dateReq)
	return args.Get(0).([]models.SaleOrderReport), args.Error(1)
}

func (m *saleorderServiceMock) GetSaleOrdersByMonth(dateReq time.Time) ([]models.SaleOrderReport, error) {
	args := m.Called(dateReq)
	return args.Get(0).([]models.SaleOrderReport), args.Error(1)
}

func (m *saleorderServiceMock) GetSaleOrdersByYear(dateReq time.Time) ([]models.SaleOrderReport, error) {
	args := m.Called(dateReq)
	return args.Get(0).([]models.SaleOrderReport), args.Error(1)
}

func (m *saleorderServiceMock) GetSaleOrderPriceAmountPass() (float64, error) {
	args := m.Called()
	if val, ok := args.Get(0).(float64); ok {
		return val, args.Error(1)
	}

	return float64(args.Get(0).(int)), args.Error(1)
}

func (m *saleorderServiceMock) GetSaleOrderPriceAmountPassByDay(dateReq time.Time) (float64, error) {
	args := m.Called(dateReq)
	if val, ok := args.Get(0).(float64); ok {
		return val, args.Error(1)
	}

	return float64(args.Get(0).(int)), args.Error(1)
}

func (m *saleorderServiceMock) GetSaleOrderPriceAmountPassByMonth(dateReq time.Time) (float64, error) {
	args := m.Called(dateReq)
	if val, ok := args.Get(0).(float64); ok {
		return val, args.Error(1)
	}

	return float64(args.Get(0).(int)), args.Error(1)
}

func (m *saleorderServiceMock) GetSaleOrderPriceAmountPassByYear(dateReq time.Time) (float64, error) {
	args := m.Called(dateReq)
	if val, ok := args.Get(0).(float64); ok {
		return val, args.Error(1)
	}

	return float64(args.Get(0).(int)), args.Error(1)
}