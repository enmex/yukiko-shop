// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldAccessType holds the string denoting the access_type field in the database.
	FieldAccessType = "access_type"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgeProductsInCart holds the string denoting the products_in_cart edge name in mutations.
	EdgeProductsInCart = "products_in_cart"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ProductsInCartTable is the table that holds the products_in_cart relation/edge.
	ProductsInCartTable = "cart_products"
	// ProductsInCartInverseTable is the table name for the CartProduct entity.
	// It exists in this package in order to avoid circular dependency with the "cartproduct" package.
	ProductsInCartInverseTable = "cart_products"
	// ProductsInCartColumn is the table column denoting the products_in_cart relation/edge.
	ProductsInCartColumn = "customer_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldFirstName,
	FieldLastName,
	FieldAccessType,
	FieldPassword,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// AccessType defines the type for the "access_type" enum field.
type AccessType string

// AccessTypeCUSTOMER is the default value of the AccessType enum.
const DefaultAccessType = AccessTypeCUSTOMER

// AccessType values.
const (
	AccessTypeADMIN    AccessType = "ADMIN"
	AccessTypeMANAGER  AccessType = "MANAGER"
	AccessTypeCUSTOMER AccessType = "CUSTOMER"
)

func (at AccessType) String() string {
	return string(at)
}

// AccessTypeValidator is a validator for the "access_type" field enum values. It is called by the builders before save.
func AccessTypeValidator(at AccessType) error {
	switch at {
	case AccessTypeADMIN, AccessTypeMANAGER, AccessTypeCUSTOMER:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for access_type field: %q", at)
	}
}
