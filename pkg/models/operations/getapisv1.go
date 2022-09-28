package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetApisV1Op struct {
	And bool `queryParam:"name=and"`
}

type GetApisV1QueryParams struct {
	Metadata map[string][]string `queryParam:"style=deepObject,explode=true,name=metadata"`
	Op       *GetApisV1Op        `queryParam:"style=deepObject,explode=true,name=op"`
}

type GetApisV1Request struct {
	QueryParams GetApisV1QueryParams
}

type GetApisV1Responses struct {
	Error                       *shared.Error
	GetApisV1200ApplicationJSON []shared.API
}

type GetApisV1Response struct {
	ContentType string
	Responses   map[int64]map[string]GetApisV1Responses
	StatusCode  int64
}
