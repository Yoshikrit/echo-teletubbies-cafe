package services

import (
	"saleorder/models"

	"time"
)

type SaleOrderDetailService interface {
	CreateSaleOrderDetail(*models.SaleOrderDetailCreate) (*models.SaleOrderDetail, error);
	GetSaleOrderDetailQtyRates() ([]models.SaleOrderDetailRate, error)
	GetSaleOrderDetailQtyRatesByDay(time.Time) ([]models.SaleOrderDetailRate, error)
	GetSaleOrderDetailQtyRatesByMonth(time.Time) ([]models.SaleOrderDetailRate, error)
	GetSaleOrderDetailQtyRatesByYear(time.Time) ([]models.SaleOrderDetailRate, error)
	GetSaleOrderDetailPriceRates() ([]models.SaleOrderDetailRate, error)
	GetSaleOrderDetailPriceRatesByDay(time.Time) ([]models.SaleOrderDetailRate, error)
	GetSaleOrderDetailPriceRatesByMonth(time.Time) ([]models.SaleOrderDetailRate, error)
	GetSaleOrderDetailPriceRatesByYear(time.Time) ([]models.SaleOrderDetailRate, error)
}