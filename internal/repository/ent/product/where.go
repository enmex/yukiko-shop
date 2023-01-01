// Code generated by entc, DO NOT EDIT.

package product

import (
	"time"
	"yukiko-shop/internal/repository/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// PhotoURL applies equality check predicate on the "photo_url" field. It's identical to PhotoURLEQ.
func PhotoURL(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhotoURL), v))
	})
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// CategoryID applies equality check predicate on the "category_id" field. It's identical to CategoryIDEQ.
func CategoryID(v uuid.UUID) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategoryID), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// PhotoURLEQ applies the EQ predicate on the "photo_url" field.
func PhotoURLEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhotoURL), v))
	})
}

// PhotoURLNEQ applies the NEQ predicate on the "photo_url" field.
func PhotoURLNEQ(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhotoURL), v))
	})
}

// PhotoURLIn applies the In predicate on the "photo_url" field.
func PhotoURLIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhotoURL), v...))
	})
}

// PhotoURLNotIn applies the NotIn predicate on the "photo_url" field.
func PhotoURLNotIn(vs ...string) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhotoURL), v...))
	})
}

// PhotoURLGT applies the GT predicate on the "photo_url" field.
func PhotoURLGT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhotoURL), v))
	})
}

// PhotoURLGTE applies the GTE predicate on the "photo_url" field.
func PhotoURLGTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhotoURL), v))
	})
}

// PhotoURLLT applies the LT predicate on the "photo_url" field.
func PhotoURLLT(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhotoURL), v))
	})
}

// PhotoURLLTE applies the LTE predicate on the "photo_url" field.
func PhotoURLLTE(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhotoURL), v))
	})
}

// PhotoURLContains applies the Contains predicate on the "photo_url" field.
func PhotoURLContains(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhotoURL), v))
	})
}

// PhotoURLHasPrefix applies the HasPrefix predicate on the "photo_url" field.
func PhotoURLHasPrefix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhotoURL), v))
	})
}

// PhotoURLHasSuffix applies the HasSuffix predicate on the "photo_url" field.
func PhotoURLHasSuffix(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhotoURL), v))
	})
}

// PhotoURLIsNil applies the IsNil predicate on the "photo_url" field.
func PhotoURLIsNil() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPhotoURL)))
	})
}

// PhotoURLNotNil applies the NotNil predicate on the "photo_url" field.
func PhotoURLNotNil() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPhotoURL)))
	})
}

// PhotoURLEqualFold applies the EqualFold predicate on the "photo_url" field.
func PhotoURLEqualFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhotoURL), v))
	})
}

// PhotoURLContainsFold applies the ContainsFold predicate on the "photo_url" field.
func PhotoURLContainsFold(v string) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhotoURL), v))
	})
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPrice), v))
	})
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPrice), v))
	})
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPrice), v...))
	})
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPrice), v...))
	})
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPrice), v))
	})
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPrice), v))
	})
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPrice), v))
	})
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPrice), v))
	})
}

// CategoryIDEQ applies the EQ predicate on the "category_id" field.
func CategoryIDEQ(v uuid.UUID) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategoryID), v))
	})
}

// CategoryIDNEQ applies the NEQ predicate on the "category_id" field.
func CategoryIDNEQ(v uuid.UUID) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCategoryID), v))
	})
}

// CategoryIDIn applies the In predicate on the "category_id" field.
func CategoryIDIn(vs ...uuid.UUID) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCategoryID), v...))
	})
}

// CategoryIDNotIn applies the NotIn predicate on the "category_id" field.
func CategoryIDNotIn(vs ...uuid.UUID) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCategoryID), v...))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Product {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Product(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// HasCategory applies the HasEdge predicate on the "category" edge.
func HasCategory() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CategoryTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCategoryWith applies the HasEdge predicate on the "category" edge with a given conditions (other predicates).
func HasCategoryWith(preds ...predicate.Category) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CategoryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CategoryTable, CategoryColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProductsInCart applies the HasEdge predicate on the "products_in_cart" edge.
func HasProductsInCart() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductsInCartTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductsInCartTable, ProductsInCartColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductsInCartWith applies the HasEdge predicate on the "products_in_cart" edge with a given conditions (other predicates).
func HasProductsInCartWith(preds ...predicate.CartProduct) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProductsInCartInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ProductsInCartTable, ProductsInCartColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Product) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		p(s.Not())
	})
}
