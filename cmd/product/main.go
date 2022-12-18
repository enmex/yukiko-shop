package main

import (
	"context"
	"strings"
	"yukiko-shop/config"
	"yukiko-shop/internal/generated/spec/product"
	"yukiko-shop/internal/repository"
	productServer "yukiko-shop/internal/server/product"
	usecase "yukiko-shop/internal/usecases/product"
	"yukiko-shop/pkg/db"
	"yukiko-shop/pkg/http"
	"yukiko-shop/pkg/logger"
	"yukiko-shop/pkg/minio"

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

	// repository of product
	productRepo := repository.NewProductRepository(dbClient, logger)

	//repository of category
	categoryRepo := repository.NewCategoryRepository(dbClient, logger)

	//minio
	minioClient, err := minio.NewClient(ctx, cfg.Minio)
	if err != nil {
		return err
	}

	//productUseCase
	productUseCase := usecase.NewProductUseCase(
		logger,
		productRepo,
		minioClient,
	)

	//categoryUseCase
	categoryUseCase := usecase.NewCategoryUseCase(
		logger,
		categoryRepo,
	)

	// Server
	srv := productServer.NewServer(
		logger,
		productUseCase,
		categoryUseCase,
	)

	options, err := productServer.NewServerOptions()
	if err != nil {
		return err
	}

	handler := spec.HandlerWithOptions(srv, options)
	httpServer := http.NewServer(ctx, strings.Split(cfg.HTTP.ProductServiceHost, ":")[1], handler)

	logger.Infoln("Product server has been started")
	err = http.StartServer(httpServer)
	if err != nil {
		return err
	}

	return nil
}
