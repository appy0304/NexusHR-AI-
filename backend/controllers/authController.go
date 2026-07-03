package controllers

import (
	"net/http"

	"simple-go-api/dao"
	"simple-go-api/dto"
	"simple-go-api/models"
	"simple-go-api/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


// Login handles POST /api/v1/auth/login
// @Summary User login
// @Description Authenticate user with email and password, returns JWT tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "Login credentials"
// @Success 200 {object} dto.AuthResponse
// @Failure 401 {object} dto.AuthError
// @Router /api/v1/auth/login [post]
func Login(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.AuthResponse{
			Success:   false,
			Message:   "Invalid request",
			Error:     "email and password are required",
			RequestID: c.GetString("requestId"),
		})
		return
	}

	response, err := services.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.AuthResponse{
			Success:   false,
			Message:   "Login failed",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{
		Success:   true,
		Message:   "Login successful",
		Data:      response,
		RequestID: c.GetString("requestId"),
	})
}

// Refresh handles POST /api/v1/auth/refresh
// @Summary Refresh access token
// @Description Use refresh token to get a new access token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.RefreshTokenRequest true "Refresh token"
// @Success 200 {object} dto.AuthResponse
// @Failure 401 {object} dto.AuthError
// @Router /api/v1/auth/refresh [post]
func Refresh(c *gin.Context) {
	var req models.RefreshTokenRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.AuthResponse{
			Success:   false,
			Message:   "Invalid request",
			Error:     "refreshToken is required",
			RequestID: c.GetString("requestId"),
		})
		return
	}

	response, err := services.RefreshToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.AuthResponse{
			Success:   false,
			Message:   "Token refresh failed",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{
		Success:   true,
		Message:   "Token refreshed successfully",
		Data:      response,
		RequestID: c.GetString("requestId"),
	})
}

// Logout handles POST /api/v1/auth/logout
// @Summary User logout
// @Description Invalidate refresh token and logout user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body models.LogoutRequest true "Logout request"
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} dto.AuthError
// @Router /api/v1/auth/logout [post]
func Logout(c *gin.Context) {
	var req models.LogoutRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.AuthResponse{
			Success:   false,
			Message:   "Invalid request",
			Error:     "refreshToken is required",
			RequestID: c.GetString("requestId"),
		})
		return
	}

	// Extract user ID from context (set by JWTAuth middleware)
	userIDStr, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.AuthResponse{
			Success:   false,
			Message:   "User not authenticated",
			Error:     "missing user ID",
			RequestID: c.GetString("requestId"),
		})
		return
	}

	userID, err := primitive.ObjectIDFromHex(userIDStr.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.AuthResponse{
			Success:   false,
			Message:   "Invalid user ID",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	if err := services.Logout(userID); err != nil {
		c.JSON(http.StatusInternalServerError, dto.AuthResponse{
			Success:   false,
			Message:   "Logout failed",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	c.JSON(http.StatusOK, dto.AuthResponse{
		Success:   true,
		Message:   "Logged out successfully",
		RequestID: c.GetString("requestId"),
	})
}

// CreateAdmin handles POST /api/v1/auth/create-admin (one-time setup)
// @Summary Create default admin user
// @Description Create the initial admin user for the system (use only once)
// @Tags auth
// @Accept json
// @Produce json
// @Param admin body object true "Admin credentials: email, password, name"
// @Success 201 {object} dto.AuthResponse
// @Failure 400 {object} dto.AuthError
// @Router /api/v1/auth/create-admin [post]
func CreateAdmin(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		Name     string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.AuthResponse{
			Success:   false,
			Message:   "Invalid request",
			Error:     "email, password, and name are required",
			RequestID: c.GetString("requestId"),
		})
		return
	}

	// Hash password
	passwordHash, err := services.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.AuthResponse{
			Success:   false,
			Message:   "Failed to process password",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	// Create auth user
	user := &models.AuthUser{
		Email:        req.Email,
		PasswordHash: passwordHash,
		Role:         "super_admin",
		IsActive:     true,
	}

	if err := dao.CreateAuthUser(user); err != nil {
		c.JSON(http.StatusBadRequest, dto.AuthResponse{
			Success:   false,
			Message:   "Failed to create admin",
			Error:     err.Error(),
			RequestID: c.GetString("requestId"),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.AuthResponse{
		Success:   true,
		Message:   "Admin user created successfully",
		RequestID: c.GetString("requestId"),
	})
}
