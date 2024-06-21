package services_test

import (
	"errors"
	"testing"

	"auth/tests/mocks/mock_repositories"
	"auth/services"
	"auth/models"
	"auth/utils/errs"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	userReqMock := &models.UserLogin{
		Email:    "C",
		Password: "A",
	}

	userClaimResMock := &models.UserClaim{
		Name:   "C",
		Role:   "A",
	}
	
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewAuthRepositoryMock()
		mockRepo.On("GetUserClaimByEmailAndPassword", userReqMock).Return(userClaimResMock, nil)
		mockRepo.On("Login", 0).Return(nil)

		service := services.NewAuthService(mockRepo)
		_, err := service.Login(userReqMock)

		mockRepo.AssertExpectations(t)
		assert.NoError(t, err)
		assert.Nil(t, err)
	})

	userReqErrorMock := &models.UserLogin{
		Email:    "",
		Password: "",
	}

	t.Run("test case : fail check null", func(t *testing.T) {
		mockRepo := mock_repositories.NewAuthRepositoryMock()

		service := services.NewAuthService(mockRepo)
		resRes, err := service.Login(userReqErrorMock)

		assert.Error(t, err)
		assert.Nil(t, resRes)
	})

	t.Run("test case : fail check null", func(t *testing.T) {
		mockRepo := mock_repositories.NewAuthRepositoryMock()
		mockRepo.On("GetUserClaimByEmailAndPassword", userReqMock).Return(&models.UserClaim{}, errors.New(""))

		service := services.NewAuthService(mockRepo)
		resRes, err := service.Login(userReqMock)

		assert.Error(t, err)
		assert.Nil(t, resRes)
	})

	t.Run("test case : fail check null", func(t *testing.T) {
		mockRepo := mock_repositories.NewAuthRepositoryMock()
		mockRepo.On("GetUserClaimByEmailAndPassword", userReqMock).Return(userClaimResMock, nil)
		mockRepo.On("Login", 0).Return(errs.NewUnexpectedError(""))

		service := services.NewAuthService(mockRepo)
		resRes, err := service.Login(userReqMock)

		mockRepo.AssertExpectations(t)
		assert.Error(t, err)
		assert.Nil(t, resRes)
	})
}

func TestLogOut(t *testing.T) {
	t.Run("test case : pass", func(t *testing.T) {
		mockRepo := mock_repositories.NewAuthRepositoryMock()
		mockRepo.On("Logout", 1).Return(nil)

		service := services.NewAuthService(mockRepo)
		err := service.Logout(1)

		mockRepo.AssertExpectations(t)
		assert.NoError(t, err)
		assert.Nil(t, err)
	})

	t.Run("test case : fail cant log out", func(t *testing.T) {
		mockRepo := mock_repositories.NewAuthRepositoryMock()
		mockRepo.On("Logout", 1).Return(errs.NewUnexpectedError(""))

		service := services.NewAuthService(mockRepo)
		err := service.Logout(1)

		expected := errs.NewUnexpectedError("")

		mockRepo.AssertExpectations(t)
		assert.Error(t, err)
		assert.Equal(t, expected, err)
	})
}

