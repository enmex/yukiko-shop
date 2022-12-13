package order

import (
	"yukiko-shop/internal/interfaces"

	"github.com/sirupsen/logrus"
)

var _ interfaces.ProductUseCase = (*ProductUseCase)(nil)

type ProductUseCase struct {
	logger *logrus.Logger
}

func NewProductUseCase(logger *logrus.Logger) *ProductUseCase {
	return &ProductUseCase{
		logger: logger,
	}
}
