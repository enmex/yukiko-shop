package minio

import "time"

type Config struct {
	Endpoint    string
	AccessKey   string
	SecretKey   string
	SslRequired bool
	BucketName  string
	UrlDuration time.Duration
}
