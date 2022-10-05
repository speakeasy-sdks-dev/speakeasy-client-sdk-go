package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type DeleteAPIEndpointPathParams struct {
	APIEndpointID string `pathParam:"style=simple,explode=false,name=apiEndpointID"`
	APIID         string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID     string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DeleteAPIEndpointRequest struct {
	PathParams DeleteAPIEndpointPathParams
}

type DeleteAPIEndpointResponses struct {
	Error *shared.Error
}

type DeleteAPIEndpointResponse struct {
	ContentType string
	Responses   map[int64]map[string]DeleteAPIEndpointResponses
	StatusCode  int64
}
