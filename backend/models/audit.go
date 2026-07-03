package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuditAction string

const (
	ActionCreate AuditAction = "CREATE"
	ActionUpdate AuditAction = "UPDATE"
	ActionDelete AuditAction = "DELETE"
	ActionLogin  AuditAction = "LOGIN"
	ActionLogout AuditAction = "LOGOUT"
	ActionView   AuditAction = "VIEW"
)

type AuditLog struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID       primitive.ObjectID `bson:"userId"`
	UserName     string             `bson:"userName"`
	Action       AuditAction        `bson:"action"`
	ResourceType string             `bson:"resourceType"` // e.g., "employee", "leave"
	ResourceID   string             `bson:"resourceId"`
	IPAddress    string             `bson:"ipAddress"`
	UserAgent    string             `bson:"userAgent"`
	Details      string             `bson:"details"` // JSON string of changes
	Timestamp    time.Time          `bson:"timestamp"`
}
