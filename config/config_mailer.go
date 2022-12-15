package config

import (
	"os"
	"yukiko-shop/pkg/mailer"
)

func NewConfigMailer() *mailer.Config {
	return &mailer.Config{
		User:     os.Getenv("GMAIL_USER"),
		Password: os.Getenv("GMAIL_PASSWORD"),
	}
}
