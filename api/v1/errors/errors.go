package errors

import (
	"fmt"
	"gotalk/internal/json"
	"net/http"
)

var STATUS_CODE = http.StatusBadRequest

var PARSING_FORM_FAILED = json.Json{
	Status: STATUS_CODE,
	Message: "Unable to parse form",
}

var INVALID_THREAD_ID = json.Json{
	Status: STATUS_CODE,
	Message: "Invalid thread id",
}

var NOT_FOUND = func(s string) json.Json {
	return json.Json{
		Status: STATUS_CODE,
		Message: fmt.Sprintf("%s not found", s),
	}
}
