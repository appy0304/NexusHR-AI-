package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LeaveType string

const (
	AnnualLeave    LeaveType = "annual"
	SickLeave      LeaveType = "sick"
	MaternityLeave LeaveType = "maternity"
	PaternityLeave LeaveType = "paternity"
	UnpaidLeave    LeaveType = "unpaid"
	CompOff        LeaveType = "comp_off"
)

type LeaveStatus string

const (
	LeavePending   LeaveStatus = "pending"
	LeaveApproved  LeaveStatus = "approved"
	LeaveRejected  LeaveStatus = "rejected"
	LeaveCancelled LeaveStatus = "cancelled"
)

type LeaveRequest struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EmployeeID      primitive.ObjectID `bson:"employeeId"`
	LeaveType       LeaveType          `bson:"leaveType"`
	Status          LeaveStatus        `bson:"status"`
	StartDate       time.Time          `bson:"startDate"`
	EndDate         time.Time          `bson:"endDate"`
	Days            float64            `bson:"days"`
	Reason          string             `bson:"reason"`
	ApprovedBy      primitive.ObjectID `bson:"approvedBy,omitempty"`
	ApprovedAt      time.Time          `bson:"approvedAt,omitempty"`
	RejectionReason string             `bson:"rejectionReason,omitempty"`
	CreatedAt       time.Time          `bson:"createdAt"`
	UpdatedAt       time.Time          `bson:"updatedAt"`
}

type LeaveBalance struct {
	EmployeeID     primitive.ObjectID `bson:"employeeId"`
	Year           int                `bson:"year"`
	TotalAllocated float64            `bson:"totalAllocated"`
	Used           float64            `bson:"used"`
	Remaining      float64            `bson:"remaining"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
}
