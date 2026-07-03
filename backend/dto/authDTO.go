package dto

import "simple-go-api/models"

// AuthResponse is a generic auth response wrapper
type AuthResponse struct {
	Success   bool                  `json:"success"`
	Message   string                `json:"message"`
	Data      *models.LoginResponse `json:"data,omitempty"`
	Error     string                `json:"error,omitempty"`
	RequestID string                `json:"requestId,omitempty"`
}

// AuthError is a standardized auth error response
type AuthError struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	Message   string `json:"message"`
	RequestID string `json:"requestId,omitempty"`
}

type LoginResponse struct {
	Token      string `json:"token"`
	EmployeeID string `json:"employeeId"`
	Email      string `json:"email"`
	Role       string `json:"role"`
}
