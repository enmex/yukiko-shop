package domain

import "errors"

var (
	UserAlreadyExistsErr     = errors.New("user already exists")
	UserNotFoundErr          = errors.New("user not found")
	WrongCredentialsErr      = errors.New("wrong credentials")
	VerifyCodeExpiredErr     = errors.New("verify code expired")
	VerifyCodeNotMatchErr    = errors.New("wrong verify code")
	CategoryNotFoundErr      = errors.New("category not found")
	ProductAlreadyExistsErr  = errors.New("product with this name already exists")
	ProductNotFoundErr       = errors.New("product not found")
	CategoryAlreadyExistsErr = errors.New("category with this name already exists")
	CartProductAlreadyExistsErr = errors.New("this product is already in the cart")
)
