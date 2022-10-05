package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetAllAPIVersionsPathParams struct {
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
}

type GetAllAPIVersionsOp struct {
	And bool `queryParam:"name=and"`
}

type GetAllAPIVersionsQueryParams struct {
	Metadata map[string][]string  `queryParam:"style=deepObject,explode=true,name=metadata"`
	Op       *GetAllAPIVersionsOp `queryParam:"style=deepObject,explode=true,name=op"`
}

type GetAllAPIVersionsRequest struct {
	PathParams  GetAllAPIVersionsPathParams
	QueryParams GetAllAPIVersionsQueryParams
}

type GetAllAPIVersionsResponses struct {
	Apis  []shared.API
	Error *shared.Error
}

type GetAllAPIVersionsResponse struct {
	ContentType string
	Responses   map[int64]map[string]GetAllAPIVersionsResponses
	StatusCode  int64
}
