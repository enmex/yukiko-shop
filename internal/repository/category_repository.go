package repository

import (
	"context"
	"strings"
	"yukiko-shop/internal/domain"
	"yukiko-shop/internal/repository/ent"
	"yukiko-shop/internal/repository/ent/category"

	"github.com/google/uuid"
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
		SetID(categoryDomain.ID).
		SetName(categoryDomain.Name)

	if categoryDomain.PhotoURL != nil {
		qb = qb.SetPhotoURL(*categoryDomain.PhotoURL)
	}

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

func (repo *CategoryRepository) GetCategories(ctx context.Context, categoryType *string) ([]*ent.Category, error) {
	qb := repo.Client.Category.Query()

	if categoryType != nil {
		if strings.EqualFold(*categoryType, "root") {
			qb = qb.Where(category.ParentCategoryIsNil())
		} else if strings.EqualFold(*categoryType, "leaf") {
			qb = qb.Where(category.Not(category.HasChildren()))
		}
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
		Where(category.ParentCategoryEQ(categoryDomain.ID)).
		All(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, domain.CategoryNotFoundErr
		}

		return nil, err
	}

	return categoriesEnt, nil
}

func (repo *CategoryRepository) GetCategoryByID(ctx context.Context, categoryID uuid.UUID) (*ent.Category, error) {
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
				pq.WithCategory()
			},
		).
		Where(category.IDEQ(categoryID)).
		Only(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, domain.CategoryNotFoundErr
		}
		return nil, err
	}
	return categoryEnt, nil
}

func (repo *CategoryRepository) UpdateCategoryPhotoUrl(ctx context.Context, categoryID uuid.UUID, photoUrl string) (*ent.Category, error) {
	categoryEnt, err := repo.Client.Category.
		UpdateOneID(categoryID).
		SetPhotoURL(photoUrl).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return categoryEnt, nil
}

func (repo *CategoryRepository) UpdateCategoriesPhotoUrl(ctx context.Context, categories []*domain.Category) error {
	tx, err := repo.Client.Tx(ctx)
	if err != nil {
		return err
	}

	for _, category := range categories {
		_, err := tx.Category.
			UpdateOneID(category.ID).
			SetPhotoURL(*category.PhotoURL).
			Save(ctx)
		if err != nil {
			return rollback(tx, err)
		}
	}

	return tx.Commit()
}

func (repo *CategoryRepository) GetCategoriesIds(ctx context.Context) ([]*ent.Category, error) {
	categoriesEnt, err := repo.Client.Category.
		Query().
		Select("id").
		All(ctx)
	if err != nil {
		return nil, err
	}

	return categoriesEnt, nil
}
