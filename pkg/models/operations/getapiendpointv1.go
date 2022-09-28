package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetAPIEndpointV1PathParams struct {
	APIEndpointID string `pathParam:"style=simple,explode=false,name=apiEndpointID"`
	APIID         string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID     string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetAPIEndpointV1Request struct {
	PathParams GetAPIEndpointV1PathParams
}

type GetAPIEndpointV1Responses struct {
	APIEndpoint *shared.APIEndpoint
	Error       *shared.Error
}

type GetAPIEndpointV1Response struct {
	ContentType string
	Responses   map[int64]map[string]GetAPIEndpointV1Responses
	StatusCode  int64
}
