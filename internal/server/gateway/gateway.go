package gateway

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	spec "yukiko-shop/internal/generated/spec/gateway"
	"yukiko-shop/internal/middleware"
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

func (s Server) GetAuthAccess(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID := middleware.GetUserIdFromContext(ctx)
	
	var url string
	if userID == nil {
		url = fmt.Sprintf("http://%s/auth/access", s.cfg.AuthServiceHost)
	} else {
		url = fmt.Sprintf("http://%s/auth/access?user=%s", s.cfg.AuthServiceHost, *userID)
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
	_, err := httpRequest.Delete(fmt.Sprintf("http://%s/products/%s", s.cfg.ProductServiceHost, string(productID)))
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	if _, err := httpRequest.Delete(fmt.Sprintf("http://%s/images/%s", s.cfg.ImageServiceHost, string(productID))); err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.EmptyJSON(w, http.StatusNoContent)
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
	url := fmt.Sprintf("http://%s/categories", s.cfg.ProductServiceHost)
	if params.Type != nil {
		url += fmt.Sprintf("?type=%s", *params.Type)
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

func (s Server) GetCategoriesChildrenCategoryID(w http.ResponseWriter, r *http.Request, categoryID spec.CategoryID) {
	res, err := httpRequest.Get(fmt.Sprintf("http://%s/categories/children/%s", s.cfg.ProductServiceHost, categoryID))
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.Reply(w, res.Code, []byte(*res.Body))
}

func (s Server) GetCategoriesCategoryID(w http.ResponseWriter, r *http.Request, categoryID spec.CategoryID) {
	res, err := httpRequest.Get(fmt.Sprintf("http://%s/categories/%s", s.cfg.ProductServiceHost, categoryID))
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

func (s Server) DeleteImagesImageID(w http.ResponseWriter, r *http.Request, imageID spec.ImageID) {
	res, err := httpRequest.Delete(fmt.Sprintf("http://%s/%s", s.cfg.ImageServiceHost, imageID))
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, spec.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	response.Reply(w, res.Code, []byte(*res.Body))
}