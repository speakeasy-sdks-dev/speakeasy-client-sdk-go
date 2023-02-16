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

type GetAllAPIVersionsResponse struct {
	Apis        []shared.API
	ContentType string
	Error       *shared.Error
	StatusCode  int
}
