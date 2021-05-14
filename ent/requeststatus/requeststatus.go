// Code generated by entc, DO NOT EDIT.

package requeststatus

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the requeststatus type in the database.
	Label = "request_status"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldRequestID holds the string denoting the request_id field in the database.
	FieldRequestID = "request_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldReason holds the string denoting the reason field in the database.
	FieldReason = "reason"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeRequest holds the string denoting the request edge name in mutations.
	EdgeRequest = "request"
	// Table holds the table name of the requeststatus in the database.
	Table = "request_status"
	// RequestTable is the table the holds the request relation/edge.
	RequestTable = "request_status"
	// RequestInverseTable is the table name for the Request entity.
	// It exists in this package in order to avoid circular dependency with the "request" package.
	RequestInverseTable = "requests"
	// RequestColumn is the table column denoting the request relation/edge.
	RequestColumn = "request_id"
)

// Columns holds all SQL columns for requeststatus fields.
var Columns = []string{
	FieldID,
	FieldCreatedBy,
	FieldRequestID,
	FieldStatus,
	FieldReason,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// StatusSubmitted is the default value of the Status enum.
const DefaultStatus = StatusSubmitted

// Status values.
const (
	StatusSubmitted   Status = "submitted"
	StatusFixRequired Status = "fix_required"
	StatusAccepted    Status = "accepted"
	StatusCompleted   Status = "completed"
	StatusRejected    Status = "rejected"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusSubmitted, StatusFixRequired, StatusAccepted, StatusCompleted, StatusRejected:
		return nil
	default:
		return fmt.Errorf("requeststatus: invalid enum value for status field: %q", s)
	}
}