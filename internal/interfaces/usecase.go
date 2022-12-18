package interfaces

import (
	"context"
	"yukiko-shop/internal/domain"
	specAuth "yukiko-shop/internal/generated/spec/auth"
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
	DeleteProduct(ctx context.Context, productID uuid.UUID) error
}
