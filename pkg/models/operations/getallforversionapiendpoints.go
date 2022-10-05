package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetAllForVersionAPIEndpointsPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetAllForVersionAPIEndpointsRequest struct {
	PathParams GetAllForVersionAPIEndpointsPathParams
}

type GetAllForVersionAPIEndpointsResponses struct {
	APIEndpoints []shared.APIEndpoint
	Error        *shared.Error
}

type GetAllForVersionAPIEndpointsResponse struct {
	ContentType string
	Responses   map[int64]map[string]GetAllForVersionAPIEndpointsResponses
	StatusCode  int64
}
