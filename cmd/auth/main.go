package main

import (
	"context"
	"fmt"
	"strings"
	"time"
	"yukiko-shop/config"
	"yukiko-shop/internal/generated/spec/auth"
	"yukiko-shop/internal/repository"
	authServer "yukiko-shop/internal/server/auth"
	authUseCase "yukiko-shop/internal/usecases/auth"
	"yukiko-shop/pkg/auth"
	"yukiko-shop/pkg/db"
	"yukiko-shop/pkg/http"
	"yukiko-shop/pkg/logger"
	"yukiko-shop/pkg/mailer"

	redisCache "yukiko-shop/pkg/redis-cache"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func main() {
	logger.Init()

	logger := &logger.Logger

	if err := run(logger); err != nil {
		logger.Fatalln(err.Error())
	}
}

func run(logger *logrus.Logger) error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	ctx := context.Background()

	// Database
	if err := db.Migrations(cfg.DB, logger); err != nil {
		return err
	}
	dbDriver, err := db.GetDriver(cfg.DB, logger)
	if err != nil {
		return err
	}
	dbClient, err := authServer.NewDBClient(dbDriver)
	if err != nil {
		return err
	}

	defer dbClient.Close()

	// repository
	repo := repository.NewUserRepository(dbClient, logger)

	//jwt auth
	jwtAuth := auth.NewJwtAuthenticate(cfg.JWT)

	//google oauth2
	//googleAuth := auth.NewGoogleAuth(cfg.Google)

	//mailer
	mailer := mailer.NewMailer(cfg.Mailer)

	//redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	redisCache := redisCache.NewRedisCache[int](redisClient, time.Hour*24)

	//useCase
	useCase := authUseCase.NewAuthUseCase(
		logger,
		cfg.JWT,
		cfg.Mailer,
		repo,
		jwtAuth,
		mailer,
		redisCache,
	)

	// Server
	srv := authServer.NewServer(
		logger,
		useCase,
		//googleAuth,
	)

	options, err := authServer.NewServerOptions()
	if err != nil {
		return err
	}

	handler := spec.HandlerWithOptions(srv, options)
	httpServer := http.NewServer(ctx, strings.Split(cfg.HTTP.AuthServiceHost, ":")[1], handler)

	logger.Infoln("Auth server has been started")
	err = http.StartServer(httpServer)
	if err != nil {
		return err
	}

	return nil
}
