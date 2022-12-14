// Package spec provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package spec

// AddProductToCartRequest defines model for AddProductToCartRequest.
type AddProductToCartRequest struct {
	Name      string  `json:"name"`
	PhotoUrl  string  `json:"photoUrl"`
	Price     float64 `json:"price"`
	ProductID string  `json:"productID"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Message string `json:"message"`
}

// GetCartResponse defines model for GetCartResponse.
type GetCartResponse struct {
	Products   []Product `json:"products"`
	TotalPrice float64   `json:"totalPrice"`
}

// Product defines model for Product.
type Product struct {
	CustomerID string  `json:"customerID"`
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	PhotoUrl   string  `json:"photoUrl"`
	Price      float64 `json:"price"`
	ProductID  string  `json:"productID"`
	Quantity   int     `json:"quantity"`
}

// UpdateCartProductRequest defines model for UpdateCartProductRequest.
type UpdateCartProductRequest struct {
	Quantity int `json:"quantity"`
}

// ProductID defines model for productID.
type ProductID string

// DeleteCartParams defines parameters for DeleteCart.
type DeleteCartParams struct {
	User string `json:"user"`
}

// GetCartParams defines parameters for GetCart.
type GetCartParams struct {
	User string `json:"user"`
}

// PostCartJSONBody defines parameters for PostCart.
type PostCartJSONBody AddProductToCartRequest

// PostCartParams defines parameters for PostCart.
type PostCartParams struct {
	User string `json:"user"`
}

// DeleteCartProductIDParams defines parameters for DeleteCartProductID.
type DeleteCartProductIDParams struct {
	User string `json:"user"`
}

// PatchCartProductIDJSONBody defines parameters for PatchCartProductID.
type PatchCartProductIDJSONBody UpdateCartProductRequest

// PatchCartProductIDParams defines parameters for PatchCartProductID.
type PatchCartProductIDParams struct {
	User string `json:"user"`
}

// PostCartJSONRequestBody defines body for PostCart for application/json ContentType.
type PostCartJSONRequestBody PostCartJSONBody

// PatchCartProductIDJSONRequestBody defines body for PatchCartProductID for application/json ContentType.
type PatchCartProductIDJSONRequestBody PatchCartProductIDJSONBody

