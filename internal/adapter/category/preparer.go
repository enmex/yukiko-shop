package adapter

import (
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/product"

	"github.com/google/uuid"
)

func PrepareCategory(categoryRequest *spec.CreateCategoryRequest) *domain.Category {
	category := &domain.Category{
		ID:       uuid.MustParse(categoryRequest.Id),
		Name:     categoryRequest.Name,
		PhotoURL: categoryRequest.PhotoUrl,
	}

	if categoryRequest.Parent != nil {
		category.ParentCategory = &domain.Category{
			Name: *categoryRequest.Parent,
		}
	}
	return category
}
