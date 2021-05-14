package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GroupOwner holds the schema definition for the GroupOwner entity.
type GroupOwner struct {
	ent.Schema
}

// Fields of the GroupOwner.
func (GroupOwner) Fields() []ent.Field {
	return []ent.Field{
		field.String("owner"),
		field.Int("group_id"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the GroupOwner.
func (GroupOwner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("group", Group.Type).
			Field("group_id").
			Unique().
			Required(),
	}
}