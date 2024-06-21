package repositories

import (
	"product/models"
)

type ProductTypeRepository interface {
	Create(*models.ProductTypeCreate) (*models.ProductTypeEntity, error)
	GetAll() ([]models.ProductTypeEntity, error)
	GetById(int) (*models.ProductTypeEntity, error)
	Update(int, *models.ProductTypeUpdate) (*models.ProductTypeEntity, error)
	DeleteById(int) (error)
	GetCount() (int64, error)
}