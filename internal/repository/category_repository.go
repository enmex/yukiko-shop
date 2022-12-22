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

func (repo *CategoryRepository) CreateCategory(ctx context.Context, categoryDomain *domain.Category) (*ent.Category, error) {
	qb := repo.Client.Category.
		Create().
		SetName(categoryDomain.Name)

	if categoryDomain.ParentCategory != nil {
		parent, err := repo.Client.Category.
			Query().
			Where(category.NameEQ(categoryDomain.ParentCategory.Name)).
			Only(ctx)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return nil, domain.CategoryNotFoundErr
			}
			return nil, err
		}

		qb = qb.SetParent(parent)
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

func (repo *CategoryRepository) GetCategories(ctx context.Context, main *bool, leaf *bool) ([]*ent.Category, error) {
	qb := repo.Client.Category.Query()

	if main != nil && *main {
		qb = qb.Where(category.ParentCategoryIsNil())
	}

	if leaf != nil && *leaf {
		qb = qb.Where(category.Not(category.HasChildren()))
	}

	categoriesEnt, err := qb.All(ctx)
	if err != nil {
		return nil, err
	}

	return categoriesEnt, nil
}

func (repo *CategoryRepository) GetCategoryChildren(ctx context.Context, categoryDomain *domain.Category) ([]*ent.Category, error) {
	categoriesEnt, err := repo.Client.Category.
		Query().WithParent().
		Where(category.HasParentWith(category.NameEQ(categoryDomain.Name))).
		All(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, domain.CategoryNotFoundErr
		}

		return nil, err
	}

	return categoriesEnt, nil
}

func (repo *CategoryRepository) GetCategoryByName(ctx context.Context, categoryName string) (*ent.Category, error) {
	categoryEnt, err := repo.Client.Category.
		Query().
		WithParent().
		WithChildren(
			func(cq *ent.CategoryQuery) {
				cq.WithParent()
			},
		).
		WithProducts(
			func(pq *ent.ProductQuery) {
				pq.WithCategory(
					func(cq *ent.CategoryQuery) {
						cq.Select("name")
					},
				)
			},
		).
		Where(category.NameEQ(categoryName)).
		Only(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, domain.CategoryNotFoundErr
		}
		return nil, err
	}
	return categoryEnt, nil
}