package dao

import (
	"context"
	"time"

	"simple-go-api/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LeaveType string
type LeaveStatus string

const (
	LeaveAnnual    LeaveType = "annual"
	LeaveSick      LeaveType = "sick"
	LeaveMaternity LeaveType = "maternity"
	LeavePaternity LeaveType = "paternity"
	LeaveUnpaid    LeaveType = "unpaid"
	LeaveCompOff   LeaveType = "comp_off"

	LeavePending   LeaveStatus = "pending"
	LeaveApproved  LeaveStatus = "approved"
	LeaveRejected  LeaveStatus = "rejected"
	LeaveCancelled LeaveStatus = "cancelled"
)

type LeaveRequest struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	EmployeeID      primitive.ObjectID `bson:"employeeId"`
	LeaveType       LeaveType          `bson:"leaveType"`
	Status          LeaveStatus        `bson:"status"`
	StartDate       time.Time          `bson:"startDate"`
	EndDate         time.Time          `bson:"endDate"`
	Days            float64            `bson:"days"`
	Reason          string             `bson:"reason"`
	ApprovedBy      primitive.ObjectID `bson:"approvedBy,omitempty"`
	ApprovedAt      time.Time          `bson:"approvedAt,omitempty"`
	RejectionReason string             `bson:"rejectionReason,omitempty"`
	CreatedAt       time.Time          `bson:"createdAt"`
	UpdatedAt       time.Time          `bson:"updatedAt"`
}

type LeaveBalance struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	EmployeeID     primitive.ObjectID `bson:"employeeId"`
	Year           int                `bson:"year"`
	TotalAllocated float64            `bson:"totalAllocated"`
	Used           float64            `bson:"used"`
	Remaining      float64            `bson:"remaining"`
	UpdatedAt      time.Time          `bson:"updatedAt"`
}

const (
	leaveRequestsCollection = "leave_requests"
	leaveBalancesCollection = "leave_balances"
)

// CreateLeaveRequest inserts a new leave request into MongoDB
func CreateLeaveRequest(leave *LeaveRequest) error {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(leaveRequestsCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	leave.CreatedAt = time.Now()
	leave.UpdatedAt = time.Now()

	_, err := collection.InsertOne(ctx, leave)
	return err
}

// GetLeaveByID finds a leave request by ID
func GetLeaveByID(id primitive.ObjectID) (*LeaveRequest, error) {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(leaveRequestsCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var leave LeaveRequest
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&leave)
	return &leave, err
}

// GetLeaveRequests finds all leave requests with filtering and pagination
func GetLeaveRequests(page, pageSize int, employeeID primitive.ObjectID, status, leaveType string) ([]LeaveRequest, int, error) {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(leaveRequestsCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{}
	if !employeeID.IsZero() {
		filter["employeeId"] = employeeID
	}
	if status != "" {
		filter["status"] = status
	}
	if leaveType != "" {
		filter["leaveType"] = leaveType
	}

	opts := options.Find().SetSkip(int64((page - 1) * pageSize)).SetLimit(int64(pageSize))

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, 0, err
	}

	var leaves []LeaveRequest
	if err := cursor.All(ctx, &leaves); err != nil {
		return nil, 0, err
	}

	// Get total count
	total, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return leaves, int(total), nil
}

// UpdateLeaveByID updates a leave request by ID
func UpdateLeaveByID(id primitive.ObjectID, leave *LeaveRequest) error {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(leaveRequestsCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	leave.UpdatedAt = time.Now()

	_, err := collection.ReplaceOne(ctx, bson.M{"_id": id}, leave)
	return err
}

// GetLeaveBalance finds the leave balance for an employee in a given year
func GetLeaveBalance(employeeID primitive.ObjectID, year int) (*LeaveBalance, error) {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(leaveBalancesCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var balance LeaveBalance
	err := collection.FindOne(ctx, bson.M{"employeeId": employeeID, "year": year}).Decode(&balance)
	return &balance, err
}

// UpdateLeaveBalance updates the leave balance for an employee
func UpdateLeaveBalance(balance *LeaveBalance) error {
	collection := config.MongoClient.Database(config.DB_NAME).Collection(leaveBalancesCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	balance.UpdatedAt = time.Now()

	_, err := collection.ReplaceOne(ctx, bson.M{"employeeId": balance.EmployeeID, "year": balance.Year}, balance)
	return err
}
