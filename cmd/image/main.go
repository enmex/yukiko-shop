package main

import (
	"context"
	"strings"
	"yukiko-shop/config"
	spec "yukiko-shop/internal/generated/spec/image"
	imageServer "yukiko-shop/internal/server/image"
	"yukiko-shop/pkg/http"
	"yukiko-shop/pkg/logger"

	"github.com/sirupsen/logrus"
	"yukiko-shop/internal/usecases/image"
	"yukiko-shop/pkg/minio"
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

	minioClient, err := minio.NewClient(ctx, cfg.Minio)
	if err != nil {
		return err
	}

	imageUseCase := image.NewImageUseCase(logger, minioClient)

	// Server
	srv := imageServer.NewServer(
		logger,
		imageUseCase,
	)

	options, err := imageServer.NewServerOptions()
	if err != nil {
		return err
	}

	handler := spec.HandlerWithOptions(srv, options)
	httpServer := http.NewServer(ctx, strings.Split(cfg.HTTP.ImageServiceHost, ":")[1], handler)

	logger.Infoln("Photo server has been started")
	err = http.StartServer(httpServer)
	if err != nil {
		return err
	}

	return nil
}
