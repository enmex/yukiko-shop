package interfaces

import (
	"context"
	"mime/multipart"
	"yukiko-shop/internal/domain"
	specAuth "yukiko-shop/internal/generated/spec/auth"
	specImage "yukiko-shop/internal/generated/spec/image"
	specProduct "yukiko-shop/internal/generated/spec/product"

	"github.com/google/uuid"
)

type AuthUseCase interface {
	SendVerifyCode(ctx context.Context, request specAuth.SendVerifyCodeRequest) error
	SignUp(ctx context.Context, user *domain.User, verifyCode int) (*specAuth.AuthResponse, error)
	SignIn(ctx context.Context, user *domain.User) (*specAuth.AuthResponse, error)
}

type ProductUseCase interface {
	CreateProduct(ctx context.Context, product *domain.Product) (*specProduct.Product, error)
	GetProduct(ctx context.Context, productID uuid.UUID) (*specProduct.Product, error)
	GetProducts(ctx context.Context, limit *int) ([]specProduct.Product, error)
	DeleteProduct(ctx context.Context, productID uuid.UUID) error
}

type CategoryUseCase interface {
	CreateCategory(ctx context.Context, category *domain.Category) (*specProduct.Category, error)
	GetCategories(ctx context.Context, main *bool, leaf *bool) ([]*specProduct.Category, error)
	GetCategoryByName(ctx context.Context, categoryName string) (*specProduct.Category, error)
	GetSubCategories(ctx context.Context, categoryName string) ([]*specProduct.Category, error)
}

type ImageUseCase interface {
	UploadImage(ctx context.Context, file multipart.File, fileHeader multipart.FileHeader) (*specImage.UploadImageResponse, error)
}
