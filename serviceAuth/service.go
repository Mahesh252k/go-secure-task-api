package serviceauth

import (
	"CRUD_API_PROJ/models"
)

type AuthService interface {
	Register(request models.User) error
	Login(userName, password string) (string, error)
}
