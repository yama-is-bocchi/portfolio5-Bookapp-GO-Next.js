// Code generated by ent, DO NOT EDIT.

package token

import (
	"Bookapp/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Token {
	return predicate.Token(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Token {
	return predicate.Token(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Token {
	return predicate.Token(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Token {
	return predicate.Token(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Token {
	return predicate.Token(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Token {
	return predicate.Token(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldUserID, v))
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldToken, v))
}

// AccesDate applies equality check predicate on the "acces_date" field. It's identical to AccesDateEQ.
func AccesDate(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldAccesDate, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Token {
	return predicate.Token(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Token {
	return predicate.Token(sql.FieldNotIn(FieldUserID, vs...))
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.Token {
	return predicate.Token(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.Token {
	return predicate.Token(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.Token {
	return predicate.Token(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.Token {
	return predicate.Token(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.Token {
	return predicate.Token(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.Token {
	return predicate.Token(sql.FieldLTE(FieldToken, v))
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.Token {
	return predicate.Token(sql.FieldContains(FieldToken, v))
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.Token {
	return predicate.Token(sql.FieldHasPrefix(FieldToken, v))
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.Token {
	return predicate.Token(sql.FieldHasSuffix(FieldToken, v))
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.Token {
	return predicate.Token(sql.FieldEqualFold(FieldToken, v))
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.Token {
	return predicate.Token(sql.FieldContainsFold(FieldToken, v))
}

// AccesDateEQ applies the EQ predicate on the "acces_date" field.
func AccesDateEQ(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldEQ(FieldAccesDate, v))
}

// AccesDateNEQ applies the NEQ predicate on the "acces_date" field.
func AccesDateNEQ(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldNEQ(FieldAccesDate, v))
}

// AccesDateIn applies the In predicate on the "acces_date" field.
func AccesDateIn(vs ...time.Time) predicate.Token {
	return predicate.Token(sql.FieldIn(FieldAccesDate, vs...))
}

// AccesDateNotIn applies the NotIn predicate on the "acces_date" field.
func AccesDateNotIn(vs ...time.Time) predicate.Token {
	return predicate.Token(sql.FieldNotIn(FieldAccesDate, vs...))
}

// AccesDateGT applies the GT predicate on the "acces_date" field.
func AccesDateGT(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldGT(FieldAccesDate, v))
}

// AccesDateGTE applies the GTE predicate on the "acces_date" field.
func AccesDateGTE(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldGTE(FieldAccesDate, v))
}

// AccesDateLT applies the LT predicate on the "acces_date" field.
func AccesDateLT(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldLT(FieldAccesDate, v))
}

// AccesDateLTE applies the LTE predicate on the "acces_date" field.
func AccesDateLTE(v time.Time) predicate.Token {
	return predicate.Token(sql.FieldLTE(FieldAccesDate, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Token) predicate.Token {
	return predicate.Token(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Token) predicate.Token {
	return predicate.Token(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Token) predicate.Token {
	return predicate.Token(sql.NotPredicates(p))
}
