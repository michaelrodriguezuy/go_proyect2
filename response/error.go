package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func InternalServerError(message string) Response {
	return NewError(message, http.StatusInternalServerError)
}

func NotFound(message string) Response {
	return NewError(message, http.StatusNotFound)
}
func Unauthorized(message string) Response {
	return NewError(message, http.StatusUnauthorized)
}
func Forbidden(message string) Response {
	return NewError(message, http.StatusForbidden)
}
func BadRequest(message string) Response {
	return NewError(message, http.StatusBadRequest)
}

func NewError(message string, status int) Response {
	return &ErrorResponse{
		Message: message,
		Status:  status,
	}
}

func (r *ErrorResponse) StatusCode() int {
	return r.Status
}

func (r *ErrorResponse) Error() string {
	return r.Message
}

func (r *ErrorResponse) Body() ([]byte, error) {
	return json.Marshal(r)
}

func (r *ErrorResponse) GetData() any {
	return nil
}
