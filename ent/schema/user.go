package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// User schema holds the schema definition for the user entity
type User struct {
	ent.Schema
}

// Fields of the Usage
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").Unique().Annotations(entsql.Annotation{
			Size: 64,
		}),
		field.Time("date_joined"),
		field.Bool("is_bot").Default(false),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages", Message.Type),
	}
}

// Indexes of the User
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id"),
	}
}
