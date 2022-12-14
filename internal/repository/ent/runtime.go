// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"
	"yukiko-shop/internal/repository/ent/cartproduct"
	"yukiko-shop/internal/repository/ent/category"
	"yukiko-shop/internal/repository/ent/product"
	"yukiko-shop/internal/repository/ent/user"
	"yukiko-shop/internal/repository/schema"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cartproductFields := schema.CartProduct{}.Fields()
	_ = cartproductFields
	// cartproductDescID is the schema descriptor for id field.
	cartproductDescID := cartproductFields[0].Descriptor()
	// cartproduct.DefaultID holds the default value on creation for the id field.
	cartproduct.DefaultID = cartproductDescID.Default.(func() uuid.UUID)
	categoryFields := schema.Category{}.Fields()
	_ = categoryFields
	// categoryDescID is the schema descriptor for id field.
	categoryDescID := categoryFields[0].Descriptor()
	// category.DefaultID holds the default value on creation for the id field.
	category.DefaultID = categoryDescID.Default.(func() uuid.UUID)
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescCreatedAt is the schema descriptor for created_at field.
	productDescCreatedAt := productFields[6].Descriptor()
	// product.DefaultCreatedAt holds the default value on creation for the created_at field.
	product.DefaultCreatedAt = productDescCreatedAt.Default.(time.Time)
	// productDescID is the schema descriptor for id field.
	productDescID := productFields[0].Descriptor()
	// product.DefaultID holds the default value on creation for the id field.
	product.DefaultID = productDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
