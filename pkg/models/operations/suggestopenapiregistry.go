// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"io"
	"net/http"
)

type SuggestOpenAPIRegistryRequest struct {
	// Suggest options
	SuggestRequestBody *shared.SuggestRequestBody `request:"mediaType=application/json"`
	NamespaceName      string                     `pathParam:"style=simple,explode=false,name=namespace_name"`
	// Tag or digest
	RevisionReference string `pathParam:"style=simple,explode=false,name=revision_reference"`
	XSessionID        string `header:"style=simple,explode=false,name=x-session-id"`
}

func (o *SuggestOpenAPIRegistryRequest) GetSuggestRequestBody() *shared.SuggestRequestBody {
	if o == nil {
		return nil
	}
	return o.SuggestRequestBody
}

func (o *SuggestOpenAPIRegistryRequest) GetNamespaceName() string {
	if o == nil {
		return ""
	}
	return o.NamespaceName
}

func (o *SuggestOpenAPIRegistryRequest) GetRevisionReference() string {
	if o == nil {
		return ""
	}
	return o.RevisionReference
}

func (o *SuggestOpenAPIRegistryRequest) GetXSessionID() string {
	if o == nil {
		return ""
	}
	return o.XSessionID
}

type SuggestOpenAPIRegistryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// An overlay containing the suggested spec modifications.
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	Schema io.ReadCloser
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SuggestOpenAPIRegistryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SuggestOpenAPIRegistryResponse) GetSchema() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *SuggestOpenAPIRegistryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SuggestOpenAPIRegistryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
