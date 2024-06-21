package mock_repositories

import (
	"github.com/stretchr/testify/mock"
	"auth/models"
)

type authRepositoryMock struct {
	mock.Mock
}

func NewAuthRepositoryMock() *authRepositoryMock {
	return &authRepositoryMock{}
}

func (m *authRepositoryMock) GetUserClaimByEmailAndPassword(userLogin *models.UserLogin) (*models.UserClaim, error) {
	args := m.Called(userLogin)
	return args.Get(0).(*models.UserClaim), args.Error(1)
}

func (m *authRepositoryMock) Login(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *authRepositoryMock) Logout(id int) error {
	args := m.Called(id)
	return args.Error(0)
}