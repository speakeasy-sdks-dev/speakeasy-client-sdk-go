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

type GenerateOpenAPISpecForAPIEndpointResponse struct {
	ContentType             string
	Error                   *shared.Error
	GenerateOpenAPISpecDiff *shared.GenerateOpenAPISpecDiff
	StatusCode              int64
}
