package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetAllAPIEndpointsPathParams struct {
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
}

type GetAllAPIEndpointsRequest struct {
	PathParams GetAllAPIEndpointsPathParams
}

type GetAllAPIEndpointsResponses struct {
	APIEndpoints []shared.APIEndpoint
	Error        *shared.Error
}

type GetAllAPIEndpointsResponse struct {
	ContentType string
	Responses   map[int64]map[string]GetAllAPIEndpointsResponses
	StatusCode  int64
}
