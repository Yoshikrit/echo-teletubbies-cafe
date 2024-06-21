package services

import (
	"user/models"
)

type UserService interface {
	CreateUser(*models.UserCreate) (*models.User, error)
	GetUsers() ([]models.User, error)
	GetUser(int) (*models.User, error)
	UpdateUser(int, *models.UserUpdate) (*models.User, error)
	DeleteUser(int) (error)
	GetUserCount() (int64, error)
}