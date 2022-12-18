package adapter

import (
	spec "yukiko-shop/internal/generated/spec/product"
	"yukiko-shop/internal/repository/ent"
)

func PresentCategory(categoryEnt *ent.Category) *spec.Category {
	return &spec.Category{
		Id: categoryEnt.ID.String(),
	}
}
