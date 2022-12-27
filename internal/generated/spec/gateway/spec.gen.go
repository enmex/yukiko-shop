// Package spec provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package spec

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

