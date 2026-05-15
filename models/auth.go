package models

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func NewUser(userName, password, role string) *User {
	return &User{
		UserName: userName,
		Password: password,
		Role:     role,
	}
}
