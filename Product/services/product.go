package services

import (
	"product/models"
)

type ProductService interface {
	CreateProduct(*models.ProductCreate) (*models.Product, error)
	GetProducts() ([]models.Product, error)
	GetProduct(int) (*models.Product, error)
	UpdateProduct(int, *models.ProductUpdate) (*models.Product, error)
	DeleteProduct(int) (error)
	GetProductCount() (int64, error)
}