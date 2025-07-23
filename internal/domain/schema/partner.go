package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Partner holds the schema definition for the Partner entity.
type Partner struct {
	ent.Schema
}

// Fields of the Partner.
func (Partner) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),

		field.String("name").
			MaxLen(100).
			NotEmpty(),

		field.String("description").
			MaxLen(500).
			Optional().Nillable(),
	}
}

// Edges of the Partner.
func (Partner) Edges() []ent.Edge {
	return nil
}
