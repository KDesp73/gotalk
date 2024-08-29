package users

import (
	"fmt"
	"gotalk/internal/utils"
)

const (
	ADMIN = "admin"
	DEFAULT = "default"
)

type User struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Email string `json:"email"`
	Key string `json:"key"`
	SignUpTime string `json:"signuptime"`
}

func (u *User) Log() {
	fmt.Println(utils.JsonToString(u))
}
