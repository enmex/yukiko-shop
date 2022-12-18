package order

import (
	"context"
	"fmt"
	adapter "yukiko-shop/internal/adapter/product"
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/product"
	"yukiko-shop/internal/interfaces"
	"yukiko-shop/pkg/minio"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var _ interfaces.ProductUseCase = (*ProductUseCase)(nil)

type ProductUseCase struct {
	logger            *logrus.Logger
	productRepository interfaces.ProductRepository
	minioClient       *minio.MinioClient
}

func NewProductUseCase(
	logger *logrus.Logger,
	productRepository interfaces.ProductRepository,
	minioClient *minio.MinioClient,
) *ProductUseCase {
	return &ProductUseCase{
		logger:            logger,
		productRepository: productRepository,
		minioClient:       minioClient,
	}
}

func (u *ProductUseCase) CreateProduct(ctx context.Context, product *domain.Product) (*spec.Product, error) {
	productEnt, err := u.productRepository.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	if _, err := u.minioClient.UploadFile(ctx, fmt.Sprintf("product-%s", productEnt.ID), product.PhotoURL, "jpg"); err != nil {
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
	if err := u.minioClient.DeleteFile(ctx, fmt.Sprintf("product-%s", productID.String())); err != nil {
		return err
	}

	return u.productRepository.DeleteProduct(ctx, productID)
}

func (u *ProductUseCase) GetProducts(ctx context.Context, limit *int) ([]spec.Product, error) {
	productsEnt, err := u.productRepository.GetProducts(ctx, limit)

	if err != nil {
		return nil, err
	}

	products := make([]spec.Product, 0, len(productsEnt))

	for _, productEnt := range productsEnt {
		products = append(products, *adapter.PresentProduct(productEnt))
	}

	return products, nil
}