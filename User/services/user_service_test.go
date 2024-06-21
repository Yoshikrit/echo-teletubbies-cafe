package services_test

import (
	"errors"
	"testing"
	"time"

	"user/tests/mocks/mock_repositories"
	"user/services"
	"user/utils/errs"
	"github.com/stretchr/testify/assert"

	"user/models"
)

func TestCreateUser(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	userReqMock := &models.UserCreate{
		Id:   1,
		Role_Id:   1,
		FName: "A",
		LName: "A",
		Email: "A",
		Password: "A",
		Sex: "M",
		TelNo: "xxxxxxxxxx",
		Salary: 100,
		Address: "A",
		WorkStatus: "W",
		BirthDate: mockDate,
	}

	userResMock := &models.User{
		Id:   1,
		Role_Id:   1,
		FName: "A",
		LName: "A",
		Email: "A",
		Sex: "M",
		TelNo: "xxxxxxxxxx",
		Salary: 100,
		Address: "A",
		WorkStatus: "W",
		BirthDate: mockDate,
	}

	userFromDBMock := &models.UserEntity{
		Id:   1,
		Role_Id:   1,
		FName: "A",
		LName: "A",
		Email: "A",
		Password: "A",
		Sex: "M",
		TelNo: "xxxxxxxxxx",
		Salary: 100,
		Address: "A",
		WorkStatus: "W",
		BirthDate: mockDate,
	}

	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewUserRepositoryMock()
		mockRepo.On("Create", userReqMock).Return(userFromDBMock, nil)

		service := services.NewUserService(mockRepo)
		userRes, err := service.CreateUser(userReqMock)

		expected := userResMock
		assert.NoError(t, err)
		assert.Equal(t, expected, userRes)
		assert.Nil(t, err)
	})

	type testCase struct {
		test_name       string
		id              int	 	
		roleId          int	 
		email           string
		password        string
		sex     		string	 
		salary     		float64	 
		workstatus     	string	
		birthdate     	time.Time	
		err 			error
	}
	cases := []testCase{
		{test_name: "test 1 fail 1 no id",        id: 0, roleId: 0, email:  "",   password: "",  sex: "",   salary: 0, workstatus: "",  birthdate: time.Time{},   err: errs.NewBadRequestError("User's Id is null")},
		{test_name: "test 2 fail 2 no role id",   id: 1, roleId: 0, email:  "",   password: "",  sex: "",   salary: 0, workstatus: "",  birthdate: time.Time{},   err: errs.NewBadRequestError("User's RoleId is null")},
		{test_name: "test 3 fail 3 no email",     id: 1, roleId: 1, email:  "",   password: "",  sex: "",   salary: 0, workstatus: "",  birthdate: time.Time{},   err: errs.NewBadRequestError("User's Email is null")},
		{test_name: "test 4 fail 4 no password",  id: 1, roleId: 1, email:  "A",  password: "",  sex: "",   salary: 0, workstatus: "",  birthdate: time.Time{},   err: errs.NewBadRequestError("User's Password is null")},
		{test_name: "test 5 fail 5 no sex",       id: 1, roleId: 1, email:  "A",  password: "A", sex: "",   salary: 0, workstatus: "",  birthdate: time.Time{},   err: errs.NewBadRequestError("User's Sex is null")},
		{test_name: "test 6 fail 6 no salary",    id: 1, roleId: 1, email:  "A",  password: "A", sex: "M",  salary: 0, workstatus: "",  birthdate: time.Time{},   err: errs.NewBadRequestError("User's Salary is null")},
		{test_name: "test 7 fail 7 no salary",    id: 1, roleId: 1, email:  "A",  password: "A", sex: "M",  salary: 1, workstatus: "",  birthdate: time.Time{},   err: errs.NewBadRequestError("User's WorkStatus is null")},
		{test_name: "test 8 fail 8 no salary",    id: 1, roleId: 1, email:  "A",  password: "A", sex: "M",  salary: 1, workstatus: "W", birthdate: time.Time{},   err: errs.NewBadRequestError("User's BirthDate is null")},
		{test_name: "test 9 fail 9 repo",         id: 1, roleId: 1, email:  "A",  password: "A", sex: "M",  salary: 1, workstatus: "W", birthdate: mockDate,   	  err: errors.New("")},
	}

	for _, tc := range cases {
		userReqFail := &models.UserCreate{
			Id:        	tc.id,
			Role_Id:   	tc.roleId,
			Email:     	tc.email,
			Password:  	tc.password,
			Sex:       	tc.sex,
			Salary:    	tc.salary,
			WorkStatus: tc.workstatus,
			BirthDate:  tc.birthdate,
		}

		t.Run(tc.test_name, func(t *testing.T) {
			mockRepo := mock_repositories.NewUserRepositoryMock()
			mockRepo.On("Create", userReqFail).Return(&models.UserEntity{}, errors.New(""))
			service := services.NewUserService(mockRepo)

			userRes, err := service.CreateUser(userReqFail)

			expected := tc.err
			assert.Error(t, err)
			assert.Equal(t, expected, err)
			assert.Nil(t, userRes)
		})
	}
}

func TestGetUsers(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	usersDBMock := []models.UserEntity{
		{
			Id:   1,
			Role_Id:   1,
			FName: "A",
			LName: "A",
			Email: "A",
			Password: "A",
			Sex: "M",
			TelNo: "xxxxxxxxxx",
			Salary: 100,
			Address: "A",
			WorkStatus: "W",
			BirthDate: mockDate,
		},
		{
			Id:   2,
			Role_Id:   1,
			FName: "A",
			LName: "A",
			Email: "A",
			Password: "A",
			Sex: "M",
			TelNo: "xxxxxxxxxx",
			Salary: 100,
			Address: "A",
			WorkStatus: "W",
			BirthDate: mockDate,
		},
	}

	usersResMock := []models.User{
		{
			Id:   1,
			Role_Id:   1,
			FName: "A",
			LName: "A",
			Email: "A",
			Sex: "M",
			TelNo: "xxxxxxxxxx",
			Salary: 100,
			Address: "A",
			WorkStatus: "W",
			BirthDate: mockDate,
		},
		{
			Id:   2,
			Role_Id:   1,
			FName: "A",
			LName: "A",
			Email: "A",
			Sex: "M",
			TelNo: "xxxxxxxxxx",
			Salary: 100,
			Address: "A",
			WorkStatus: "W",
			BirthDate: mockDate,
		},
	}

	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewUserRepositoryMock()
		mockRepo.On("GetAll").Return(usersDBMock, nil)

		service := services.NewUserService(mockRepo)
		usersRes, err := service.GetUsers()

		assert.NoError(t, err)
		assert.Equal(t, usersResMock, usersRes)
		assert.Nil(t, err)
	})

	t.Run("test case : repository error", func(t *testing.T) {
		mockRepo := mock_repositories.NewUserRepositoryMock()
		mockRepo.On("GetAll").Return([]models.UserEntity{}, errors.New(""))

		service := services.NewUserService(mockRepo)
		usersRes, err := service.GetUsers()

		expected := errors.New("")

		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, usersRes)
	})
}

func TestGetUser(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	userDBMock := &models.UserEntity{
		Id:   1,
		Role_Id:   1,
		FName: "A",
		LName: "A",
		Email: "A",
		Password: "A",
		Sex: "M",
		TelNo: "xxxxxxxxxx",
		Salary: 100,
		Address: "A",
		WorkStatus: "W",
		BirthDate: mockDate,
	}
	userResMock := &models.User{
		Id:   1,
		Role_Id:   1,
		FName: "A",
		LName: "A",
		Email: "A",
		Sex: "M",
		TelNo: "xxxxxxxxxx",
		Salary: 100,
		Address: "A",
		WorkStatus: "W",
		BirthDate: mockDate,
	}

	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewUserRepositoryMock()
		mockRepo.On("GetById", 1).Return(userDBMock, nil)

		service := services.NewUserService(mockRepo)
		userResponse, err := service.GetUser(1)

		assert.NoError(t, err)
		assert.Equal(t, userResMock, userResponse)
		assert.Nil(t, err)
	})

	t.Run("test case : repository error", func(t *testing.T) {
		mockRepo := mock_repositories.NewUserRepositoryMock()
		mockRepo.On("GetById", 1).Return(&models.UserEntity{}, errors.New(""))

		service := services.NewUserService(mockRepo)
		prodRes, err := service.GetUser(1)

		expected := errors.New("")

		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, prodRes)
	})
}

func TestUpdateUser(t *testing.T) {
	mockDate := time.Date(2023, time.December, 28, 12, 30, 0, 0, time.UTC)
	userReqMock := &models.UserUpdate{
		Role_Id:   1,
		FName: "A",
		LName: "A",
		Email: "A",
		Password: "A",
		Sex: "M",
		TelNo: "xxxxxxxxxx",
		Salary: 100,
		Address: "A",
		WorkStatus: "W",
		BirthDate: mockDate,
	}
	userErrorReqMock := &models.UserUpdate{
		Role_Id:   0,
		FName: "",
		LName: "",
		Email: "",
		Password: "",
		Sex: "",
		TelNo: "",
		Salary: 0,
		Address: "",
		WorkStatus: "",
		BirthDate: time.Time{},
	}
	userDBMock := &models.UserEntity{
		Id:   1,
		Role_Id:   1,
		FName: "A",
		LName: "A",
		Email: "A",
		Password: "A",
		Sex: "M",
		TelNo: "xxxxxxxxxx",
		Salary: 100,
		Address: "A",
		WorkStatus: "W",
		BirthDate: mockDate,
	}
	userResMock := &models.User{
		Id:   1,
		Role_Id:   1,
		FName: "A",
		LName: "A",
		Email: "A",
		Sex: "M",
		TelNo: "xxxxxxxxxx",
		Salary: 100,
		Address: "A",
		WorkStatus: "W",
		BirthDate: mockDate,
	}

	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewUserRepositoryMock()
		mockRepo.On("Update", 1, userReqMock).Return(userDBMock, nil)

		service := services.NewUserService(mockRepo)
		userRes, err := service.UpdateUser(1, userReqMock)

		assert.NoError(t, err)
		assert.Equal(t, userResMock, userRes)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : fail null", func(t *testing.T) {
		mockRepo := mock_repositories.NewUserRepositoryMock()

		service := services.NewUserService(mockRepo)
		userRes, err := service.UpdateUser(1, userErrorReqMock)

		expected := errs.NewBadRequestError("User's Data is null")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, userRes)
	})

	t.Run("test case : fail repository", func(t *testing.T) {
		mockRepo := mock_repositories.NewUserRepositoryMock()
		mockRepo.On("Update", 1, userReqMock).Return(&models.UserEntity{}, errors.New(""))
		service := services.NewUserService(mockRepo)

		userTypeRes, err := service.UpdateUser(1, userReqMock)

		expected := errors.New("")
		assert.Error(t, err)
		assert.Equal(t, expected, err)
		assert.Nil(t, userTypeRes)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewUserRepositoryMock()
		mockRepo.On("DeleteById", 1).Return(nil)

		service := services.NewUserService(mockRepo)
		err := service.DeleteUser(1)

		assert.NoError(t, err)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : repository error", func(t *testing.T) {
		mockRepo := mock_repositories.NewUserRepositoryMock()
		mockRepo.On("DeleteById", 1).Return(errors.New(""))

		service := services.NewUserService(mockRepo)
		err := service.DeleteUser(1)

		expected := errors.New("")
		assert.Equal(t, expected, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetUserCount(t *testing.T) {
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewUserRepositoryMock()
		mockRepo.On("GetCount").Return(int64(5), nil)

		service := services.NewUserService(mockRepo)
		count, err := service.GetUserCount()

		mockRepo.AssertExpectations(t)
		assert.NoError(t, err)
		assert.Equal(t, int64(5), count)
		assert.Nil(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("test case : repository error", func(t *testing.T) {
		mockRepo := mock_repositories.NewUserRepositoryMock()
		mockRepo.On("GetCount").Return(int64(0), errors.New(""))

		service := services.NewUserService(mockRepo)
		count, err := service.GetUserCount()

		expected := errors.New("")
		assert.Equal(t, expected, err)
		assert.Equal(t, int64(0), count)
		mockRepo.AssertExpectations(t)
	})
}