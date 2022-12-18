package adapter

import (
	spec "yukiko-shop/internal/generated/spec/product"
	"yukiko-shop/internal/repository/ent"
)

func PresentProduct(productEnt *ent.Product) *spec.Product {
	return &spec.Product{
		Id:           productEnt.ID.String(),
		Name:         productEnt.Name,
		Description:  productEnt.Description,
		PhotoUrl:     productEnt.PhotoURL,
		Price:        productEnt.Price,
		CategoryName: productEnt.Edges.Category.Name,
	}
}
