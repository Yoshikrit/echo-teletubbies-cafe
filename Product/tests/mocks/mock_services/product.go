package mock_services

import (
	"product/models"
	"github.com/stretchr/testify/mock"
)

type prodServiceMock struct {
	mock.Mock
}

func NewProductServiceMock() *prodServiceMock {
	return &prodServiceMock{}
}

func (m *prodServiceMock) CreateProduct(prodReq *models.ProductCreate) (*models.Product, error) {
	args := m.Called(prodReq)
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *prodServiceMock) GetProducts() ([]models.Product, error) {
	args := m.Called()
	if args.Get(0) != nil {
        return args.Get(0).([]models.Product), args.Error(1)
    }
    return nil, args.Error(1)
}

func (m *prodServiceMock) GetProduct(id int) (*models.Product, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
        return args.Get(0).(*models.Product), args.Error(1)
    }
    return nil, args.Error(1)
}

func (m *prodServiceMock) UpdateProduct(id int, reqProd *models.ProductUpdate) (*models.Product, error) {
	args := m.Called(id, reqProd)
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *prodServiceMock) DeleteProduct(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *prodServiceMock) GetProductCount() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}