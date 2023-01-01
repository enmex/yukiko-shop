package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("email").Unique(),
		field.String("first_name"),
		field.String("last_name"),
		field.Enum("access_type").Values(
			"ADMIN",
			"MANAGER",
			"CUSTOMER",
		).Default("CUSTOMER"),
		field.String("password"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products_in_cart", CartProduct.Type),
	}
}
