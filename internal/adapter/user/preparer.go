package adapter

import (
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/auth"
)

func PrepareUserSignIn(userRequest *spec.SignInRequest) *domain.User {
	return &domain.User{
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}
}

func PrepareUserSignUp(userRequest *spec.SignUpRequest) (*domain.User, int) {
	return &domain.User{
		Email:     userRequest.Email,
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
		Password:  userRequest.Password,
	}, userRequest.Code
}
