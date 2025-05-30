// Code generated by ent, DO NOT EDIT.

package refreshtoken

import (
	"Leech-ru/pkg/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLTE(FieldID, id))
}

// Jti applies equality check predicate on the "jti" field. It's identical to JtiEQ.
func Jti(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldJti, v))
}

// JtiEQ applies the EQ predicate on the "jti" field.
func JtiEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldJti, v))
}

// JtiNEQ applies the NEQ predicate on the "jti" field.
func JtiNEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldJti, v))
}

// JtiIn applies the In predicate on the "jti" field.
func JtiIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIn(FieldJti, vs...))
}

// JtiNotIn applies the NotIn predicate on the "jti" field.
func JtiNotIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotIn(FieldJti, vs...))
}

// JtiGT applies the GT predicate on the "jti" field.
func JtiGT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGT(FieldJti, v))
}

// JtiGTE applies the GTE predicate on the "jti" field.
func JtiGTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGTE(FieldJti, v))
}

// JtiLT applies the LT predicate on the "jti" field.
func JtiLT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLT(FieldJti, v))
}

// JtiLTE applies the LTE predicate on the "jti" field.
func JtiLTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLTE(FieldJti, v))
}

// JtiContains applies the Contains predicate on the "jti" field.
func JtiContains(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContains(FieldJti, v))
}

// JtiHasPrefix applies the HasPrefix predicate on the "jti" field.
func JtiHasPrefix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasPrefix(FieldJti, v))
}

// JtiHasSuffix applies the HasSuffix predicate on the "jti" field.
func JtiHasSuffix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasSuffix(FieldJti, v))
}

// JtiEqualFold applies the EqualFold predicate on the "jti" field.
func JtiEqualFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEqualFold(FieldJti, v))
}

// JtiContainsFold applies the ContainsFold predicate on the "jti" field.
func JtiContainsFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContainsFold(FieldJti, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.RefreshToken {
	return predicate.RefreshToken(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.RefreshToken {
	return predicate.RefreshToken(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RefreshToken) predicate.RefreshToken {
	return predicate.RefreshToken(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RefreshToken) predicate.RefreshToken {
	return predicate.RefreshToken(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RefreshToken) predicate.RefreshToken {
	return predicate.RefreshToken(sql.NotPredicates(p))
}
