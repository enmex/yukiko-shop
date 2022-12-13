package interfaces

import (
	"context"
	spec "yukiko-shop/internal/generated/spec/auth"
)

type AuthUseCase interface {
	SignUp(ctx context.Context, request spec.SignUpRequest) (*spec.AuthResponse, error)
	SignIn(ctx context.Context, request spec.SignInRequest) (*spec.AuthResponse, error)
}