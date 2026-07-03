package services

import (
	"context"
	"errors"
	"time"

	"simple-go-api/dao"
	"simple-go-api/models"

	"simple-go-api/config"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	jwtSecret          = "EMPLOYEE-MANAGEMENT-PLATFORM-2026-SECRET-KEY-CHANGE-IN-PRODUCTION"
	accessTokenExpiry  = 15 * time.Minute
	refreshTokenExpiry = 7 * 24 * time.Hour // 7 days
)

// Login authenticates a user with email and password
func Login(email, password string) (*models.LoginResponse, error) {
	// Step 1: Find user by email
	user, err := dao.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Step 2: Check if user is active
	if !user.IsActive {
		return nil, errors.New("account is deactivated. Contact HR admin.")
	}

	// Step 3: Verify password using bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Step 4: Generate JWT tokens
	accessToken, refreshToken, err := generateTokens(user)
	if err != nil {
		return nil, errors.New("failed to generate tokens")
	}

	// Step 5: Store refresh token hash in database
	storeRefreshToken(user.ID, refreshToken)

	// Step 6: Update last login time
	dao.UpdateLastLogin(user.ID)

	// Step 7: Build response
	userInfo := models.UserInfo{
		ID:    user.ID.Hex(),
		Email: user.Email,
		Role:  user.Role,
		Name:  getEmployeeName(user.EmployeeID),
	}

	return &models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(accessTokenExpiry.Seconds()),
		User:         userInfo,
	}, nil
}

// RefreshToken validates a refresh token and issues new tokens
func RefreshToken(oldRefreshToken string) (*models.LoginResponse, error) {
	// Step 1: Validate the refresh token
	claims, err := validateToken(oldRefreshToken, "refresh")
	if err != nil {
		return nil, errors.New("invalid or expired refresh token")
	}

	// Step 2: Get user ID from token
	userIDStr, ok := claims["user_id"].(string)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		return nil, errors.New("invalid user ID in token")
	}

	// Step 3: Find user in database
	user, err := dao.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Step 4: Check if user is still active
	if !user.IsActive {
		return nil, errors.New("account is deactivated")
	}

	// Step 5: Generate new tokens
	accessToken, newRefreshToken, err := generateTokens(user)
	if err != nil {
		return nil, errors.New("failed to generate tokens")
	}

	// Step 6: Store new refresh token
	storeRefreshToken(user.ID, newRefreshToken)

	// Step 7: Build response
	userInfo := models.UserInfo{
		ID:    user.ID.Hex(),
		Email: user.Email,
		Role:  user.Role,
		Name:  getEmployeeName(user.EmployeeID),
	}

	return &models.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    int64(accessTokenExpiry.Seconds()),
		User:         userInfo,
	}, nil
}

// Logout invalidates the refresh token
func Logout(userID primitive.ObjectID) error {
	return dao.InvalidateRefreshToken(userID)
}

// generateTokens creates both access and refresh tokens
func generateTokens(user *models.AuthUser) (string, string, error) {
	// Access token claims
	accessClaims := jwt.MapClaims{
		"user_id": user.ID.Hex(),
		"email":   user.Email,
		"role":    user.Role,
		"type":    "access",
		"exp":     time.Now().Add(accessTokenExpiry).Unix(),
		"iat":     time.Now().Unix(),
	}

	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := access_token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", "", err
	}

	// Refresh token claims
	refreshClaims := jwt.MapClaims{
		"user_id": user.ID.Hex(),
		"email":   user.Email,
		"type":    "refresh",
		"exp":     time.Now().Add(refreshTokenExpiry).Unix(),
		"iat":     time.Now().Unix(),
	}

	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := refresh_token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// validateToken validates a JWT token and checks its type
func validateToken(tokenString, expectedType string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Check token type
	if tokenType, ok := claims["type"].(string); ok && tokenType != expectedType {
		return nil, errors.New("invalid token type")
	}

	return claims, nil
}

// storeRefreshToken saves the refresh token hash in the database
func storeRefreshToken(userID primitive.ObjectID, refreshToken string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	collection := config.MongoClient.Database(config.DB_NAME).Collection("auth_users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection.UpdateOne(
		ctx,
		bson.M{"_id": userID},
		bson.M{
			"$set": bson.M{
				"refreshTokenHash":      string(hash),
				"refreshTokenExpiredAt": time.Now().Add(refreshTokenExpiry),
			},
		},
	)
}

// getEmployeeName fetches the employee's name from the employee collection
func getEmployeeName(employeeID primitive.ObjectID) string {
	if employeeID.IsZero() {
		return ""
	}

	collection := config.MongoClient.Database(config.DB_NAME).Collection("employees")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var result struct {
		FirstName string `bson:"firstName"`
		LastName  string `bson:"lastName"`
	}

	err := collection.FindOne(ctx, bson.M{"_id": employeeID}).Decode(&result)
	if err != nil {
		return ""
	}

	return result.FirstName + " " + result.LastName
}

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a password with a hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
