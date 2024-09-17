// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"io"
	"net/http"
)

type Schema struct {
	// This field accepts []byte data or io.Reader implementations, such as *os.File.
	Content  any    `multipartForm:"content"`
	FileName string `multipartForm:"name=schema"`
}

func (o *Schema) GetContent() any {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *Schema) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

// SuggestOpenAPIRequestBody - The schema file to upload provided as a multipart/form-data file segment.
type SuggestOpenAPIRequestBody struct {
	Opts   *shared.SuggestOptsOld `multipartForm:"name=opts,json"`
	Schema Schema                 `multipartForm:"file"`
}

func (o *SuggestOpenAPIRequestBody) GetOpts() *shared.SuggestOptsOld {
	if o == nil {
		return nil
	}
	return o.Opts
}

func (o *SuggestOpenAPIRequestBody) GetSchema() Schema {
	if o == nil {
		return Schema{}
	}
	return o.Schema
}

type SuggestOpenAPIRequest struct {
	// The schema file to upload provided as a multipart/form-data file segment.
	RequestBody SuggestOpenAPIRequestBody `request:"mediaType=multipart/form-data"`
	XSessionID  string                    `header:"style=simple,explode=false,name=x-session-id"`
}

func (o *SuggestOpenAPIRequest) GetRequestBody() SuggestOpenAPIRequestBody {
	if o == nil {
		return SuggestOpenAPIRequestBody{}
	}
	return o.RequestBody
}

func (o *SuggestOpenAPIRequest) GetXSessionID() string {
	if o == nil {
		return ""
	}
	return o.XSessionID
}

type SuggestOpenAPIResponse struct {
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

func (o *SuggestOpenAPIResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SuggestOpenAPIResponse) GetSchema() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.Schema
}

func (o *SuggestOpenAPIResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SuggestOpenAPIResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
