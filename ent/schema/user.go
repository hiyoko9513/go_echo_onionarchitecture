package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hiyoko-echo/ent/util"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(util.ID("")).
			DefaultFunc(func() util.ID {
				return util.NewULID()
			}).
			Immutable().
			Unique(),
		field.String("name"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}
