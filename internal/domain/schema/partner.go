package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Partner struct {
	ent.Schema
}

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

func (Partner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("links", PartnerLink.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}
