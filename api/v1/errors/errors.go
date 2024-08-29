package errors

import (
	"fmt"
	"gotalk/internal/json"
	"net/http"
)

var FAILED = func(s string) json.Json {
	return json.Json{
		Status: http.StatusInternalServerError,
		Message: fmt.Sprintf("Failed %s", s),
	}
}

var INVALID = func (s string) json.Json {
	return json.Json{
		Status: http.StatusBadRequest,
		Message: fmt.Sprintf("Invalid %s", s),
	}
}

var DUPLICATE = func (s string) json.Json {
	return json.Json{
		Status: http.StatusConflict,
		Message: fmt.Sprintf("%s already exists", s),
	}
}

var NOT_SET = func(s string) json.Json {
	return json.Json{
		Status: http.StatusBadRequest,
		Message: fmt.Sprintf("%s not set", s),
	}
}

var NOT_FOUND = func(s string) json.Json {
	return json.Json{
		Status: http.StatusNotFound,
		Message: fmt.Sprintf("%s not found", s),
	}
}

var UNAUTHORIZED = func() json.Json {
	return json.Json{
		Status: http.StatusUnauthorized,
		Message: "Unauthorized",
	}
}
