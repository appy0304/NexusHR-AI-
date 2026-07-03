package dto

import (
	"time"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateEmployeeRequest represents the request body for creating an employee
type CreateEmployeeRequest struct {
	EmployeeID       string              `json:"employeeId" validate:"required,max=50"`
	FirstName        string              `json:"firstName" validate:"required,max=100"`
	LastName         string              `json:"lastName" validate:"required,max=100"`
	Email            string              `json:"email" validate:"required,email,max=150"`
	Phone            string              `json:"phone" validate:"max=20"`
	DateOfBirth      time.Time           `json:"dateOfBirth"`
	Department       string              `json:"department" validate:"required,max=100"`
	Designation      string              `json:"designation" validate:"required,max=100"`
	ManagerID        string              `json:"managerId"` // Sent as string, converted to ObjectID
	Salary           float64             `json:"salary" validate:"gte=0"`
	JoiningDate      time.Time           `json:"joiningDate"`
	EmploymentStatus string              `json:"employmentStatus" validate:"required,oneof=active inactive terminated onboarding leave"`
	Skills           []string            `json:"skills"`
	Address          AddressDTO          `json:"address"`
	EmergencyContact EmergencyContactDTO `json:"emergencyContact"`
}

// UpdateEmployeeRequest represents the request body for updating an employee
type UpdateEmployeeRequest struct {
	FirstName        *string              `json:"firstName"`
	LastName         *string              `json:"lastName"`
	Email            *string              `json:"email"`
	Phone            *string              `json:"phone"`
	DateOfBirth      *time.Time           `json:"dateOfBirth"`
	Department       *string              `json:"department"`
	Designation      *string              `json:"designation"`
	ManagerID        *string              `json:"managerId"`
	Salary           *float64             `json:"salary"`
	JoiningDate      *time.Time           `json:"joiningDate"`
	EmploymentStatus *string              `json:"employmentStatus"`
	Skills           []string             `json:"skills"`
	Address          *AddressDTO          `json:"address"`
	EmergencyContact *EmergencyContactDTO `json:"emergencyContact"`
}

// AddressDTO represents address in request/response
type AddressDTO struct {
	Street  string `json:"street" validate:"max=200"`
	City    string `json:"city" validate:"required,max=100"`
	State   string `json:"state" validate:"required,max=100"`
	ZipCode string `json:"zipCode" validate:"required,max=20"`
	Country string `json:"country" validate:"required,max=100"`
}

// EmergencyContactDTO represents emergency contact in request/response
type EmergencyContactDTO struct {
	Name         string `json:"name" validate:"required,max=150"`
	Relationship string `json:"relationship" validate:"required,max=50"`
	Phone        string `json:"phone" validate:"required,max=20"`
	Email        string `json:"email" validate:"email,max=150"`
}

// EmployeeResponse represents the response for a single employee
type EmployeeResponse struct {
	ID               string              `json:"id"`
	EmployeeID       string              `json:"employeeId"`
	FirstName        string              `json:"firstName"`
	LastName         string              `json:"lastName"`
	Email            string              `json:"email"`
	Phone            string              `json:"phone"`
	DateOfBirth      time.Time           `json:"dateOfBirth"`
	Department       string              `json:"department"`
	Designation      string              `json:"designation"`
	ManagerID        string              `json:"managerId"`
	ManagerName      string              `json:"managerName,omitempty"` // Populated from related employee
	Salary           float64             `json:"salary"`
	JoiningDate      time.Time           `json:"joiningDate"`
	EmploymentStatus string              `json:"employmentStatus"`
	Skills           []string            `json:"skills"`
	Address          AddressDTO          `json:"address"`
	EmergencyContact EmergencyContactDTO `json:"emergencyContact"`
	CreatedAt        time.Time           `json:"createdAt"`
	UpdatedAt        time.Time           `json:"updatedAt"`
}

// EmployeeListResponse represents the paginated response
type EmployeeListResponse struct {
	TotalCount int64              `json:"totalCount"`
	Page       int                `json:"page"`
	PageSize   int                `json:"pageSize"`
	TotalPages int                `json:"totalPages"`
	Data       []EmployeeResponse `json:"data"`
}
