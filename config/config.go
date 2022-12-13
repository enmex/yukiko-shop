package config

import (
	"yukiko-shop/pkg/auth"
	"yukiko-shop/pkg/db"
)

type Config struct {
	DB *db.Config
	JWT *auth.Config
	HTTP *ConfigHTTP
}

func NewConfig() (*Config, error) {
	jwtConf, err := NewJWTConfig()

	if err != nil {
		return nil, err
	}

	return &Config{
		DB: NewDBConfig(),
		JWT: jwtConf,
		HTTP: NewConfigHTTP(),
	}, nil
}