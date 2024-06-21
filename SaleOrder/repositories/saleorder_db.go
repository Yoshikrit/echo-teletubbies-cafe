package repositories

import (
	"saleorder/models"
	"saleorder/utils/errs"
	
	"time"
    "gorm.io/gorm"
)

type saleorderRepositoryDB struct {
	db *gorm.DB
}

func NewSaleOrderRepositoryDB(db *gorm.DB) SaleOrderRepository {
	return saleorderRepositoryDB{db: db}
}

func (r saleorderRepositoryDB) Create(saleorderReq *models.SaleOrderCreate) (*models.SaleOrderEntity, error) {
	var userFromDB models.UserEntity
	if err := r.db.First(&userFromDB, saleorderReq.CreatedUser).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, errs.NewNotFoundError("No User from User Id")
		}
		return nil, errs.NewUnexpectedError(err.Error())
	}

	var saleorderEntity models.SaleOrderEntity
	saleorderEntity.CreatedUser = saleorderReq.CreatedUser
	saleorderEntity.CreatedDate = time.Now()
	saleorderEntity.TotalPrice = saleorderReq.TotalPrice
	saleorderEntity.Status = saleorderReq.Status
	saleorderEntity.PayMethod = saleorderReq.PayMethod

	if err := r.db.Create(&saleorderEntity).Error; err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return &saleorderEntity, nil
}

func (r saleorderRepositoryDB) GetAll() ([]models.SaleOrderReport, error) {
	var saleordersFromDB []models.SaleOrderEntity
	if err := r.db.Find(&saleordersFromDB).Error; err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	var saleordersReport []models.SaleOrderReport
	for i, saleorderFromDB := range saleordersFromDB {

		var userFromDB models.UserEntity
		if err := r.db.First(&userFromDB, saleorderFromDB.CreatedUser).Error; err != nil {
			if gorm.ErrRecordNotFound == err {
				return nil, errs.NewNotFoundError(err.Error())
			}
			return nil, errs.NewUnexpectedError(err.Error())
		}

		var payChannelFromDB models.PaymentMethodEntity
		if err := r.db.First(&payChannelFromDB, saleorderFromDB.PayMethod).Error; err != nil {
			if gorm.ErrRecordNotFound == err {
				return nil, errs.NewNotFoundError(err.Error())
			}
			return nil, errs.NewUnexpectedError(err.Error())
		}
		
		name := userFromDB.FName + " " + userFromDB.LName
		saleorderRes := models.SaleOrderReport{
			Seq:      		i + 1,
			Id: 		    saleorderFromDB.Id,
			User:           name,
			Date: 		    saleorderFromDB.CreatedDate,
			TotalPrice: 	saleorderFromDB.TotalPrice,
			Status: 	    saleorderFromDB.Status,
			PayMethodName: 	payChannelFromDB.Name,
		}
		saleordersReport = append(saleordersReport, saleorderRes)
	}

	return saleordersReport, nil
}

func (r saleorderRepositoryDB) GetAllByDay(dateReq time.Time) ([]models.SaleOrderReport, error) {
	var saleordersFromDB []models.SaleOrderEntity
	if err := r.db.Where("DATE(saleorder_created_at) = ?", dateReq.Format("2006-01-02")).Find(&saleordersFromDB).Error; err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	var saleordersReport []models.SaleOrderReport
	for i, saleorderFromDB := range saleordersFromDB {

		var userFromDB models.UserEntity
		if err := r.db.First(&userFromDB, saleorderFromDB.CreatedUser).Error; err != nil {
			if gorm.ErrRecordNotFound == err {
				return nil, errs.NewNotFoundError(err.Error())
			}
			return nil, errs.NewUnexpectedError(err.Error())
		}

		var payChannelFromDB models.PaymentMethodEntity
		if err := r.db.First(&payChannelFromDB, saleorderFromDB.PayMethod).Error; err != nil {
			if gorm.ErrRecordNotFound == err {
				return nil, errs.NewNotFoundError(err.Error())
			}
			return nil, errs.NewUnexpectedError(err.Error())
		}
		
		name := userFromDB.FName + " " + userFromDB.LName
		saleorderRes := models.SaleOrderReport{
			Seq:      		i + 1,
			Id: 		    saleorderFromDB.Id,
			User:           name,
			Date: 		    saleorderFromDB.CreatedDate,
			TotalPrice: 	saleorderFromDB.TotalPrice,
			Status: 	    saleorderFromDB.Status,
			PayMethodName: 	payChannelFromDB.Name,
		}
		saleordersReport = append(saleordersReport, saleorderRes)
	}

	return saleordersReport, nil
}

func (r saleorderRepositoryDB) GetAllByMonth(dateReq time.Time) ([]models.SaleOrderReport, error) {
	var saleordersFromDB []models.SaleOrderEntity

	month := dateReq.Month()
	year := dateReq.Year()

	if err := r.db.Where("EXTRACT(MONTH FROM saleorder_created_at) = ? AND EXTRACT(YEAR FROM saleorder_created_at) = ?", month, year).Find(&saleordersFromDB).Error; err != nil {
        return nil, errs.NewUnexpectedError(err.Error())
    }

	var saleordersReport []models.SaleOrderReport
	for i, saleorderFromDB := range saleordersFromDB {

		var userFromDB models.UserEntity
		if err := r.db.First(&userFromDB, saleorderFromDB.CreatedUser).Error; err != nil {
			if gorm.ErrRecordNotFound == err {
				return nil, errs.NewNotFoundError(err.Error())
			}
			return nil, errs.NewUnexpectedError(err.Error())
		}

		var payChannelFromDB models.PaymentMethodEntity
		if err := r.db.First(&payChannelFromDB, saleorderFromDB.PayMethod).Error; err != nil {
			if gorm.ErrRecordNotFound == err {
				return nil, errs.NewNotFoundError(err.Error())
			}
			return nil, errs.NewUnexpectedError(err.Error())
		}
		
		name := userFromDB.FName + " " + userFromDB.LName
		saleorderRes := models.SaleOrderReport{
			Seq:      		i + 1,
			Id: 		    saleorderFromDB.Id,
			User:           name,
			Date: 		    saleorderFromDB.CreatedDate,
			TotalPrice: 	saleorderFromDB.TotalPrice,
			Status: 	    saleorderFromDB.Status,
			PayMethodName: 	payChannelFromDB.Name,
		}
		saleordersReport = append(saleordersReport, saleorderRes)
	}

	return saleordersReport, nil
}

func (r saleorderRepositoryDB) GetAllByYear(dateReq time.Time) ([]models.SaleOrderReport, error) {
	var saleordersFromDB []models.SaleOrderEntity

	year := dateReq.Year()

	if err := r.db.Where("EXTRACT(YEAR FROM saleorder_created_at) = ?", year).Find(&saleordersFromDB).Error; err != nil {
        return nil, errs.NewUnexpectedError(err.Error())
    }

	var saleordersReport []models.SaleOrderReport
	for i, saleorderFromDB := range saleordersFromDB {

		var userFromDB models.UserEntity
		if err := r.db.First(&userFromDB, saleorderFromDB.CreatedUser).Error; err != nil {
			if gorm.ErrRecordNotFound == err {
				return nil, errs.NewNotFoundError(err.Error())
			}
			return nil, errs.NewUnexpectedError(err.Error())
		}

		var payChannelFromDB models.PaymentMethodEntity
		if err := r.db.First(&payChannelFromDB, saleorderFromDB.PayMethod).Error; err != nil {
			if gorm.ErrRecordNotFound == err {
				return nil, errs.NewNotFoundError(err.Error())
			}
			return nil, errs.NewUnexpectedError(err.Error())
		}
		
		name := userFromDB.FName + " " + userFromDB.LName
		saleorderRes := models.SaleOrderReport{
			Seq:      		i + 1,
			Id: 		    saleorderFromDB.Id,
			User:           name,
			Date: 		    saleorderFromDB.CreatedDate,
			TotalPrice: 	saleorderFromDB.TotalPrice,
			Status: 	    saleorderFromDB.Status,
			PayMethodName: 	payChannelFromDB.Name,
		}
		saleordersReport = append(saleordersReport, saleorderRes)
	}
	
	return saleordersReport, nil
}

func (r saleorderRepositoryDB) GetTotalPricePass() (float64, error) {
	var totalAmount float64
	err := r.db.Table("saleorder").
		Where("saleorder_status = ?", "Pass").
		Select("COALESCE(SUM(saleorder_total_price), 0)").
		Row().
		Scan(&totalAmount)
	if err != nil {
		return 0, errs.NewUnexpectedError(err.Error())
	}
	
	return totalAmount, nil
}

func (r saleorderRepositoryDB) GetTotalPricePassByDay(dateReq time.Time) (float64, error) {
	var totalAmount float64
	err := r.db.Table("saleorder").
		Where("DATE(saleorder_created_at) = ? AND saleorder_status = ?", dateReq.Format("2006-01-02"), "Pass").
		Select("COALESCE(SUM(saleorder_total_price), 0)").
		Row().
		Scan(&totalAmount)
	if err != nil {
		return 0, errs.NewUnexpectedError(err.Error())
	}
	
	return totalAmount, nil
}

func (r saleorderRepositoryDB) GetTotalPricePassByMonth(dateReq time.Time) (float64, error) {
	month := dateReq.Month()
	year := dateReq.Year()

	var totalAmount float64
	err := r.db.Table("saleorder").
		Where("EXTRACT(MONTH FROM saleorder_created_at) = ? AND EXTRACT(YEAR FROM saleorder_created_at) = ? AND saleorder_status = ?", month, year, "Pass").
		Select("COALESCE(SUM(saleorder_total_price), 0)").
		Row().
		Scan(&totalAmount)
	if err != nil {
		return 0, errs.NewUnexpectedError(err.Error())
	}
	
	return totalAmount, nil
}

func (r saleorderRepositoryDB) GetTotalPricePassByYear(dateReq time.Time) (float64, error) {
	year := dateReq.Year()

	var totalAmount float64
	err := r.db.Table("saleorder").
		Where("EXTRACT(YEAR FROM saleorder_created_at) = ? AND saleorder_status = ?", year, "Pass").
		Select("COALESCE(SUM(saleorder_total_price), 0)").
		Row().
		Scan(&totalAmount)
	if err != nil {
		return 0, errs.NewUnexpectedError(err.Error())
	}
	
	return totalAmount, nil
}