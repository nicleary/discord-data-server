package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User schema holds the schema definition for the user entity
type User struct {
	ent.Schema
}

// Fields of the Usage
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").Annotations(entsql.Annotation{
			Size: 64,
		}),
	}
}

// Edges of the User
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages", Message.Type),
	}
}
