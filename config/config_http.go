package config

import "os"

type ConfigHTTP struct {
	AuthHTTPPort string
}

func NewConfigHTTP() *ConfigHTTP {
	return &ConfigHTTP{
		AuthHTTPPort: os.Getenv("AUTH_HTTP_PORT"),
	}
}