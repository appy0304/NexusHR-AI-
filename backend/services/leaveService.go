package services

import (
	"errors"
	"fmt"
	"time"

	"simple-go-api/dao"
	"simple-go-api/dto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateLeaveRequest creates a new leave request
func CreateLeaveRequest(req dto.CreateLeaveRequest) (*dto.LeaveResponse, error) {
	// Validate dates
	if req.EndDate.Before(req.StartDate) {
		return nil, errors.New("end date cannot be before start date")
	}

	// Calculate number of days
	days := req.EndDate.Sub(req.StartDate).Hours() / 24
	if days < 1 {
		return nil, errors.New("minimum leave duration is 1 day")
	}

	// Check leave balance for annual leave
	if req.LeaveType == "annual" {
		balance, err := dao.GetLeaveBalance(req.EmployeeID, time.Now().Year())
		if err != nil {
			return nil, errors.New("leave balance not found")
		}
		if balance.Remaining < days {
			return nil, fmt.Errorf("insufficient leave balance. Remaining: %.0f, Requested: %.0f", balance.Remaining, days)
		}
	}

	// Create leave request
	leave := &dao.LeaveRequest{
		EmployeeID: req.EmployeeID,
		// LeaveType:  req.LeaveType,
		LeaveType: dao.LeaveType(req.LeaveType),

		Status:    "pending",
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
		Days:      days,
		Reason:    req.Reason,
	}

	err := dao.CreateLeaveRequest(leave)
	if err != nil {
		return nil, err
	}

	// Convert to response
	return convertLeaveToResponse(leave), nil
}

// GetLeaveRequest retrieves a single leave request by ID
func GetLeaveRequest(id string) (*dto.LeaveResponse, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid leave request ID")
	}

	leave, err := dao.GetLeaveByID(oid)
	if err != nil {
		return nil, err
	}

	return convertLeaveToResponse(leave), nil
}

// GetLeaveRequests retrieves all leave requests with pagination and filtering
func GetLeaveRequests(page, pageSize int, employeeID, status, leaveType string) (*dto.PaginatedLeaveResponse, error) {
	var employeeOID primitive.ObjectID
	var err error
	if employeeID != "" {
		employeeOID, err = primitive.ObjectIDFromHex(employeeID)
		if err != nil {
			return nil, errors.New("invalid employee ID")
		}
	}

	leaves, total, err := dao.GetLeaveRequests(page, pageSize, employeeOID, status, leaveType)
	if err != nil {
		return nil, err
	}

	items := make([]dto.LeaveResponse, len(leaves))
	for i, leave := range leaves {
		items[i] = *convertLeaveToResponse(&leave)
	}

	return &dto.PaginatedLeaveResponse{
		Items:      items,
		TotalCount: total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: (total + pageSize - 1) / pageSize,
	}, nil
}

// UpdateLeaveRequest approves or rejects a leave request
func UpdateLeaveRequest(id string, req dto.UpdateLeaveRequest) (*dto.LeaveResponse, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid leave request ID")
	}

	leave, err := dao.GetLeaveByID(oid)
	if err != nil {
		return nil, err
	}

	// Only pending requests can be updated
	if leave.Status != "pending" {
		return nil, errors.New("only pending leave requests can be updated")
	}

	// Update status
	leave.Status = dao.LeaveStatus(req.Status)
	leave.UpdatedAt = time.Now()

	if req.Status == "approved" {
		// Deduct from leave balance
		if leave.LeaveType == "annual" {
			balance, err := dao.GetLeaveBalance(leave.EmployeeID, time.Now().Year())
			if err == nil {
				balance.Used += leave.Days
				balance.Remaining -= leave.Days
				balance.UpdatedAt = time.Now()
				dao.UpdateLeaveBalance(balance)
			}
		}
	} else if req.Status == "rejected" {
		leave.RejectionReason = req.RejectionReason
	}

	err = dao.UpdateLeaveByID(oid, leave)
	if err != nil {
		return nil, err
	}

	return convertLeaveToResponse(leave), nil
}

// GetLeaveBalance retrieves the leave balance for an employee
func GetLeaveBalance(employeeID string, year int) (*dto.LeaveBalanceResponse, error) {
	oid, err := primitive.ObjectIDFromHex(employeeID)
	if err != nil {
		return nil, errors.New("invalid employee ID")
	}

	balance, err := dao.GetLeaveBalance(oid, year)
	if err != nil {
		return nil, err
	}

	return &dto.LeaveBalanceResponse{
		EmployeeID:     employeeID,
		Year:           year,
		TotalAllocated: balance.TotalAllocated,
		Used:           balance.Used,
		Remaining:      balance.Remaining,
		UpdatedAt:      balance.UpdatedAt,
	}, nil
}

// convertLeaveToResponse converts a DAO leave request to a DTO response
func convertLeaveToResponse(leave *dao.LeaveRequest) *dto.LeaveResponse {
	return &dto.LeaveResponse{
		ID:              leave.ID.Hex(),
		EmployeeID:      leave.EmployeeID.Hex(),
		LeaveType:       string(leave.LeaveType),
		Status:          string(leave.Status),
		StartDate:       leave.StartDate,
		EndDate:         leave.EndDate,
		Days:            leave.Days,
		Reason:          leave.Reason,
		ApprovedBy:      leave.ApprovedBy.Hex(),
		ApprovedAt:      leave.ApprovedAt,
		RejectionReason: leave.RejectionReason,
		CreatedAt:       leave.CreatedAt,
		UpdatedAt:       leave.UpdatedAt,
	}
}
