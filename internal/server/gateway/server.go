package gateway

import (
	"yukiko-shop/config"
	spec "yukiko-shop/internal/generated/spec/gateway"
	"yukiko-shop/internal/middleware"
	"yukiko-shop/pkg/auth"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/legacy"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger *logrus.Logger
	cfg    *config.ConfigHTTP
}

type OpenAPI struct {
	*openapi3.T
	routers.Router
}

func NewServer(logger *logrus.Logger, cfg *config.ConfigHTTP) *Server {
	return &Server{
		logger: logger,
		cfg:    cfg,
	}
}

func NewServerOptions(
	authenticator auth.Authenticator,
	jwtAuthenticator auth.JwtAuthenticator,
) (spec.ChiServerOptions, error) {
	r := chi.NewRouter()
	corsHandler := cors.AllowAll()
	r.Use(corsHandler.Handler)
	return spec.ChiServerOptions{
		BaseRouter: r,
		Middlewares: []spec.MiddlewareFunc{
			middleware.AuthMiddleware(authenticator, jwtAuthenticator),
		},
	}, nil
}

func NewOpenAPI() (*OpenAPI, error) {
	var api OpenAPI
	var err error

	api.T, err = spec.GetSwagger()
	if err != nil {
		return nil, err
	}
	api.Servers = openapi3.Servers{&openapi3.Server{URL: "/api/v1"}, &openapi3.Server{URL: "/"}}
	api.Router, err = legacy.NewRouter(api.T)
	if err != nil {
		return nil, err
	}
	return &api, nil
}
