package mock_services

import (
	"product/models"
	"github.com/stretchr/testify/mock"
)

type prodTypeServiceMock struct {
	mock.Mock
}

func NewProductTypeServiceMock() *prodTypeServiceMock {
	return &prodTypeServiceMock{}
}

func (m *prodTypeServiceMock) CreateProductType(prodTypeReq *models.ProductTypeCreate) (*models.ProductType, error) {
	args := m.Called(prodTypeReq)
	return args.Get(0).(*models.ProductType), args.Error(1)
}

func (m *prodTypeServiceMock) GetProductTypes() ([]models.ProductType, error) {
	args := m.Called()
	if args.Get(0) != nil {
        return args.Get(0).([]models.ProductType), args.Error(1)
    }
    return nil, args.Error(1)
}

func (m *prodTypeServiceMock) GetProductType(id int) (*models.ProductType, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
        return args.Get(0).(*models.ProductType), args.Error(1)
    }
    return nil, args.Error(1)
}

func (m *prodTypeServiceMock) UpdateProductType(id int, reqProdType *models.ProductTypeUpdate) (*models.ProductType, error) {
	args := m.Called(id, reqProdType)
	return args.Get(0).(*models.ProductType), args.Error(1)
}

func (m *prodTypeServiceMock) DeleteProductType(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *prodTypeServiceMock) GetProductTypeCount() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}