package interfaces

import (
	"context"
	"yukiko-shop/internal/domain"
	"yukiko-shop/internal/repository/ent"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*ent.User, error)
	GetUserByEmail(ctx context.Context, email string) (*ent.User, error)
}

type ProductRepository struct {
}
