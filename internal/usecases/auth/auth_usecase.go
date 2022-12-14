package usecases

import (
	"context"
	"fmt"
	"strings"
	"time"
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/auth"
	"yukiko-shop/internal/interfaces"
	"yukiko-shop/internal/utils"
	"yukiko-shop/pkg/auth"
	"yukiko-shop/pkg/mailer"

	redisCache "yukiko-shop/pkg/redis-cache"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
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
		"Код от магазинчика Yukiko shop",
		fmt.Sprintf("А вот и ваше это самое: %d", code),
		u.mailCfg.User,
		request.Email,
	)
}

func (u *AuthUseCase) SignUp(ctx context.Context, user *domain.User, verifyCode int) (*spec.AuthResponse, error) {
	userEnt, err := u.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if userEnt != nil {
		return nil, domain.UserAlreadyExistsErr
	}

	code, err := u.redisCache.Get(ctx, user.Email)
	if err != nil {
		return nil, domain.VerifyCodeExpiredErr
	}

	if *code != verifyCode {
		return nil, domain.VerifyCodeNotMatchErr
	}

	if err := u.redisCache.Delete(ctx, user.Email); err != nil {
		return nil, err
	}

	user.Password = utils.CryptString(user.Password, string(u.authCfg.Secret))

	userEnt, err = u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	accessIssuedAt := time.Now()
	accessExpiresAt := time.Now().Add(u.authCfg.AccessExpirationTime)

	accessClaims := auth.Claims{
		UserID:     userEnt.ID,
		AccessType: string(userEnt.AccessType),
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  accessIssuedAt.Unix(),
			ExpiresAt: accessExpiresAt.Unix(),
		},
	}

	jwtAccessToken, err := u.jwtAuth.GenerateAccessToken(accessClaims)
	if err != nil {
		return nil, err
	}

	refreshIssuedAt := time.Now()
	refreshExpiresAt := time.Now().Add(u.authCfg.RefreshExpirationTime)

	refreshClaims := auth.Claims{
		UserID:     userEnt.ID,
		AccessType: string(userEnt.AccessType),
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  refreshIssuedAt.Unix(),
			ExpiresAt: refreshExpiresAt.Unix(),
		},
	}

	jwtRefreshToken, err := u.jwtAuth.GenerateAccessToken(refreshClaims)
	if err != nil {
		return nil, err
	}

	return &spec.AuthResponse{
		Access: spec.Token{
			Token:     jwtAccessToken,
			ExpiresAt: accessExpiresAt.UnixMilli(),
		},
		Refresh: spec.Token{
			Token:     jwtRefreshToken,
			ExpiresAt: refreshExpiresAt.UnixMilli(),
		},
	}, nil
}

func (u *AuthUseCase) SignIn(ctx context.Context, user *domain.User) (*spec.AuthResponse, error) {
	userEnt, err := u.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	if userEnt == nil {
		return nil, domain.UserNotFoundErr
	}

	if !strings.EqualFold(utils.CryptString(user.Password, string(u.authCfg.Secret)), userEnt.Password) {
		return nil, domain.WrongCredentialsErr
	}

	accessIssuedAt := time.Now()
	accessExpiresAt := time.Now().Add(u.authCfg.AccessExpirationTime)

	accessClaims := auth.Claims{
		UserID:     userEnt.ID,
		AccessType: string(userEnt.AccessType),
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  accessIssuedAt.Unix(),
			ExpiresAt: accessExpiresAt.Unix(),
		},
	}

	jwtAccessToken, err := u.jwtAuth.GenerateAccessToken(accessClaims)
	if err != nil {
		return nil, err
	}

	refreshIssuedAt := time.Now()
	refreshExpiresAt := time.Now().Add(u.authCfg.RefreshExpirationTime)

	refreshClaims := auth.Claims{
		UserID:     userEnt.ID,
		AccessType: string(userEnt.AccessType),
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  refreshIssuedAt.Unix(),
			ExpiresAt: refreshExpiresAt.Unix(),
		},
	}

	jwtRefreshToken, err := u.jwtAuth.GenerateAccessToken(refreshClaims)
	if err != nil {
		return nil, err
	}

	return &spec.AuthResponse{
		Access: spec.Token{
			Token:     jwtAccessToken,
			ExpiresAt: accessExpiresAt.UnixMilli(),
		},
		Refresh: spec.Token{
			Token:     jwtRefreshToken,
			ExpiresAt: refreshExpiresAt.UnixMilli(),
		},
	}, nil
}

func (u *AuthUseCase) GetAccessType(ctx context.Context, user *domain.User) (*spec.GetAccessTypeResponse, error) {
	accessType, err := u.userRepo.GetUserAccessType(ctx, user)
	if err != nil {
		return nil, err
	}

	return &spec.GetAccessTypeResponse{
		AccessType: spec.GetAccessTypeResponseAccessType(*accessType),
	}, nil
}

func (u *AuthUseCase) RefreshToken(ctx context.Context, user *domain.User) (*spec.AuthResponse, error) {
	userEnt, err := u.userRepo.GetUserByID(ctx, user)
	if err!= nil {
        return nil, err
    }

	accessIssuedAt := time.Now()
	accessExpiresAt := time.Now().Add(u.authCfg.AccessExpirationTime)

	accessClaims := auth.Claims{
		UserID:     userEnt.ID,
		AccessType: string(userEnt.AccessType),
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  accessIssuedAt.Unix(),
			ExpiresAt: accessExpiresAt.Unix(),
		},
	}

	jwtAccessToken, err := u.jwtAuth.GenerateAccessToken(accessClaims)
	if err != nil {
		return nil, err
	}

	refreshIssuedAt := time.Now()
	refreshExpiresAt := time.Now().Add(u.authCfg.RefreshExpirationTime)

	refreshClaims := auth.Claims{
		UserID:     userEnt.ID,
		AccessType: string(userEnt.AccessType),
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  refreshIssuedAt.Unix(),
			ExpiresAt: refreshExpiresAt.Unix(),
		},
	}

	jwtRefreshToken, err := u.jwtAuth.GenerateAccessToken(refreshClaims)
	if err != nil {
		return nil, err
	}

	return &spec.AuthResponse{
		Access: spec.Token{
			Token:     jwtAccessToken,
			ExpiresAt: accessExpiresAt.UnixMilli(),
		},
		Refresh: spec.Token{
			Token:     jwtRefreshToken,
			ExpiresAt: refreshExpiresAt.UnixMilli(),
		},
	}, nil
}