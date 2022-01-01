package utils

import (
	"encoding/json"
	"net/http"

	"github.com/killtheverse/nitd-results/app/models"
)

// ResponseWriter writes response based on data provided
func ResponseWriter(rw http.ResponseWriter, statusCode int, message string, data interface{}) error {
	rw.WriteHeader(statusCode)
	httpResponse := models.NewResponse(statusCode, message, data)
	err := json.NewEncoder(rw).Encode(httpResponse)
	return err
}