package mock_services

import (
	"user/models"
	"github.com/stretchr/testify/mock"
)

type roleServiceMock struct {
	mock.Mock
}

func NewRoleServiceMock() *roleServiceMock {
	return &roleServiceMock{}
}

func (m *roleServiceMock) CreateRole(roleReq *models.RoleCreate) (*models.Role, error) {
	args := m.Called(roleReq)
	return args.Get(0).(*models.Role), args.Error(1)
}

func (m *roleServiceMock) GetRoles() ([]models.Role, error) {
	args := m.Called()
	if args.Get(0) != nil {
        return args.Get(0).([]models.Role), args.Error(1)
    }
    return nil, args.Error(1)
}

func (m *roleServiceMock) GetRole(id int) (*models.Role, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
        return args.Get(0).(*models.Role), args.Error(1)
    }
    return nil, args.Error(1)
}

func (m *roleServiceMock) UpdateRole(id int, roleReq *models.RoleUpdate) (*models.Role, error) {
	args := m.Called(id, roleReq)
	return args.Get(0).(*models.Role), args.Error(1)
}

func (m *roleServiceMock) DeleteRole(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *roleServiceMock) GetRoleCount() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}