package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type FindAPIEndpointPathParams struct {
	APIID       string `pathParam:"style=simple,explode=false,name=apiID"`
	DisplayName string `pathParam:"style=simple,explode=false,name=displayName"`
	VersionID   string `pathParam:"style=simple,explode=false,name=versionID"`
}

type FindAPIEndpointRequest struct {
	PathParams FindAPIEndpointPathParams
}

type FindAPIEndpointResponses struct {
	APIEndpoint *shared.APIEndpoint
	Error       *shared.Error
}

type FindAPIEndpointResponse struct {
	ContentType string
	Responses   map[int64]map[string]FindAPIEndpointResponses
	StatusCode  int64
}
