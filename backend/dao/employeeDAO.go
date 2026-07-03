package dao

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"simple-go-api/config"
	"simple-go-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateEmployee inserts a new employee into the database
func CreateEmployee(employee models.Employee) (models.Employee, error) {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(config.EMPLOYEES_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if EmployeeID already exists
	var existing models.Employee
	err := collection.FindOne(ctx, bson.M{"employeeId": employee.EmployeeID}).Decode(&existing)
	if err == nil {
		return models.Employee{}, errors.New("employee ID already exists")
	}

	// Check if email already exists
	err = collection.FindOne(ctx, bson.M{"email": employee.Email}).Decode(&existing)
	if err == nil {
		return models.Employee{}, errors.New("email already exists")
	}

	result, err := collection.InsertOne(ctx, employee)
	if err != nil {
		return models.Employee{}, err
	}

	employee.ID = result.InsertedID.(primitive.ObjectID)
	return employee, nil
}

// GetEmployee finds a single employee by ID
func GetEmployee(id primitive.ObjectID) (models.Employee, error) {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(config.EMPLOYEES_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var employee models.Employee
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&employee)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Employee{}, errors.New("employee not found")
		}
		return models.Employee{}, err
	}

	return employee, nil
}

// GetEmployees retrieves all employees with pagination, filtering, sorting, and search
func GetEmployees(filter models.EmployeeFilter) (models.EmployeeListResponse, error) {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(config.EMPLOYEES_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Build filter query
	query := bson.M{}

	if filter.Department != "" {
		query["department"] = filter.Department
	}
	if filter.EmploymentStatus != "" {
		query["employmentStatus"] = filter.EmploymentStatus
	}
	if filter.ManagerID != "" {
		managerObjectID, err := primitive.ObjectIDFromHex(filter.ManagerID)
		if err == nil {
			query["managerId"] = managerObjectID
		}
	}
	if !filter.JoinDateFrom.IsZero() {
		query["joiningDate"] = bson.M{"$gte": filter.JoinDateFrom}
	}
	if !filter.JoinDateTo.IsZero() {
		if query["joiningDate"] == nil {
			query["joiningDate"] = bson.M{}
		}
		if gd, ok := query["joiningDate"].(bson.M); ok {
			gd["$lte"] = filter.JoinDateTo
		}
	}
	if filter.Search != "" {
		// Use text index for full-text search
		query["$text"] = bson.M{"$search": filter.Search}
	}
	if filter.Skills != "" {
		skills := strings.Split(filter.Skills, ",")
		query["skills"] = bson.M{"$in": skills}
	}

	// Count total documents matching filter
	totalCount, err := collection.CountDocuments(ctx, query)
	if err != nil {
		return models.EmployeeListResponse{}, err
	}

	// Build sort
	sortField := filter.SortBy
	if sortField == "" {
		sortField = "createdAt"
	}
	sortOrder := 1 // asc
	if strings.ToLower(filter.SortOrder) == "desc" {
		sortOrder = -1
	}
	sortOpt := options.Find().SetSort(bson.M{sortField: sortOrder})

	// Calculate skip and limit for pagination
	skip := (filter.Page - 1) * filter.PageSize
	sortOpt.SetSkip(int64(skip))
	sortOpt.SetLimit(int64(filter.PageSize))

	// Execute query
	cursor, err := collection.Find(ctx, query, sortOpt)
	if err != nil {
		return models.EmployeeListResponse{}, err
	}
	defer cursor.Close(ctx)

	var employees []models.Employee
	if err = cursor.All(ctx, &employees); err != nil {
		return models.EmployeeListResponse{}, err
	}

	// Calculate total pages
	totalPages := int(totalCount) / filter.PageSize
	if int(totalCount)%filter.PageSize > 0 {
		totalPages++
	}

	return models.EmployeeListResponse{
		TotalCount: totalCount,
		Page:       filter.Page,
		PageSize:   filter.PageSize,
		TotalPages: totalPages,
		Data:       employees,
	}, nil
}

// UpdateEmployee updates an existing employee
func UpdateEmployee(id primitive.ObjectID, updates models.Employee) (models.Employee, error) {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(config.EMPLOYEES_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if employee exists
	var existing models.Employee
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&existing)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Employee{}, errors.New("employee not found")
		}
		return models.Employee{}, err
	}

	// Build update document (only non-zero fields)
	update := bson.M{
		"$set": bson.M{
			"firstName":        updates.FirstName,
			"lastName":         updates.LastName,
			"email":            updates.Email,
			"phone":            updates.Phone,
			"dateOfBirth":      updates.DateOfBirth,
			"department":       updates.Department,
			"designation":      updates.Designation,
			"managerId":        updates.ManagerID,
			"salary":           updates.Salary,
			"joiningDate":      updates.JoiningDate,
			"employmentStatus": updates.EmploymentStatus,
			"skills":           updates.Skills,
			"address":          updates.Address,
			"emergencyContact": updates.EmergencyContact,
			"updatedAt":        time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return models.Employee{}, err
	}

	// Return updated document
	var updated models.Employee
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&updated)
	if err != nil {
		return models.Employee{}, err
	}

	return updated, nil
}

// DeleteEmployee deletes an employee by ID
func DeleteEmployee(id primitive.ObjectID) error {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(config.EMPLOYEES_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("employee not found")
	}

	return nil
}

// GetEmployeeByEmployeeID finds an employee by their human-readable ID
func GetEmployeeByEmployeeID(employeeID string) (models.Employee, error) {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(config.EMPLOYEES_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var employee models.Employee
	err := collection.FindOne(ctx, bson.M{"employeeId": employeeID}).Decode(&employee)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Employee{}, fmt.Errorf("employee with ID %s not found", employeeID)
		}
		return models.Employee{}, err
	}

	return employee, nil
}

// GetEmployeeByEmail finds an employee by email
func GetEmployeeByEmail(email string) (models.Employee, error) {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(config.EMPLOYEES_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var employee models.Employee
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&employee)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Employee{}, fmt.Errorf("employee with email %s not found", email)
		}
		return models.Employee{}, err
	}

	return employee, nil
}
