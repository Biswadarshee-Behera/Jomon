package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GroupBudget holds the schema definition for the GroupBudget entity.
type GroupBudget struct {
	ent.Schema
}

// Fields of the GroupBudget.
func (GroupBudget) Fields() []ent.Field {
	return []ent.Field{
		field.Int("amount"),
		field.Int("group_id"),
		field.Text("comment").
			Nillable().
			Optional(),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the GroupBudget.
func (GroupBudget) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("group", Group.Type).
			Field("group_id").
			Unique().
			Required(),
	}
}