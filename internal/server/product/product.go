package product

import (
	"encoding/json"
	"net/http"
	"strings"
	adapter "yukiko-shop/internal/adapter/product"
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/product"
	"yukiko-shop/pkg/response"

	"github.com/google/uuid"
)

func (s Server) PostProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request spec.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	product := adapter.PrepareProduct(&request)

	productResponse, err := s.productUseCase.CreateProduct(ctx, product)
	if err != nil {
		var statusCode int
		if strings.EqualFold(err.Error(), domain.ProductAlreadyExistsErr.Error()) || strings.EqualFold(err.Error(), domain.CategoryNotFoundErr.Error()) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}
		response.JSON(w, statusCode, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, spec.CreateProductResponse{
		Id:           productResponse.Id,
		Name:         productResponse.Name,
		Description:  productResponse.Description,
		PhotoUrl:     productResponse.PhotoUrl,
		Price:        productResponse.Price,
		CategoryName: productResponse.CategoryName,
	})
}

func (s Server) GetProductsProductID(w http.ResponseWriter, r *http.Request, productID spec.ProductID) {
	ctx := r.Context()

	productUuid, err := uuid.Parse(string(productID))
	if err != nil {
		response.JSON(w, http.StatusBadRequest, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	productResponse, err := s.productUseCase.GetProduct(ctx, productUuid)
	if err != nil {
		var statusCode int
		if strings.EqualFold(err.Error(), domain.ProductNotFoundErr.Error()) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}

		response.JSON(w, statusCode, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, spec.GetProductResponse{
		Id:           productResponse.Id,
		Name:         productResponse.Name,
		Description:  productResponse.Description,
		PhotoUrl:     productResponse.PhotoUrl,
		Price:        productResponse.Price,
		CategoryName: productResponse.CategoryName,
	})
}

func (s Server) DeleteProductsProductID(w http.ResponseWriter, r *http.Request, productID spec.ProductID) {
	ctx := r.Context()

	productUuid, err := uuid.Parse(string(productID))
	if err != nil {
		response.JSON(w, http.StatusBadRequest, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	if err := s.productUseCase.DeleteProduct(ctx, productUuid); err != nil {
		var statusCode int
		if strings.EqualFold(err.Error(), domain.ProductNotFoundErr.Error()) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}

		response.JSON(w, statusCode, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.EmptyJSON(w, http.StatusNoContent)
}