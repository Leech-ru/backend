package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PartnerLink holds the schema definition for the PartnerLink entity.
type PartnerLink struct {
	ent.Schema
}

func (PartnerLink) Fields() []ent.Field {
	return []ent.Field{
		field.String("label").
			NotEmpty(),

		field.String("href").
			NotEmpty(),
	}
}

func (PartnerLink) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("partner", Partner.Type).
			Ref("links").
			Unique().
			Required(),
	}
}
