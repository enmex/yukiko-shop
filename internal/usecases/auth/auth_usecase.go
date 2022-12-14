package usecases

import (
	"context"
	"strings"
	"time"
	adapter "yukiko-shop/internal/adapter/user"
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/auth"
	"yukiko-shop/internal/interfaces"
	"yukiko-shop/internal/utils"
	"yukiko-shop/pkg/auth"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

var _ interfaces.AuthUseCase = (*AuthUseCase)(nil)

type AuthUseCase struct {
	logger   *logrus.Logger
	cfg      *auth.Config
	userRepo interfaces.UserRepository
	jwtAuth  *auth.JwtAuthenticate
}

func NewAuthUseCase(logger *logrus.Logger, cfg *auth.Config, userRepo interfaces.UserRepository, jwtAuth *auth.JwtAuthenticate) *AuthUseCase {
	return &AuthUseCase{
		logger:   logger,
		cfg:      cfg,
		userRepo: userRepo,
		jwtAuth:  jwtAuth,
	}
}

func (u *AuthUseCase) SignUp(ctx context.Context, request spec.SignUpRequest) (*spec.AuthResponse, error) {
	userEnt, err := u.userRepo.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	if userEnt != nil {
		return nil, domain.UserAlreadyExistsErr
	}

	password := utils.CryptString(request.Password, string(u.cfg.Secret))

	userDomain := domain.User{
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Password:  password,
	}

	userEnt, err = u.userRepo.CreateUser(ctx, &userDomain)
	if err != nil {
		return nil, err
	}

	issuedAt := time.Now()
	expiresAt := time.Now().Add(u.cfg.ExpirationTime)

	claims := auth.AccessClaims{
		UserID: userEnt.ID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  issuedAt.Unix(),
			ExpiresAt: expiresAt.Unix(),
		},
	}

	jwtToken, err := u.jwtAuth.GenerateAccessToken(claims)

	if err != nil {
		return nil, err
	}

	return &spec.AuthResponse{
		Auth: spec.UserToken{
			Access: spec.Token{
				Token:     jwtToken,
				ExpiresAt: expiresAt.UnixNano(),
			},
			Refresh: spec.Token{
				Token:     jwtToken,
				ExpiresAt: expiresAt.UnixNano(),
			},
		},
		Profile: *adapter.ConvertEntToDomain(userEnt),
	}, nil
}

func (u *AuthUseCase) SignIn(ctx context.Context, request spec.SignInRequest) (*spec.AuthResponse, error) {
	userEnt, err := u.userRepo.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}
	if userEnt == nil {
		return nil, domain.UserNotFoundErr
	}

	if !strings.EqualFold(utils.CryptString(request.Password, string(u.cfg.Secret)), userEnt.Password) {
		return nil, domain.WrongCredentialsErr
	}

	issuedAt := time.Now()
	expiresAt := time.Now().Add(u.cfg.ExpirationTime)

	claims := auth.AccessClaims{
		UserID: userEnt.ID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  issuedAt.Unix(),
			ExpiresAt: expiresAt.Unix(),
		},
	}

	jwtToken, err := u.jwtAuth.GenerateAccessToken(claims)

	if err != nil {
		return nil, err
	}

	return &spec.AuthResponse{
		Auth: spec.UserToken{
			Access: spec.Token{
				Token:     jwtToken,
				ExpiresAt: expiresAt.UnixMilli(),
			},
			Refresh: spec.Token{
				Token:     jwtToken,
				ExpiresAt: expiresAt.UnixMilli(),
			},
		},
		Profile: *adapter.ConvertEntToDomain(userEnt),
	}, nil
}
