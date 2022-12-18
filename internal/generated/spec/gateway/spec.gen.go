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
	"BzGklT0cAAA=",
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

