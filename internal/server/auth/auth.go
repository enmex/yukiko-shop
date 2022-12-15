package auth

import (
	"encoding/json"
	"net/http"
	"strings"
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/auth"
	"yukiko-shop/pkg/response"
)

func (s Server) PostAuthSendVerifyCode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request spec.SendVerifyCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message:   err.Error(),
		})
		return
	}

	if err := s.authUseCase.SendVerifyCode(ctx, request); err != nil {
		if strings.EqualFold(err.Error(), domain.VerifyCodeExpiredErr.Error()) || strings.EqualFold(err.Error(), domain.VerifyCodeNotMatchErr.Error()) {
			response.JSON(w, http.StatusBadRequest, spec.ErrorResponse{
				Message:   err.Error(),
			})
			return
		}
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message:   err.Error(),
		})
		return
	}
}

func (s Server) PostAuthSignIn(w http.ResponseWriter, r *http.Request) {
	var request spec.SignInRequest
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message:   err.Error(),
		})
		return
	}

	res, err := s.authUseCase.SignIn(ctx, request)
	if err != nil {
		if strings.EqualFold(err.Error(), domain.UserNotFoundErr.Error()) {
			response.JSON(w, http.StatusBadRequest, spec.ErrorResponse{
				Message:   err.Error(),
			})
			return
		}

		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message:   err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, res)
}

func (s Server) PostAuthSignUp(w http.ResponseWriter, r *http.Request) {
	var request spec.SignUpRequest
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message:   err.Error(),
		})
		return
	}

	res, err := s.authUseCase.SignUp(ctx, request)
	if err != nil {
		if strings.EqualFold(err.Error(), domain.UserAlreadyExistsErr.Error()) {
			response.JSON(w, http.StatusBadRequest, spec.ErrorResponse{
				Message:   err.Error(),
			})
			return
		}

		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message:   err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, res)
}
