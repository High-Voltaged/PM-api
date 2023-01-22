package schema

import (
	"api/types"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProjectTask holds the schema definition for the ProjectTask entity.
type ProjectTask struct {
	ent.Schema
}

// Fields of the ProjectTask.
func (ProjectTask) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			MinLen(4).MaxLen(10),
		field.String("summary").
			MinLen(10).MaxLen(100),
		field.Enum("status").
			Values(types.StatusValues...).
			Default(types.Status.IN_PROGRESS),
		field.Enum("priority").
			Values(types.PriorityValues...).
			Default(types.Priority.MEDIUM),
		field.Time("due_date").
			SchemaType(map[string]string{
				dialect.Postgres: "date",
			}).
			Default(time.Now),
		field.Int("project").
			Optional(),
	}
}

// Edges of the ProjectTask.
func (ProjectTask) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project_id", Project.Type).
			Ref("project_tasks").
			Field("project").
			Unique(),
	}
}
