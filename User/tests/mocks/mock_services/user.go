package mock_services

import (
	"user/models"
	"github.com/stretchr/testify/mock"
)

type userServiceMock struct {
	mock.Mock
}

func NewUserServiceMock() *userServiceMock {
	return &userServiceMock{}
}

func (m *userServiceMock) CreateUser(userReq *models.UserCreate) (*models.User, error) {
	args := m.Called(userReq)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *userServiceMock) GetUsers() ([]models.User, error) {
	args := m.Called()
	if args.Get(0) != nil {
        return args.Get(0).([]models.User), args.Error(1)
    }
    return nil, args.Error(1)
}

func (m *userServiceMock) GetUser(id int) (*models.User, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
        return args.Get(0).(*models.User), args.Error(1)
    }
    return nil, args.Error(1)
}

func (m *userServiceMock) UpdateUser(id int, userReq *models.UserUpdate) (*models.User, error) {
	args := m.Called(id, userReq)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *userServiceMock) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *userServiceMock) GetUserCount() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}