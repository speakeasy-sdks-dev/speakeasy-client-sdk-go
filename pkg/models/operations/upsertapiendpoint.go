package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type UpsertAPIEndpointPathParams struct {
	APIEndpointID string `pathParam:"style=simple,explode=false,name=apiEndpointID"`
	APIID         string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID     string `pathParam:"style=simple,explode=false,name=versionID"`
}

type UpsertAPIEndpointRequest struct {
	PathParams UpsertAPIEndpointPathParams
	Request    shared.APIEndpoint `request:"mediaType=application/json"`
}

type UpsertAPIEndpointResponses struct {
	APIEndpoint *shared.APIEndpoint
	Error       *shared.Error
}

type UpsertAPIEndpointResponse struct {
	ContentType string
	Responses   map[int64]map[string]UpsertAPIEndpointResponses
	StatusCode  int64
}
