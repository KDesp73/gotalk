package users

import (
	"fmt"
	"gotalk/internal/utils"
)

const (
	USER_ADMIN = 1
	USER_DEFAULT = 2
)

type User struct {
	Type int `json:"type"`
	Name string `json:"name"`
	Email string `json:"email"`
	Key string `json:"key"`
	SignUpTime string `json:"signuptime"`
}

func (u *User) Log() {
	fmt.Println(utils.JsonToString(u))
}
