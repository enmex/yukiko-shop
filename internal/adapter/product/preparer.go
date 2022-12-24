package adapter

import (
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/product"

	"github.com/google/uuid"
)

func PrepareProduct(productRequest *spec.CreateProductRequest) *domain.Product {
	return &domain.Product{
		ID:   uuid.MustParse(productRequest.Id),
		Name: productRequest.Name,
		Category: domain.Category{
			Name: productRequest.CategoryName,
		},
		Description: productRequest.Description,
		Price:       productRequest.Price,
		PhotoURL:    &productRequest.PhotoUrl,
	}
}
