// Package spec provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package spec

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Message string `json:"message"`
}

// UploadImageResponse defines model for UploadImageResponse.
type UploadImageResponse struct {
	Id       string `json:"id"`
	PhotoUrl string `json:"photoUrl"`
}

// ImageID defines model for imageID.
type ImageID string

