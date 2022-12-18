package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type Category struct {
	ent.Schema
}

// Fields of the User.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name").Unique(),
		field.UUID("parent_category", uuid.UUID{}).Optional(),
	}
}

// Edges of the User.
func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Category.Type).
			From("parent").
			Unique().
			Field("parent_category"),
		edge.To("products", Product.Type),
	}
}
