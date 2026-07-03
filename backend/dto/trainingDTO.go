package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateTrainingProgramRequest is the request body for creating a training program
type CreateTrainingProgramRequest struct {
	Title        string `json:"title" binding:"required"`
	Description  string `json:"description"`
	Type         string `json:"type" binding:"required,oneof=mandatory technical soft_skills compliance"`
	Duration     int    `json:"duration" binding:"required"`
	IsMandatory  bool   `json:"isMandatory"`
	DeadlineDays int    `json:"deadlineDays"`
}

// AssignTrainingRequest is the request body for assigning training to employees
type AssignTrainingRequest struct {
	ProgramID primitive.ObjectID `json:"programId" binding:"required"`
	EmployeeIDs []primitive.ObjectID `json:"employeeIds" binding:"required"`
}

// TrainingRecordResponse is the response body for training records
type TrainingRecordResponse struct {
	ID           string    `json:"id"`
	EmployeeID   string    `json:"employeeId"`
	ProgramID    string    `json:"programId"`
	ProgramTitle string    `json:"programTitle"`
	Status       string    `json:"status"`
	AssignedDate time.Time `json:"assignedDate"`
	StartDate    time.Time `json:"startDate,omitempty"`
	CompletedDate time.Time `json:"completedDate,omitempty"`
	Score        int       `json:"score,omitempty"`
	CertificateURL string  `json:"certificateUrl,omitempty"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}