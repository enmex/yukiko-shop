package main

import (
	"context"
	"strings"
	"yukiko-shop/config"
	"yukiko-shop/internal/generated/spec/gateway"
	gatewayServer "yukiko-shop/internal/server/gateway"
	"yukiko-shop/pkg/auth"
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

	// Server
	srv := gatewayServer.NewServer(
		logger,
		cfg.HTTP,
	)

	//jwt auth
	jwtAuth := auth.NewJwtAuthenticate(cfg.JWT)

	api, err := gatewayServer.NewOpenAPI()

	if err != nil {
		return err
	}

	authenticator := auth.NewAuthenticate(api)

	options, err := gatewayServer.NewServerOptions(authenticator, jwtAuth)
	if err != nil {
		return err
	}

	handler := spec.HandlerWithOptions(srv, options)
	httpServer := http.NewServer(ctx, strings.Split(cfg.HTTP.GatewayServiceHost, ":")[1], handler)

	logger.Infoln("API gateway has been opened for client requests")
	err = http.StartServer(httpServer)
	if err != nil {
		return err
	}

	return nil
}
