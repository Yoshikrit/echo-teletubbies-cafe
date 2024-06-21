package services

import (
	"auth/models"
)

type AuthService interface {
	Login(*models.UserLogin) (*models.Response, error)
	Logout(int) error
}