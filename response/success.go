package response

import (
	"encoding/json"
	"net/http"
)

type SuccessResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    any    //cualquier cosa
}

func OK(message string, data any) Response {
	return success(message, http.StatusOK, data)
}
func Created(message string, data any) Response {
	return success(message, http.StatusCreated, data)
}
func Accepted(message string, data any) Response {
	return success(message, http.StatusAccepted, data)
}
func NonAuthoritativeInfo(message string, data any) Response {
	return success(message, http.StatusNonAuthoritativeInfo, data)
}
func NoContent(message string, data any) Response {
	return success(message, http.StatusNoContent, data)
}
func ResetContent(message string, data any) Response {
	return success(message, http.StatusResetContent, data)
}
func PartialContent(message string, data any) Response {
	return success(message, http.StatusPartialContent, data)
}

func success(message string, status int, data any) Response {
	return &SuccessResponse{
		Message: message,
		Status:  status,
		Data:    data,
	}
}

func (s *SuccessResponse) StatusCode() int {
	return s.Status
}

func (s *SuccessResponse) Error() string {
	return ""
}

func (s *SuccessResponse) Body() ([]byte, error) {
	return json.Marshal(s)
}

func (s *SuccessResponse) GetData() any {
	return s.Data
}
