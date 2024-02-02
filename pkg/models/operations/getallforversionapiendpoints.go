// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type GetAllForVersionAPIEndpointsRequest struct {
	// The ID of the Api to retrieve ApiEndpoints for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The version ID of the Api to retrieve ApiEndpoints for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

func (o *GetAllForVersionAPIEndpointsRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *GetAllForVersionAPIEndpointsRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

type GetAllForVersionAPIEndpointsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *shared.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Classes []shared.APIEndpoint
}

func (o *GetAllForVersionAPIEndpointsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAllForVersionAPIEndpointsResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetAllForVersionAPIEndpointsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAllForVersionAPIEndpointsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetAllForVersionAPIEndpointsResponse) GetClasses() []shared.APIEndpoint {
	if o == nil {
		return nil
	}
	return o.Classes
}
