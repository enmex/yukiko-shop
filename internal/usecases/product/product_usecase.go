package product

import (
	"context"
	"fmt"
	adapter "yukiko-shop/internal/adapter/product"
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/product"
	"yukiko-shop/internal/interfaces"
	"yukiko-shop/pkg/minio"
	"yukiko-shop/pkg/scheduler"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var _ interfaces.ProductUseCase = (*ProductUseCase)(nil)

type ProductUseCase struct {
	logger            *logrus.Logger
	productRepository interfaces.ProductRepository
	minioClient       *minio.MinioClient
	errors            chan error
}

func NewProductUseCase(
	logger *logrus.Logger,
	minioClient *minio.MinioClient,
	productRepository interfaces.ProductRepository,
) *ProductUseCase {
	return &ProductUseCase{
		logger:            logger,
		productRepository: productRepository,
		minioClient:       minioClient,
		errors:            make(chan error),
	}
}

func (u *ProductUseCase) StartScheduler(ctx context.Context, cfg *scheduler.Config) {
	scheduler := scheduler.NewScheduler(cfg, func(ctx context.Context) error {
		productsEnt, err := u.productRepository.GetProductsIds(ctx)
		if err != nil {
			return err
		}

		var products []*domain.Product
		for _, categoryEnt := range productsEnt {
			url, err := u.minioClient.GetObject(ctx, fmt.Sprintf("image_%s.jpg", categoryEnt.ID))
			if err != nil {
				return err
			}

			products = append(products, &domain.Product{
				ID:       categoryEnt.ID,
				PhotoURL: url,
			})
		}

		if err := u.productRepository.UpdateProductsPhotoUrl(ctx, products); err != nil {
			return err
		}

		u.logger.Infoln("Product photo urls updated successfully")
		return nil
	})

	go scheduler.Start(ctx)
	go func() {
		err := <-scheduler.Error()
		u.errors <- err
	}()
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

func (u *ProductUseCase) ReadError() error {
	return <-u.errors
}
