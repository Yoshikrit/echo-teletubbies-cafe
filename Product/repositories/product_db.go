package repositories

import (
	"product/models"

	"gorm.io/gorm"

	"product/utils/errs"
	"time"
)

type productRepositoryDB struct {
	db *gorm.DB
}

func NewProductRepositoryDB(db *gorm.DB) ProductRepository {
	return productRepositoryDB{db: db}
}

func (r productRepositoryDB) Create(prodReq *models.ProductCreate) (*models.ProductEntity, error) {
	var existProdEntity models.ProductEntity
	if err := r.db.Where("prod_name = ?", prodReq.Name).First(&existProdEntity).Error; err == nil {
		return nil, errs.NewConflictError("Product with the same name already exists")
	}

	var prodTypeFromDB models.ProductTypeEntity
	if err := r.db.First(&prodTypeFromDB, prodReq.ProdType_Id).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, errs.NewNotFoundError("ProductType is not exists")
		}
		return nil, errs.NewUnexpectedError(err.Error())
	}

	var userFromDB models.UserEntity
	if err := r.db.First(&userFromDB, prodReq.CreatedUser).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, errs.NewNotFoundError("CreatedUser's Id is not exists")
		}
		return nil, errs.NewUnexpectedError(err.Error())
	}

	var prodEntity models.ProductEntity
	prodEntity.Id = prodReq.Id
	prodEntity.ProdType_Id = prodReq.ProdType_Id
	prodEntity.Name = prodReq.Name
	prodEntity.Desc = prodReq.Desc
	prodEntity.Price = prodReq.Price
	prodEntity.Discount = prodReq.Discount
	prodEntity.CreatedAt = time.Now()
	prodEntity.CreatedUser = prodReq.CreatedUser
	prodEntity.UpdatedAt = time.Now()
	prodEntity.UpdatedUser = prodReq.CreatedUser

	err4 := r.db.Create(&prodEntity).Error
	if err4 != nil {
		return nil, errs.NewUnexpectedError(err4.Error())
	}

	return &prodEntity, nil
}

func (r productRepositoryDB) GetAll() ([]models.ProductEntity, error) {
	var prodFromDB []models.ProductEntity
	err := r.db.Find(&prodFromDB).Error
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return prodFromDB, nil
}

func (r productRepositoryDB) GetById(id int) (*models.ProductEntity, error) {
	var prodFromDB models.ProductEntity
	err := r.db.First(&prodFromDB, id).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, errs.NewNotFoundError(err.Error())
		}
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return &prodFromDB, nil
}

func (r productRepositoryDB) Update(id int, updateProd *models.ProductUpdate) (*models.ProductEntity, error) {
	prodFromDB, err := r.GetById(id)
	if err != nil {
		return nil, errs.NewNotFoundError(err.Error())
	}

	var prodTypeFromDB models.ProductTypeEntity
	if err := r.db.First(&prodTypeFromDB, prodFromDB.ProdType_Id).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, errs.NewNotFoundError("ProductType is not exists")
		}
		return nil, errs.NewUnexpectedError(err.Error())
	}

	var userFromDB models.UserEntity
	if err := r.db.First(&userFromDB, prodFromDB.UpdatedUser).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, errs.NewNotFoundError("UpdatedUser's Id is not exists")
		}
		return nil, errs.NewUnexpectedError(err.Error())
	}

	if err := r.db.Save(prodFromDB).Error; err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return prodFromDB, nil
}

func (r productRepositoryDB) DeleteById(id int) error {
	_, err := r.GetById(id)
	if err != nil {
		return errs.NewNotFoundError(err.Error())
	}

	if err := r.db.Delete(&models.ProductEntity{}, id).Error; err != nil {
		return errs.NewUnexpectedError(err.Error())
	}
	return nil
}

func (r productRepositoryDB) GetCount() (int64, error) {
	var count int64
	err := r.db.Table("product").Count(&count).Error
	if err != nil {
		return 0, errs.NewUnexpectedError(err.Error())
	}
	return count, nil
}
