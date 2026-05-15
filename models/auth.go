package models

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func NewUser(userName, password string) *User {
	return &User{
		UserName: userName,
		Password: password,
	}
}
