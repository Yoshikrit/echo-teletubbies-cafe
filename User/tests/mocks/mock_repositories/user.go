package mock_repositories

import (
	"github.com/stretchr/testify/mock"
	"user/models"
)

type userRepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock() *userRepositoryMock {
	return &userRepositoryMock{}
}

func (m *userRepositoryMock) Create(userReq *models.UserCreate) (*models.UserEntity, error) {
	args := m.Called(userReq)
	return args.Get(0).(*models.UserEntity), args.Error(1)
}

func (m *userRepositoryMock) GetAll() ([]models.UserEntity, error) {
	args := m.Called()
	return args.Get(0).([]models.UserEntity), args.Error(1)
}

func (m *userRepositoryMock) GetById(id int) (*models.UserEntity, error) {
	args := m.Called(id)
	return args.Get(0).(*models.UserEntity), args.Error(1)
}

func (m *userRepositoryMock) Update(id int, updateUser *models.UserUpdate) (*models.UserEntity, error) {
	args := m.Called(id, updateUser)
	return args.Get(0).(*models.UserEntity), args.Error(1)
}

func (m *userRepositoryMock) DeleteById(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *userRepositoryMock) GetCount() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

