package dao

import (
	"context"
	"errors"
	"time"

	"simple-go-api/config"
	"simple-go-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const authCollection = "auth_users"

// FindByEmail finds an auth user by email
func FindByEmail(email string) (*models.AuthUser, error) {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(authCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.AuthUser
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// FindByID finds an auth user by ID
func FindByID(id primitive.ObjectID) (*models.AuthUser, error) {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(authCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.AuthUser
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// FindByRefreshToken finds a user by their refresh token
func FindByRefreshToken(tokenHash string) (*models.AuthUser, error) {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(authCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.AuthUser
	err := collection.FindOne(ctx, bson.M{"refreshTokenHash": tokenHash, "refreshTokenExpiredAt": bson.M{"$gt": time.Now()}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("invalid refresh token")
		}
		return nil, err
	}

	return &user, nil
}

// CreateAuthUser creates a new auth user (called during user registration or seed)
func CreateAuthUser(user *models.AuthUser) error {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(authCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if email already exists
	existing, err := FindByEmail(user.Email)
	if err == nil && existing != nil {
		return errors.New("email already registered")
	}

	user.IsActive = true
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err = collection.InsertOne(ctx, user)
	return err
}

// UpdateLastLogin updates the last login timestamp
func UpdateLastLogin(userID primitive.ObjectID) error {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(authCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": userID},
		bson.M{
			"$set": bson.M{
				"lastLoginAt": time.Now(),
				"updatedAt":   time.Now(),
			},
		},
	)

	return err
}

// InvalidateRefreshToken marks a refresh token as expired
func InvalidateRefreshToken(userID primitive.ObjectID) error {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(authCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": userID},
		bson.M{
			"$set": bson.M{
				"refreshTokenHash":      "",
				"refreshTokenExpiredAt": time.Time{},
				"updatedAt":             time.Now(),
			},
		},
	)

	return err
}

// CreateDefaultAdmin creates a default admin user for development
// Call this once during setup
func CreateDefaultAdmin() {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(authCollection)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if admin already exists
	count, err := collection.CountDocuments(ctx, bson.M{"role": "super_admin"})
	if err == nil && count > 0 {
		return // Admin already exists
	}

	// Import bcrypt here - see authService.go for hash function
	// For now, this function is a placeholder
	// The actual admin creation should be done via authService
	_ = collection
	_ = ctx
	_ = cancel
}
