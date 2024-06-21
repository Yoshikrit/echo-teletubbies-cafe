package mock_repositories

import (
	"github.com/stretchr/testify/mock"
	"product/models"
)

type productRepositoryMock struct {
	mock.Mock
}

func NewProductRepositoryMock() *productRepositoryMock {
	return &productRepositoryMock{}
}

func (m *productRepositoryMock) Create(prodReq *models.ProductCreate) (*models.ProductEntity, error) {
	args := m.Called(prodReq)
	return args.Get(0).(*models.ProductEntity), args.Error(1)
}

func (m *productRepositoryMock) GetAll() ([]models.ProductEntity, error) {
	args := m.Called()
	return args.Get(0).([]models.ProductEntity), args.Error(1)
}

func (m *productRepositoryMock) GetById(id int) (*models.ProductEntity, error) {
	args := m.Called(id)
	return args.Get(0).(*models.ProductEntity), args.Error(1)
}

func (m *productRepositoryMock) Update(id int, updateProd *models.ProductUpdate) (*models.ProductEntity, error) {
	args := m.Called(id, updateProd)
	return args.Get(0).(*models.ProductEntity), args.Error(1)
}

func (m *productRepositoryMock) DeleteById(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *productRepositoryMock) GetCount() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}
