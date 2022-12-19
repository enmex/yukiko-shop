package repository

import (
	"context"
	"strings"
	"yukiko-shop/internal/domain"
	"yukiko-shop/internal/repository/ent"
	"yukiko-shop/internal/repository/ent/category"

	"github.com/sirupsen/logrus"
)

type CategoryRepository struct {
	Client *ent.Client
	log    *logrus.Logger
}

func NewCategoryRepository(client *ent.Client, log *logrus.Logger) *CategoryRepository {
	return &CategoryRepository{
		Client: client,
		log:    log,
	}
}

func (repo *CategoryRepository) CreateCategory(ctx context.Context, category *domain.Category) (*ent.Category, error) {
	qb := repo.Client.Category.
		Create().
		SetName(category.Name)

	if category.ParentCategory != nil {
		qb = qb.SetParentCategory(category.ParentCategory.ID)
	}

	categoryEnt, err := qb.Save(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return nil, domain.CategoryAlreadyExistsErr
		}
		return nil, err
	}

	return categoryEnt, nil
}

func (repo *CategoryRepository) GetCategories(ctx context.Context, main *bool) ([]*ent.Category, error) {
	qb := repo.Client.Category.Query()

	if main != nil && *main {
		qb = qb.Where(category.ParentCategoryIsNil())
	}

	categoriesEnt, err := qb.All(ctx)
	if err != nil {
		return nil, err
	}

	return categoriesEnt, nil
}