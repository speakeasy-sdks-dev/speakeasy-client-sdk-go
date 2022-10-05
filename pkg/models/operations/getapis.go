package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetApisOp struct {
	And bool `queryParam:"name=and"`
}

type GetApisQueryParams struct {
	Metadata map[string][]string `queryParam:"style=deepObject,explode=true,name=metadata"`
	Op       *GetApisOp          `queryParam:"style=deepObject,explode=true,name=op"`
}

type GetApisRequest struct {
	QueryParams GetApisQueryParams
}

type GetApisResponses struct {
	Apis  []shared.API
	Error *shared.Error
}

type GetApisResponse struct {
	ContentType string
	Responses   map[int64]map[string]GetApisResponses
	StatusCode  int64
}
