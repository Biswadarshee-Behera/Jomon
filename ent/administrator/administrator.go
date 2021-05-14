// Code generated by entc, DO NOT EDIT.

package administrator

const (
	// Label holds the string label denoting the administrator type in the database.
	Label = "administrator"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTrapID holds the string denoting the trap_id field in the database.
	FieldTrapID = "trap_id"
	// Table holds the table name of the administrator in the database.
	Table = "administrators"
)

// Columns holds all SQL columns for administrator fields.
var Columns = []string{
	FieldID,
	FieldTrapID,
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