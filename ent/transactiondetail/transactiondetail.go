// Code generated by entc, DO NOT EDIT.

package transactiondetail

import (
	"time"
)

const (
	// Label holds the string label denoting the transactiondetail type in the database.
	Label = "transaction_detail"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldRequestID holds the string denoting the request_id field in the database.
	FieldRequestID = "request_id"
	// FieldGroupID holds the string denoting the group_id field in the database.
	FieldGroupID = "group_id"
	// FieldTarget holds the string denoting the target field in the database.
	FieldTarget = "target"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeTransaction holds the string denoting the transaction edge name in mutations.
	EdgeTransaction = "transaction"
	// EdgeRequest holds the string denoting the request edge name in mutations.
	EdgeRequest = "request"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// Table holds the table name of the transactiondetail in the database.
	Table = "transaction_details"
	// TransactionTable is the table the holds the transaction relation/edge.
	TransactionTable = "transactions"
	// TransactionInverseTable is the table name for the Transaction entity.
	// It exists in this package in order to avoid circular dependency with the "transaction" package.
	TransactionInverseTable = "transactions"
	// TransactionColumn is the table column denoting the transaction relation/edge.
	TransactionColumn = "transaction_detail_transaction"
	// RequestTable is the table the holds the request relation/edge.
	RequestTable = "transaction_details"
	// RequestInverseTable is the table name for the Request entity.
	// It exists in this package in order to avoid circular dependency with the "request" package.
	RequestInverseTable = "requests"
	// RequestColumn is the table column denoting the request relation/edge.
	RequestColumn = "request_id"
	// GroupTable is the table the holds the group relation/edge.
	GroupTable = "transaction_details"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "groups"
	// GroupColumn is the table column denoting the group relation/edge.
	GroupColumn = "group_id"
)

// Columns holds all SQL columns for transactiondetail fields.
var Columns = []string{
	FieldID,
	FieldAmount,
	FieldRequestID,
	FieldGroupID,
	FieldTarget,
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
	// DefaultAmount holds the default value on creation for the "amount" field.
	DefaultAmount int
	// DefaultTarget holds the default value on creation for the "target" field.
	DefaultTarget string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
