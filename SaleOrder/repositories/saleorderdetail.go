package repositories

import (
	"saleorder/models"

	"time"
)

type SaleOrderDetailRepository interface {
	Create(*models.SaleOrderDetailCreate) (*models.SaleOrderDetailEntity, error)
	GetAllQtyRates() ([]models.SaleOrderDetailRate, error)
	GetAllQtyRatesByDay(time.Time) ([]models.SaleOrderDetailRate, error)
	GetAllQtyRatesByMonth(time.Time) ([]models.SaleOrderDetailRate, error)
	GetAllQtyRatesByYear(time.Time) ([]models.SaleOrderDetailRate, error)
	GetAllPriceRates() ([]models.SaleOrderDetailRate, error)
	GetAllPriceRatesByDay(time.Time) ([]models.SaleOrderDetailRate, error)
	GetAllPriceRatesByMonth(time.Time) ([]models.SaleOrderDetailRate, error)
	GetAllPriceRatesByYear(time.Time) ([]models.SaleOrderDetailRate, error)
}