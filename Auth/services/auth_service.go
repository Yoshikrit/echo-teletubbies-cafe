package services

import (
	"auth/models"
	"auth/repositories"
	"auth/utils/errs"
	"auth/utils/logs"
	
	"github.com/dgrijalva/jwt-go"
	"time"
    "os"
)

type authService struct {
	authRepo repositories.AuthRepository
}

func NewAuthService(authRepo repositories.AuthRepository) AuthService {
	return authService{authRepo: authRepo}
}

func (s authService) Login(userReq *models.UserLogin) (*models.Response, error) {
	if userReq.Email == "" || userReq.Password == ""{
		logs.Error("Service : User's Email or Password is null")
		return nil, errs.NewBadRequestError("User's Email or Password is null")
	}
	userClaimRes, err := s.authRepo.GetUserClaimByEmailAndPassword(userReq)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	claims := jwt.StandardClaims{
		Subject:   userReq.Email,
		Issuer:    userClaimRes.Name,
		IssuedAt:  time.Now().UTC().Unix(),
		Audience:  userClaimRes.Role,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	if err := s.authRepo.Login(userClaimRes.Id); err != nil {
		logs.Error("Service: Log in Failed")
		return nil, errs.NewUnexpectedError(err.Error())
	}

	response := models.Response{
		Id: userClaimRes.Id,
		Name: userClaimRes.Name,
		Role: userClaimRes.Role,
		JWT: token,
	}

	logs.Info("Service: Log in Successfully")
	return &response, nil
}

func  (s authService) Logout(id int) error {
	if err := s.authRepo.Logout(id); err != nil {
		logs.Error("Service: Log out Failed")
		return errs.NewUnexpectedError(err.Error())
	}

	logs.Info("Service: Log out Successfully")
	return nil
}