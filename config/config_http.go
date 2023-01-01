package config

import "os"

type ConfigHTTP struct {
	AuthServiceHost    string
	ProductServiceHost string
	GatewayServiceHost string
	ImageServiceHost   string
	CartServiceHost string
}

func NewConfigHTTP() *ConfigHTTP {
	return &ConfigHTTP{
		AuthServiceHost:    os.Getenv("AUTH_SERVICE_HOST"),
		ProductServiceHost: os.Getenv("PRODUCT_SERVICE_HOST"),
		GatewayServiceHost: os.Getenv("GATEWAY_SERVICE_HOST"),
		ImageServiceHost:   os.Getenv("IMAGE_SERVICE_HOST"),
		CartServiceHost: os.Getenv("CART_SERVICE_HOST"),
	}
}
