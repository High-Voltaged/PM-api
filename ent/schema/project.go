package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			MinLen(6).MaxLen(16),
		field.String("description").
			Optional().
			MinLen(20).MaxLen(100),
		field.Time("start_at"),
		field.Time("end_at"),
		field.UUID("creator", uuid.UUID{}),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
