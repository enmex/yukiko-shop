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

func (u *CategoryUseCase) GetCategories(ctx context.Context, main *bool, leaf *bool) ([]string, error) {
	categoriesEnt, err := u.categoryRepository.GetCategories(ctx, main, leaf)
	if err != nil {
		return nil, err
	}

	var categories []string
	for _, category := range categoriesEnt {
		categories = append(categories, category.Name)
	}

	return categories, nil
}

func (u *CategoryUseCase) GetSubCategories(ctx context.Context, categoryName string) ([]string, error) {
	categoriesEnt, err := u.categoryRepository.GetCategoryChildren(ctx, &domain.Category{
		Name: categoryName,
	})
	if err != nil {
		return nil, err
	}

	var categories []string
	for _, categoryEnt := range categoriesEnt {
		categories = append(categories, categoryEnt.Name)
	}

	return categories, nil
}

func (u *CategoryUseCase) GetCategoryByName(ctx context.Context, categoryName string) (*spec.Category, error) {
	categoryEnt, err := u.categoryRepository.GetCategoryByName(ctx, categoryName)
	if err != nil {
		return nil, err
	}

	return adapter.PresentCategory(categoryEnt), nil
}