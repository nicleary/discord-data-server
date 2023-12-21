package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("contents").Annotations(entsql.Annotation{
			Size: 10,
		}).Optional(),
		field.String("something"),
		field.String("something else"),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return nil
}
