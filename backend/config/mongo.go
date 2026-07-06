package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	DB_NAME = "employeeDB"

	EMPLOYEES_COLLECTION = "employees"
	AUDIT_COLLECTION     = "audit_logs"
	DOCUMENTS_COLLECTION = "documents"
	AUTH_COLLECTION      = "auth_users"
)

var MongoClient *mongo.Client

func ConnectDB() {
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}
	clientOptions := options.Client().ApplyURI(mongoURI).
		SetMaxPoolSize(50). // Enterprise: handle many concurrent connections
		SetMinPoolSize(10).
		SetMaxConnIdleTime(30 * time.Second)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	MongoClient = client
	fmt.Println("MongoDB Connected Successfully")

	CreateIndexes()
}

// CreateIndexes creates all necessary database indexes for performance
func CreateIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	collection := MongoClient.Database(DB_NAME).Collection(EMPLOYEES_COLLECTION)

	// 1. Unique index on EmployeeID
	employeeIDIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "employeeId", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(ctx, employeeIDIndex)
	if err != nil {
		fmt.Printf("Error creating employeeId index: %v\n", err)
	}

	// 2. Unique index on Email
	emailIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err = collection.Indexes().CreateOne(ctx, emailIndex)
	if err != nil {
		fmt.Printf("Error creating email index: %v\n", err)
	}

	// 3. Compound index on Department + EmploymentStatus
	deptStatusIndex := mongo.IndexModel{
		Keys: bson.D{
			{Key: "department", Value: 1},
			{Key: "employmentStatus", Value: 1},
		},
	}
	_, err = collection.Indexes().CreateOne(ctx, deptStatusIndex)
	if err != nil {
		fmt.Printf("Error creating department+status index: %v\n", err)
	}

	// 4. Index on ManagerID
	managerIndex := mongo.IndexModel{
		Keys: bson.D{{Key: "managerId", Value: 1}},
	}
	_, err = collection.Indexes().CreateOne(ctx, managerIndex)
	if err != nil {
		fmt.Printf("Error creating managerId index: %v\n", err)
	}

	// 5. Index on JoiningDate
	joiningDateIndex := mongo.IndexModel{
		Keys: bson.D{{Key: "joiningDate", Value: 1}},
	}
	_, err = collection.Indexes().CreateOne(ctx, joiningDateIndex)
	if err != nil {
		fmt.Printf("Error creating joiningDate index: %v\n", err)
	}

	// 6. Text index for full-text search
	textIndex := mongo.IndexModel{
		Keys: bson.D{
			{Key: "firstName", Value: "text"},
			{Key: "lastName", Value: "text"},
			{Key: "email", Value: "text"},
			{Key: "skills", Value: "text"},
		},
	}
	_, err = collection.Indexes().CreateOne(ctx, textIndex)
	if err != nil {
		fmt.Printf("Error creating text index: %v\n", err)
	}

	fmt.Println("All indexes created successfully")
}
