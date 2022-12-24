package config

import (
	"yukiko-shop/pkg/auth"
	"yukiko-shop/pkg/db"
	"yukiko-shop/pkg/mailer"
	"yukiko-shop/pkg/minio"
	"yukiko-shop/pkg/scheduler"
)

type Config struct {
	DB        *db.Config
	JWT       *auth.Config
	HTTP      *ConfigHTTP
	Mailer    *mailer.Config
	Redis     *ConfigRedis
	Minio     *minio.Config
	Scheduler *scheduler.Config
}

func NewConfig() (*Config, error) {
	jwtConf, err := NewJWTConfig()
	if err != nil {
		return nil, err
	}

	minioConf, err := NewConfigMinio()
	if err != nil {
		return nil, err
	}

	schedulerConf, err := NewConfigScheduler()
	if err != nil {
		return nil, err
	}

	return &Config{
		DB:        NewDBConfig(),
		JWT:       jwtConf,
		HTTP:      NewConfigHTTP(),
		Mailer:    NewConfigMailer(),
		Redis:     NewConfigRedis(),
		Minio:     minioConf,
		Scheduler: schedulerConf,
	}, nil
}
