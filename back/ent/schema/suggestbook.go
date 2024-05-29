package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SuggestBook holds the schema definition for the SuggestBook entity.
type SuggestBook struct {
	ent.Schema
}

// Fields of the SuggestBook.
func (SuggestBook) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
		Unique(),
		field.String("kind").
			NotEmpty(),
		field.Int("price").
		Positive(),
		field.String("memo"),
	}
}

// Edges of the SuggestBook.
func (SuggestBook) Edges() []ent.Edge {
	return nil
}
