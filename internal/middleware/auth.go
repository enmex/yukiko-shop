package middleware

import (
	"context"
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

			accessRoles, err := authenticator.GetRequiredAccessRoles(r)
			if err != nil {
				response.JSON(w, http.StatusInternalServerError, Error{
					Message: err.Error(),
				})
				return
			}
			if accessRoles == nil {
				handler.ServeHTTP(w, r)
				return
			}

			jwt, err := authenticator.GetToken(r)
			if err != nil || jwt == nil {
				response.JSON(w, http.StatusUnauthorized, Error{
					Message: "invalid security token",
				})
				return
			}

			// Получаем данные пользователя из токена
			var accessClaims *auth.AccessClaims
			cleanToken := auth.GetBearer(*jwt)
			accessClaims, err = jwtAuth.ParseAccessToken(cleanToken)
			if err != nil {
				response.JSON(w, http.StatusUnauthorized, Error{
					Message: "invalid security token",
				})
				return
			}

			//Проверяем, что у пользователя есть доступ к операции
			if !slices.Contains(accessRoles, accessClaims.AccessType) {
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
	claims, ok := ctx.Value(auth.UserContextKey).(auth.AccessClaims)

	if !ok {
		return nil
	}

	return &claims.UserID
}
