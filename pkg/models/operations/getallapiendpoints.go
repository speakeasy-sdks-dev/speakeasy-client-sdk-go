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

type GetAllAPIEndpointsResponse struct {
	APIEndpoints []shared.APIEndpoint
	ContentType  string
	Error        *shared.Error
	StatusCode   int
}
