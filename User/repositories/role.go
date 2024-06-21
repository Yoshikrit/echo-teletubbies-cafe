package repositories

import (
	"user/models"
)

type RoleRepository interface {
	Create(*models.RoleCreate) (*models.RoleEntity, error)
	GetAll() ([]models.RoleEntity, error)
	GetById(int) (*models.RoleEntity, error)
	Update(int, *models.RoleUpdate) (*models.RoleEntity, error)
	DeleteById(int) (error)
	GetCount() (int64, error)
}