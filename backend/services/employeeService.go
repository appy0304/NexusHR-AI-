package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"simple-go-api/dao"
	"simple-go-api/dto"
	"simple-go-api/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateEmployee creates a new employee with validation
func CreateEmployee(req dto.CreateEmployeeRequest) (models.Employee, error) {
	// Validate required fields
	if strings.TrimSpace(req.FirstName) == "" {
		return models.Employee{}, errors.New("first name is required")
	}
	if strings.TrimSpace(req.LastName) == "" {
		return models.Employee{}, errors.New("last name is required")
	}
	if strings.TrimSpace(req.Department) == "" {
		return models.Employee{}, errors.New("department is required")
	}
	if strings.TrimSpace(req.Designation) == "" {
		return models.Employee{}, errors.New("designation is required")
	}
	if strings.TrimSpace(req.EmploymentStatus) == "" {
		return models.Employee{}, errors.New("employment status is required")
	}

	// Validate employment status value
	validStatuses := []string{"active", "inactive", "terminated", "onboarding", "leave"}
	statusValid := false
	for _, s := range validStatuses {
		if req.EmploymentStatus == s {
			statusValid = true
			break
		}
	}
	if !statusValid {
		return models.Employee{}, errors.New("invalid employment status")
	}

	// Validate salary
	if req.Salary < 0 {
		return models.Employee{}, errors.New("salary cannot be negative")
	}

	// Generate EmployeeID if not provided
	employeeID := strings.TrimSpace(req.EmployeeID)
	if employeeID == "" {
		employeeID = generateEmployeeID()
	}

	// Parse ManagerID if provided
	var managerID primitive.ObjectID
	if req.ManagerID != "" {
		var err error
		managerID, err = primitive.ObjectIDFromHex(req.ManagerID)
		if err != nil {
			return models.Employee{}, errors.New("invalid manager ID format")
		}
	}

	now := time.Now()

	employee := models.Employee{
		EmployeeID:       employeeID,
		FirstName:        strings.TrimSpace(req.FirstName),
		LastName:         strings.TrimSpace(req.LastName),
		Email:            strings.ToLower(strings.TrimSpace(req.Email)),
		Phone:            strings.TrimSpace(req.Phone),
		DateOfBirth:      req.DateOfBirth,
		Department:       strings.TrimSpace(req.Department),
		Designation:      strings.TrimSpace(req.Designation),
		ManagerID:        managerID,
		Salary:           req.Salary,
		JoiningDate:      req.JoiningDate,
		EmploymentStatus: req.EmploymentStatus,
		Skills:           req.Skills,
		Address: models.Address{
			Street:  req.Address.Street,
			City:    req.Address.City,
			State:   req.Address.State,
			ZipCode: req.Address.ZipCode,
			Country: req.Address.Country,
		},
		EmergencyContact: models.EmergencyContact{
			Name:         req.EmergencyContact.Name,
			Relationship: req.EmergencyContact.Relationship,
			Phone:        req.EmergencyContact.Phone,
			Email:        req.EmergencyContact.Email,
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	return dao.CreateEmployee(employee)
}

// GetEmployee retrieves a single employee by ID
func GetEmployee(id string) (models.Employee, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Employee{}, errors.New("invalid employee ID format")
	}
	return dao.GetEmployee(objectID)
}

// GetEmployees retrieves all employees with pagination, filtering, sorting, and search
func GetEmployees(filter models.EmployeeFilter) (models.EmployeeListResponse, error) {
	// Set defaults
	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.PageSize < 1 || filter.PageSize > 100 {
		filter.PageSize = 20
	}
	if filter.SortBy == "" {
		filter.SortBy = "createdAt"
	}
	if filter.SortOrder == "" {
		filter.SortOrder = "desc"
	}

	// // Parse ManagerID if provided
	// if filter.ManagerID != "" {
	//     mid, err := primitive.ObjectIDFromHex(filter.ManagerID)
	//     if err == nil {
	//         filter.ManagerID = mid
	//     }
	// }

	return dao.GetEmployees(filter)
}

// UpdateEmployee updates an existing employee
func UpdateEmployee(id string, req dto.UpdateEmployeeRequest) (models.Employee, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Employee{}, errors.New("invalid employee ID format")
	}

	// Get existing employee
	existing, err := dao.GetEmployee(objectID)
	if err != nil {
		return models.Employee{}, err
	}

	// Build updated employee (only update provided fields)
	updated := existing

	if req.FirstName != nil {
		if strings.TrimSpace(*req.FirstName) == "" {
			return models.Employee{}, errors.New("first name cannot be empty")
		}
		updated.FirstName = strings.TrimSpace(*req.FirstName)
	}
	if req.LastName != nil {
		if strings.TrimSpace(*req.LastName) == "" {
			return models.Employee{}, errors.New("last name cannot be empty")
		}
		updated.LastName = strings.TrimSpace(*req.LastName)
	}
	if req.Email != nil {
		updated.Email = strings.ToLower(strings.TrimSpace(*req.Email))
	}
	if req.Phone != nil {
		updated.Phone = strings.TrimSpace(*req.Phone)
	}
	if req.DateOfBirth != nil {
		updated.DateOfBirth = *req.DateOfBirth
	}
	if req.Department != nil {
		if strings.TrimSpace(*req.Department) == "" {
			return models.Employee{}, errors.New("department cannot be empty")
		}
		updated.Department = strings.TrimSpace(*req.Department)
	}
	if req.Designation != nil {
		if strings.TrimSpace(*req.Designation) == "" {
			return models.Employee{}, errors.New("designation cannot be empty")
		}
		updated.Designation = strings.TrimSpace(*req.Designation)
	}
	if req.ManagerID != nil {
		if *req.ManagerID != "" {
			mid, err := primitive.ObjectIDFromHex(*req.ManagerID)
			if err != nil {
				return models.Employee{}, errors.New("invalid manager ID format")
			}
			updated.ManagerID = mid
		} else {
			updated.ManagerID = primitive.ObjectID{}
		}
	}
	if req.Salary != nil {
		if *req.Salary < 0 {
			return models.Employee{}, errors.New("salary cannot be negative")
		}
		updated.Salary = *req.Salary
	}
	if req.JoiningDate != nil {
		updated.JoiningDate = *req.JoiningDate
	}
	if req.EmploymentStatus != nil {
		validStatuses := []string{"active", "inactive", "terminated", "onboarding", "leave"}
		statusValid := false
		for _, s := range validStatuses {
			if *req.EmploymentStatus == s {
				statusValid = true
				break
			}
		}
		if !statusValid {
			return models.Employee{}, errors.New("invalid employment status")
		}
		updated.EmploymentStatus = *req.EmploymentStatus
	}
	if req.Skills != nil {
		updated.Skills = req.Skills
	}
	if req.Address != nil {
		updated.Address = models.Address{
			Street:  req.Address.Street,
			City:    req.Address.City,
			State:   req.Address.State,
			ZipCode: req.Address.ZipCode,
			Country: req.Address.Country,
		}
	}
	if req.EmergencyContact != nil {
		updated.EmergencyContact = models.EmergencyContact{
			Name:         req.EmergencyContact.Name,
			Relationship: req.EmergencyContact.Relationship,
			Phone:        req.EmergencyContact.Phone,
			Email:        req.EmergencyContact.Email,
		}
	}

	updated.UpdatedAt = time.Now()

	return dao.UpdateEmployee(objectID, updated)
}

// DeleteEmployee deletes an employee by ID
func DeleteEmployee(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid employee ID format")
	}
	return dao.DeleteEmployee(objectID)
}

// generateEmployeeID generates a unique employee ID
func generateEmployeeID() string {
	// Query the last employee ID and increment
	// For simplicity in this phase, use timestamp-based
	now := time.Now()
	return fmt.Sprintf("EMP-%s-%04d", now.Format("20060102"), now.UnixNano()%10000)
}

// func sprintf(format string, a ...interface{}) string {
//     // Simple sprintf replacement
//     result := format
//     for _, a := range a {
//         result = sprintfReplace(result, "%d", a)
//     }
//     return result
// }

// func sprintfReplace(s, format string, a interface{}) string {
//     // Simplified - in production use fmt.Sprintf
//     return s
// }
