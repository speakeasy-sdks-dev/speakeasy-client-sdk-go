package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetAllAPIEndpointsV1PathParams struct {
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
}

type GetAllAPIEndpointsV1Request struct {
	PathParams GetAllAPIEndpointsV1PathParams
}

type GetAllAPIEndpointsV1Responses struct {
	APIEndpoint []shared.APIEndpoint
	Error       *shared.Error
	RawResponse []byte
}

type GetAllAPIEndpointsV1Response struct {
	ContentType string
	Responses   map[int64]map[string]GetAllAPIEndpointsV1Responses
	StatusCode  int64
}
