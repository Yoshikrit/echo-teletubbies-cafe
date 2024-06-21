package services

import (
	"user/models"
	"user/repositories"

	"user/utils/errs"
	"user/utils/logs"
)

type roleService struct {
	roleRepo repositories.RoleRepository
}

func NewRoleService(roleRepo repositories.RoleRepository) RoleService {
	return roleService{roleRepo: roleRepo}
}

func (s roleService) CreateRole(roleReq *models.RoleCreate) (*models.Role, error) {
	if roleReq.Id == 0{
		return nil, errs.NewBadRequestError("Role's Id is null")
	}
	if roleReq.Name == ""{
		return nil, errs.NewBadRequestError("Role's Name is null")
	}

	roleEntityRes, err := s.roleRepo.Create(roleReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	
	roleRes := models.Role{
		Id:       		roleEntityRes.Id,
		Name:     		roleEntityRes.Name,
	}

	logs.Info("Service: Create Role Successfully")
	return &roleRes, nil
}

func (s roleService) GetRoles() ([]models.Role, error) {
	rolesFromDB, err := s.roleRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	var rolesRes []models.Role
	for _, roleFromDB := range rolesFromDB {
		roleRes := models.Role{
			Id:       		roleFromDB.Id,
			Name:     		roleFromDB.Name,
		}
		rolesRes = append(rolesRes, roleRes)
	}

	logs.Info("Service: Get Roles Successfully")
	return rolesRes, nil
}

func (s roleService) GetRole(id int) (*models.Role, error) {
	roleFromDB, err := s.roleRepo.GetById(id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	roleRes := models.Role{
		Id:       		roleFromDB.Id,
		Name:     		roleFromDB.Name,
	}

	logs.Info("Service: Get Role Successfully")
	return &roleRes, nil
}

func (s roleService) UpdateRole(id int, roleReq *models.RoleUpdate) (*models.Role, error) {
	if roleReq.Name == ""{
		return nil, errs.NewBadRequestError("Role's Name is null")
	}
	roleFromDB, err := s.roleRepo.Update(id, roleReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	roleRes := models.Role{
		Id:       		roleFromDB.Id,
		Name:     		roleFromDB.Name,
	}

	logs.Info("Service: Update Role Successfully")
	return &roleRes, nil
}

func (s roleService) DeleteRole(id int) (error) {
	err := s.roleRepo.DeleteById(id)
	if err != nil {
		logs.Error(err)
		return err
	}
	
	logs.Info("Service: Delete Role Successfully")
	return nil
}

func (s roleService) GetRoleCount() (int64, error) {
	count, err := s.roleRepo.GetCount()
	if err != nil {
		logs.Error(err)
		return 0, err
	}

	logs.Info("Service: Get Role's Count Successfully")
	return count, nil
}

