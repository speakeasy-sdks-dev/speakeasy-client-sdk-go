package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type RunPluginRequest struct {
	Filters  *shared.Filters `queryParam:"serialization=json,name=filters"`
	PluginID string          `pathParam:"style=simple,explode=false,name=pluginID"`
}

type RunPluginResponse struct {
	BoundedRequests []shared.BoundedRequest
	ContentType     string
	Error           *shared.Error
	StatusCode      int
	RawResponse     *http.Response
}
