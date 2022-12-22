package adapter

import (
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/product"
)

func PrepareCategory(categoryRequest *spec.CreateCategoryRequest) *domain.Category {
	category := &domain.Category{
		Name: categoryRequest.Name,
	}

	if categoryRequest.Parent != nil {
		category.ParentCategory = &domain.Category{
			Name: *categoryRequest.Parent,
		}
	}
	return category
}