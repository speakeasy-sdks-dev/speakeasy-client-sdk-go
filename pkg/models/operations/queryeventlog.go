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

type QueryEventLogResponses struct {
	BoundedRequests []shared.BoundedRequest
	Error           *shared.Error
}

type QueryEventLogResponse struct {
	ContentType string
	Responses   map[int64]map[string]QueryEventLogResponses
	StatusCode  int64
}
