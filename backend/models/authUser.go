package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AuthUser represents the authentication user (separate from Employee)
type AuthUser struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email        string             `bson:"email" validate:"required,email"`
	PasswordHash string             `bson:"passwordHash" validate:"required"`
	Role         string             `bson:"role" validate:"required,oneof=super_admin hr_admin manager employee"`
	EmployeeID   primitive.ObjectID `bson:"employeeId"`
	IsActive     bool               `bson:"isActive"`
	LastLoginAt  time.Time          `bson:"lastLoginAt"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
}

// LoginRequest represents the login request body
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	AccessToken  string   `json:"accessToken"`
	RefreshToken string   `json:"refreshToken"`
	ExpiresIn    int64    `json:"expiresIn"`
	User         UserInfo `json:"user"`
}

// UserInfo represents minimal user info returned after login
type UserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Name  string `json:"name"`
}

// RefreshTokenRequest represents the refresh token request
type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// LogoutRequest represents the logout request
type LogoutRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}
