package photo

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	spec "yukiko-shop/internal/generated/spec/image"
	"yukiko-shop/internal/interfaces"
)

type Server struct {
	logger       *logrus.Logger
	imageUseCase interfaces.ImageUseCase
}

func NewServer(
	logger *logrus.Logger,
	imageUseCase interfaces.ImageUseCase,
) *Server {
	return &Server{
		logger:       logger,
		imageUseCase: imageUseCase,
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
