package errors

import (
	"fmt"
	"gotalk/internal/json"
	"net/http"
)

var STATUS_BAD_REQUEST = http.StatusBadRequest
var STATUS_CONFLICT = http.StatusConflict

var PARSING_FORM_FAILED = json.Json{
	Status: STATUS_BAD_REQUEST,
	Message: "Unable to parse form",
}

var INVALID_THREAD_ID = json.Json{
	Status: STATUS_BAD_REQUEST,
	Message: "Invalid thread id",
}

var INVALID_USER_ID = json.Json{
	Status: STATUS_BAD_REQUEST,
	Message: "Invalid user id",
}

var INVALID_EMAIL = json.Json{
	Status: STATUS_BAD_REQUEST,
	Message: "Invalid email",
}

var DUPLICATE_EMAIL = json.Json{
	Status: STATUS_CONFLICT,
	Message: "Email already exists",
}

var NOT_FOUND = func(s string) json.Json {
	return json.Json{
		Status: STATUS_BAD_REQUEST,
		Message: fmt.Sprintf("%s not found", s),
	}
}
