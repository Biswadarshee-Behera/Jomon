// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/traPtitech/Jomon/ent/group"
	"github.com/traPtitech/Jomon/ent/groupbudget"
)

// GroupBudget is the model entity for the GroupBudget schema.
type GroupBudget struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount int `json:"amount,omitempty"`
	// Comment holds the value of the "comment" field.
	Comment *string `json:"comment,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupBudgetQuery when eager-loading is set.
	Edges              GroupBudgetEdges `json:"edges"`
	group_group_budget *uuid.UUID
}

// GroupBudgetEdges holds the relations/edges for other nodes in the graph.
type GroupBudgetEdges struct {
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// Transaction holds the value of the transaction edge.
	Transaction []*Transaction `json:"transaction,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupBudgetEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[0] {
		if e.Group == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// TransactionOrErr returns the Transaction value or an error if the edge
// was not loaded in eager-loading.
func (e GroupBudgetEdges) TransactionOrErr() ([]*Transaction, error) {
	if e.loadedTypes[1] {
		return e.Transaction, nil
	}
	return nil, &NotLoadedError{edge: "transaction"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupBudget) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case groupbudget.FieldAmount:
			values[i] = new(sql.NullInt64)
		case groupbudget.FieldComment:
			values[i] = new(sql.NullString)
		case groupbudget.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case groupbudget.FieldID:
			values[i] = new(uuid.UUID)
		case groupbudget.ForeignKeys[0]: // group_group_budget
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type GroupBudget", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupBudget fields.
func (gb *GroupBudget) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case groupbudget.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				gb.ID = *value
			}
		case groupbudget.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				gb.Amount = int(value.Int64)
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
		case groupbudget.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field group_group_budget", values[i])
			} else if value.Valid {
				gb.group_group_budget = new(uuid.UUID)
				*gb.group_group_budget = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryGroup queries the "group" edge of the GroupBudget entity.
func (gb *GroupBudget) QueryGroup() *GroupQuery {
	return (&GroupBudgetClient{config: gb.config}).QueryGroup(gb)
}

// QueryTransaction queries the "transaction" edge of the GroupBudget entity.
func (gb *GroupBudget) QueryTransaction() *TransactionQuery {
	return (&GroupBudgetClient{config: gb.config}).QueryTransaction(gb)
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
	_tx, ok := gb.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroupBudget is not a transactional entity")
	}
	gb.config.driver = _tx.drv
	return gb
}

// String implements the fmt.Stringer.
func (gb *GroupBudget) String() string {
	var builder strings.Builder
	builder.WriteString("GroupBudget(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gb.ID))
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", gb.Amount))
	builder.WriteString(", ")
	if v := gb.Comment; v != nil {
		builder.WriteString("comment=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("created_at=")
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
