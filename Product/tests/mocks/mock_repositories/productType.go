package mock_repositories

import (
	"github.com/stretchr/testify/mock"
	"product/models"
)

type productTypeRepositoryMock struct {
	mock.Mock
}

func NewProductTypeRepositoryMock() *productTypeRepositoryMock {
	return &productTypeRepositoryMock{}
}

func (m *productTypeRepositoryMock) Create(prodTypeReq *models.ProductTypeCreate) (*models.ProductTypeEntity, error) {
	args := m.Called(prodTypeReq)
	return args.Get(0).(*models.ProductTypeEntity), args.Error(1)
}

func (m *productTypeRepositoryMock) GetAll() ([]models.ProductTypeEntity, error) {
	args := m.Called()
	return args.Get(0).([]models.ProductTypeEntity), args.Error(1)
}

func (m *productTypeRepositoryMock) GetById(id int) (*models.ProductTypeEntity, error) {
	args := m.Called(id)
	return args.Get(0).(*models.ProductTypeEntity), args.Error(1)
}

func (m *productTypeRepositoryMock) Update(id int, updateProdType *models.ProductTypeUpdate) (*models.ProductTypeEntity, error) {
	args := m.Called(id, updateProdType)
	return args.Get(0).(*models.ProductTypeEntity), args.Error(1)
}

func (m *productTypeRepositoryMock) DeleteById(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *productTypeRepositoryMock) GetCount() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}
