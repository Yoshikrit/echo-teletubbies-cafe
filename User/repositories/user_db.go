package repositories

import (
	"user/models"
	"user/utils/errs"
    "gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) UserRepository {
	return userRepositoryDB{db: db}
}


func (r userRepositoryDB) Create(userReq *models.UserCreate) (*models.UserEntity, error) {
	var existProdEntity models.UserEntity
	err1 := r.db.Where("user_email = ?", userReq.Email).First(&existProdEntity).Error
	if err1 == nil {
		return nil, errs.NewConflictError("User with the same email already exists")
	}

	encryptPassword, err2 := bcrypt.GenerateFromPassword([]byte(userReq.Password), 10)
	if err2 != nil {
		return nil, errs.NewUnprocessableError(err2.Error())
	}

	var userEntity models.UserEntity
	userEntity.Id = userReq.Id
	userEntity.Role_Id = userReq.Role_Id
	userEntity.FName = userReq.FName
	userEntity.LName = userReq.LName
	userEntity.Email = userReq.Email
	userEntity.Password = string(encryptPassword)
	userEntity.Sex = userReq.Sex
	userEntity.TelNo = userReq.TelNo
	userEntity.Salary = userReq.Salary
	userEntity.Address = userReq.Address
	userEntity.WorkStatus = userReq.WorkStatus
	userEntity.BirthDate = userReq.BirthDate

	var roleFromDB models.RoleEntity
	err3 := r.db.First(&roleFromDB, userEntity.Role_Id).Error
	if err3 != nil {
		if gorm.ErrRecordNotFound == err3 {
			return nil, errs.NewNotFoundError("No Role from Role Id")
		}
		return nil, errs.NewUnexpectedError(err3.Error())
	}

	err4 := r.db.Create(&userEntity).Error
	if err4 != nil {
		return nil, errs.NewUnexpectedError(err4.Error())
	}

	return &userEntity, nil
}

func (r userRepositoryDB) GetAll() ([]models.UserEntity, error) {
	var userFromDB []models.UserEntity
	err := r.db.Find(&userFromDB).Error
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return userFromDB, nil
}

func (r userRepositoryDB) GetById(id int) (*models.UserEntity, error) {
	var userFromDB models.UserEntity
	err := r.db.First(&userFromDB, id).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, errs.NewNotFoundError(err.Error())
		}
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return &userFromDB, nil
}

func (r userRepositoryDB) Update(id int, updateUser *models.UserUpdate) (*models.UserEntity, error) {
	userFromDB, err1 := r.GetById(id)
	if err1 != nil {
		return nil, errs.NewNotFoundError(err1.Error())
	}

	encryptPassword, err2 := bcrypt.GenerateFromPassword([]byte(updateUser.Password), 10)
	if err2 != nil {
		return nil, errs.NewUnprocessableError(err2.Error())
	}

	if updateUser.Role_Id != 0 {
		userFromDB.Role_Id = updateUser.Role_Id
	}
	
	if updateUser.FName != "" {
		userFromDB.FName = updateUser.FName
	}
	
	if updateUser.LName != "" {
		userFromDB.LName = updateUser.LName
	}
	
	if updateUser.Email != "" {
		userFromDB.Email = updateUser.Email
	}
	
	if updateUser.Password != "" {
		userFromDB.Password = string(encryptPassword)
	}
	
	if updateUser.Sex != "" {
		userFromDB.Sex = updateUser.Sex
	}
	
	if updateUser.TelNo != "" {
		userFromDB.TelNo = updateUser.TelNo
	}
	
	if updateUser.Salary != 0 {
		userFromDB.Salary = updateUser.Salary
	}
	
	if updateUser.Address != "" {
		userFromDB.Address = updateUser.Address
	}
	
	if updateUser.WorkStatus != "" {
		userFromDB.WorkStatus = updateUser.WorkStatus
	}
	
	if !updateUser.BirthDate.IsZero() {
		userFromDB.BirthDate = updateUser.BirthDate
	}

	var roleFromDB models.RoleEntity
	err3 := r.db.First(&roleFromDB, userFromDB.Role_Id).Error
	if err3 != nil {
		if gorm.ErrRecordNotFound == err3 {
			return nil, errs.NewNotFoundError("No Role from Role Id")
		}
		return nil, errs.NewUnexpectedError(err3.Error())
	}

	if err4 := r.db.Save(userFromDB).Error; err4 != nil {
		return nil, errs.NewUnexpectedError(err4.Error())
	}

	return userFromDB, nil
}

func (r userRepositoryDB) DeleteById(id int) error {
	_, err1 := r.GetById(id)
	if err1 != nil {
		return errs.NewNotFoundError(err1.Error())
	}

	if err2 := r.db.Delete(&models.UserEntity{}, id).Error; err2 != nil {
		return errs.NewUnexpectedError(err2.Error())
	}
	return nil
}

func (r userRepositoryDB) GetCount() (int64, error) {
	var count int64
	err := r.db.Table("user").Count(&count).Error
	if err != nil {
		return 0, errs.NewUnexpectedError(err.Error())
	}
	return count, nil
}






