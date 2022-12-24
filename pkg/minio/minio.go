package minio

import (
	"context"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	conf   *Config
	client minio.Client
}

func NewClient(ctx context.Context, conf *Config) (*MinioClient, error) {
	minioClient, err := minio.New(conf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.AccessKey, conf.SecretKey, ""),
		Secure: conf.SslRequired,
	})

	if err != nil {
		return nil, err
	}

	location := "us-east-1"

	minioClient.MakeBucket(ctx, conf.BucketName, minio.MakeBucketOptions{Region: location})

	return &MinioClient{
		conf:   conf,
		client: *minioClient,
	}, nil
}

func (mc MinioClient) UploadFile(ctx context.Context, objectName string, file multipart.File, fileHeader multipart.FileHeader, objectType string) (*string, error) {
	_, err := mc.client.PutObject(
		ctx,
		mc.conf.BucketName,
		objectName,
		file,
		fileHeader.Size,
		minio.PutObjectOptions{
			ContentType: objectType,
		},
	)
	if err != nil {
		return nil, err
	}

	url, err := mc.client.PresignedGetObject(ctx, mc.conf.BucketName, objectName, mc.conf.UrlDuration, nil)
	if err != nil {
		return nil, err
	}

	urlString := url.String()

	return &urlString, err
}

func (mc MinioClient) DeleteFile(ctx context.Context, objectName string) error {
	err := mc.client.RemoveObject(ctx, mc.conf.BucketName, objectName, minio.RemoveObjectOptions{
		ForceDelete: true,
	})
	return err
}

func (mc MinioClient) GetObject(ctx context.Context, objectName string) (*string, error) {
	url, err := mc.client.PresignedGetObject(ctx, mc.conf.BucketName, objectName, mc.conf.UrlDuration, nil)
	if err != nil {
		return nil, err
	}

	urlString := url.String()
	return &urlString, nil
}
