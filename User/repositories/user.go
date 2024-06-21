package repositories

import (
	"user/models"
)

type UserRepository interface {
	Create(*models.UserCreate) (*models.UserEntity, error)
	GetAll() ([]models.UserEntity, error)
	GetById(int) (*models.UserEntity, error)
	Update(int, *models.UserUpdate) (*models.UserEntity, error)
	DeleteById(int) (error)
	GetCount() (int64, error)
}