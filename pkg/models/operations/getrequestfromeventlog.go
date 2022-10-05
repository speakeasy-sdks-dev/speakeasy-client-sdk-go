package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetRequestFromEventLogPathParams struct {
	RequestID string `pathParam:"style=simple,explode=false,name=requestID"`
}

type GetRequestFromEventLogRequest struct {
	PathParams GetRequestFromEventLogPathParams
}

type GetRequestFromEventLogResponses struct {
	Error            *shared.Error
	UnboundedRequest *shared.UnboundedRequest
}

type GetRequestFromEventLogResponse struct {
	ContentType string
	Responses   map[int64]map[string]GetRequestFromEventLogResponses
	StatusCode  int64
}
