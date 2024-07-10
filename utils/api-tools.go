package utils

import (
	"encoding/json"
	"net/http"
)

const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
)

type APIResponse struct {
	Status       bool        `json:"status"`
	Result       interface{} `json:"result,omitempty"`
	ErrorMessage string      `json:"error,omitempty"`
}

func JsonResponse(w http.ResponseWriter, status int, data interface{}, errorMessage string) {
	response := APIResponse{}
	if status >= 500 {
		response.Status = false
		response.Result = nil
		response.ErrorMessage = http.StatusText(status)
	} else if status != http.StatusOK && status != http.StatusCreated {
		response.Status = false
		response.Result = nil
		response.ErrorMessage = errorMessage
	} else {
		response.Status = true
		response.Result = data
		response.ErrorMessage = ""
	}
	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
