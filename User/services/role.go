package services

import (
	"user/models"
)

type RoleService interface {
	CreateRole(*models.RoleCreate) (*models.Role, error)
	GetRoles() ([]models.Role, error)
	GetRole(int) (*models.Role, error)
	UpdateRole(int, *models.RoleUpdate) (*models.Role, error)
	DeleteRole(int) (error)
	GetRoleCount() (int64, error)
}