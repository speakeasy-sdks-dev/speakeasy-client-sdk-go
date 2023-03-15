package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GetRequestFromEventLogRequest struct {
	RequestID string `pathParam:"style=simple,explode=false,name=requestID"`
}

type GetRequestFromEventLogResponse struct {
	ContentType      string
	Error            *shared.Error
	StatusCode       int
	RawResponse      *http.Response
	UnboundedRequest *shared.UnboundedRequest
}
