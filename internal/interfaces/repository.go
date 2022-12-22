package interfaces

import (
	"context"
	"yukiko-shop/internal/domain"
	"yukiko-shop/internal/repository/ent"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*ent.User, error)
	GetUserByEmail(ctx context.Context, email string) (*ent.User, error)
}

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *domain.Product) (*ent.Product, error)
	GetProduct(ctx context.Context, productID uuid.UUID) (*ent.Product, error)
	GetProducts(ctx context.Context, limit *int) ([]*ent.Product, error)
	DeleteProduct(ctx context.Context, productID uuid.UUID) error
}

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *domain.Category) (*ent.Category, error)
	GetCategories(ctx context.Context, main *bool, leaf *bool) ([]*ent.Category, error)
	GetCategoryByName(ctx context.Context, name string) (*ent.Category, error)
	GetCategoryChildren(ctx context.Context, category *domain.Category) ([]*ent.Category, error)
}
