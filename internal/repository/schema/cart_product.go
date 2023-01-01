package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type CartProduct struct {
	ent.Schema
}

// Fields of the Product.
func (CartProduct) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name").Unique(),
		field.UUID("product_id", uuid.UUID{}),
		field.UUID("customer_id", uuid.UUID{}),
		field.String("photo_url").Optional(),
		field.Float("price"),
		field.Int("quantity"),
	}
}

// Edges of the Product.
func (CartProduct) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", User.Type).
			Ref("products_in_cart").
			Unique().
			Field("customer_id").
			Required(),
		edge.From("product", Product.Type).
			Ref("products_in_cart").
			Unique().
			Field("product_id").
			Required(),
	}
}
