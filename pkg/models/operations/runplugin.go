package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type RunPluginPathParams struct {
	PluginID string `pathParam:"style=simple,explode=false,name=pluginID"`
}

type RunPluginQueryParams struct {
	Filters *shared.Filters `queryParam:"serialization=json,name=filters"`
}

type RunPluginRequest struct {
	PathParams  RunPluginPathParams
	QueryParams RunPluginQueryParams
}

type RunPluginResponse struct {
	BoundedRequests []shared.BoundedRequest
	ContentType     string
	Error           *shared.Error
	StatusCode      int
}
