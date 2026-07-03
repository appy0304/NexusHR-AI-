package middleware

import (
	"net/http"
	"strings"

	"simple-go-api/dto"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const jwtSecret = "EMPLOYEE-MANAGEMENT-PLATFORM-2026-SECRET-KEY-CHANGE-IN-PRODUCTION"

// JWTAuth validates JWT access tokens from the Authorization header
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Step 1: Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, dto.AuthResponse{
				Success:   false,
				Message:   "Unauthorized",
				Error:     "missing authorization header",
				RequestID: c.GetString("requestId"),
			})
			c.Abort()
			return
		}

		// Step 2: Extract token (format: "Bearer <token>")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, dto.AuthResponse{
				Success:   false,
				Message:   "Unauthorized",
				Error:     "invalid token format. Use: Bearer <token>",
				RequestID: c.GetString("requestId"),
			})
			c.Abort()
			return
		}

		// Step 3: Parse and validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, gin.Error{Err: http.ErrAbortHandler}
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, dto.AuthResponse{
				Success:   false,
				Message:   "Unauthorized",
				Error:     "invalid or expired token",
				RequestID: c.GetString("requestId"),
			})
			c.Abort()
			return
		}

		// Step 4: Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, dto.AuthResponse{
				Success:   false,
				Message:   "Unauthorized",
				Error:     "invalid token claims",
				RequestID: c.GetString("requestId"),
			})
			c.Abort()
			return
		}

		// Step 5: Check token type (must be access token)
		if tokenType, ok := claims["type"].(string); ok && tokenType != "access" {
			c.JSON(http.StatusUnauthorized, dto.AuthResponse{
				Success:   false,
				Message:   "Unauthorized",
				Error:     "invalid token type",
				RequestID: c.GetString("requestId"),
			})
			c.Abort()
			return
		}

		// Step 6: Store user info in context for controllers
		if userID, ok := claims["user_id"].(string); ok {
			c.Set("userId", userID)
		}
		if email, ok := claims["email"].(string); ok {
			c.Set("userEmail", email)
		}
		if role, ok := claims["role"].(string); ok {
			c.Set("userRole", role)
		}

		// Step 7: Continue to next handler
		c.Next()
	}
}

// RequireRole checks if the authenticated user has at least one of the required roles
func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user role from context (set by JWTAuth middleware)
		userRole, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusForbidden, dto.AuthResponse{
				Success:   false,
				Message:   "Forbidden",
				Error:     "user role not found",
				RequestID: c.GetString("requestId"),
			})
			c.Abort()
			return
		}

		role := userRole.(string)

		// Check if user's role is in the allowed list
		hasPermission := false
		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				hasPermission = true
				break
			}
		}

		if !hasPermission {
			c.JSON(http.StatusForbidden, dto.AuthResponse{
				Success:   false,
				Message:   "Forbidden",
				Error:     "you do not have permission to access this resource",
				RequestID: c.GetString("requestId"),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// OptionalAuth tries to authenticate but does not fail if no token is provided
// Use this for endpoints that work with or without authentication
func OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.Next()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, gin.Error{Err: http.ErrAbortHandler}
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.Next()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.Next()
			return
		}

		if tokenType, ok := claims["type"].(string); ok && tokenType != "access" {
			c.Next()
			return
		}

		if userID, ok := claims["user_id"].(string); ok {
			c.Set("userId", userID)
		}
		if email, ok := claims["email"].(string); ok {
			c.Set("userEmail", email)
		}
		if role, ok := claims["role"].(string); ok {
			c.Set("userRole", role)
		}

		c.Next()
	}
}
