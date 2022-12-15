package usecases

import (
	"context"
	"fmt"
	"strings"
	"time"
	adapter "yukiko-shop/internal/adapter/user"
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/auth"
	"yukiko-shop/internal/interfaces"
	"yukiko-shop/internal/utils"
	"yukiko-shop/pkg/auth"
	"yukiko-shop/pkg/mailer"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	redisCache "yukiko-shop/pkg/redis-cache"
)

var _ interfaces.AuthUseCase = (*AuthUseCase)(nil)

type AuthUseCase struct {
	logger     *logrus.Logger
	authCfg    *auth.Config
	mailCfg    *mailer.Config
	userRepo   interfaces.UserRepository
	mailer     *mailer.Mailer
	redisCache *redisCache.RedisCache[int]
	jwtAuth    *auth.JwtAuthenticate
}

func NewAuthUseCase(
	logger *logrus.Logger,
	authCfg *auth.Config,
	mailCfg *mailer.Config,
	userRepo interfaces.UserRepository,
	jwtAuth *auth.JwtAuthenticate,
	mailer *mailer.Mailer,
	redisCache *redisCache.RedisCache[int]) *AuthUseCase {
	return &AuthUseCase{
		logger:     logger,
		authCfg:    authCfg,
		mailCfg:    mailCfg,
		userRepo:   userRepo,
		mailer:     mailer,
		redisCache: redisCache,
		jwtAuth:    jwtAuth,
	}
}

func (u *AuthUseCase) SendVerifyCode(ctx context.Context, request spec.SendVerifyCodeRequest) error {
	code := utils.GenerateInLine(100000, 999999)

	if err := u.redisCache.Set(ctx, request.Email, code); err != nil {
		return err
	}

	return u.mailer.SendMail(
		"Code verification Yukiko shop",
		fmt.Sprintf("Your verify code is %d", code),
		u.mailCfg.User,
		request.Email,
	)
}

func (u *AuthUseCase) SignUp(ctx context.Context, request spec.SignUpRequest) (*spec.AuthResponse, error) {
	userEnt, err := u.userRepo.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}

	if userEnt != nil {
		return nil, domain.UserAlreadyExistsErr
	}

	code, err := u.redisCache.Get(ctx, request.Email)
	if err != nil {
		return nil, domain.VerifyCodeExpiredErr
	}

	if *code != request.Code {
		return nil, domain.VerifyCodeNotMatchErr
	}

	password := utils.CryptString(request.Password, string(u.authCfg.Secret))

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
	expiresAt := time.Now().Add(u.authCfg.ExpirationTime)

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

	if !strings.EqualFold(utils.CryptString(request.Password, string(u.authCfg.Secret)), userEnt.Password) {
		return nil, domain.WrongCredentialsErr
	}

	issuedAt := time.Now()
	expiresAt := time.Now().Add(u.authCfg.ExpirationTime)

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
