package schema

import (
	"Leech-ru/internal/domain/types"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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

		field.String("ozon_link").
			Optional().Nillable(),

		field.String("wildberries_link").
			Optional().Nillable(),

		field.Bool("is_hidden"),
	}

}

// Edges of the Cosmetics.
func (Cosmetics) Edges() []ent.Edge {
	return nil
}

func (Cosmetics) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("is_hidden"),
	}
}
