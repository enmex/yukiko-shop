// Code generated by entc, DO NOT EDIT.

package category

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the category type in the database.
	Label = "category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPhotoURL holds the string denoting the photo_url field in the database.
	FieldPhotoURL = "photo_url"
	// FieldParentCategory holds the string denoting the parent_category field in the database.
	FieldParentCategory = "parent_category"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeProducts holds the string denoting the products edge name in mutations.
	EdgeProducts = "products"
	// Table holds the table name of the category in the database.
	Table = "categories"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "categories"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "parent_category"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "categories"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "parent_category"
	// ProductsTable is the table that holds the products relation/edge.
	ProductsTable = "products"
	// ProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductsInverseTable = "products"
	// ProductsColumn is the table column denoting the products relation/edge.
	ProductsColumn = "category_id"
)

// Columns holds all SQL columns for category fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPhotoURL,
	FieldParentCategory,
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
