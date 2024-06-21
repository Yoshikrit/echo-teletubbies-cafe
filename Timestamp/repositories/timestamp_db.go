package repositories

import (
	"timestamp/models"
	"timestamp/utils/errs"
	
	"time"
    "gorm.io/gorm"
)

type timestampRepositoryDB struct {
	db *gorm.DB
}

func NewTimestampRepositoryDB(db *gorm.DB) TimestampRepository {
	return timestampRepositoryDB{db: db}
}

func (r timestampRepositoryDB) GetAll() ([]models.TimestampReport, error) {
	var timestampsFromDB []models.TimestampEntity
	if err := r.db.Find(&timestampsFromDB).Error; err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	var timestampsReport []models.TimestampReport
	for i, timestampFromDB := range timestampsFromDB {

		var userFromDB models.UserEntity
		err := r.db.First(&userFromDB, timestampFromDB.UserId).Error
		if err != nil {
			if gorm.ErrRecordNotFound == err {
				return nil, errs.NewNotFoundError(err.Error())
			}
			return nil, errs.NewUnexpectedError(err.Error())
		}

		name := userFromDB.FName + " " + userFromDB.LName
		timestampRes := models.TimestampReport{
			Seq:      		i + 1,
			UserName: 		name,
			LoginAt: 		timestampFromDB.LoginAt,
			LogoutAt: 		timestampFromDB.LogoutAt,
			Hour: 			timestampFromDB.Hour,
		}
		timestampsReport = append(timestampsReport, timestampRes)
	}

	return timestampsReport, nil
}

func (r timestampRepositoryDB) GetAllByDay(dateReq time.Time) ([]models.TimestampReport, error) {
	var timestampsFromDB []models.TimestampEntity
	if err := r.db.Where("DATE(timestamp_login_at) = ?", dateReq.Format("2006-01-02")).Find(&timestampsFromDB).Error; err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	var timestampsReport []models.TimestampReport
	for i, timestampFromDB := range timestampsFromDB {
		var userFromDB models.UserEntity
		err := r.db.First(&userFromDB, timestampFromDB.UserId).Error
		if err != nil {
			if gorm.ErrRecordNotFound == err {
				return nil, errs.NewNotFoundError(err.Error())
			}
			return nil, errs.NewUnexpectedError(err.Error())
		}

		name := userFromDB.FName + " " + userFromDB.LName
		timestampRes := models.TimestampReport{
			Seq:      		i + 1,
			UserName: 		name,
			LoginAt: 		timestampFromDB.LoginAt,
			LogoutAt: 		timestampFromDB.LogoutAt,
			Hour: 			timestampFromDB.Hour,
		}
		timestampsReport = append(timestampsReport, timestampRes)
	}
	
	return timestampsReport, nil
}

func (r timestampRepositoryDB) GetAllByMonth(dateReq time.Time) ([]models.TimestampReport, error) {
	var timestampsFromDB []models.TimestampEntity

	month := dateReq.Month()
	year := dateReq.Year()

	if err := r.db.Where("EXTRACT(MONTH FROM timestamp_login_at) = ? AND EXTRACT(YEAR FROM timestamp_login_at) = ?", month, year).Find(&timestampsFromDB).Error; err != nil {
        return nil, errs.NewUnexpectedError(err.Error())
    }

	var timestampsReport []models.TimestampReport
	for i, timestampFromDB := range timestampsFromDB {
		var userFromDB models.UserEntity
		err := r.db.First(&userFromDB, timestampFromDB.UserId).Error
		if err != nil {
			if gorm.ErrRecordNotFound == err {
				return nil, errs.NewNotFoundError(err.Error())
			}
			return nil, errs.NewUnexpectedError(err.Error())
		}

		name := userFromDB.FName + " " + userFromDB.LName
		timestampRes := models.TimestampReport{
			Seq:      		i + 1,
			UserName: 		name,
			LoginAt: 		timestampFromDB.LoginAt,
			LogoutAt: 		timestampFromDB.LogoutAt,
			Hour: 			timestampFromDB.Hour,
		}
		timestampsReport = append(timestampsReport, timestampRes)
	}

	return timestampsReport, nil
}

func (r timestampRepositoryDB) GetAllByYear(dateReq time.Time) ([]models.TimestampReport, error) {
	var timestampsFromDB []models.TimestampEntity

	year := dateReq.Year()

	if err := r.db.Where("EXTRACT(YEAR FROM timestamp_login_at) = ?", year).Find(&timestampsFromDB).Error; err != nil {
        return nil, errs.NewUnexpectedError(err.Error())
    }

	var timestampsReport []models.TimestampReport
	for i, timestampFromDB := range timestampsFromDB {
		var userFromDB models.UserEntity
		err := r.db.First(&userFromDB, timestampFromDB.UserId).Error
		if err != nil {
			if gorm.ErrRecordNotFound == err {
				return nil, errs.NewNotFoundError(err.Error())
			}
			return nil, errs.NewUnexpectedError(err.Error())
		}
		
		name := userFromDB.FName + " " + userFromDB.LName
		timestampRes := models.TimestampReport{
			Seq:      		i + 1,
			UserName: 		name,
			LoginAt: 		timestampFromDB.LoginAt,
			LogoutAt: 		timestampFromDB.LogoutAt,
			Hour: 			timestampFromDB.Hour,
		}
		timestampsReport = append(timestampsReport, timestampRes)
	}
	
	return timestampsReport, nil
}
