package cart

import (
	"context"
	adapter "yukiko-shop/internal/adapter/cart"
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/cart"
	"yukiko-shop/internal/interfaces"

	"github.com/sirupsen/logrus"
)

type CartUseCase struct {
	logger         *logrus.Logger
	cartRepository interfaces.CartRepository
}

func NewCartUseCase(logger *logrus.Logger, cartRepository interfaces.CartRepository) *CartUseCase {
	return &CartUseCase{
		logger:         logger,
		cartRepository: cartRepository,
	}
}

func (u *CartUseCase) AddProductToCart(ctx context.Context, cartProduct *domain.CartProduct) error {
	return u.cartRepository.AddProductToCart(ctx, cartProduct)
}

func (u *CartUseCase) DeleteProductFromCart(ctx context.Context, cartProduct *domain.CartProduct) error {
	return u.cartRepository.DeleteProductFromCart(ctx, cartProduct)
}

func (u *CartUseCase) GetCart(ctx context.Context, user *domain.User) (*spec.GetCartResponse, error) {
	productsEnt, err := u.cartRepository.GetProductsFromCart(ctx, user)
	if err != nil {
		return nil, err
	}

	var price float64 = 0.0
	for _, product := range productsEnt {
		price += product.Price
	}

	var products []spec.Product
	for _, product := range productsEnt {
		products = append(products, *adapter.PresentCartProduct(product))
	}

	return &spec.GetCartResponse{
		Products:   products,
		TotalPrice: price,
	}, nil
}

func (u *CartUseCase) UpdateProductQuantity(ctx context.Context, cartProduct *domain.CartProduct) error {
	return u.cartRepository.UpdateProductInCart(ctx, cartProduct)
}

func (u *CartUseCase) ClearCart(ctx context.Context, user *domain.User) error {
	return u.cartRepository.DeleteAllProductsFromCart(ctx, user)
}
