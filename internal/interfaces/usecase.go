package interfaces

import (
	"context"
	spec "yukiko-shop/internal/generated/spec/auth"
)

type AuthUseCase interface {
	SendVerifyCode(ctx context.Context, request spec.SendVerifyCodeRequest) error
	SignUp(ctx context.Context, request spec.SignUpRequest) (*spec.AuthResponse, error)
	SignIn(ctx context.Context, request spec.SignInRequest) (*spec.AuthResponse, error)
}

type ProductUseCase interface {
}
