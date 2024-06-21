package services

import (
	"saleorder/models"
	"saleorder/repositories"
	"saleorder/utils/errs"
	"saleorder/utils/logs"

	"time"
)

type saleorderDetailService struct {
	saleorderDetailRepo repositories.SaleOrderDetailRepository
}

func NewSaleOrderDetailService(saleorderDetailRepo repositories.SaleOrderDetailRepository) SaleOrderDetailService {
	return saleorderDetailService{saleorderDetailRepo: saleorderDetailRepo}
}

func (s saleorderDetailService) CreateSaleOrderDetail(saleOrderDetailReq *models.SaleOrderDetailCreate) (*models.SaleOrderDetail, error) {
	if saleOrderDetailReq.SO_Id == 0 {
		return nil, errs.NewBadRequestError("SaleOrderDetail's SaleOrder Id is null")
	}
	if saleOrderDetailReq.Prod_Id == 0 {
		return nil, errs.NewBadRequestError("SaleOrderDetail's User Id is null")
	}
	if saleOrderDetailReq.Quantity == 0 {
		return nil, errs.NewBadRequestError("SaleOrderDetail's Quantity is null")
	}
	if saleOrderDetailReq.Price == 0 {
		return nil, errs.NewBadRequestError("SaleOrderDetail's Price is null")
	}

	saleOrderDetailEntityRes, err := s.saleorderDetailRepo.Create(saleOrderDetailReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	saleOrderDetailRes := models.SaleOrderDetail{
		Seq:      		saleOrderDetailEntityRes.Seq,
		SO_Id:    		saleOrderDetailEntityRes.SO_Id,
		Prod_Id:    	saleOrderDetailEntityRes.Prod_Id,
		Quantity:     	saleOrderDetailEntityRes.Quantity,
		Price:         	saleOrderDetailEntityRes.Price,
		Discount:      	saleOrderDetailEntityRes.Discount,
	}

	logs.Info("Service: Create SaleOrderDetail Successfully")
	return &saleOrderDetailRes, nil
}

func (s saleorderDetailService) GetSaleOrderDetailQtyRates() ([]models.SaleOrderDetailRate, error) {
	saleordersQtyRate, err := s.saleorderDetailRepo.GetAllQtyRates()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	logs.Info("Service: Get SaleOrderDetail Qty Rate Successfully")
	return saleordersQtyRate, nil
}

func (s saleorderDetailService) GetSaleOrderDetailQtyRatesByDay(dayReq time.Time) ([]models.SaleOrderDetailRate, error) {
	saleordersQtyRate, err := s.saleorderDetailRepo.GetAllQtyRatesByDay(dayReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	logs.Info("Service: Get SaleOrderDetail Qty Rate By Day Successfully")
	return saleordersQtyRate, nil
}

func (s saleorderDetailService) GetSaleOrderDetailQtyRatesByMonth(dayReq time.Time) ([]models.SaleOrderDetailRate, error) {
	saleordersQtyRate, err := s.saleorderDetailRepo.GetAllQtyRatesByMonth(dayReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	logs.Info("Service: Get SaleOrderDetail Qty Rate By Month Successfully")
	return saleordersQtyRate, nil
}

func (s saleorderDetailService) GetSaleOrderDetailQtyRatesByYear(dayReq time.Time) ([]models.SaleOrderDetailRate, error) {
	saleordersQtyRate, err := s.saleorderDetailRepo.GetAllQtyRatesByYear(dayReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	logs.Info("Service: Get SaleOrderDetail Qty Rate By Year Successfully")
	return saleordersQtyRate, nil
}

func (s saleorderDetailService) GetSaleOrderDetailPriceRates() ([]models.SaleOrderDetailRate, error) {
	saleordersPriceRate, err := s.saleorderDetailRepo.GetAllPriceRates()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	logs.Info("Service: Get SaleOrderDetail Price Rate Successfully")
	return saleordersPriceRate, nil
}

func (s saleorderDetailService) GetSaleOrderDetailPriceRatesByDay(dayReq time.Time) ([]models.SaleOrderDetailRate, error) {
	saleordersPriceRate, err := s.saleorderDetailRepo.GetAllPriceRatesByDay(dayReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	logs.Info("Service: Get SaleOrderDetail Price Rate By Day Successfully")
	return saleordersPriceRate, nil
}

func (s saleorderDetailService) GetSaleOrderDetailPriceRatesByMonth(dayReq time.Time) ([]models.SaleOrderDetailRate, error) {
	saleordersPriceRate, err := s.saleorderDetailRepo.GetAllPriceRatesByMonth(dayReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	logs.Info("Service: Get SaleOrderDetail Price Rate By Month Successfully")
	return saleordersPriceRate, nil
}

func (s saleorderDetailService) GetSaleOrderDetailPriceRatesByYear(dayReq time.Time) ([]models.SaleOrderDetailRate, error) {
	saleordersPriceRate, err := s.saleorderDetailRepo.GetAllPriceRatesByYear(dayReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	logs.Info("Service: Get SaleOrderDetail Price Rate By Year Successfully")
	return saleordersPriceRate, nil
}