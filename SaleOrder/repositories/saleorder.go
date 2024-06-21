package repositories

import (
	"saleorder/models"
	"time"
)

type SaleOrderRepository interface {
	Create(*models.SaleOrderCreate) (*models.SaleOrderEntity, error)
	GetAll() ([]models.SaleOrderReport, error)
	GetAllByDay(time.Time) ([]models.SaleOrderReport, error)
	GetAllByMonth(time.Time) ([]models.SaleOrderReport, error)
	GetAllByYear(time.Time) ([]models.SaleOrderReport, error)
	GetTotalPricePass() (float64, error)
	GetTotalPricePassByDay(time.Time) (float64, error)
	GetTotalPricePassByMonth(time.Time) (float64, error)
	GetTotalPricePassByYear(time.Time) (float64, error)
}