package repository

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type UserRepo struct {
	UserName string `gorm:"primaryKey"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`
}

func CreateUserTable() {
	db := GetDB()
	err := db.AutoMigrate(&UserRepo{})
	if err != nil {
		panic(err)
	}
}

func CreateUser(user *UserRepo) error {
	pass := base64.StdEncoding.EncodeToString([]byte(user.Password))
	user.Password = pass
	result := GetDB().Create(user)
	return result.Error
}

func GetUser(userName string) (*UserRepo, error) {
	var user UserRepo
	result := GetDB().First(&user, "user_name = ?", userName)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserRoles(userName string) (string, error) {
	var user UserRepo
	result := GetDB().First(&user, "user_name = ?", userName)
	if result.Error != nil {
		return "", result.Error
	}
	return user.Role, nil
}

func ValidateUser(userName, password string) bool {
	user, err := GetUser(userName)
	if err != nil {
		return false
	}
	pass := base64.StdEncoding.EncodeToString([]byte(password))
	return user.Password == pass
}

func ValidateToken(token, username string) (bool, error) {
	user, _ := GetTaskByID(username)
	usr, err := GetUser(user.UserName)
	if err != nil {
		return false, err
	}
	tok := sha256.Sum256([]byte(usr.UserName + usr.Password))
	tok1 := hex.EncodeToString(tok[:])
	fmt.Println(tok1, tok1)
	fmt.Print(token, token)
	if tok1 != token {
		return false, nil
	}
	return true, nil
}
