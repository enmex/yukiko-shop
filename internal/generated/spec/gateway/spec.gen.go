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

	"H4sIAAAAAAAC/8xVX28bRRD/KtbA45FzaYWqe7vYF3qSdQ4XuwIqy1rOa3uL73a7u0eJKktJEIgHEIhP",
	"wBuvIWA1aqn5CrPfCO1eiP/U+UPlSLxY3p2Z3/7mNzczLyDjueAFLbSC4AVIqgQvFHWHsNTj9OLCnjNe",
	"aFpo+5cIMWEZ0YwX/lPFC3unsjHNif33vqRDCOA9fwHuV1blr4BOp1MPBlRlkgmLBQHgL+YYZzXzLc7N",
	"Ec7N9ziHqQeRlFxuncwq6iY2v5kTnH9gf2r4xvI6wVN8BdbzAmSTUkJyQaVmlY6k1OObmHQVlR3+JS1s",
	"skLyIZvQ28Q4JpI+K5mkAwieVI8tIHoe6ENBIQD+xVOa6Y1artKl1tzgA2eiRZlb2N2w2U+jT7rRQQc8",
	"eBy24mbYidtJfy+MW1ETPOgmYbfzqJ3Gn7vjXjvdjZvNKAEPknanv9fuJva+0U72WnGj41xazTj5uN+M",
	"9juP+tGnjShqutjGZ41W3Ohf2MGDOOlEaRK2+gdR+jhK+1GattOl3JSWrBjZ3HKqFBk56mu2NZ0WWS6C",
	"Nol1wEZFXKT0WUmV3iBWTthkw2seCKLUcy4Ht6DiMJYiruLRFe/AY8ik0gnJ6UbrhFxj/O8pLB5bgr4h",
	"s+qzfzujrwWTVIUu2SGXOdEQACv0Rw/gEoUVmo5sF3ig/4W5nmnl5i3B96YeuFbamqhssMK5LNkAvDvT",
	"3oHfugAXyV4hOskyqtRNg+dyUEk6lFSNb+m/PqmqtxYoPTeAWTHkFnB1EIf7cW1ENH1ODmv4u/kRX+EM",
	"3+CfeGqVZdpOy2Un8OArKlUVXN+p79yzfLmgBREMAri/U9+574TRY5evbwenr1y3O1141WZr++BnPLPL",
	"wBzhOb7Ec3NkvsGZOanh3zjH1+YHfIlzPMNTc4IzfG1+quFZzRzjuTl2N3/hDBwN6dZVPIAA9rnSdoFU",
	"kwYqkajSu3xwuLVNtzrGpqu10LKk7mJp+X9Yv3cV5qXf2jL34EG9fnPQ+tL1QJV5TuTh2/qemu/w3Ip4",
	"hbq29mSk7NfkOrhn0RaV7IprKvkrzvCPqjD2tW1XsivusJKLRfD/reSqvu9QSQtHpe1iCJ6sl6/FMzKp",
	"VXbwoJQTCGCstQh8f2JtY6508LD+sO7DtDf9JwAA///Wz51p6AoAAA==",
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
