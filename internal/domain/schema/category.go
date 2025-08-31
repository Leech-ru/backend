package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),

		field.UUID("image_id", uuid.UUID{}),

		field.String("name").
			NotEmpty(),
	}
}

func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cosmetics", Cosmetics.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
