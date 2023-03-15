package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GetAllAPIVersionsOp struct {
	And bool `queryParam:"name=and"`
}

type GetAllAPIVersionsRequest struct {
	APIID    string               `pathParam:"style=simple,explode=false,name=apiID"`
	Metadata map[string][]string  `queryParam:"style=deepObject,explode=true,name=metadata"`
	Op       *GetAllAPIVersionsOp `queryParam:"style=deepObject,explode=true,name=op"`
}

type GetAllAPIVersionsResponse struct {
	Apis        []shared.API
	ContentType string
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
