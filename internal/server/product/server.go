package auth

import (
	spec "yukiko-shop/internal/generated/spec/auth"
	"yukiko-shop/internal/interfaces"
	"yukiko-shop/internal/repository/ent"

	"entgo.io/ent/dialect/sql"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger       *logrus.Logger
	orderUseCase interfaces.ProductUseCase
}

func NewServer(logger *logrus.Logger, orderUseCase interfaces.ProductUseCase) *Server {
	return &Server{
		logger:       logger,
		orderUseCase: orderUseCase,
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

func NewDBClient(driver *sql.Driver) (*ent.Client, error) {
	client := ent.NewClient(ent.Driver(driver))
	return client, nil
}
