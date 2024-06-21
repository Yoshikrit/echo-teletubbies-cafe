package services

import (
	"user/models"
	"user/repositories"

	"user/utils/errs"
	"user/utils/logs"
)

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return userService{userRepo: userRepo}
}

func (s userService) CreateUser(userReq *models.UserCreate) (*models.User, error) {
	if userReq.Id == 0 {
		return nil, errs.NewBadRequestError("User's Id is null")
	}
	if userReq.Role_Id == 0 {
		return nil, errs.NewBadRequestError("User's RoleId is null")
	}
	if userReq.Email == "" {
		return nil, errs.NewBadRequestError("User's Email is null")
	}
	if userReq.Password == "" {
		return nil, errs.NewBadRequestError("User's Password is null")
	}
	if userReq.Sex == "" {
		return nil, errs.NewBadRequestError("User's Sex is null")
	}
	if userReq.Salary == 0 {
		return nil, errs.NewBadRequestError("User's Salary is null")
	}
	if userReq.WorkStatus == "" {
		return nil, errs.NewBadRequestError("User's WorkStatus is null")
	}
	if userReq.BirthDate.IsZero(){
		return nil, errs.NewBadRequestError("User's BirthDate is null")
	}

	userEntityRes, err := s.userRepo.Create(userReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	userRes := models.User{
		Id:      		userEntityRes.Id,
		Role_Id: 		userEntityRes.Role_Id,
		FName:   		userEntityRes.FName,
		LName:   		userEntityRes.LName,
		Email:   		userEntityRes.Email,
		Sex:     		userEntityRes.Sex,
		TelNo:   		userEntityRes.TelNo,
		Salary:  		userEntityRes.Salary,
		Address: 		userEntityRes.Address,
		WorkStatus:  	userEntityRes.WorkStatus,
		BirthDate: 		userEntityRes.BirthDate,
	}

	logs.Info("Service: Create User Successfully")
	return &userRes, nil
}

func (s userService) GetUsers() ([]models.User, error) {
	usersFromDB, err := s.userRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	var usersRes []models.User
	for _, userFromDB := range usersFromDB {
		userRes := models.User{
			Id:      		userFromDB.Id,
			Role_Id: 		userFromDB.Role_Id,
			FName:   		userFromDB.FName,
			LName:   		userFromDB.LName,
			Email:   		userFromDB.Email,
			Sex:     		userFromDB.Sex,
			TelNo:   		userFromDB.TelNo,
			Salary:  		userFromDB.Salary,
			Address: 		userFromDB.Address,
			WorkStatus:  	userFromDB.WorkStatus,
			BirthDate: 		userFromDB.BirthDate,
		}
		usersRes = append(usersRes, userRes)
	}

	logs.Info("Service: Get Users Successfully")
	return usersRes, nil
}

func (s userService) GetUser(id int) (*models.User, error) {
	userFromDB, err := s.userRepo.GetById(id)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	userRes := models.User{
		Id:      		userFromDB.Id,
		Role_Id: 		userFromDB.Role_Id,
		FName:   		userFromDB.FName,
		LName:   		userFromDB.LName,
		Email:   		userFromDB.Email,
		Sex:     		userFromDB.Sex,
		TelNo:   		userFromDB.TelNo,
		Salary:  		userFromDB.Salary,
		Address: 		userFromDB.Address,
		WorkStatus:  	userFromDB.WorkStatus,
		BirthDate: 		userFromDB.BirthDate,
	}

	logs.Info("Service: Get User Successfully")
	return &userRes, nil
}

func (s userService) UpdateUser(id int, userReq *models.UserUpdate) (*models.User, error) {
	if (userReq.Role_Id == 0 && userReq.FName == "" && userReq.LName == "" && userReq.Email == "" && userReq.Password == "" && userReq.Sex == "" && userReq.TelNo == "" && userReq.Salary == 0 && userReq.Address == "" && userReq.WorkStatus == "" && userReq.BirthDate.IsZero()) {
		return nil, errs.NewBadRequestError("User's Data is null")
	}
	
	userFromDB, err := s.userRepo.Update(id, userReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	userRes := models.User{
		Id:      		userFromDB.Id,
		Role_Id: 		userFromDB.Role_Id,
		FName:   		userFromDB.FName,
		LName:   		userFromDB.LName,
		Email:   		userFromDB.Email,
		Sex:     		userFromDB.Sex,
		TelNo:   		userFromDB.TelNo,
		Salary:  		userFromDB.Salary,
		Address: 		userFromDB.Address,
		WorkStatus:  	userFromDB.WorkStatus,
		BirthDate: 		userFromDB.BirthDate,
	}

	logs.Info("Service: Update User Successfully")
	return &userRes, nil
}

func (s userService) DeleteUser(id int) error {
	err := s.userRepo.DeleteById(id)
	if err != nil {
		logs.Error(err)
		return err
	}

	logs.Info("Service: Delete User Successfully")
	return nil
}

func (s userService) GetUserCount() (int64, error) {
	count, err := s.userRepo.GetCount()
	if err != nil {
		logs.Error(err)
		return 0, err
	}

	logs.Info("Service: Get User's Count Successfully")
	return count, nil
}
