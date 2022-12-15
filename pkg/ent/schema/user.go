package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "user",
		},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name").MaxLen(32).Default(""),
		field.String("email").MaxLen(32).Default(""),
		field.String("password").MaxLen(256).Default(""),
		field.String("salt").MaxLen(16).Default(""),
		field.Uint8("role").Default(0),
		field.Int64("registed_at").Immutable().Default(0),
		field.Int64("last_login_at").Default(0),
		field.Int64("created_at").Immutable().Default(0),
		field.Int64("updated_at").Default(0),
		field.Int64("deleted_at").Default(0),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
