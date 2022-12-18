package order

import (
	"context"
	adapter "yukiko-shop/internal/adapter/product"
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/product"
	"yukiko-shop/internal/interfaces"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var _ interfaces.ProductUseCase = (*ProductUseCase)(nil)

type ProductUseCase struct {
	logger            *logrus.Logger
	productRepository interfaces.ProductRepository
}

func NewProductUseCase(logger *logrus.Logger, productRepository interfaces.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		logger:            logger,
		productRepository: productRepository,
	}
}

func (u *ProductUseCase) CreateProduct(ctx context.Context, product *domain.Product) (*spec.Product, error) {
	productEnt, err := u.productRepository.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return adapter.PresentProduct(productEnt), nil
}

func (u *ProductUseCase) GetProduct(ctx context.Context, productID uuid.UUID) (*spec.Product, error) {
	productEnt, err := u.productRepository.GetProduct(ctx, productID)
	if err != nil {
		return nil, err
	}

	return adapter.PresentProduct(productEnt), nil
}

func (u *ProductUseCase) DeleteProduct(ctx context.Context, productID uuid.UUID) error {
	return u.productRepository.DeleteProduct(ctx, productID)
}
