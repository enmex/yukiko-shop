package product

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	adapter "yukiko-shop/internal/adapter/category"
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/product"
	"yukiko-shop/internal/interfaces"
	"yukiko-shop/pkg/minio"
	"yukiko-shop/pkg/scheduler"
)

type CategoryUseCase struct {
	logger             *logrus.Logger
	minioClient        *minio.MinioClient
	categoryRepository interfaces.CategoryRepository
	errors             chan error
}

func NewCategoryUseCase(
	logger *logrus.Logger,
	minioClient *minio.MinioClient,
	categoryRepository interfaces.CategoryRepository) *CategoryUseCase {

	return &CategoryUseCase{
		logger:             logger,
		categoryRepository: categoryRepository,
		minioClient:        minioClient,
		errors:             make(chan error),
	}
}

func (u *CategoryUseCase) StartScheduler(ctx context.Context, cfg *scheduler.Config) {
	scheduler := scheduler.NewScheduler(cfg, func(ctx context.Context) error {
		categoriesEnt, err := u.categoryRepository.GetCategoriesIds(ctx)
		if err != nil {
			return err
		}

		var categories []*domain.Category
		for _, categoryEnt := range categoriesEnt {
			url, err := u.minioClient.GetObject(ctx, fmt.Sprintf("image_%s.jpg", categoryEnt.ID))
			if err != nil {
				return err
			}

			categories = append(categories, &domain.Category{
				ID:       categoryEnt.ID,
				PhotoURL: url,
			})
		}

		if err := u.categoryRepository.UpdateCategoriesPhotoUrl(ctx, categories); err != nil {
			return err
		}

		u.logger.Infoln("Category photo urls updated successfully")
		return nil
	})

	go scheduler.Start(ctx)
	go func() {
		err := <-scheduler.Error()
		u.errors <- err
	}()
}

func (u *CategoryUseCase) CreateCategory(ctx context.Context, category *domain.Category) (*spec.Category, error) {
	categoryEnt, err := u.categoryRepository.CreateCategory(ctx, category)
	if err != nil {
		return nil, err
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
			Name:     categoryEnt.Name,
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
			Name:     categoryEnt.Name,
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

func (u *CategoryUseCase) ReadError() error {
	return <-u.errors
}
