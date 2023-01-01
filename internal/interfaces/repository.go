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
	GetUserByID(ctx context.Context, user *domain.User) (*ent.User, error)
	GetUserAccessType(ctx context.Context, user *domain.User) (*string, error)
}

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *domain.Product) (*ent.Product, error)
	GetProduct(ctx context.Context, productID uuid.UUID) (*ent.Product, error)
	GetProducts(ctx context.Context, limit *int) ([]*ent.Product, error)
	DeleteProduct(ctx context.Context, productID uuid.UUID) error
	UpdateProductPhotoUrl(ctx context.Context, productID uuid.UUID, photoUrl string) (*ent.Product, error)
	UpdateProductsPhotoUrl(ctx context.Context, products []*domain.Product) error
	GetProductsIds(ctx context.Context) ([]*ent.Product, error)
}

type CategoryRepository interface {
	CreateCategory(ctx context.Context, category *domain.Category) (*ent.Category, error)
	GetCategories(ctx context.Context, categoryType *string) ([]*ent.Category, error)
	GetCategoryByID(ctx context.Context, categoryID uuid.UUID) (*ent.Category, error)
	GetCategoryChildren(ctx context.Context, category *domain.Category) ([]*ent.Category, error)
	UpdateCategoryPhotoUrl(ctx context.Context, categoryID uuid.UUID, photoUrl string) (*ent.Category, error)
	UpdateCategoriesPhotoUrl(ctx context.Context, categories []*domain.Category) error
	GetCategoriesIds(ctx context.Context) ([]*ent.Category, error)
}

type CartRepository interface {
	AddProductToCart(ctx context.Context, product *domain.CartProduct) error
	DeleteProductFromCart(ctx context.Context, product *domain.CartProduct) error
	GetProductsFromCart(ctx context.Context, user *domain.User) ([]*ent.CartProduct, error)
	UpdateProductInCart(ctx context.Context, product *domain.CartProduct) error
	DeleteAllProductsFromCart(ctx context.Context, user *domain.User) error
}
