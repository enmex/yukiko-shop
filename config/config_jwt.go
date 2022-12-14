package config

import (
	"os"
	"time"
	"yukiko-shop/pkg/auth"
)

func NewJWTConfig() (*auth.Config, error) {
	expirationTime, err := time.ParseDuration(os.Getenv("EXPIRATION_TIME"))

	if err != nil {
		return nil, err
	}

	return &auth.Config{
		Secret:         []byte(os.Getenv("HASH_SALT")),
		ExpirationTime: expirationTime,
	}, nil
}
