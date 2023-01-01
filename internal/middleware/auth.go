package middleware

import (
	"context"
	"errors"
	"net/http"
	"time"
	"yukiko-shop/pkg/auth"
	"yukiko-shop/pkg/response"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

func AuthMiddleware(authenticator auth.Authenticator, jwtAuth auth.JwtAuthenticator) func(handler http.HandlerFunc) http.HandlerFunc {
	return func(handler http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			// Получаем список ролей, имеющих доступ к операции
			requiredAccessRoles, err := authenticator.GetRequiredAccessRoles(r)
			if err != nil {
				response.JSON(w, http.StatusInternalServerError, Error{
					Message: err.Error(),
				})
				return
			}

			// Если ролей нет, то операция доступна всем
			if requiredAccessRoles == nil {
				accessClaims, _ := getAccessClaims(r, authenticator, jwtAuth)
				if accessClaims == nil {
					handler.ServeHTTP(w, r)
				} else {
					handler.ServeHTTP(w, r.WithContext(accessClaims.WithContext(ctx)))
				}
				return
			}

			accessClaims, err := getAccessClaims(r, authenticator, jwtAuth)
			if err != nil {
				response.JSON(w, http.StatusUnauthorized, Error{
					Message: "no access",
				})
				return
			}

			//Проверяем, что у пользователя есть доступ к операции
			if !slices.Contains(requiredAccessRoles, accessClaims.AccessType) {
				response.JSON(w, http.StatusForbidden, Error{
					Message: "no access",
				})
				return
			}

			// Проверяем, что токен валиден для этого пользователя
			if time.Now().Unix() > accessClaims.ExpiresAt {
				response.JSON(w, http.StatusUnauthorized, Error{
					Message: "session expired",
				})
				return
			}

			handler.ServeHTTP(w, r.WithContext(accessClaims.WithContext(ctx)))
		}
	}
}

type Error struct {
	Message string
}

func GetUserIdFromContext(ctx context.Context) *uuid.UUID {
	claims, ok := ctx.Value(auth.UserContextKey).(auth.Claims)

	if !ok {
		return nil
	}

	return &claims.UserID
}

func getAccessClaims(r *http.Request, authenticator auth.Authenticator, jwtAuth auth.JwtAuthenticator) (*auth.Claims, error) {
	jwt, err := authenticator.GetToken(r)
	if err != nil || jwt == nil {
		return nil, errors.New("invalid security token")
	}

	var accessClaims *auth.Claims
	cleanToken := auth.GetBearer(*jwt)
	accessClaims, err = jwtAuth.ParseAccessToken(cleanToken)
	if err != nil {
		return nil, errors.New("invalid security token")
	}

	return accessClaims, nil
}
