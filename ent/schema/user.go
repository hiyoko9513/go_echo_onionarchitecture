package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"hiyoko-echo/ent/util"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
// updateやcreateをmixinするとカラムの順序を制御出来ない
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{}
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
		field.Time("created_at").Immutable().Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}
