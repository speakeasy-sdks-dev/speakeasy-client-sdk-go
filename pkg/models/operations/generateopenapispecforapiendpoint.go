package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GenerateOpenAPISpecForAPIEndpointPathParams struct {
	APIEndpointID string `pathParam:"style=simple,explode=false,name=apiEndpointID"`
	APIID         string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID     string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GenerateOpenAPISpecForAPIEndpointRequest struct {
	PathParams GenerateOpenAPISpecForAPIEndpointPathParams
}

type GenerateOpenAPISpecForAPIEndpointResponses struct {
	Error                   *shared.Error
	GenerateOpenAPISpecDiff *shared.GenerateOpenAPISpecDiff
}

type GenerateOpenAPISpecForAPIEndpointResponse struct {
	ContentType string
	Responses   map[int64]map[string]GenerateOpenAPISpecForAPIEndpointResponses
	StatusCode  int64
}
