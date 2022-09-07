package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type UpsertAPIEndpointV1PathParams struct {
	APIEndpointID string `pathParam:"style=simple,explode=false,name=apiEndpointID"`
	APIID         string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID     string `pathParam:"style=simple,explode=false,name=versionID"`
}

type UpsertAPIEndpointV1Request struct {
	PathParams UpsertAPIEndpointV1PathParams
	Request    shared.APIEndpoint `request:"mediaType=application/json"`
}

type UpsertAPIEndpointV1Responses struct {
	APIEndpoint *shared.APIEndpoint
	Error       *shared.Error
	RawResponse []byte
}

type UpsertAPIEndpointV1Response struct {
	ContentType string
	Responses   map[int64]map[string]UpsertAPIEndpointV1Responses
	StatusCode  int64
}
