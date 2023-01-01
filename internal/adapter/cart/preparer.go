package adapter

import (
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/cart"

	"github.com/google/uuid"
)

func PrepareCartProduct(request *spec.AddProductToCartRequest, customerID uuid.UUID) *domain.CartProduct {
	return &domain.CartProduct{
		Name:       request.Name,
		Price:      request.Price,
		Quantity:   1,
		ProductID:  uuid.MustParse(request.ProductID),
		CustomerID: customerID,
		PhotoUrl:   request.PhotoUrl,
	}
}
