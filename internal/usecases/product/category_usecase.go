package product

import (
	"context"
	adapter "yukiko-shop/internal/adapter/category"
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/product"
	"yukiko-shop/internal/interfaces"

	"github.com/sirupsen/logrus"
)

type CategoryUseCase struct {
	logger             *logrus.Logger
	categoryRepository interfaces.CategoryRepository
}

func NewCategoryUseCase(logger *logrus.Logger, categoryRepository interfaces.CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{
		logger:             logger,
		categoryRepository: categoryRepository,
	}
}

func (u *CategoryUseCase) CreateCategory(ctx context.Context, category *domain.Category) (*spec.Category, error) {
	categoryEnt, err := u.categoryRepository.CreateCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	return adapter.PresentCategory(categoryEnt), nil
}
