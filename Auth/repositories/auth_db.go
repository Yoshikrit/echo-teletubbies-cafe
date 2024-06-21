package repositories

import (
	"auth/models"
	"auth/utils/errs"
	"strconv"
	"time"

	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authRepositoryDB struct {
	db *gorm.DB
}

func NewAuthRepositoryDB(db *gorm.DB) AuthRepository {
	return authRepositoryDB{db: db}
}

func (r authRepositoryDB) GetUserClaimByEmailAndPassword(userReq *models.UserLogin) (*models.UserClaim, error) {
	var userFromDB models.UserEntity
	result1 := r.db.Where("user_email = ?", userReq.Email).First(&userFromDB)
	if result1.Error != nil {
		if errors.Is(result1.Error, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("Email or Password is not valid")
		}
		return nil, errs.NewUnexpectedError("Database error: " + result1.Error.Error())
	}

	err := bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(userReq.Password))
	if err != nil {
		return nil, errs.NewNotFoundError("Email or Password is not valid")
	}

	var roleFromDB models.RoleEntity
	result2 := r.db.Where("role_code = ?", userFromDB.Role_Id).First(&roleFromDB)
	if result2.Error != nil {
		if errors.Is(result2.Error, gorm.ErrRecordNotFound) {
			return nil, errs.NewNotFoundError("Role Id is not valid " + strconv.Itoa(roleFromDB.Id) + roleFromDB.Name)
		}
		return nil, errs.NewUnexpectedError("Database error: " + result2.Error.Error())
	}

	var userClaim models.UserClaim
	userClaim.Id = userFromDB.Id
	userClaim.Name = userFromDB.FName + " " + userFromDB.LName
	userClaim.Role = roleFromDB.Name

	return &userClaim, nil
}

func (r authRepositoryDB) Login(userId int) error {
	var userFromDB models.UserEntity
	if err := r.db.First(&userFromDB, userId).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return errs.NewNotFoundError("No User from User Id : " + strconv.Itoa(userId))
		}
		return errs.NewUnexpectedError(err.Error())
	}

	var timestampCreate models.TimestampEntity
	timestampCreate.UserId = userId
	timestampCreate.LoginAt = time.Now()

	if err := r.db.Create(&timestampCreate).Error; err != nil {
		return errs.NewUnexpectedError(err.Error())
	}

	return nil
}

func (r authRepositoryDB) Logout(userId int) error {
	var userFromDB models.UserEntity
	if err := r.db.First(&userFromDB, userId).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return errs.NewNotFoundError("No User from User Id : " + strconv.Itoa(userId))
		}
		return errs.NewUnexpectedError(err.Error())
	}

	var timestampFromDB models.TimestampEntity
	if err := r.db.
		Where("timestamp_user_code = ?", userId).
		Order("timestamp_code desc").
		First(&timestampFromDB).
		Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return errs.NewNotFoundError("No Timestamp with User Id : " + strconv.Itoa(userId))
		}
		return errs.NewNotFoundError(err.Error())
	}
	timestampFromDB.LogoutAt = time.Now()

	layout := "2006-01-02 15:04:05"
	loginTimeStr := timestampFromDB.LoginAt.Format(layout)
	logoutTimeStr := timestampFromDB.LogoutAt.Format(layout)
	loginTime, _ := time.Parse(layout, loginTimeStr)
	logoutTime, _ := time.Parse(layout, logoutTimeStr)

	timestampFromDB.Hour = int(logoutTime.Sub(loginTime).Hours())

	if err := r.db.Save(timestampFromDB).Error; err != nil {
		return errs.NewUnexpectedError(err.Error())
	}

	return nil
}
