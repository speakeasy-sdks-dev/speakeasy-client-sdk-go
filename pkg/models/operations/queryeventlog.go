package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type QueryEventLogQueryParams struct {
	Filters *shared.Filters `queryParam:"serialization=json,name=filters"`
}

type QueryEventLogRequest struct {
	QueryParams QueryEventLogQueryParams
}

type QueryEventLogResponse struct {
	BoundedRequests []shared.BoundedRequest
	ContentType     string
	Error           *shared.Error
	StatusCode      int
}
