package repository

import (
	"context"
	"strings"
	"yukiko-shop/internal/domain"
	"yukiko-shop/internal/interfaces"
	"yukiko-shop/internal/repository/ent"
	"yukiko-shop/internal/repository/ent/user"

	"github.com/sirupsen/logrus"
)

var _ interfaces.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	Client *ent.Client
	log    *logrus.Logger
}

func NewUserRepository(client *ent.Client, log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Client: client,
		log:    log,
	}
}

func (repo *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*ent.User, error) {
	userEnt, err := repo.Client.User.
		Create().
		SetEmail(user.Email).
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetPassword(user.Password).
		SetAccessType("ADMIN").
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return userEnt, nil
}

func (repo *UserRepository) GetUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	userEnt, err := repo.Client.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, nil
		}
		return nil, err
	}

	return userEnt, nil
}
