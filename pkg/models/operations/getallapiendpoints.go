package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GetAllAPIEndpointsRequest struct {
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
}

type GetAllAPIEndpointsResponse struct {
	APIEndpoints []shared.APIEndpoint
	ContentType  string
	Error        *shared.Error
	StatusCode   int
	RawResponse  *http.Response
}
