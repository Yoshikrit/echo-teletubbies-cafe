package repositories

import (
	"product/models"
)

type ProductRepository interface {
	Create(*models.ProductCreate) (*models.ProductEntity, error)
	GetAll() ([]models.ProductEntity, error)
	GetById(int) (*models.ProductEntity, error)
	Update(int, *models.ProductUpdate) (*models.ProductEntity, error)
	DeleteById(int) (error)
	GetCount() (int64, error)
}