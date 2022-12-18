package auth

import (
	"context"
	"net/http"

	"github.com/getkin/kin-openapi/routers"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type AccessClaims struct {
	UserID             uuid.UUID `json:"user_id"`
	AccessType         string    `json:"access_type"`
	Token              string    `json:"token"`
	jwt.StandardClaims `json:"claims"`
}

type ContextKey string

const UserContextKey ContextKey = "authorized_user"

type Authenticator interface {
	GetToken(request *http.Request) (*string, error)
	GetRequiredAccessRoles(request *http.Request) ([]string, error)
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

func (a *Authenticate) GetRequiredAccessRoles(r *http.Request) ([]string, error) {
	route, _, _ := a.router.FindRoute(r)

	security := route.Operation.Security
	if security == nil {
		return nil, nil
	}

	var accessRoles []string
	for _, s := range *security {
		if roles, ok := s["bearerAuth"]; ok {
			accessRoles = append(accessRoles, roles...)
		}
	}

	return accessRoles, nil
}
