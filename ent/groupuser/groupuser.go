// Code generated by entc, DO NOT EDIT.

package groupuser

import (
	"time"
)

const (
	// Label holds the string label denoting the groupuser type in the database.
	Label = "group_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldGroupID holds the string denoting the group_id field in the database.
	FieldGroupID = "group_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// Table holds the table name of the groupuser in the database.
	Table = "group_users"
	// GroupTable is the table the holds the group relation/edge.
	GroupTable = "group_users"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "groups"
	// GroupColumn is the table column denoting the group relation/edge.
	GroupColumn = "group_id"
)

// Columns holds all SQL columns for groupuser fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldGroupID,
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
