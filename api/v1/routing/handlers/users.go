package handlers

import (
	"fmt"
	"gotalk/api/v1/errors"
	"gotalk/api/v1/response"
	"gotalk/internal/json"
	"gotalk/internal/state"
	"gotalk/internal/users"
	"gotalk/internal/utils"
	"net/http"
	"strings"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		prefix := "Bearer "

		fmt.Println("Here")
		if !strings.HasPrefix(authorization, prefix) {
			response.Error(w, errors.UNAUTHORIZED())
			return
		}

		encodedToken := strings.TrimPrefix(authorization, prefix)
		encodedToken = strings.TrimSpace(encodedToken)
		userid := r.PathValue("userid")
		userid = strings.TrimSpace(userid)

		fmt.Printf("userid: %s\nbearer: %s\n", userid, encodedToken)

		if userid == "" {
			response.Error(w, errors.NOT_SET("User id"))
			return
		}

		if !state.Instance.Users.IsAdmin(encodedToken) && userid != encodedToken {
			response.Error(w, errors.UNAUTHORIZED())
			return
		}

		succ := state.Instance.Users.RemoveUser(userid)

		if !succ {
			response.Error(w, errors.FAILED("deleting user"))
			return
		}

		response.Success(w, json.Json{
			Status: 204,
			Message: "User deleted successfully",
		})
	
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	j := json.Json {
		Status: 200,
		Message: "Users retreived successfully",
		Data: json.NestedJson {},
	}

	for _, user := range state.Instance.Users.Items {
		j.Data.Users= append(j.Data.Users, *user)
	}

	response.Success(w, j)
}

func IsAdmin(w http.ResponseWriter, r *http.Request){
	user := r.PathValue("user")

	response.Success(w, json.Json{
		Status: 200,
		Message: fmt.Sprintf("%s you are an admin!", user),
	})
}

func UndoSudo(w http.ResponseWriter, r *http.Request) {
	userid := r.PathValue("userid")
	userid = strings.TrimSpace(userid)

	succ := state.Instance.Users.Sudo(userid, true)

	if !succ {
		response.Error(w, errors.FAILED("revoking admin privileges"))
		return
	}

	response.Success(w, json.Json{
		Status: 200,
		Message: "Admin privileges revoked",
	})
}

func Sudo(w http.ResponseWriter, r *http.Request) {
	userid := r.PathValue("userid")
	userid = strings.TrimSpace(userid)

	succ := state.Instance.Users.Sudo(userid, false)

	if !succ {
		response.Error(w, errors.FAILED("granding admin privileges"))
		return
	}

	response.Success(w, json.Json{
		Status: 200,
		Message: "Admin privileges granted",
	})
}

func Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		response.Error(w, errors.FAILED("parsing form"))
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	if strings.TrimSpace(name) == "" {
		response.Error(w, errors.NOT_SET("Name"))
		return
	}
	if strings.TrimSpace(email) == "" {
		response.Error(w, errors.NOT_SET("Email"))
		return
	}
	if !utils.IsValidEmail(email) {
		response.Error(w, errors.INVALID("email"))
		return
	}
	if state.Instance.Users.NameExists(name) {
		response.Error(w, errors.DUPLICATE("name"))
		return
	}
	if state.Instance.Users.EmailExists(email) {
		response.Error(w, errors.DUPLICATE("email"))
		return
	}

	key := state.Instance.Users.PushUser(&users.User{
		Name: name,
		Email: email,
		Type: users.DEFAULT,
		SignUpTime: utils.CurrentTimestamp(),
	})

	response.Success(w, json.Json{
		Status: 201,
		Message: "Registration complete",
		Data: json.NestedJson{
			Key: key,
		},
	})
}
