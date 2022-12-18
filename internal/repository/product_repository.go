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

	productEnt, err := repo.Client.Product.
		Create().
		SetName(product.Name).
		SetDescription(product.Description).
		SetCategory(categoryEnt).
		Save(ctx)

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
