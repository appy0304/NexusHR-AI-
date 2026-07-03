package controllers

import (
	"net/http"
	"strconv"
	"time"

	"simple-go-api/dto"
	"simple-go-api/models"
	"simple-go-api/services"

	"github.com/gin-gonic/gin"
)

// CreateEmployee handles POST /api/v1/employees
// @Summary Create a new employee
// @Description Create a new employee record with full validation
// @Tags employees
// @Accept json
// @Produce json
// @Param employee body dto.CreateEmployeeRequest true "Employee creation payload"
// @Success 201 {object} dto.StandardResponse
// @Failure 400 {object} dto.StandardResponse
// @Failure 500 {object} dto.StandardResponse
// @Router /api/v1/employees [post]
func CreateEmployee(c *gin.Context) {
	var req dto.CreateEmployeeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.StandardResponse{
			Success:   false,
			Message:   "Invalid request payload",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	employee, err := services.CreateEmployee(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.StandardResponse{
			Success:   false,
			Message:   "Failed to create employee",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.StandardResponse{
		Success:   true,
		Message:   "Employee created successfully",
		Data:      employee,
		RequestID: c.GetString("requestId"),
	})
}

// GetEmployee handles GET /api/v1/employees/:id
// @Summary Get employee by ID
// @Description Retrieve a single employee record by their MongoDB ID
// @Tags employees
// @Produce json
// @Param id path string true "Employee MongoDB ID"
// @Success 200 {object} dto.StandardResponse
// @Failure 404 {object} dto.StandardResponse
// @Router /api/v1/employees/{id} [get]
func GetEmployee(c *gin.Context) {
	id := c.Param("id")

	employee, err := services.GetEmployee(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.StandardResponse{
			Success:   false,
			Message:   "Employee not found",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	c.JSON(http.StatusOK, dto.StandardResponse{
		Success:   true,
		Message:   "Employee retrieved successfully",
		Data:      employee,
		RequestID: c.GetString("requestId"),
	})
}

// GetEmployees handles GET /api/v1/employees
// @Summary Get all employees (paginated)
// @Description Retrieve all employees with pagination, filtering, sorting, and search
// @Tags employees
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Items per page" default(20) maximum(100)
// @Param department query string false "Filter by department"
// @Param status query string false "Filter by employment status"
// @Param search query string false "Full-text search across name, email, skills"
// @Param sort query string false "Sort field" default(created_at)
// @Param order query string false "Sort order" default(desc)
// @Success 200 {object} dto.StandardResponse
// @Failure 500 {object} dto.StandardResponse
// @Router /api/v1/employees [get]
func GetEmployees(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	department := c.Query("department")
	status := c.Query("status")
	search := c.Query("search")
	sortBy := c.DefaultQuery("sort", "createdAt")
	order := c.DefaultQuery("order", "desc")
	managerID := c.Query("managerId")
	joinDateFrom := c.Query("joinDateFrom")
	joinDateTo := c.Query("joinDateTo")
	skills := c.Query("skills")

	filter := models.EmployeeFilter{
		Page:             page,
		PageSize:         pageSize,
		Department:       department,
		EmploymentStatus: status,
		Search:           search,
		SortBy:           sortBy,
		SortOrder:        order,
		ManagerID:        managerID,
		Skills:           skills,
	}

	// Parse date filters if provided
	if joinDateFrom != "" {
		if t, err := parseTime(joinDateFrom); err == nil {
			filter.JoinDateFrom = t
		}
	}
	if joinDateTo != "" {
		if t, err := parseTime(joinDateTo); err == nil {
			filter.JoinDateTo = t
		}
	}

	result, err := services.GetEmployees(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.StandardResponse{
			Success:   false,
			Message:   "Failed to fetch employees",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	c.JSON(http.StatusOK, dto.StandardResponse{
		Success:   true,
		Message:   "Employees retrieved successfully",
		Data:      result,
		RequestID: c.GetString("requestId"),
	})
}

// UpdateEmployee handles PUT /api/v1/employees/:id
// @Summary Update employee
// @Description Update an existing employee record (partial update supported)
// @Tags employees
// @Accept json
// @Produce json
// @Param id path string true "Employee MongoDB ID"
// @Param employee body dto.UpdateEmployeeRequest true "Employee update payload"
// @Success 200 {object} dto.StandardResponse
// @Failure 400 {object} dto.StandardResponse
// @Failure 404 {object} dto.StandardResponse
// @Router /api/v1/employees/{id} [put]
func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateEmployeeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.StandardResponse{
			Success:   false,
			Message:   "Invalid request payload",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	employee, err := services.UpdateEmployee(id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.StandardResponse{
			Success:   false,
			Message:   "Failed to update employee",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	c.JSON(http.StatusOK, dto.StandardResponse{
		Success:   true,
		Message:   "Employee updated successfully",
		Data:      employee,
		RequestID: c.GetString("requestId"),
	})
}

// DeleteEmployee handles DELETE /api/v1/employees/:id
// @Summary Delete employee
// @Description Delete an employee by their MongoDB ID
// @Tags employees
// @Produce json
// @Param id path string true "Employee MongoDB ID"
// @Success 200 {object} dto.StandardResponse
// @Failure 400 {object} dto.StandardResponse
// @Failure 404 {object} dto.StandardResponse
// @Router /api/v1/employees/{id} [delete]
func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	err := services.DeleteEmployee(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.StandardResponse{
			Success:   false,
			Message:   "Failed to delete employee",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	c.JSON(http.StatusOK, dto.StandardResponse{
		Success:   true,
		Message:   "Employee deleted successfully",
		RequestID: c.GetString("requestId"),
	})
}

func parseTime(s string) (time.Time, error) {
	return time.Parse("2006-01-02", s)
}
