package config

import "os"

type ConfigHTTP struct {
	AuthServiceHost    string
	ProductServiceHost string
	GatewayServiceHost string
}

func NewConfigHTTP() *ConfigHTTP {
	return &ConfigHTTP{
		AuthServiceHost:    os.Getenv("AUTH_SERVICE_HOST"),
		ProductServiceHost: os.Getenv("PRODUCT_SERVICE_HOST"),
		GatewayServiceHost: os.Getenv("GATEWAY_SERVICE_HOST"),
	}
}
