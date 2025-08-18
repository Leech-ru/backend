package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MainPage holds the schema definition for the MainPage entity.
type MainPage struct {
	ent.Schema
}

// Fields of the MainPage.
func (MainPage) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Unique(),

		field.UUID("image_id", uuid.UUID{}),

		field.String("title").
			NotEmpty(),

		field.String("content").
			NotEmpty(),

		field.String("Href").
			NotEmpty(),

		field.Bool("IsHidden").
			Default(false),

		field.Bool("fluid").
			Default(false),
	}
}

// Edges of the MainPage.
func (MainPage) Edges() []ent.Edge {
	return nil
}
