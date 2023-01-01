package adapter

import (
	spec "yukiko-shop/internal/generated/spec/cart"
	"yukiko-shop/internal/repository/ent"
)

func PresentCartProduct(cartProduct *ent.CartProduct) *spec.Product {
	return &spec.Product{
		Id:         cartProduct.ID.String(),
		Name:       cartProduct.Name,
		Price:      cartProduct.Price,
		Quantity:   cartProduct.Quantity,
		PhotoUrl:   cartProduct.PhotoURL,
		CustomerID: cartProduct.CustomerID.String(),
		ProductID:  cartProduct.ProductID.String(),
	}
}
