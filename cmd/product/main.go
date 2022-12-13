package main

import (
	"context"
	"yukiko-shop/config"
	"yukiko-shop/internal/generated/spec/auth"
	productServer "yukiko-shop/internal/server/product"
	productUseCase "yukiko-shop/internal/usecases/product"
	"yukiko-shop/pkg/db"
	"yukiko-shop/pkg/http"
	"yukiko-shop/pkg/logger"

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
	dbClient, err := productServer.NewDBClient(dbDriver)
	if err != nil {
		return err
	}

	defer dbClient.Close()

	// repository
	//repo := repository.NewUserRepository(dbClient, logger)

	//jwt auth
	//jwtAuth := auth.NewJwtAuthenticate(cfg.JWT)

	//google oauth2
	//googleAuth := auth.NewGoogleAuth(cfg.Google)

	//useCase
	useCase := productUseCase.NewProductUseCase(
		logger,
	)

	// Server
	srv := productServer.NewServer(
		logger,
		useCase,
		//googleAuth,
	)

	options, err := productServer.NewServerOptions()
	if err != nil {
		return err
	}

	handler := spec.HandlerWithOptions(srv, options)
	httpServer := http.NewServer(ctx, cfg.HTTP.AuthHTTPPort, handler)

	logger.Infoln("Auth server has been started")
	err = http.StartServer(httpServer)
	if err != nil {
		return err
	}

	return nil
}