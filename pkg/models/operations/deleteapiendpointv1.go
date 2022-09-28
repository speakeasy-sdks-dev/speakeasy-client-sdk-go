package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type DeleteAPIEndpointV1PathParams struct {
	APIEndpointID string `pathParam:"style=simple,explode=false,name=apiEndpointID"`
	APIID         string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID     string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DeleteAPIEndpointV1Request struct {
	PathParams DeleteAPIEndpointV1PathParams
}

type DeleteAPIEndpointV1Responses struct {
	Error *shared.Error
}

type DeleteAPIEndpointV1Response struct {
	ContentType string
	Responses   map[int64]map[string]DeleteAPIEndpointV1Responses
	StatusCode  int64
}
