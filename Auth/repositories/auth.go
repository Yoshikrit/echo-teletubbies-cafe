package repositories

import (
	"auth/models"
)

type AuthRepository interface {
	GetUserClaimByEmailAndPassword(*models.UserLogin) (*models.UserClaim, error)
	Login(int) error
	Logout(int) error
}