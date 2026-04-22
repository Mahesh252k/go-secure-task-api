package serviceauth

import (
	"CRUD_API_PROJ/models"
	"CRUD_API_PROJ/repository"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"gorm.io/gorm"
)

type serviceConcrete struct {
	DB gorm.DB
}

func NewServiceConcrete() AuthService {
	return &serviceConcrete{
		DB: *repository.GetDB(),
	}
}

func (s *serviceConcrete) Register(request models.User) error {
	err := repository.CreateUser(&repository.UserRepo{
		UserName: request.UserName,
		Password: request.Password,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceConcrete) Login(userName, password string) (string, error) {
	user, err := repository.GetUser(userName)

	if err != nil {
		return "", err
	}

	pass, err := base64.StdEncoding.DecodeString(user.Password)
	if err != nil {
		return "", err
	}

	if string(pass) != password {
		return "", fmt.Errorf("invalid credentials")
	}

	token := sha256.Sum256([]byte(userName + user.Password))
	tok := hex.EncodeToString(token[:])
	return tok, nil
}
