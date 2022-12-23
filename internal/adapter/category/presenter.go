package adapter

import (
	adapter "yukiko-shop/internal/adapter/product"
	spec "yukiko-shop/internal/generated/spec/product"
	"yukiko-shop/internal/repository/ent"
)

func PresentCategory(categoryEnt *ent.Category) *spec.Category {
	childrenEnt, err := categoryEnt.Edges.ChildrenOrErr()
	var children []spec.Category
	if err == nil {
		for _, childEnt := range childrenEnt {
			children = append(children, spec.Category{
				Id:   childEnt.ID.String(),
				Name: childEnt.Name,
				Parent: &spec.Category{
					Name: childEnt.Edges.Parent.Name,
				},
			})
		}
	}

	parentEnt, err := categoryEnt.Edges.ParentOrErr()
	var parent *spec.Category
	if err == nil {
		parent = &spec.Category{
			Name: parentEnt.Name,
		}
	}

	var products []spec.Product
	productsEnt, err := categoryEnt.Edges.ProductsOrErr()
	if err == nil {
		for _, productEnt := range productsEnt {
			products = append(products, *adapter.PresentProduct(productEnt))
		}
	}

	return &spec.Category{
		Id:       categoryEnt.ID.String(),
		Name:     categoryEnt.Name,
		PhotoUrl: &categoryEnt.PhotoURL,
		Children: children,
		Parent:   parent,
		Products: products,
	}
}
