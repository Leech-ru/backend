package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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

		field.UUID("image_id", uuid.UUID{}),

		field.String("title").
			NotEmpty(),

		field.String("description").
			Optional().Nillable(),

		field.String("applicationMethod").
			Optional().Nillable(),

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

func (Cosmetics) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).
			Ref("cosmetics").
			Unique().
			Required(),
	}
}

func (Cosmetics) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("is_hidden"),
	}
}
