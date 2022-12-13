package auth

import "time"

type Config struct {
	Secret         []byte        `json:"secret"`
	ExpirationTime time.Duration `json:"expiration_time"`
}
