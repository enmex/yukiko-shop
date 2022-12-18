package gateway

import (
	"encoding/json"
	"fmt"
	"net/http"
	spec "yukiko-shop/internal/generated/spec/gateway"
	httpRequest "yukiko-shop/pkg/request"
	"yukiko-shop/pkg/response"
)

func (s Server) PostAuthSendVerifyCode(w http.ResponseWriter, r *http.Request) {
	var request spec.SendVerifyCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	res, err := httpRequest.Post(fmt.Sprintf("http://%s/auth/sendVerifyCode", s.cfg.AuthServiceHost), request)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.Reply(w, res.Code, []byte(*res.Body))
}

func (s Server) PostAuthSignUp(w http.ResponseWriter, r *http.Request) {
	var request spec.SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	res, err := httpRequest.Post(fmt.Sprintf("http://%s/auth/signUp", s.cfg.AuthServiceHost), request)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
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
			Message: err.Error(),
		})
		return
	}

	res, err := httpRequest.Post(fmt.Sprintf("http://%s/auth/signIn", s.cfg.AuthServiceHost), request)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.Reply(w, res.Code, []byte(*res.Body))
}

func (s Server) PostProducts(w http.ResponseWriter, r *http.Request) {
	var request spec.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	res, err := httpRequest.Post(fmt.Sprintf("http://%s/products", s.cfg.ProductServiceHost), request)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.Reply(w, res.Code, []byte(*res.Body))
}

func (s Server) DeleteProductsProductID(w http.ResponseWriter, r *http.Request, productID spec.ProductID) {
	res, err := httpRequest.Delete(fmt.Sprintf("http://%s/products/%s", s.cfg.ProductServiceHost, string(productID)))
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.Reply(w, res.Code, []byte(*res.Body))
}

func (s Server) GetProductsProductID(w http.ResponseWriter, r *http.Request, productID spec.ProductID) {
	res, err := httpRequest.Get(fmt.Sprintf("http://%s/products/%s", s.cfg.ProductServiceHost, string(productID)))
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.Reply(w, res.Code, []byte(*res.Body))
}

func (s Server) GetProducts(w http.ResponseWriter, r *http.Request, params spec.GetProductsParams) {
	var url string
	if params.Limit != nil {
		url = fmt.Sprintf("http://%s/products?limit=%d", s.cfg.ProductServiceHost, *params.Limit)
	} else {
		url = fmt.Sprintf("http://%s/products", s.cfg.ProductServiceHost)
	}

	res, err := httpRequest.Get(url)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.Reply(w, res.Code, []byte(*res.Body))
}