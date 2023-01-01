package repository

import (
	"context"
	"strings"
	"yukiko-shop/internal/domain"
	"yukiko-shop/internal/repository/ent"
	"yukiko-shop/internal/repository/ent/cartproduct"
	"yukiko-shop/internal/repository/ent/product"
	"yukiko-shop/internal/repository/ent/user"

	"github.com/sirupsen/logrus"
)

type CartRepository struct {
	Client *ent.Client
	log    *logrus.Logger
}

func NewCartRepository(client *ent.Client, log *logrus.Logger) *CartRepository {
	return &CartRepository{
		Client: client,
		log:    log,
	}
}

func (repo *CartRepository) AddProductToCart(ctx context.Context, productDomain *domain.CartProduct) error {
	customer, err := repo.Client.User.
		Query().
		Where(user.IDEQ(productDomain.CustomerID)).
		Only(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return domain.UserNotFoundErr
		}
		return err
	}

	productEnt, err := repo.Client.Product.
		Query().
		Where(product.IDEQ(productDomain.ProductID)).
		Only(ctx)
	if err != nil {
		return domain.ProductNotFoundErr
	}

	if _, err := repo.Client.CartProduct.
		Create().
		SetCustomer(customer).
		SetProduct(productEnt).
		SetName(productDomain.Name).
		SetPrice(productDomain.Price).
		SetPhotoURL(productDomain.PhotoUrl).
		SetQuantity(productDomain.Quantity).
		Save(ctx); err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				return domain.CartProductAlreadyExistsErr
			}
			return err
	}

	return nil
}
	
func (repo *CartRepository) DeleteProductFromCart(ctx context.Context, productDomain *domain.CartProduct) error {
	productEnt, err := repo.Client.CartProduct.
		Query().
		Where(cartproduct.And(cartproduct.ProductIDEQ(productDomain.ProductID), cartproduct.CustomerIDEQ(productDomain.CustomerID))).
		Only(ctx)
	if err != nil {
		return domain.ProductNotFoundErr
	}

	return repo.Client.CartProduct.DeleteOne(productEnt).Exec(ctx)
}
	
func (repo *CartRepository) GetProductsFromCart(ctx context.Context, user *domain.User) ([]*ent.CartProduct, error) {
	productsEnt, err := repo.Client.CartProduct.
		Query().
		Where(cartproduct.CustomerIDEQ(user.ID)).
		All(ctx)
	if err != nil {
		if strings.EqualFold(err.Error(), "not found") {
			return nil, domain.UserNotFoundErr
		}
		return nil, err
	}

	return productsEnt, nil
}
	
func (repo *CartRepository) UpdateProductInCart(ctx context.Context, product *domain.CartProduct) error {
	productEnt, err := repo.Client.CartProduct.
		Query().
		Where(cartproduct.And(cartproduct.CustomerID(product.CustomerID), cartproduct.ProductIDEQ(product.ProductID))).
		Only(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return domain.ProductNotFoundErr
		}
		return err
	}

    return repo.Client.CartProduct.
		UpdateOne(productEnt).
		SetQuantity(product.Quantity).
		Exec(ctx)
}
	
func (repo *CartRepository) DeleteAllProductsFromCart(ctx context.Context, user *domain.User) error {
	_, err := repo.Client.CartProduct.
		Delete().
		Where(cartproduct.CustomerIDEQ(user.ID)).
		Exec(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil
		}
        return err
    }

    return nil
}