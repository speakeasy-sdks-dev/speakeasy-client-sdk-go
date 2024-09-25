// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

// Op - Configuration for filter operations
type Op struct {
	// Whether to AND or OR the filters
	And bool `queryParam:"name=and"`
}

func (o *Op) GetAnd() bool {
	if o == nil {
		return false
	}
	return o.And
}

type GetAllAPIVersionsRequest struct {
	// The ID of the Api to retrieve.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// Metadata to filter Apis on
	Metadata map[string][]string `queryParam:"style=deepObject,explode=true,name=metadata"`
	// Configuration for filter operations
	Op *Op `queryParam:"style=deepObject,explode=true,name=op"`
}

func (o *GetAllAPIVersionsRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *GetAllAPIVersionsRequest) GetMetadata() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *GetAllAPIVersionsRequest) GetOp() *Op {
	if o == nil {
		return nil
	}
	return o.Op
}

type GetAllAPIVersionsResponse struct {
	// OK
	Apis []shared.API
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *shared.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetAllAPIVersionsResponse) GetApis() []shared.API {
	if o == nil {
		return nil
	}
	return o.Apis
}

func (o *GetAllAPIVersionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllAPIVersionsResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetAllAPIVersionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllAPIVersionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
