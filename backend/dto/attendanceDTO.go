package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CheckInRequest is the request body for employee check-in
type CheckInRequest struct {
	EmployeeID primitive.ObjectID `json:"employeeId" binding:"required"`
}

// CheckOutRequest is the request body for employee check-out
type CheckOutRequest struct {
	EmployeeID primitive.ObjectID `json:"employeeId" binding:"required"`
}

// AttendanceRecordResponse is the response body for attendance records
type AttendanceRecordResponse struct {
	ID            string    `json:"id"`
	EmployeeID    string    `json:"employeeId"`
	Date          time.Time `json:"date"`
	CheckIn       time.Time `json:"checkIn,omitempty"`
	CheckOut      time.Time `json:"checkOut,omitempty"`
	WorkingHours  float64   `json:"workingHours"`
	Status        string    `json:"status"`
	Notes         string    `json:"notes"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}