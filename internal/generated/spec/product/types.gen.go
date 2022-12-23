// Package spec provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package spec

// Category defines model for Category.
type Category struct {
	Children []Category `json:"children"`
	Id       string     `json:"id"`
	Name     string     `json:"name"`
	Parent   *Category  `json:"parent,omitempty"`
	PhotoUrl *string    `json:"photoUrl,omitempty"`
	Products []Product  `json:"products"`
}

// CreateCategoryRequest defines model for CreateCategoryRequest.
type CreateCategoryRequest struct {
	Name     string  `json:"name"`
	Parent   *string `json:"parent,omitempty"`
	PhotoUrl *string `json:"photoUrl,omitempty"`
}

// CreateProductRequest defines model for CreateProductRequest.
type CreateProductRequest struct {
	CategoryName string  `json:"categoryName"`
	Description  string  `json:"description"`
	Name         string  `json:"name"`
	PhotoUrl     string  `json:"photoUrl"`
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

// GetCategoriesResponse defines model for GetCategoriesResponse.
type GetCategoriesResponse struct {
	Categories []Category `json:"categories"`
}

// GetCategoryResponse defines model for GetCategoryResponse.
type GetCategoryResponse struct {
	// Embedded struct due to allOf(#/components/schemas/Category)
	Category `yaml:",inline"`
	// Embedded fields due to inline allOf schema
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

// GetSubCategoriesResponse defines model for GetSubCategoriesResponse.
type GetSubCategoriesResponse struct {
	Categories []string `json:"categories"`
}

// Product defines model for Product.
type Product struct {
	CategoryName string  `json:"categoryName"`
	Description  string  `json:"description"`
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	PhotoUrl     *string `json:"photoUrl,omitempty"`
	Price        float64 `json:"price"`
}

// CategoryName defines model for categoryName.
type CategoryName string

// ProductID defines model for productID.
type ProductID string

// GetCategoriesParams defines parameters for GetCategories.
type GetCategoriesParams struct {
	Main *bool `json:"main,omitempty"`
	Leaf *bool `json:"leaf,omitempty"`
}

// PostCategoriesJSONBody defines parameters for PostCategories.
type PostCategoriesJSONBody CreateCategoryRequest

// GetProductsParams defines parameters for GetProducts.
type GetProductsParams struct {
	Limit *int `json:"limit,omitempty"`
}

// PostProductsJSONBody defines parameters for PostProducts.
type PostProductsJSONBody CreateProductRequest

// PostCategoriesJSONRequestBody defines body for PostCategories for application/json ContentType.
type PostCategoriesJSONRequestBody PostCategoriesJSONBody

// PostProductsJSONRequestBody defines body for PostProducts for application/json ContentType.
type PostProductsJSONRequestBody PostProductsJSONBody

