package product

import (
	"encoding/json"
	"net/http"
	"strings"
	categoryAdapter "yukiko-shop/internal/adapter/category"
	productAdapter "yukiko-shop/internal/adapter/product"
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/product"
	"yukiko-shop/pkg/response"

	"github.com/google/uuid"
)

func (s Server) PostProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request spec.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		s.logger.Warnf("POST /products Error: %s", err.Error())
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	product := productAdapter.PrepareProduct(&request)

	productResponse, err := s.productUseCase.CreateProduct(ctx, product)
	if err != nil {
		var statusCode int
		if strings.EqualFold(err.Error(), domain.ProductAlreadyExistsErr.Error()) || strings.EqualFold(err.Error(), domain.CategoryNotFoundErr.Error()) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}
		s.logger.Warnf("POST /products Error: %s", err.Error())
		response.JSON(w, statusCode, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, spec.CreateProductResponse{
		Id:           productResponse.Id,
		Name:         productResponse.Name,
		Description:  productResponse.Description,
		PhotoUrl:     *productResponse.PhotoUrl,
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
		s.logger.Infof("GET /products/%s Error: %s", productResponse.Id, err.Error())
		response.JSON(w, statusCode, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, spec.GetProductResponse{
		Id:           productResponse.Id,
		Name:         productResponse.Name,
		Description:  productResponse.Description,
		PhotoUrl:     *productResponse.PhotoUrl,
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
		s.logger.Infof("DELETE /products/%s Error: %s", string(productID), err.Error())
		response.JSON(w, statusCode, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.EmptyJSON(w, http.StatusNoContent)
}

func (s Server) GetProducts(w http.ResponseWriter, r *http.Request, params spec.GetProductsParams) {
	ctx := r.Context()

	products, err := s.productUseCase.GetProducts(ctx, params.Limit)
	if err != nil {
		s.logger.Infof("GET /products Error: %s", err.Error())
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, spec.GetProductsResponse{
		Products: products,
	})
}

func (s Server) PostCategories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request spec.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	_, err := s.categoryUseCase.CreateCategory(ctx, categoryAdapter.PrepareCategory(&request))
	if err != nil {
		var statusCode int
		if strings.EqualFold(err.Error(), domain.CategoryAlreadyExistsErr.Error()) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}

		s.logger.Warnf("POST /categories Error: %s", err.Error())
		response.JSON(w, statusCode, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.EmptyJSON(w, http.StatusOK)
}

func (s Server) GetCategories(w http.ResponseWriter, r *http.Request, params spec.GetCategoriesParams) {
	ctx := r.Context()

	categories, err := s.categoryUseCase.GetCategories(ctx, params.Main, params.Leaf)
	if err != nil {
		s.logger.Warnf("GET /categories?main=%s Error: %s", *params.Main, err.Error())
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, categories)
}

func (s Server) GetCategoriesChildrenCategoryName(w http.ResponseWriter, r *http.Request, categoryName spec.CategoryName) {
	ctx := r.Context()

	categories, err := s.categoryUseCase.GetSubCategories(ctx, string(categoryName))
	if err != nil {
		s.logger.Warnf("GET /categories/children/%s Error: %s", categoryName, err.Error())
		var statusCode int
		if strings.EqualFold(err.Error(), domain.CategoryNotFoundErr.Error()) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}
		response.JSON(w, statusCode, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, categories)
}

func (s Server) GetCategoriesCategoryName(w http.ResponseWriter, r *http.Request, categoryName spec.CategoryName) {
	ctx := r.Context()

	category, err := s.categoryUseCase.GetCategoryByName(ctx, string(categoryName))
	if err != nil {
		s.logger.Warnf("GET /categories/%s Error: %s", categoryName, err.Error())
		var statusCode int
		if strings.EqualFold(err.Error(), domain.CategoryNotFoundErr.Error()) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}
		response.JSON(w, statusCode, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, category)
}
