package services

import (
	"product/models"
)

type ProductTypeService interface {
	CreateProductType(*models.ProductTypeCreate) (*models.ProductType, error)
	GetProductTypes() ([]models.ProductType, error)
	GetProductType(int) (*models.ProductType, error)
	UpdateProductType(int, *models.ProductTypeUpdate) (*models.ProductType, error)
	DeleteProductType(int) (error)
	GetProductTypeCount() (int64, error)
}