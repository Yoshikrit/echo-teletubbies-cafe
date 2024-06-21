package services

import (
	"saleorder/models"

	"time"
)

type SaleOrderService interface {
	CreateSaleOrder(*models.SaleOrderCreate) (*models.SaleOrder, error);
	GetSaleOrders() ([]models.SaleOrderReport, error)
	GetSaleOrdersByDay(time.Time) ([]models.SaleOrderReport, error)
	GetSaleOrdersByMonth(time.Time) ([]models.SaleOrderReport, error)
	GetSaleOrdersByYear(time.Time) ([]models.SaleOrderReport, error)
	GetSaleOrderPriceAmountPass() (float64, error)
	GetSaleOrderPriceAmountPassByDay(time.Time) (float64, error)
	GetSaleOrderPriceAmountPassByMonth(time.Time) (float64, error)
	GetSaleOrderPriceAmountPassByYear(time.Time) (float64, error)
}

