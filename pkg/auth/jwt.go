package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

var _ JwtAuthenticator = (*JwtAuthenticate)(nil)

type JwtAuthenticator interface {
	GenerateAccessToken(accessClaims Claims) (string, error)
	ParseAccessToken(jwtToken string) (*Claims, error)
}

type JwtAuthenticate struct {
	conf *Config
}

func NewJwtAuthenticate(conf *Config) *JwtAuthenticate {
	return &JwtAuthenticate{
		conf: conf,
	}
}

func (j *JwtAuthenticate) GenerateAccessToken(accessClaims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	jwtToken, err := token.SignedString(j.conf.Secret)
	if err != nil {
		return "", fmt.Errorf("can't create jwt token: %w", err)
	}

	return jwtToken, nil
}

func (j *JwtAuthenticate) ParseAccessToken(jwtToken string) (*Claims, error) {
	ac := Claims{}
	token, err := jwt.ParseWithClaims(jwtToken, &ac, func(token *jwt.Token) (interface{}, error) {
		return j.conf.Secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		claims.Token = jwtToken
		return claims, nil
	}
	return nil, err
}
