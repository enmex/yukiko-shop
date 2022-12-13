package config

import (
	"os"
	"yukiko-shop/pkg/db"
)

func NewDBConfig() *db.Config {
	return &db.Config{
		Port:               os.Getenv("DB_PORT"),
		Host:               os.Getenv("DB_HOST"),
		User:               os.Getenv("DB_USER"),
		Password:           os.Getenv("DB_PASSWORD"),
		DriverName:         os.Getenv("DB_DRIVER_NAME"),
		DatabaseName:       os.Getenv("DB_NAME"),
		Schema:             os.Getenv("DB_SCHEMA"),
		MigrationDirectory: os.Getenv("MIGRATIONS_DIRECTORY"),
		SslMode:            os.Getenv("DB_SSL_MODE"),
	}
}
