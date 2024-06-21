package mock_services

import (
	"auth/models"
	"github.com/stretchr/testify/mock"
)

type authServiceMock struct {
	mock.Mock
}

func NewAuthServiceMock() *authServiceMock {
	return &authServiceMock{}
}

func (m *authServiceMock) Login(userLoginReq *models.UserLogin) (*models.Response, error) {
	args := m.Called(userLoginReq)
	return args.Get(0).(*models.Response), args.Error(1)
}

func (m *authServiceMock) Logout(id int) error {
	args := m.Called(id)
	return args.Error(0)
}