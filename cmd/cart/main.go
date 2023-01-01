package main

import (
	"context"
	"strings"
	"yukiko-shop/config"
	"yukiko-shop/internal/generated/spec/cart"
	"yukiko-shop/internal/repository"
	cartServer "yukiko-shop/internal/server/cart"
	cartUseCase "yukiko-shop/internal/usecases/cart"
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
	dbClient, err := cartServer.NewDBClient(dbDriver)
	if err != nil {
		return err
	}

	defer dbClient.Close()

	// repository
	repo := repository.NewCartRepository(dbClient, logger)

	//useCase
	useCase := cartUseCase.NewCartUseCase(
		logger,
		repo,
	)

	// Server
	srv := cartServer.NewServer(
		logger,
		useCase,
	)

	options, err := cartServer.NewServerOptions()
	if err != nil {
		return err
	}

	handler := spec.HandlerWithOptions(srv, options)
	httpServer := http.NewServer(ctx, strings.Split(cfg.HTTP.CartServiceHost, ":")[1], handler)

	logger.Infoln("Cart server has been started")
	err = http.StartServer(httpServer)
	if err != nil {
		return err
	}

	return nil
}
