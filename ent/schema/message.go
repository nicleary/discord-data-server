package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("contents").Annotations(entsql.Annotation{
			Size: 8192,
		}),
		field.Time("sent_at"),
		field.Int("sender_id"),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sender", User.Type).Ref("messages").Unique().Field("sender_id").Required(),
	}
}

// Indexes of the User
func (Message) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sent_at"),
		index.Fields("sender_id"),
	}
}
