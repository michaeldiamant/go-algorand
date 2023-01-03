// Package data provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package data

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	. "github.com/algorand/go-algorand/daemon/algod/api/server/v2/generated/model"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get a LedgerStateDelta object for a given round
	// (GET /v2/deltas/{round})
	GetLedgerStateDelta(ctx echo.Context, round uint64) error
	// Removes minimum sync round restriction from the ledger.
	// (DELETE /v2/ledger/sync)
	UnsetSyncRound(ctx echo.Context) error
	// Returns the minimum sync round the ledger is keeping in cache.
	// (GET /v2/ledger/sync)
	GetSyncRound(ctx echo.Context) error
	// Given a round, tells the ledger to keep that round in its cache.
	// (POST /v2/ledger/sync/{round})
	SetSyncRound(ctx echo.Context, round uint64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetLedgerStateDelta converts echo context to params.
func (w *ServerInterfaceWrapper) GetLedgerStateDelta(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameterWithLocation("simple", false, "round", runtime.ParamLocationPath, ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetLedgerStateDelta(ctx, round)
	return err
}

// UnsetSyncRound converts echo context to params.
func (w *ServerInterfaceWrapper) UnsetSyncRound(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UnsetSyncRound(ctx)
	return err
}

// GetSyncRound converts echo context to params.
func (w *ServerInterfaceWrapper) GetSyncRound(ctx echo.Context) error {
	var err error

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSyncRound(ctx)
	return err
}

// SetSyncRound converts echo context to params.
func (w *ServerInterfaceWrapper) SetSyncRound(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameterWithLocation("simple", false, "round", runtime.ParamLocationPath, ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set(Api_keyScopes, []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.SetSyncRound(ctx, round)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface, m ...echo.MiddlewareFunc) {
	RegisterHandlersWithBaseURL(router, si, "", m...)
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/v2/deltas/:round", wrapper.GetLedgerStateDelta, m...)
	router.DELETE(baseURL+"/v2/ledger/sync", wrapper.UnsetSyncRound, m...)
	router.GET(baseURL+"/v2/ledger/sync", wrapper.GetSyncRound, m...)
	router.POST(baseURL+"/v2/ledger/sync/:round", wrapper.SetSyncRound, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"YWj3h0//PwAA//9gJncgTcsAAA==",
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