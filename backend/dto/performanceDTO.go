package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreatePerformanceReviewRequest is the request body for creating a performance review
type CreatePerformanceReviewRequest struct {
	EmployeeID      primitive.ObjectID `json:"employeeId" binding:"required"`
	ReviewPeriod    string             `json:"reviewPeriod" binding:"required"`
	WorkQuality     int                `json:"workQuality" binding:"required,min=1,max=5"`
	Punctuality     int                `json:"punctuality" binding:"required,min=1,max=5"`
	Teamwork        int                `json:"teamwork" binding:"required,min=1,max=5"`
	Communication   int                `json:"communication" binding:"required,min=1,max=5"`
	Leadership      int                `json:"leadership" binding:"required,min=1,max=5"`
	TechnicalSkills int                `json:"technicalSkills" binding:"required,min=1,max=5"`
	Comments        string             `json:"comments"`
}

// PerformanceReviewResponse is the response body for performance reviews
type PerformanceReviewResponse struct {
	ID              string  `json:"id"`
	EmployeeID      string  `json:"employeeId"`
	ReviewerID      string  `json:"reviewerId"`
	ReviewPeriod    string  `json:"reviewPeriod"`
	WorkQuality     int     `json:"workQuality"`
	Punctuality     int     `json:"punctuality"`
	Teamwork        int     `json:"teamwork"`
	Communication   int     `json:"communication"`
	Leadership      int     `json:"leadership"`
	TechnicalSkills int     `json:"technicalSkills"`
	OverallScore    float64 `json:"overallScore"`
	Comments        string  `json:"comments"`
	ReviewDate      string  `json:"reviewDate"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
}
