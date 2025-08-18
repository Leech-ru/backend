package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// News holds the schema definition for the News entity.
type News struct {
	ent.Schema
}

// Fields of the News.
func (News) Fields() []ent.Field {
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

		field.Bool("is_hidden"),
	}
}

// Edges of the News.
func (News) Edges() []ent.Edge {
	return nil
}
