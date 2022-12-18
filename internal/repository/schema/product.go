package schema

import (
	"time"

	"entgo.io/ent"
	//"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("name").Unique(),
		field.String("description"),
		field.String("photo_url").Optional(),
		field.Float("price"),
		field.UUID("category_id", uuid.UUID{}),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("category", Category.Type).
			Ref("products").
			Unique().
			Field("category_id").
			Required(),
	}
}
