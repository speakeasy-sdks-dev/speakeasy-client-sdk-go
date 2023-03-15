package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type QueryEventLogRequest struct {
	Filters *shared.Filters `queryParam:"serialization=json,name=filters"`
}

type QueryEventLogResponse struct {
	BoundedRequests []shared.BoundedRequest
	ContentType     string
	Error           *shared.Error
	StatusCode      int
	RawResponse     *http.Response
}
