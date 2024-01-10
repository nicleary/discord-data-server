package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
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
		field.String("message_id").Unique(),
		field.String("channel_id"),
		field.Int("in_reply_to_id").Optional(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sender", User.Type).
			Ref("messages").
			Unique().
			Field("sender_id").
			Required(),
		edge.To("responders", Message.Type).
			From("in_reply_to").
			Field("in_reply_to_id").
			Unique(),
	}
}

// Indexes of the User
func (Message) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sent_at"),
		index.Fields("sender_id"),
		index.Fields("message_id"),
		index.Fields("channel_id"),
	}
}
