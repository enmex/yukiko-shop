package gateway

import (
	"yukiko-shop/config"
	spec "yukiko-shop/internal/generated/spec/gateway"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger *logrus.Logger
	cfg *config.ConfigHTTP
}

func NewServer(logger *logrus.Logger, cfg *config.ConfigHTTP) *Server {
	return &Server{
		logger: logger,
		cfg: cfg,
	}
}

func NewServerOptions() (spec.ChiServerOptions, error) {
	r := chi.NewRouter()
	corsHandler := cors.AllowAll()
	r.Use(corsHandler.Handler)
	return spec.ChiServerOptions{
		BaseRouter:  r,
		Middlewares: []spec.MiddlewareFunc{},
	}, nil
}
