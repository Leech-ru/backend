package schema

import (
	"Leech-ru/internal/domain/types"
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
			Default(uuid.New).
			Immutable().
			Unique(),

		field.String("email").
			Unique().
			NotEmpty(),

		field.String("password").
			NotEmpty().
			Sensitive(),

		field.String("name").
			NotEmpty().
			MaxLen(100),

		field.String("surname").
			NotEmpty().
			MaxLen(100),

		field.Int("role").GoType(types.Role(0)).
			GoType(types.Role(0)).
			Default(int(types.RoleUser)),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("refresh_tokens", RefreshToken.Type).
			Unique(),
	}
}
