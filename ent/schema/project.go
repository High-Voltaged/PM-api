package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
		field.Time("start_at").
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}),
		field.Time("end_at").
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}),
		field.Int("creator"),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
		edge.To("project_tasks", ProjectTask.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
