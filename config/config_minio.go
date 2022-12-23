package config

import (
	"os"
	"strconv"
	"time"
	"yukiko-shop/pkg/minio"
)

func NewConfigMinio() (*minio.Config, error) {
	sslReqStr := os.Getenv("MINIO_SSL_REQUIRED")

	sslReq, err := strconv.ParseBool(sslReqStr)

	if err != nil {
		return nil, err
	}

	duration, err := time.ParseDuration(os.Getenv("MINIO_URL_DURATION"))
	if err != nil {
		return nil, err
	}

	return &minio.Config{
		Endpoint:    os.Getenv("MINIO_ENDPOINT"),
		AccessKey:   os.Getenv("MINIO_ACCESS_KEY"),
		SecretKey:   os.Getenv("MINIO_SECRET_KEY"),
		BucketName:  os.Getenv("MINIO_BUCKET_NAME"),
		UrlDuration: duration,
		SslRequired: sslReq,
	}, nil
}
