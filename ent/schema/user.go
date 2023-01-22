package schema

import (
	"api/types"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MinLen(4).MaxLen(10),
		field.String("email").
			Unique(),
		field.Enum("role").
			Values(types.UserRoleValues...).
			Default(types.UserRole.USER),
		field.String("password").
			Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("projects", Project.Type).
			Ref("users"),
	}
}
