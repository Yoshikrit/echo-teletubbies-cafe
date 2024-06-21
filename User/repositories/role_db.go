package repositories

import (
	"user/models"
    "gorm.io/gorm"

	"user/utils/errs"
)

type roleRepositoryDB struct {
	db *gorm.DB
}

func NewRoleRepositoryDB(db *gorm.DB) RoleRepository {
	return roleRepositoryDB{db: db}
}

func (r roleRepositoryDB) Create(roleReq *models.RoleCreate) (*models.RoleEntity, error) {
	var existRoleEntity models.RoleEntity
	err1 := r.db.Where("role_name = ?", roleReq.Name).First(&existRoleEntity).Error
	if err1 == nil {
		return nil, errs.NewConflictError("Role with the same name already exists")
	}

	var roleEntity models.RoleEntity
	roleEntity.Name = roleReq.Name

	err2 := r.db.Create(&roleEntity).Error
	if err2 != nil {
		return nil, errs.NewUnexpectedError(err2.Error())
	}

	return &roleEntity, nil
}


func (r roleRepositoryDB) GetAll() ([]models.RoleEntity, error) {
	var roleFromDB []models.RoleEntity
	err := r.db.Find(&roleFromDB).Error
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return roleFromDB, nil
}

func (r roleRepositoryDB) GetById(id int) (*models.RoleEntity, error) {
	var roleFromDB models.RoleEntity
	err := r.db.First(&roleFromDB, id).Error
	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, errs.NewNotFoundError(err.Error())
		}
		return nil, errs.NewUnexpectedError(err.Error())
	}

	return &roleFromDB, nil
}

func (r roleRepositoryDB) Update(id int, updateRole *models.RoleUpdate) (*models.RoleEntity, error) {
	roleFromDB, err1 := r.GetById(id)
	if err1 != nil {
		return nil, errs.NewNotFoundError(err1.Error())
	}

	roleFromDB.Name = updateRole.Name

	if err2 := r.db.Save(roleFromDB).Error; err2 != nil {
		return nil, errs.NewUnexpectedError(err2.Error())
	}

	return roleFromDB, nil
}

func (r roleRepositoryDB) DeleteById(id int) error {
	_, err1 := r.GetById(id)
	if err1 != nil {
		return errs.NewNotFoundError(err1.Error())
	}

	if err2 := r.db.Delete(&models.RoleEntity{}, id).Error; err2 != nil {
		return errs.NewUnexpectedError(err2.Error())
	}
	return nil
}

func (r roleRepositoryDB) GetCount() (int64, error) {
	var count int64
	err := r.db.Table("role").Count(&count).Error
	if err != nil {
		return 0, errs.NewUnexpectedError(err.Error())
	}
	return count, nil
}






