package mock_repositories

import (
	"github.com/stretchr/testify/mock"
	"user/models"
)

type roleRepositoryMock struct {
	mock.Mock
}

func NewRoleRepositoryMock() *roleRepositoryMock {
	return &roleRepositoryMock{}
}

func (m *roleRepositoryMock) Create(roleReq *models.RoleCreate) (*models.RoleEntity, error) {
	args := m.Called(roleReq)
	return args.Get(0).(*models.RoleEntity), args.Error(1)
}

func (m *roleRepositoryMock) GetAll() ([]models.RoleEntity, error) {
	args := m.Called()
	return args.Get(0).([]models.RoleEntity), args.Error(1)
}

func (m *roleRepositoryMock) GetById(id int) (*models.RoleEntity, error) {
	args := m.Called(id)
	return args.Get(0).(*models.RoleEntity), args.Error(1)
}

func (m *roleRepositoryMock) Update(id int, updateRole *models.RoleUpdate) (*models.RoleEntity, error) {
	args := m.Called(id, updateRole)
	return args.Get(0).(*models.RoleEntity), args.Error(1)
}

func (m *roleRepositoryMock) DeleteById(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *roleRepositoryMock) GetCount() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}
