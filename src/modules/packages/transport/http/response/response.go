package response

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents the structure for error response
type ErrorResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

// SuccessResponse represents the structure for success response
type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// sendErrorResponse sends a JSON error response
func SendErrorResponse(w http.ResponseWriter, statusCode int, message string, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errorResponse := ErrorResponse{
		Status:  statusCode,
		Message: message,
		Errors:  err,
	}
	json.NewEncoder(w).Encode(errorResponse)
}

// sendSuccessResponse sends a JSON success response
func SendSuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	successResponse := SuccessResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(successResponse)
}
