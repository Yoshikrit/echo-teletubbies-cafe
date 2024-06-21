package services

import (
	"saleorder/models"
	"saleorder/repositories"
	"saleorder/utils/errs"
	"saleorder/utils/logs"

	"time"
)

type saleorderService struct {
	saleorderRepo repositories.SaleOrderRepository
}

func NewSaleOrderService(saleorderRepo repositories.SaleOrderRepository) SaleOrderService {
	return saleorderService{saleorderRepo: saleorderRepo}
}

func (s saleorderService) CreateSaleOrder(saleOrderReq *models.SaleOrderCreate) (*models.SaleOrder, error) {
	if saleOrderReq.CreatedUser == 0 {
		return nil, errs.NewBadRequestError("SaleOrder's User Id is null")
	}
	if saleOrderReq.Status == "" {
		return nil, errs.NewBadRequestError("SaleOrder's Status is null")
	}

	saleOrderEntityRes, err := s.saleorderRepo.Create(saleOrderReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	saleOrderRes := models.SaleOrder{
		Id:      		saleOrderEntityRes.Id,
		CreatedUser:    saleOrderEntityRes.CreatedUser,
		CreatedDate:    saleOrderEntityRes.CreatedDate,
		TotalPrice:     saleOrderEntityRes.TotalPrice,
		Status:         saleOrderEntityRes.Status,
		PayMethod:      saleOrderEntityRes.PayMethod,
	}

	logs.Info("Service: Create SaleOrder Successfully")
	return &saleOrderRes, nil
}

func (s saleorderService) GetSaleOrders() ([]models.SaleOrderReport, error) {
	saleordersReport, err := s.saleorderRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	logs.Info("Service: Get SaleOrders Successfully")
	return saleordersReport, nil
}

func (s saleorderService) GetSaleOrdersByDay(dateReq time.Time) ([]models.SaleOrderReport, error) {
	saleordersReport, err := s.saleorderRepo.GetAllByDay(dateReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	logs.Info("Service: Get SaleOrders By Day Successfully")
	return saleordersReport, nil
}

func (s saleorderService) GetSaleOrdersByMonth(dateReq time.Time) ([]models.SaleOrderReport, error) {
	saleordersReport, err := s.saleorderRepo.GetAllByMonth(dateReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	logs.Info("Service: Get SaleOrders By Month Successfully")
	return saleordersReport, nil
}

func (s saleorderService) GetSaleOrdersByYear(dateReq time.Time) ([]models.SaleOrderReport, error) {
	saleordersReport, err := s.saleorderRepo.GetAllByYear(dateReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	logs.Info("Service: Get SaleOrders By Year Successfully")
	return saleordersReport, nil
}

func (s saleorderService) GetSaleOrderPriceAmountPass() (float64, error) {
	totalAmount, err := s.saleorderRepo.GetTotalPricePass()
	if err != nil {
		logs.Error(err)
		return 0, err
	}

	logs.Info("Service: Get Total Price That Pass Successfully")
	return totalAmount, nil
}

func (s saleorderService) GetSaleOrderPriceAmountPassByDay(dateReq time.Time) (float64, error) {
	totalAmount, err := s.saleorderRepo.GetTotalPricePassByDay(dateReq)
	if err != nil {
		logs.Error(err)
		return 0, err
	}

	logs.Info("Service: Get Total Price That Pass By Day Successfully")
	return totalAmount, nil
}

func (s saleorderService) GetSaleOrderPriceAmountPassByMonth(dateReq time.Time) (float64, error) {
	totalAmount, err := s.saleorderRepo.GetTotalPricePassByMonth(dateReq)
	if err != nil {
		logs.Error(err)
		return 0, err
	}

	logs.Info("Service: Get Total Price That Pass By Month Successfully")
	return totalAmount, nil
}

func (s saleorderService) GetSaleOrderPriceAmountPassByYear(dateReq time.Time) (float64, error) {
	totalAmount, err := s.saleorderRepo.GetTotalPricePassByYear(dateReq)
	if err != nil {
		logs.Error(err)
		return 0, err
	}

	logs.Info("Service: Get Total Price That Pass By Year Successfully")
	return totalAmount, nil
}