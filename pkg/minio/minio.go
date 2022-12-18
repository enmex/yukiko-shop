package minio

import (
	"context"

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

	clear(ctx, conf.BucketName, *minioClient)

	minioClient.MakeBucket(ctx, conf.BucketName, minio.MakeBucketOptions{Region: location})

	return &MinioClient{
		conf:   conf,
		client: *minioClient,
	}, nil
}

func (mc MinioClient) UploadFile(ctx context.Context, objectName, objectPath, contentType string) (*minio.UploadInfo, error) {
	info, err := mc.client.FPutObject(
		ctx,
		mc.conf.BucketName,
		objectName,
		objectPath,
		minio.PutObjectOptions{ContentType: contentType},
	)

	return &info, err
}

func (mc MinioClient) DownloadFile(ctx context.Context, objectName string) (*minio.Object, error) {
	object, err := mc.client.GetObject(ctx, mc.conf.BucketName, objectName, minio.GetObjectOptions{})

	return object, err
}

func (mc MinioClient) DeleteFile(ctx context.Context, objectName string) error {
	err := mc.client.RemoveObject(ctx, mc.conf.BucketName, objectName, minio.RemoveObjectOptions{
		ForceDelete: true,
	})
	return err
}

func clear(ctx context.Context, bucketName string, client minio.Client) error {
	err := <-client.RemoveObjects(ctx, bucketName, nil, minio.RemoveObjectsOptions{})
	if err.Err != nil {
		return err.Err
	}

	return client.RemoveBucket(ctx, bucketName)
}
