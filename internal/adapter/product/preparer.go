package adapter

import (
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/product"
)

func PrepareProduct(productRequest *spec.CreateProductRequest) *domain.Product {
	return &domain.Product{
		Name: productRequest.Name,
		Category: domain.Category{
			Name: productRequest.CategoryName,
		},
		Description: productRequest.Description,
		Price:       productRequest.Price,
		PhotoURL:    &productRequest.Path,
	}
}
