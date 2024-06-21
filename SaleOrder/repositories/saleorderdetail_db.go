package repositories

import (
	"saleorder/models"
	"saleorder/utils/errs"
	
	"strconv"
    "gorm.io/gorm"
	"time"
)

type saleorderDetailRepositoryDB struct {
	db *gorm.DB
}

func NewSaleOrderDetailRepositoryDB(db *gorm.DB) SaleOrderDetailRepository {
	return saleorderDetailRepositoryDB{db: db}
}

func (r saleorderDetailRepositoryDB) Create(saleorderDetailReq *models.SaleOrderDetailCreate) (*models.SaleOrderDetailEntity, error) {
	var saleorderFromDB models.SaleOrderEntity
	if err := r.db.First(&saleorderFromDB, saleorderDetailReq.SO_Id).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, errs.NewNotFoundError("No SaleOrder from SaleOrder's Id : " + strconv.Itoa(saleorderDetailReq.SO_Id))
		}
		return nil, errs.NewUnexpectedError(err.Error())
	}

	var productFromDB models.ProductEntity
	if err := r.db.First(&productFromDB, saleorderDetailReq.Prod_Id).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, errs.NewNotFoundError("No Product from Product's Id : " + strconv.Itoa(saleorderDetailReq.Prod_Id))
		}
		return nil, errs.NewUnexpectedError(err.Error())
	}

	var saleorderDetailEntity models.SaleOrderDetailEntity
	saleorderDetailEntity.SO_Id = saleorderDetailReq.SO_Id
	saleorderDetailEntity.Prod_Id = saleorderDetailReq.Prod_Id
	saleorderDetailEntity.Quantity = saleorderDetailReq.Quantity
	saleorderDetailEntity.Price = saleorderDetailReq.Price
	saleorderDetailEntity.Discount = saleorderDetailReq.Discount

	if err := r.db.Create(&saleorderDetailEntity).Error; err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return &saleorderDetailEntity, nil
}

func (r saleorderDetailRepositoryDB) GetAllQtyRates() ([]models.SaleOrderDetailRate, error) {
	var qtyRatesFromDB []models.SaleOrderDetailRate
	err := r.db.Table("saleorderdetail").
		Select("sod_prod_code, SUM(sod_quantity) as quantity, SUM(sod_price) as price").
		Group("sod_prod_code").
		Order("quantity DESC").
		Scan(&qtyRatesFromDB).Error

	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return qtyRatesFromDB, nil
}

func (r saleorderDetailRepositoryDB) GetAllQtyRatesByDay(dateReq time.Time) ([]models.SaleOrderDetailRate, error) {
	var qtyRatesFromDB []models.SaleOrderDetailRate
	err := r.db.Table("saleorderdetail").
		Select("saleorderdetail.sod_prod_code, SUM(saleorderdetail.sod_quantity) as quantity, SUM(saleorderdetail.sod_price) as price").
		Joins("INNER JOIN saleorder ON saleorderdetail.sod_so_code = saleorder.saleorder_code").
		Where("DATE(saleorder.saleorder_created_at) = ?", dateReq.Format("2006-01-02")).
		Group("saleorderdetail.sod_prod_code").
		Order("quantity DESC").
		Scan(&qtyRatesFromDB).Error

	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return qtyRatesFromDB, nil
}

func (r saleorderDetailRepositoryDB) GetAllQtyRatesByMonth(dateReq time.Time) ([]models.SaleOrderDetailRate, error) {
	var qtyRatesFromDB []models.SaleOrderDetailRate
	err := r.db.Table("saleorderdetail").
		Select("saleorderdetail.sod_prod_code, SUM(saleorderdetail.sod_quantity) as quantity, SUM(saleorderdetail.sod_price) as price").
		Joins("INNER JOIN saleorder ON saleorderdetail.sod_so_code = saleorder.saleorder_code").
		Where("EXTRACT(MONTH FROM saleorder.saleorder_created_at) = ? AND EXTRACT(YEAR FROM saleorder.saleorder_created_at) = ?", dateReq.Month(), dateReq.Year()).
		Group("saleorderdetail.sod_prod_code").
		Order("quantity DESC").
		Scan(&qtyRatesFromDB).Error

	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return qtyRatesFromDB, nil
}

func (r saleorderDetailRepositoryDB) GetAllQtyRatesByYear(dateReq time.Time) ([]models.SaleOrderDetailRate, error) {
	var qtyRatesFromDB []models.SaleOrderDetailRate
	err := r.db.Table("saleorderdetail").
		Select("saleorderdetail.sod_prod_code, SUM(saleorderdetail.sod_quantity) as quantity, SUM(saleorderdetail.sod_price) as price").
		Joins("INNER JOIN saleorder ON saleorderdetail.sod_so_code = saleorder.saleorder_code").
		Where("EXTRACT(YEAR FROM saleorder.saleorder_created_at) = ?", dateReq.Year()).
		Group("saleorderdetail.sod_prod_code").
		Order("quantity DESC").
		Scan(&qtyRatesFromDB).Error

	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return qtyRatesFromDB, nil
}

func (r saleorderDetailRepositoryDB) GetAllPriceRates() ([]models.SaleOrderDetailRate, error) {
	var priceRatesFromDB []models.SaleOrderDetailRate
	err := r.db.Table("saleorderdetail").
		Select("sod_prod_code, SUM(sod_quantity) as quantity, SUM(sod_price) as price").
		Group("sod_prod_code").
		Order("price DESC").
		Scan(&priceRatesFromDB).Error

	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return priceRatesFromDB, nil
}

func (r saleorderDetailRepositoryDB) GetAllPriceRatesByDay(dateReq time.Time) ([]models.SaleOrderDetailRate, error) {
	var priceRatesFromDB []models.SaleOrderDetailRate
	err := r.db.Table("saleorderdetail").
		Select("saleorderdetail.sod_prod_code, SUM(saleorderdetail.sod_quantity) as quantity, SUM(saleorderdetail.sod_price) as price").
		Joins("INNER JOIN saleorder ON saleorderdetail.sod_so_code = saleorder.saleorder_code").
		Where("DATE(saleorder.saleorder_created_at) = ?", dateReq.Format("2006-01-02")).
		Group("saleorderdetail.sod_prod_code").
		Order("price DESC").
		Scan(&priceRatesFromDB).Error

	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return priceRatesFromDB, nil
}

func (r saleorderDetailRepositoryDB) GetAllPriceRatesByMonth(dateReq time.Time) ([]models.SaleOrderDetailRate, error) {
	var priceRatesFromDB []models.SaleOrderDetailRate
	err := r.db.Table("saleorderdetail").
		Select("saleorderdetail.sod_prod_code, SUM(saleorderdetail.sod_quantity) as quantity, SUM(saleorderdetail.sod_price) as price").
		Joins("INNER JOIN saleorder ON saleorderdetail.sod_so_code = saleorder.saleorder_code").
		Where("EXTRACT(MONTH FROM saleorder.saleorder_created_at) = ? AND EXTRACT(YEAR FROM saleorder.saleorder_created_at) = ?", dateReq.Month(), dateReq.Year()).
		Group("saleorderdetail.sod_prod_code").
		Order("price DESC").
		Scan(&priceRatesFromDB).Error

	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return priceRatesFromDB, nil
}

func (r saleorderDetailRepositoryDB) GetAllPriceRatesByYear(dateReq time.Time) ([]models.SaleOrderDetailRate, error) {
	var priceRatesFromDB []models.SaleOrderDetailRate
	err := r.db.Table("saleorderdetail").
		Select("saleorderdetail.sod_prod_code, SUM(saleorderdetail.sod_quantity) as quantity, SUM(saleorderdetail.sod_price) as price").
		Joins("INNER JOIN saleorder ON saleorderdetail.sod_so_code = saleorder.saleorder_code").
		Where("EXTRACT(YEAR FROM saleorder.saleorder_created_at) = ?", dateReq.Year()).
		Group("saleorderdetail.sod_prod_code").
		Order("price DESC").
		Scan(&priceRatesFromDB).Error

	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return priceRatesFromDB, nil
}