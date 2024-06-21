package services_test

import (
	"errors"
	"testing"

	"user/tests/mocks/mock_repositories"
	"user/services"
	"user/utils/errs"
	"github.com/stretchr/testify/assert"

	"user/models"
)

func TestCreateRole(t *testing.T) {
	roleReqMock := &models.RoleCreate{
		Id:   1,
		Name: "A",
	}

	roleResMock := &models.Role{
		Id:   1,
		Name: "A",
	}

	roleFromDBMock := &models.RoleEntity{
		Id:   1,
		Name: "A",
	}

	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewRoleRepositoryMock()
		mockRepo.On("Create", roleReqMock).Return(roleFromDBMock, nil)

		service := services.NewRoleService(mockRepo)
		roleRes, err := service.CreateRole(roleReqMock)

		expected := roleResMock
		assert.NoError(t, err)
		assert.Equal(t, expected, roleRes)
		assert.Nil(t, err)
	})

	type testCase struct {
		test_name       string
		id              int	 
		name 			string	 
		err 			error
	}
	cases := []testCase{
		{test_name: "test case : fail no id",        id: 0, name: "",  err: errs.NewBadRequestError("Role's Id is null")},
		{test_name: "test case : fail no name",      id: 1, name: "",  err: errs.NewBadRequestError("Role's Name is null")},
		{test_name: "test case : fail repo",         id: 1, name: "A", err: errors.New("")},
	}

	for _, tc := range cases {
		roleReqFail := &models.RoleCreate{
			Id:        tc.id,
			Name: 	   tc.name,
		}

		t.Run(tc.test_name, func(t *testing.T) {
			mockRepo := mock_repositories.NewRoleRepositoryMock()
			mockRepo.On("Create", roleReqFail).Return(&models.RoleEntity{}, errors.New(""))
			service := services.NewRoleService(mockRepo)

			roleRes, err := service.CreateRole(roleReqFail)

			expected := tc.err
			assert.Error(t, err)
			assert.Equal(t, expected, err)
			assert.Nil(t, roleRes)
		})
	}
}

func TestGetRoles(t *testing.T) {
	rolesDBMock := []models.RoleEntity{
		{
			Id:   1,
			Name: "A",
		},
		{
			Id:   2,
			Name: "A",
		},
	}

	rolesResMock := []models.Role{
		{
			Id:   1,
			Name: "A",
		},
		{
			Id:   2,
			Name: "A",
		},
	}

	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewRoleRepositoryMock()
		mockRepo.On("GetAll").Return(rolesDBMock, nil)

		service := services.NewRoleService(mockRepo)
		rolesRes, err := service.GetRoles()

		assert.NoError(t, err)
		assert.Equal(t, rolesResMock, rolesRes)
		assert.Nil(t, err)
	})

	t.Run("test case : repository error", func(t *testing.T) {
		mockRepo := mock_repositories.NewRoleRepositoryMock()
		mockRepo.On("GetAll").Return([]models.RoleEntity{}, errors.New(""))

		service := services.NewRoleService(mockRepo)
		rolesRes, err := service.GetRoles()

		expected := errors.New("")

		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, rolesRes)
	})
}

func TestGetRole(t *testing.T) {
	roleDBMock := &models.RoleEntity{
		Id:   1,
		Name: "A",
	}
	roleResMock := &models.Role{
		Id:   1,
		Name: "A",
	}

	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewRoleRepositoryMock()
		mockRepo.On("GetById", 1).Return(roleDBMock, nil)

		service := services.NewRoleService(mockRepo)
		roleResponse, err := service.GetRole(1)

		assert.NoError(t, err)
		assert.Equal(t, roleResMock, roleResponse)
		assert.Nil(t, err)
	})

	t.Run("test case : repository error", func(t *testing.T) {
		mockRepo := mock_repositories.NewRoleRepositoryMock()
		mockRepo.On("GetById", 1).Return(&models.RoleEntity{}, errors.New(""))

		service := services.NewRoleService(mockRepo)
		roleRes, err := service.GetRole(1)

		expected := errors.New("")

		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, roleRes)
	})
}

func TestUpdateRole(t *testing.T) {
	roleReqMock := &models.RoleUpdate{
		Name: "A",
	}
	roleDBMock := &models.RoleEntity{
		Id:   1,
		Name: "A",
	}
	roleResMock := &models.Role{
		Id:   1,
		Name: "A",
	}

	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewRoleRepositoryMock()
		mockRepo.On("Update", 1, roleReqMock).Return(roleDBMock, nil)

		service := services.NewRoleService(mockRepo)
		roleRes, err := service.UpdateRole(1, roleReqMock)

		assert.NoError(t, err)
		assert.Equal(t, roleResMock, roleRes)
		assert.Nil(t, err)
	})

	type testCase struct {
		test_name       string
		id              int	 
		name            string 
		err 			error
	}
	cases := []testCase{
		{test_name: "test case : fail no name",   id: 1, name: "", err: errs.NewBadRequestError("Role's Name is null")},
		{test_name: "test case : fail repo",      id: 1, name: "A", err: errors.New("")},
	}

	for _, tc := range cases {
		roleReqFail := &models.RoleUpdate{
			Name:   tc.name,
		}

		t.Run(tc.test_name, func(t *testing.T) {
			mockRepo := mock_repositories.NewRoleRepositoryMock()
			mockRepo.On("Update", 1, roleReqFail).Return(&models.RoleEntity{}, errors.New(""))
			service := services.NewRoleService(mockRepo)

			roleTypeRes, err := service.UpdateRole(1, roleReqFail)

			expected := tc.err
			assert.Error(t, err)
			assert.Equal(t, expected, err)
			assert.Nil(t, roleTypeRes)
		})
	}
}

func TestDeleteRole(t *testing.T) {
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewRoleRepositoryMock()
		mockRepo.On("DeleteById", 1).Return(nil)

		service := services.NewRoleService(mockRepo)
		err := service.DeleteRole(1)

		assert.NoError(t, err)
		assert.Nil(t, err)
	})

	t.Run("test case : repository error", func(t *testing.T) {
		mockRepo := mock_repositories.NewRoleRepositoryMock()
		mockRepo.On("DeleteById", 1).Return(errors.New(""))

		service := services.NewRoleService(mockRepo)
		err := service.DeleteRole(1)

		expected := errors.New("")
		assert.Equal(t, expected, err)
	})
}

func TestGetRoleCount(t *testing.T) {
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewRoleRepositoryMock()
		mockRepo.On("GetCount").Return(int64(5), nil)

		service := services.NewRoleService(mockRepo)
		count, err := service.GetRoleCount()

		mockRepo.AssertExpectations(t)
		assert.NoError(t, err)
		assert.Equal(t, int64(5), count)
		assert.Nil(t, err)
	})

	t.Run("test case : repository error", func(t *testing.T) {
		mockRepo := mock_repositories.NewRoleRepositoryMock()
		mockRepo.On("GetCount").Return(int64(0), errors.New(""))

		service := services.NewRoleService(mockRepo)
		count, err := service.GetRoleCount()

		expected := errors.New("")
		assert.Equal(t, expected, err)
		assert.Equal(t, int64(0), count)
	})
}