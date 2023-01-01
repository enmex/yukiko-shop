package image

import (
	"context"
	"fmt"
	"mime/multipart"
	"strings"
	spec "yukiko-shop/internal/generated/spec/image"
	"yukiko-shop/internal/interfaces"
	"yukiko-shop/pkg/minio"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

var _ interfaces.ImageUseCase = (*ImageUseCase)(nil)

type ImageUseCase struct {
	logger      *logrus.Logger
	minioClient *minio.MinioClient
}

func NewImageUseCase(
	logger *logrus.Logger,
	minioClient *minio.MinioClient) *ImageUseCase {
	return &ImageUseCase{
		logger:      logger,
		minioClient: minioClient,
	}
}

func (u *ImageUseCase) UploadImage(ctx context.Context, file multipart.File, fileHeader multipart.FileHeader) (*spec.UploadImageResponse, error) {
	splits := strings.Split(fileHeader.Filename, ".")
	fileType := splits[len(splits)-1]

	id := uuid.New()

	url, err := u.minioClient.UploadFile(ctx, fmt.Sprintf("image_%s.%s", id, fileType), file, fileHeader, fileType)

	if err != nil {
		return nil, err
	}

	return &spec.UploadImageResponse{
		Id:       id.String(),
		PhotoUrl: *url,
	}, nil
}

func (u *ImageUseCase) DeleteImage(ctx context.Context, imageID uuid.UUID) error {
	return u.minioClient.DeleteFile(ctx, fmt.Sprintf("image_%s.jpg", imageID.String()))
}
