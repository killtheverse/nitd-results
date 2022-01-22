package utils

import (
	"encoding/json"
	"net/http"

	"github.com/killtheverse/nitd-results/app/models"
)

// ResponseWriter writes response of type - Response
func ResponseWriter(rw http.ResponseWriter, statusCode int, message string, student models.Student) error {
	rw.WriteHeader(statusCode)
	httpResponse := models.NewResponse(statusCode, message, student)
	err := json.NewEncoder(rw).Encode(httpResponse)
	return err
}

// ErrorResponseWriter writes response of type - ErrorResponse
func ErrorResponseWriter(rw http.ResponseWriter, statusCode int, message string, data interface{}) error {
	rw.WriteHeader(statusCode)
	httpResponse := models.NewErrorResponse(statusCode, message, data)
	err := json.NewEncoder(rw).Encode(httpResponse)
	return err
}

// PaginatedResponseWriter writes response of type - PaginatedResponse
func PaginatedResponseWriter(rw http.ResponseWriter, statusCode int, message string, count int, next string, prev string, students []models.Student) error {
	rw.WriteHeader(statusCode)
	httpResponse := models.NewPaginatedResponse(statusCode, count, message, next, prev, students)
	err := json.NewEncoder(rw).Encode(httpResponse)
	return err
}