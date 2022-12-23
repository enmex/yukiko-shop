package product

import (
	"context"
	"fmt"
	adapter "yukiko-shop/internal/adapter/category"
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/product"
	"yukiko-shop/internal/interfaces"
	"yukiko-shop/pkg/minio"

	"github.com/sirupsen/logrus"
)

type CategoryUseCase struct {
	logger             *logrus.Logger
	minioClient        *minio.MinioClient
	categoryRepository interfaces.CategoryRepository
}

func NewCategoryUseCase(logger *logrus.Logger, minioClient *minio.MinioClient, categoryRepository interfaces.CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{
		logger:             logger,
		categoryRepository: categoryRepository,
		minioClient:        minioClient,
	}
}

func (u *CategoryUseCase) CreateCategory(ctx context.Context, category *domain.Category) (*spec.Category, error) {
	categoryEnt, err := u.categoryRepository.CreateCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	if category.PhotoURL != nil {
		fileName := fmt.Sprintf("category_%s.jpg", categoryEnt.ID)
		url, err := u.minioClient.UploadFile(ctx, fileName, *category.PhotoURL, "jpg")
		if err != nil {
			return nil, err
		}

		categoryEnt, err = u.categoryRepository.UpdateCategoryPhotoUrl(ctx, categoryEnt.ID, *url)
		if err != nil {
			return nil, err
		}
	}

	return adapter.PresentCategory(categoryEnt), nil
}

func (u *CategoryUseCase) GetCategories(ctx context.Context, main *bool, leaf *bool) ([]*spec.Category, error) {
	categoriesEnt, err := u.categoryRepository.GetCategories(ctx, main, leaf)
	if err != nil {
		return nil, err
	}

	var categories []*spec.Category
	for _, categoryEnt := range categoriesEnt {
		categories = append(categories, &spec.Category{
			Name: categoryEnt.Name,
			PhotoUrl: &categoryEnt.PhotoURL,
		})
	}

	return categories, nil
}

func (u *CategoryUseCase) GetSubCategories(ctx context.Context, categoryName string) ([]*spec.Category, error) {
	categoriesEnt, err := u.categoryRepository.GetCategoryChildren(ctx, &domain.Category{
		Name: categoryName,
	})
	if err != nil {
		return nil, err
	}

	var categories []*spec.Category
	for _, categoryEnt := range categoriesEnt {
		categories = append(categories, &spec.Category{
			Name: categoryEnt.Name,
			PhotoUrl: &categoryEnt.PhotoURL,
		})
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
