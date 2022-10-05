package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetAPIEndpointPathParams struct {
	APIEndpointID string `pathParam:"style=simple,explode=false,name=apiEndpointID"`
	APIID         string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID     string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetAPIEndpointRequest struct {
	PathParams GetAPIEndpointPathParams
}

type GetAPIEndpointResponses struct {
	APIEndpoint *shared.APIEndpoint
	Error       *shared.Error
}

type GetAPIEndpointResponse struct {
	ContentType string
	Responses   map[int64]map[string]GetAPIEndpointResponses
	StatusCode  int64
}
