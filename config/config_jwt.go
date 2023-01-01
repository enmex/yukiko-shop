package config

import (
	"os"
	"time"
	"yukiko-shop/pkg/auth"
)

func NewJWTConfig() (*auth.Config, error) {
	accessExpirationTime, err := time.ParseDuration(os.Getenv("ACCESS_EXPIRATION_TIME"))
	if err != nil {
		return nil, err
	}

	refreshExpirationTime, err := time.ParseDuration(os.Getenv("REFRESH_EXPIRATION_TIME"))
	if err!= nil {
        return nil, err
    }

	return &auth.Config{
		Secret:         []byte(os.Getenv("HASH_SALT")),
		AccessExpirationTime: accessExpirationTime,
		RefreshExpirationTime: refreshExpirationTime,
	}, nil
}
