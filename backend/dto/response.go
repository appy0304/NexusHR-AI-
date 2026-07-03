package dto

import (
	"time"
)

// StandardResponse is the unified API response format
type StandardResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	RequestID string      `json:"requestId,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

// PaginatedResponse wraps paginated data
type PaginatedResponse struct {
	Items      interface{} `json:"items"`
	TotalCount int64       `json:"totalCount"`
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	TotalPages int         `json:"totalPages"`
}

// ErrorResponse for standardized errors
type ErrorResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
	Code      string `json:"code,omitempty"`
}

// Success creates a standardized success response
func Success(message string, data interface{}, requestID string) StandardResponse {
	return StandardResponse{
		Success:   true,
		Message:   message,
		Data:      data,
		RequestID: requestID,
		Timestamp: time.Now(),
	}
}

// Fail creates a standardized error response
func Fail(message string, err error, requestID string) StandardResponse {
	errorText := ""
	if err != nil {
		errorText = err.Error()
	}

	return StandardResponse{
		Success:   false,
		Message:   message,
		Error:     errorText,
		RequestID: requestID,
		Timestamp: time.Now(),
	}
}
