package auth

import (
	"encoding/json"
	"net/http"
	"strings"
	adapter "yukiko-shop/internal/adapter/user"
	"yukiko-shop/internal/domain"
	spec "yukiko-shop/internal/generated/spec/auth"
	"yukiko-shop/pkg/response"
)

func (s Server) PostAuthSendVerifyCode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request spec.SendVerifyCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	if err := s.authUseCase.SendVerifyCode(ctx, request); err != nil {
		var statusCode int
		if strings.EqualFold(err.Error(), domain.VerifyCodeExpiredErr.Error()) || strings.EqualFold(err.Error(), domain.VerifyCodeNotMatchErr.Error()) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}
		s.logger.Infof("/auth/sendVerifyCode Error: %s", err.Error())
		response.JSON(w, statusCode, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
}

func (s Server) PostAuthSignIn(w http.ResponseWriter, r *http.Request) {
	var request spec.SignInRequest
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	res, err := s.authUseCase.SignIn(ctx, adapter.PrepareUserSignIn(&request))
	if err != nil {
		var statusCode int
		if strings.EqualFold(err.Error(), domain.UserNotFoundErr.Error()) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}
		s.logger.Infof("/auth/signIn Error: %s", err.Error())
		response.JSON(w, statusCode, spec.ErrorResponse{
			Message: err.Error(),
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
			Message: err.Error(),
		})
		return
	}

	user, code := adapter.PrepareUserSignUp(&request)

	res, err := s.authUseCase.SignUp(ctx, user, code)
	if err != nil {
		var statusCode int
		if strings.EqualFold(err.Error(), domain.UserAlreadyExistsErr.Error()) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}
		s.logger.Infof("/auth/signUp Error: %s", err.Error())
		response.JSON(w, statusCode, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.JSON(w, http.StatusOK, res)
}
