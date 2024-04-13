// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/sdkerrors"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type GetRevisionsRequest struct {
	NamespaceName string `pathParam:"style=simple,explode=false,name=namespace_name"`
	// Token to retrieve the next page of results
	NextPageToken *string `queryParam:"style=form,explode=true,name=next_page_token"`
}

func (o *GetRevisionsRequest) GetNamespaceName() string {
	if o == nil {
		return ""
	}
	return o.NamespaceName
}

func (o *GetRevisionsRequest) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}

type GetRevisionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *sdkerrors.Error
	// OK
	GetRevisionsResponse *shared.GetRevisionsResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetRevisionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRevisionsResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetRevisionsResponse) GetGetRevisionsResponse() *shared.GetRevisionsResponse {
	if o == nil {
		return nil
	}
	return o.GetRevisionsResponse
}

func (o *GetRevisionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRevisionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
