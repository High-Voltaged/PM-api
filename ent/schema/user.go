package schema

import (
	"api/types"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name").MinLen(4).MaxLen(10),
		field.String("email").
			Unique(),
		field.Enum("role").
			GoType(types.UserRole("")),
		field.String("password").
			Sensitive(),
		field.Time("created_at").
			Default(time.Now()),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("projects", Project.Type).
			Ref("users"),
	}
}
