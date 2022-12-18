// Package spec provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package spec

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// AuthResponse defines model for AuthResponse.
type AuthResponse struct {
	Auth    UserToken `json:"auth"`
	Profile User      `json:"profile"`
}

// Category defines model for Category.
type Category struct {
	ChildrenCategories *[]Category `json:"childrenCategories,omitempty"`
	Id                 string      `json:"id"`
	Name               string      `json:"name"`
	ParentCategory     *Category   `json:"parentCategory,omitempty"`
	Products           *[]Product  `json:"products,omitempty"`
}

// CreateCategoryRequest defines model for CreateCategoryRequest.
type CreateCategoryRequest struct {
	Name           string  `json:"name"`
	ParentCategory *string `json:"parentCategory,omitempty"`
}

// CreateProductRequest defines model for CreateProductRequest.
type CreateProductRequest struct {
	CategoryName string  `json:"categoryName"`
	Description  string  `json:"description"`
	Name         string  `json:"name"`
	Path         string  `json:"path"`
	Price        float64 `json:"price"`
}

// CreateProductResponse defines model for CreateProductResponse.
type CreateProductResponse struct {
	CategoryName string  `json:"categoryName"`
	Description  string  `json:"description"`
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	PhotoUrl     string  `json:"photoUrl"`
	Price        float64 `json:"price"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Message string `json:"message"`
}

// GetProductResponse defines model for GetProductResponse.
type GetProductResponse struct {
	CategoryName string  `json:"categoryName"`
	Description  string  `json:"description"`
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	PhotoUrl     string  `json:"photoUrl"`
	Price        float64 `json:"price"`
}

// GetProductsResponse defines model for GetProductsResponse.
type GetProductsResponse struct {
	Products []Product `json:"products"`
}

// Product defines model for Product.
type Product struct {
	CategoryName string  `json:"categoryName"`
	Description  string  `json:"description"`
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	PhotoUrl     string  `json:"photoUrl"`
	Price        float64 `json:"price"`
}

// SendVerifyCodeRequest defines model for SendVerifyCodeRequest.
type SendVerifyCodeRequest struct {
	Email string `json:"email"`
}

// SignInRequest defines model for SignInRequest.
type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignUpRequest defines model for SignUpRequest.
type SignUpRequest struct {
	Code      int    `json:"code"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

// Token defines model for Token.
type Token struct {
	ExpiresAt int64  `json:"expiresAt"`
	Token     string `json:"token"`
}

// User defines model for User.
type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	Id        string `json:"id"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

// UserToken defines model for UserToken.
type UserToken struct {
	Access  Token `json:"access"`
	Refresh Token `json:"refresh"`
}

// ProductID defines model for productID.
type ProductID string

// PostAuthSendVerifyCodeJSONBody defines parameters for PostAuthSendVerifyCode.
type PostAuthSendVerifyCodeJSONBody SendVerifyCodeRequest

// PostAuthSignInJSONBody defines parameters for PostAuthSignIn.
type PostAuthSignInJSONBody SignInRequest

// PostAuthSignUpJSONBody defines parameters for PostAuthSignUp.
type PostAuthSignUpJSONBody SignUpRequest

// PostCategoriesJSONBody defines parameters for PostCategories.
type PostCategoriesJSONBody CreateCategoryRequest

// GetProductsParams defines parameters for GetProducts.
type GetProductsParams struct {
	Limit *int `json:"limit,omitempty"`
}

// PostProductsJSONBody defines parameters for PostProducts.
type PostProductsJSONBody CreateProductRequest

// PostAuthSendVerifyCodeJSONRequestBody defines body for PostAuthSendVerifyCode for application/json ContentType.
type PostAuthSendVerifyCodeJSONRequestBody PostAuthSendVerifyCodeJSONBody

// PostAuthSignInJSONRequestBody defines body for PostAuthSignIn for application/json ContentType.
type PostAuthSignInJSONRequestBody PostAuthSignInJSONBody

// PostAuthSignUpJSONRequestBody defines body for PostAuthSignUp for application/json ContentType.
type PostAuthSignUpJSONRequestBody PostAuthSignUpJSONBody

// PostCategoriesJSONRequestBody defines body for PostCategories for application/json ContentType.
type PostCategoriesJSONRequestBody PostCategoriesJSONBody

// PostProductsJSONRequestBody defines body for PostProducts for application/json ContentType.
type PostProductsJSONRequestBody PostProductsJSONBody

