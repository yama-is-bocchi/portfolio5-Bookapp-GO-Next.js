package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
		NotEmpty(),
		field.Int("user_id"),
		field.String("kind").
			NotEmpty(),
		field.String("memo"),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("books").
			Field("user_id").
			Unique().
			Required(),
	}
}
