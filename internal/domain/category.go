package domain

import "github.com/google/uuid"

type Category struct {
	ID                 uuid.UUID
	Name               string
	PhotoURL           *string
	ParentCategory     *Category
	ChildrenCategories []Category
	Products           []Product
}
