package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type FindAPIEndpointV1PathParams struct {
	APIID       string `pathParam:"style=simple,explode=false,name=apiID"`
	DisplayName string `pathParam:"style=simple,explode=false,name=displayName"`
	VersionID   string `pathParam:"style=simple,explode=false,name=versionID"`
}

type FindAPIEndpointV1Request struct {
	PathParams FindAPIEndpointV1PathParams
}

type FindAPIEndpointV1Responses struct {
	APIEndpoint *shared.APIEndpoint
	Error       *shared.Error
	RawResponse []byte
}

type FindAPIEndpointV1Response struct {
	ContentType string
	Responses   map[int64]map[string]FindAPIEndpointV1Responses
	StatusCode  int64
}
