package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Address struct {
	Street  string `json:"street" bson:"street"`
	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	ZipCode string `json:"zipCode" bson:"zipCode"`
	Country string `json:"country" bson:"country"`
}

type EmergencyContact struct {
	Name         string `json:"name" bson:"name"`
	Relationship string `json:"relationship" bson:"relationship"`
	Phone        string `json:"phone" bson:"phone"`
	Email        string `json:"email" bson:"email"`
}

type Employee struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	EmployeeID       string             `json:"employeeId" bson:"employeeId"`
	FirstName        string             `json:"firstName" bson:"firstName"`
	LastName         string             `json:"lastName" bson:"lastName"`
	Email            string             `json:"email" bson:"email"`
	Phone            string             `json:"phone" bson:"phone"`
	DateOfBirth      time.Time          `json:"dateOfBirth" bson:"dateOfBirth"`
	Department       string             `json:"department" bson:"department"`
	Designation      string             `json:"designation" bson:"designation"`
	ManagerID        primitive.ObjectID `json:"managerId" bson:"managerId,omitempty"`
	Salary           float64            `json:"salary" bson:"salary"`
	JoiningDate      time.Time          `json:"joiningDate" bson:"joiningDate"`
	EmploymentStatus string             `json:"employmentStatus" bson:"employmentStatus"`
	Skills           []string           `json:"skills" bson:"skills"`
	Address          Address            `json:"address" bson:"address"`
	EmergencyContact EmergencyContact   `json:"emergencyContact" bson:"emergencyContact"`
	CreatedAt        time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt        time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type EmployeeFilter struct {
	Page             int
	PageSize         int
	Department       string
	EmploymentStatus string
	Search           string
	SortBy           string
	SortOrder        string
	ManagerID        string
	JoinDateFrom     time.Time
	JoinDateTo       time.Time
	Skills           string
}

type EmployeeListResponse struct {
	TotalCount int64      `json:"totalCount"`
	Page       int        `json:"page"`
	PageSize   int        `json:"pageSize"`
	TotalPages int        `json:"totalPages"`
	Data       []Employee `json:"data"`
}
