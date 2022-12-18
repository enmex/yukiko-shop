package adapter

import (
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/product"
)

func PrepareCategory(categoryRequest *spec.CreateCategoryRequest) *domain.Category {
	category := &domain.Category{
		Name: categoryRequest.Name,
	}

	if categoryRequest.ParentCategory != nil {
		category.ParentCategory.Name = *categoryRequest.ParentCategory
	}
	return category
}
