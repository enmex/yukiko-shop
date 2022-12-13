package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/getkin/kin-openapi/routers"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type AccessClaims struct {
	UserID             uuid.UUID `json:"user_id"`
	Token              string    `json:"token"`
	jwt.StandardClaims `json:"claims"`
}

type ContextKey string

const UserContextKey ContextKey = "authorized_user"

type Authenticator interface {
	GetToken(request *http.Request) (*string, error)
	IsSecure(request *http.Request) (bool, error)
}

type Authenticate struct {
	router routers.Router
}

func NewAuthenticate(openAPIRouter routers.Router) Authenticator {
	return &Authenticate{
		router: openAPIRouter,
	}
}

func (a *Authenticate) GetToken(request *http.Request) (*string, error) {
	token := request.Header.Get("Authorization")
	if len(token) > 0 {
		return &token, nil
	}

	return nil, nil
}

func (ac AccessClaims) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, UserContextKey, ac)
}

func (a *Authenticate) IsSecure(r *http.Request) (bool, error) {
	route, _, err := a.router.FindRoute(r)

	if err != nil {
		return false, err
	}

	if strings.Contains(route.Path, "auth") {
		return false, nil
	}

	return true, nil
}
