package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateLeaveRequest is the request body for creating a leave request
type CreateLeaveRequest struct {
	EmployeeID primitive.ObjectID `json:"employeeId" binding:"required"`
	LeaveType  string             `json:"leaveType" binding:"required,oneof=annual sick maternity paternity unpaid comp_off"`
	StartDate  time.Time          `json:"startDate" binding:"required"`
	EndDate    time.Time          `json:"endDate" binding:"required"`
	Reason     string             `json:"reason" binding:"required"`
}

// UpdateLeaveRequest is the request body for updating a leave request (approve/reject)
type UpdateLeaveRequest struct {
	Status          string             `json:"status" binding:"required,oneof=approved rejected cancelled"`
	RejectionReason string             `json:"rejectionReason"`
}

// LeaveResponse is the response body for leave requests
type LeaveResponse struct {
	ID              string    `json:"id"`
	EmployeeID      string    `json:"employeeId"`
	LeaveType       string    `json:"leaveType"`
	Status          string    `json:"status"`
	StartDate       time.Time `json:"startDate"`
	EndDate         time.Time `json:"endDate"`
	Days            float64   `json:"days"`
	Reason          string    `json:"reason"`
	ApprovedBy      string    `json:"approvedBy,omitempty"`
	ApprovedAt      time.Time `json:"approvedAt,omitempty"`
	RejectionReason string    `json:"rejectionReason,omitempty"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// LeaveBalanceResponse is the response body for leave balance
type LeaveBalanceResponse struct {
	EmployeeID     string  `json:"employeeId"`
	Year           int     `json:"year"`
	TotalAllocated float64 `json:"totalAllocated"`
	Used           float64 `json:"used"`
	Remaining      float64 `json:"remaining"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// PaginatedLeaveResponse is the response body for paginated leave requests
type PaginatedLeaveResponse struct {
	Items      []LeaveResponse `json:"items"`
	TotalCount int             `json:"totalCount"`
	Page       int             `json:"page"`
	PageSize   int             `json:"pageSize"`
	TotalPages int             `json:"totalPages"`
}