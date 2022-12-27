package product

import (
	"context"
	adapter "yukiko-shop/internal/adapter/category"
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/product"
	"yukiko-shop/internal/interfaces"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CategoryUseCase struct {
	logger             *logrus.Logger
	categoryRepository interfaces.CategoryRepository
}

func NewCategoryUseCase(
	logger *logrus.Logger,
	categoryRepository interfaces.CategoryRepository) *CategoryUseCase {

	return &CategoryUseCase{
		logger:             logger,
		categoryRepository: categoryRepository,
	}
}

func (u *CategoryUseCase) CreateCategory(ctx context.Context, category *domain.Category) error {
	_, err := u.categoryRepository.CreateCategory(ctx, category)
	if err != nil {
		return err
	}

	return nil
}

func (u *CategoryUseCase) GetCategories(ctx context.Context, categoryType *spec.GetCategoriesParamsType) ([]*spec.Category, error) {
	var categoryTypeString *string
	if categoryType != nil {
		str := string(*categoryType)
		categoryTypeString = &str
	}

	categoriesEnt, err := u.categoryRepository.GetCategories(ctx, categoryTypeString)
	if err != nil {
		return nil, err
	}

	var categories []*spec.Category
	for _, categoryEnt := range categoriesEnt {
		categories = append(categories, &spec.Category{
			Id:       categoryEnt.ID.String(),
			Name:     categoryEnt.Name,
			PhotoUrl: &categoryEnt.PhotoURL,
		})
	}

	return categories, nil
}

func (u *CategoryUseCase) GetSubCategories(ctx context.Context, categoryID uuid.UUID) ([]*spec.Category, error) {
	categoriesEnt, err := u.categoryRepository.GetCategoryChildren(ctx, &domain.Category{
		ID: categoryID,
	})
	if err != nil {
		return nil, err
	}

	var categories []*spec.Category
	for _, categoryEnt := range categoriesEnt {
		categories = append(categories, &spec.Category{
			Id:       categoryEnt.ID.String(),
			Name:     categoryEnt.Name,
			PhotoUrl: &categoryEnt.PhotoURL,
		})
	}

	return categories, nil
}

func (u *CategoryUseCase) GetCategoryByID(ctx context.Context, categoryID uuid.UUID) (*spec.Category, error) {
	categoryEnt, err := u.categoryRepository.GetCategoryByID(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	return adapter.PresentCategory(categoryEnt), nil
}
