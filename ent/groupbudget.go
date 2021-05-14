// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/traPtitech/Jomon/ent/group"
	"github.com/traPtitech/Jomon/ent/groupbudget"
)

// GroupBudget is the model entity for the GroupBudget schema.
type GroupBudget struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount int `json:"amount,omitempty"`
	// GroupID holds the value of the "group_id" field.
	GroupID int `json:"group_id,omitempty"`
	// Comment holds the value of the "comment" field.
	Comment *string `json:"comment,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupBudgetQuery when eager-loading is set.
	Edges GroupBudgetEdges `json:"edges"`
}

// GroupBudgetEdges holds the relations/edges for other nodes in the graph.
type GroupBudgetEdges struct {
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupBudgetEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[0] {
		if e.Group == nil {
			// The edge group was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupBudget) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case groupbudget.FieldID, groupbudget.FieldAmount, groupbudget.FieldGroupID:
			values[i] = new(sql.NullInt64)
		case groupbudget.FieldComment:
			values[i] = new(sql.NullString)
		case groupbudget.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GroupBudget", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupBudget fields.
func (gb *GroupBudget) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case groupbudget.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gb.ID = int(value.Int64)
		case groupbudget.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				gb.Amount = int(value.Int64)
			}
		case groupbudget.FieldGroupID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				gb.GroupID = int(value.Int64)
			}
		case groupbudget.FieldComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				gb.Comment = new(string)
				*gb.Comment = value.String
			}
		case groupbudget.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gb.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryGroup queries the "group" edge of the GroupBudget entity.
func (gb *GroupBudget) QueryGroup() *GroupQuery {
	return (&GroupBudgetClient{config: gb.config}).QueryGroup(gb)
}

// Update returns a builder for updating this GroupBudget.
// Note that you need to call GroupBudget.Unwrap() before calling this method if this GroupBudget
// was returned from a transaction, and the transaction was committed or rolled back.
func (gb *GroupBudget) Update() *GroupBudgetUpdateOne {
	return (&GroupBudgetClient{config: gb.config}).UpdateOne(gb)
}

// Unwrap unwraps the GroupBudget entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gb *GroupBudget) Unwrap() *GroupBudget {
	tx, ok := gb.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroupBudget is not a transactional entity")
	}
	gb.config.driver = tx.drv
	return gb
}

// String implements the fmt.Stringer.
func (gb *GroupBudget) String() string {
	var builder strings.Builder
	builder.WriteString("GroupBudget(")
	builder.WriteString(fmt.Sprintf("id=%v", gb.ID))
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", gb.Amount))
	builder.WriteString(", group_id=")
	builder.WriteString(fmt.Sprintf("%v", gb.GroupID))
	if v := gb.Comment; v != nil {
		builder.WriteString(", comment=")
		builder.WriteString(*v)
	}
	builder.WriteString(", created_at=")
	builder.WriteString(gb.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// GroupBudgets is a parsable slice of GroupBudget.
type GroupBudgets []*GroupBudget

func (gb GroupBudgets) config(cfg config) {
	for _i := range gb {
		gb[_i].config = cfg
	}
}