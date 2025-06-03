package schema

import (
	"Leech-ru/internal/domain/types"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Cosmetics holds the schema definition for the Cosmetics entity.
type Cosmetics struct {
	ent.Schema
}

// Fields of the Cosmetics.
func (Cosmetics) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),

		field.Int("category").
			GoType(types.Category(0)).
			NonNegative(),

		field.String("title").
			NotEmpty(),

		field.String("description").
			Optional().Nillable().
			Default(""),

		field.String("applicationMethod").
			Optional().Nillable().
			Default(""),

		field.Int("volume").
			Optional().Nillable().
			Positive(),
	}

}

// Edges of the Cosmetics.
func (Cosmetics) Edges() []ent.Edge {
	return nil
}
