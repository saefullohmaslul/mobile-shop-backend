package response

import (
	"net/http"
)

// NotFound is response empty data
func NotFound(message string, errors []Error) {
	response := Response{
		Status:  http.StatusNotFound,
		Message: message,
		Data:    nil,
		Errors:  errors,
	}
	panic(response)
}

// BadRequest is response for bad request
func BadRequest(message string, errors []Error) {
	response := Response{
		Status:  http.StatusBadRequest,
		Message: message,
		Data:    nil,
		Errors:  errors,
	}
	panic(response)
}

// Unauthorized is response for unauthorize request
func Unauthorized(message string, errors []Error) {
	response := Response{
		Status:  http.StatusUnauthorized,
		Message: message,
		Data:    nil,
		Errors:  errors,
	}
	panic(response)
}

// Conflict is response for conflict data
func Conflict(message string, errors []Error) {
	response := Response{
		Status:  http.StatusConflict,
		Message: message,
		Data:    nil,
		Errors:  errors,
	}
	panic(response)
}

// InternalServerError is response for internal server error
func InternalServerError(message string, errors []Error) {
	response := Response{
		Status:  http.StatusInternalServerError,
		Message: message,
		Data:    nil,
		Errors:  errors,
	}
	panic(response)
}
