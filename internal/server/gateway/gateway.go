package gateway

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (s Server) PostCategories(w http.ResponseWriter, r *http.Request) {
	var request spec.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	res, err := httpRequest.Post(fmt.Sprintf("http://%s/categories", s.cfg.ProductServiceHost), request)
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.Reply(w, res.Code, []byte(*res.Body))
}

func (s Server) GetCategories(w http.ResponseWriter, r *http.Request, params spec.GetCategoriesParams) {
	var url string
	if params.Main != nil {
		url = fmt.Sprintf("http://%s/categories?main=%t", s.cfg.ProductServiceHost, *params.Main)
	} else if params.Leaf != nil {
		url = fmt.Sprintf("http://%s/categories?leaf=%t", s.cfg.ProductServiceHost, *params.Leaf)
	} else {
		url = fmt.Sprintf("http://%s/categories", s.cfg.ProductServiceHost)
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

func (s Server) GetCategoriesChildrenCategoryName(w http.ResponseWriter, r *http.Request, categoryName spec.CategoryName) {
	res, err := httpRequest.Get(fmt.Sprintf("http://%s/categories/children/%s", s.cfg.ProductServiceHost, categoryName))
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.Reply(w, res.Code, []byte(*res.Body))
}

func (s Server) GetCategoriesCategoryName(w http.ResponseWriter, r *http.Request, categoryName spec.CategoryName) {
	res, err := httpRequest.Get(fmt.Sprintf("http://%s/categories/%s", s.cfg.ProductServiceHost, categoryName))
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.Reply(w, res.Code, []byte(*res.Body))
}

func (s Server) PostImages(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header["Content-Type"]
	res, err := http.Post(fmt.Sprintf("http://%s/images", s.cfg.ImageServiceHost), contentType[0], r.Body)
	if err != nil {
		s.logger.Warn(err)
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	defer res.Body.Close()
	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		s.logger.Warn(err)
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.Reply(w, res.StatusCode, resData)
}
