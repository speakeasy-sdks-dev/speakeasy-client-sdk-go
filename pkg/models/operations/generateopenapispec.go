// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GenerateOpenAPISpecRequest struct {
	// The ID of the Api to generate an OpenAPI specification for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The version ID of the Api to generate an OpenAPI specification for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

func (o *GenerateOpenAPISpecRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *GenerateOpenAPISpecRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

type GenerateOpenAPISpecResponse struct {
	ContentType string
	// Default error response
	Error *shared.Error
	// OK
	GenerateOpenAPISpecDiff *shared.GenerateOpenAPISpecDiff
	StatusCode              int
	RawResponse             *http.Response
}

func (o *GenerateOpenAPISpecResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GenerateOpenAPISpecResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GenerateOpenAPISpecResponse) GetGenerateOpenAPISpecDiff() *shared.GenerateOpenAPISpecDiff {
	if o == nil {
		return nil
	}
	return o.GenerateOpenAPISpecDiff
}

func (o *GenerateOpenAPISpecResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GenerateOpenAPISpecResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
