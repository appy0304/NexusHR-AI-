package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentType string

const (
	Resume      DocumentType = "resume"
	OfferLetter DocumentType = "offer_letter"
	Certificate DocumentType = "certificate"
	IdentityDoc DocumentType = "identity_doc"
	Contract    DocumentType = "contract"
)

type Document struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	EmployeeID   primitive.ObjectID `bson:"employeeId"`
	DocumentType DocumentType       `bson:"documentType"`
	FileName     string             `bson:"fileName"`
	FileType     string             `bson:"fileType"`
	FileSize     int64              `bson:"fileSize"`
	S3Key        string             `bson:"s3Key"`
	S3Bucket     string             `bson:"s3Bucket"`
	Version      int                `bson:"version"`
	UploadedBy   primitive.ObjectID `bson:"uploadedBy"`
	UploadedAt   time.Time          `bson:"uploadedAt"`
	IsDeleted    bool               `bson:"isDeleted"`
}
