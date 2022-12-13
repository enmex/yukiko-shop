package http

import (
	"net"
)

const DefaultPort = "5050"

type Config struct {
	Port string `json:"port"`
}

// Addr returns server address in format ":<port>"
func (c *Config) Addr() string {
	return net.JoinHostPort("", c.Port)
}
