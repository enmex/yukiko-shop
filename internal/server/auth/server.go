package auth

import (
	spec "yukiko-shop/internal/generated/spec/auth"
	"yukiko-shop/internal/interfaces"
	"yukiko-shop/internal/repository/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger     *logrus.Logger
	authUseCase interfaces.AuthUseCase
}

func NewServer(logger *logrus.Logger, authUseCase interfaces.AuthUseCase) *Server {
	return &Server{
		logger: logger,
		authUseCase: authUseCase,
	}
}

func NewServerOptions() (spec.ChiServerOptions, error) {
	r := chi.NewRouter()

	return spec.ChiServerOptions{
		BaseRouter:  r,
		Middlewares: []spec.MiddlewareFunc{},
	}, nil
}

func NewDBClient(driver *sql.Driver) (*ent.Client, error) {
	client := ent.NewClient(ent.Driver(driver))
	return client, nil
}