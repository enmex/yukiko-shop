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
	GetAccessType(ctx context.Context, user *domain.User) (*specAuth.GetAccessTypeResponse, error)
}

type ProductUseCase interface {
	CreateProduct(ctx context.Context, product *domain.Product) error
	GetProduct(ctx context.Context, productID uuid.UUID) (*specProduct.Product, error)
	GetProducts(ctx context.Context, limit *int) ([]specProduct.Product, error)
	DeleteProduct(ctx context.Context, productID uuid.UUID) error
}

type CategoryUseCase interface {
	CreateCategory(ctx context.Context, category *domain.Category) error
	GetCategories(ctx context.Context, categoryType *specProduct.GetCategoriesParamsType) ([]*specProduct.Category, error)
	GetCategoryByID(ctx context.Context, categoryID uuid.UUID) (*specProduct.Category, error)
	GetSubCategories(ctx context.Context, categoryID uuid.UUID) ([]*specProduct.Category, error)
}

type ImageUseCase interface {
	UploadImage(ctx context.Context, file multipart.File, fileHeader multipart.FileHeader) (*specImage.UploadImageResponse, error)
	DeleteImage(ctx context.Context, imageID uuid.UUID) error
}
