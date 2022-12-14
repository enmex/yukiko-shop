package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"
	spec "yukiko-shop/internal/generated/spec/gateway"
	httpRequest "yukiko-shop/pkg/request"
	"yukiko-shop/pkg/response"
)

func (s Server) PostAuthSignUp(w http.ResponseWriter, r *http.Request) {
	var request spec.SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			ErrorCode: spec.ErrorResponseErrorCodeINTERNALSERVERERROR,
			Message: err.Error(),
		})
		return
	}

	res, err := httpRequest.Post(fmt.Sprintf("http://%s/auth/signUp", s.cfg.AuthServiceHost,), request)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			ErrorCode: spec.ErrorResponseErrorCodeINTERNALSERVERERROR,
			Message: err.Error(),
		})
		return
	}

	response.Reply(w, res.Code, []byte(*res.Body))
}

func (s Server) PostAuthSignIn(w http.ResponseWriter, r *http.Request) {
	var request spec.SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			ErrorCode: spec.ErrorResponseErrorCodeINTERNALSERVERERROR,
			Message: err.Error(),
		})
		return
	}

	res, err := httpRequest.Post(fmt.Sprintf("http://%s/auth/signIn", s.cfg.AuthServiceHost,), request)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			ErrorCode: spec.ErrorResponseErrorCodeINTERNALSERVERERROR,
			Message: err.Error(),
		})
		return
	}

	response.Reply(w, res.Code, []byte(*res.Body))
}
