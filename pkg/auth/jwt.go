package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

var _ JwtAuthenticator = (*JwtAuthenticate)(nil)

type JwtAuthenticator interface {
	GenerateAccessToken(accessClaims AccessClaims) (string, error)
	ParseAccessToken(jwtToken string) (*AccessClaims, error)
}

type JwtAuthenticate struct {
	conf *Config
}

func NewJwtAuthenticate(conf *Config) *JwtAuthenticate {
	return &JwtAuthenticate{
		conf: conf,
	}
}

func (j *JwtAuthenticate) GenerateAccessToken(accessClaims AccessClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	jwtToken, err := token.SignedString(j.conf.Secret)
	if err != nil {
		return "", fmt.Errorf("can't create jwt token: %w", err)
	}

	return jwtToken, nil
}

func (j *JwtAuthenticate) ParseAccessToken(jwtToken string) (*AccessClaims, error) {
	ac := AccessClaims{}
	token, err := jwt.ParseWithClaims(jwtToken, &ac, func(token *jwt.Token) (interface{}, error) {
		return j.conf.Secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AccessClaims); ok && token.Valid {
		claims.Token = jwtToken
		return claims, nil
	}
	return nil, err
}
