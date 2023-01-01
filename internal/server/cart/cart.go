package cart

import (
	"encoding/json"
	"net/http"
	"strings"
	adapter "yukiko-shop/internal/adapter/cart"
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/cart"
	"yukiko-shop/pkg/response"

	"github.com/google/uuid"
)

func (s Server) DeleteCart(w http.ResponseWriter, r *http.Request, params spec.DeleteCartParams) {
	ctx := r.Context()
	userID := uuid.MustParse(params.User)

	if err := s.cartUseCase.ClearCart(ctx, &domain.User{
		ID: userID,
	}); err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.EmptyJSON(w, http.StatusNoContent)
}

func (s Server) GetCart(w http.ResponseWriter, r *http.Request, params spec.GetCartParams) {
	ctx := r.Context()
	userID := uuid.MustParse(params.User)

	res, err := s.cartUseCase.GetCart(ctx, &domain.User{
		ID: userID,
	})
	if err != nil {
		var statusCode int
		if strings.EqualFold(err.Error(), domain.UserNotFoundErr.Error()) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}
		response.JSON(w, statusCode, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, res)
}

func (s Server) PostCart(w http.ResponseWriter, r *http.Request, params spec.PostCartParams) {
	ctx := r.Context()
	userID := uuid.MustParse(params.User)

	var request spec.AddProductToCartRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	if err := s.cartUseCase.AddProductToCart(ctx, adapter.PrepareCartProduct(&request, userID)); err != nil {
		var statusCode int
		if strings.EqualFold(err.Error(), domain.UserNotFoundErr.Error()) || 
				strings.EqualFold(err.Error(), domain.ProductNotFoundErr.Error()) || 
				strings.EqualFold(err.Error(), domain.CartProductAlreadyExistsErr.Error()) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}
		response.JSON(w, statusCode, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.EmptyJSON(w, http.StatusCreated)
}

func (s Server) DeleteCartProductID(w http.ResponseWriter, r *http.Request, productID spec.ProductID, params spec.DeleteCartProductIDParams) {
	ctx := r.Context()
	userID := uuid.MustParse(params.User)
	productUuid := uuid.MustParse(string(productID))

	if err := s.cartUseCase.DeleteProductFromCart(ctx, &domain.CartProduct{
		ProductID:         productUuid,
		CustomerID: userID,
	}); err != nil {
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

func (s Server) PatchCartProductID(w http.ResponseWriter, r *http.Request, productID spec.ProductID, params spec.PatchCartProductIDParams) {
	ctx := r.Context()
	userID := uuid.MustParse(params.User)
	productUuid := uuid.MustParse(string(productID))

	var request *spec.UpdateCartProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	
	if err := s.cartUseCase.UpdateProductQuantity(ctx, &domain.CartProduct{
		ProductID:         productUuid,
		CustomerID: userID,
		Quantity:   request.Quantity,
	}); err != nil {
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

	response.EmptyJSON(w, http.StatusOK)
}
