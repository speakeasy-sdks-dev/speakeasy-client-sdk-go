// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

// GetApisOp - Configuration for filter operations
type GetApisOp struct {
	// Whether to AND or OR the filters
	And bool `queryParam:"name=and"`
}

type GetApisRequest struct {
	// Metadata to filter Apis on
	Metadata map[string][]string `queryParam:"style=deepObject,explode=true,name=metadata"`
	// Configuration for filter operations
	Op *GetApisOp `queryParam:"style=deepObject,explode=true,name=op"`
}

type GetApisResponse struct {
	// OK
	Apis        []shared.API
	ContentType string
	// Default error response
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
