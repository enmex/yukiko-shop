package config

import (
	"yukiko-shop/pkg/auth"
	"yukiko-shop/pkg/db"
	"yukiko-shop/pkg/mailer"
)

type Config struct {
	DB     *db.Config
	JWT    *auth.Config
	HTTP   *ConfigHTTP
	Mailer *mailer.Config
	Redis  *ConfigRedis
}

func NewConfig() (*Config, error) {
	jwtConf, err := NewJWTConfig()

	if err != nil {
		return nil, err
	}

	return &Config{
		DB:     NewDBConfig(),
		JWT:    jwtConf,
		HTTP:   NewConfigHTTP(),
		Mailer: NewConfigMailer(),
		Redis:  NewConfigRedis(),
	}, nil
}
