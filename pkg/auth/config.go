package auth

import "time"

type Config struct {
	Secret         []byte        `json:"secret"`
	AccessExpirationTime time.Duration `json:"access_expiration_time"`
	RefreshExpirationTime time.Duration `json:"refresh_expiration_time"`
}
