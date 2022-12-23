package repository

import (
	"context"
	"strings"
	"yukiko-shop/internal/domain"
	"yukiko-shop/internal/interfaces"
	"yukiko-shop/internal/repository/ent"
	"yukiko-shop/internal/repository/ent/category"
	"yukiko-shop/internal/repository/ent/product"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var _ interfaces.ProductRepository = (*ProductRepository)(nil)

var (
	defaultRowsLimit = 35
)

type ProductRepository struct {
	Client *ent.Client
	log    *logrus.Logger
}

func NewProductRepository(client *ent.Client, log *logrus.Logger) *ProductRepository {
	return &ProductRepository{
		Client: client,
		log:    log,
	}
}

func (repo *ProductRepository) CreateProduct(ctx context.Context, product *domain.Product) (*ent.Product, error) {
	categoryEnt, err := repo.Client.Category.
		Query().
		Where(category.NameEQ(product.Category.Name)).
		Only(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, domain.CategoryNotFoundErr
		}
		return nil, err
	}

	qb := repo.Client.Product.
		Create().
		SetName(product.Name).
		SetDescription(product.Description).
		SetCategory(categoryEnt).
		SetPrice(product.Price)

	if product.PhotoURL != nil {
		qb = qb.SetPhotoURL(*product.PhotoURL)
	}

	productEnt, err := qb.Save(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return nil, domain.ProductAlreadyExistsErr
		}
		return nil, err
	}

	return productEnt, nil
}

func (repo *ProductRepository) GetProduct(ctx context.Context, productID uuid.UUID) (*ent.Product, error) {
	productEnt, err := repo.Client.Product.
		Query().
		WithCategory(
			func(cq *ent.CategoryQuery) {
				cq.Select("name")
			},
		).
		Where(product.IDEQ(productID)).
		Only(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, domain.ProductNotFoundErr
		}
		return nil, err
	}

	return productEnt, nil
}

func (repo *ProductRepository) DeleteProduct(ctx context.Context, productID uuid.UUID) error {
	rows, err := repo.Client.Product.
		Delete().
		Where(product.IDEQ(productID)).
		Exec(ctx)

	if rows == 0 {
		return domain.ProductNotFoundErr
	}

	return err
}

func (repo *ProductRepository) GetProducts(ctx context.Context, limit *int) ([]*ent.Product, error) {
	qb := repo.Client.Product.Query().WithCategory()

	if limit != nil {
		qb = qb.Limit(*limit)
	} else {
		qb = qb.Limit(defaultRowsLimit)
	}

	productsEnt, err := qb.All(ctx)

	if err != nil {
		return nil, err
	}

	return productsEnt, nil
}
