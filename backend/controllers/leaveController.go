package controllers

import (
	"net/http"
	"strconv"

	"simple-go-api/dto"
	"simple-go-api/services"

	"github.com/gin-gonic/gin"
)

// CreateLeave handles POST /api/v1/leaves
func CreateLeave(c *gin.Context) {
	var req dto.CreateLeaveRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.StandardResponse{
			Success:   false,
			Message:   "Invalid request payload",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	leave, err := services.CreateLeaveRequest(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.StandardResponse{
			Success:   false,
			Message:   "Failed to create leave request",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.StandardResponse{
		Success:   true,
		Message:   "Leave request created successfully",
		Data:      leave,
		RequestID: c.GetString("requestId"),
	})
}

// GetLeave handles GET /api/v1/leaves/:id
func GetLeave(c *gin.Context) {
	id := c.Param("id")

	leave, err := services.GetLeaveRequest(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.StandardResponse{
			Success:   false,
			Message:   "Leave request not found",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	c.JSON(http.StatusOK, dto.StandardResponse{
		Success:   true,
		Message:   "Leave request retrieved successfully",
		Data:      leave,
		RequestID: c.GetString("requestId"),
	})
}

// GetLeaves handles GET /api/v1/leaves
func GetLeaves(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	employeeID := c.Query("employeeId")
	status := c.Query("status")
	leaveType := c.Query("leaveType")

	result, err := services.GetLeaveRequests(page, pageSize, employeeID, status, leaveType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.StandardResponse{
			Success:   false,
			Message:   "Failed to fetch leave requests",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	c.JSON(http.StatusOK, dto.StandardResponse{
		Success:   true,
		Message:   "Leave requests retrieved successfully",
		Data:      result,
		RequestID: c.GetString("requestId"),
	})
}

// UpdateLeave handles PUT /api/v1/leaves/:id
func UpdateLeave(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateLeaveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.StandardResponse{
			Success:   false,
			Message:   "Invalid request payload",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	leave, err := services.UpdateLeaveRequest(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.StandardResponse{
			Success:   false,
			Message:   "Failed to update leave request",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	c.JSON(http.StatusOK, dto.StandardResponse{
		Success:   true,
		Message:   "Leave request updated successfully",
		Data:      leave,
		RequestID: c.GetString("requestId"),
	})
}

// GetLeaveBalance handles GET /api/v1/leaves/balance
func GetLeaveBalance(c *gin.Context) {
	employeeID := c.Query("employeeId")
	year, _ := strconv.Atoi(c.DefaultQuery("year", "2026"))

	balance, err := services.GetLeaveBalance(employeeID, year)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.StandardResponse{
			Success:   false,
			Message:   "Failed to fetch leave balance",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	c.JSON(http.StatusOK, dto.StandardResponse{
		Success:   true,
		Message:   "Leave balance retrieved successfully",
		Data:      balance,
		RequestID: c.GetString("requestId"),
	})
}