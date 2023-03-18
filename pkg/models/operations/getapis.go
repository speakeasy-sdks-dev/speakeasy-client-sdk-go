package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

// GetApisOp
// Configuration for filter operations
type GetApisOp struct {
	And bool `queryParam:"name=and"`
}

type GetApisRequest struct {
	Metadata map[string][]string `queryParam:"style=deepObject,explode=true,name=metadata"`
	Op       *GetApisOp          `queryParam:"style=deepObject,explode=true,name=op"`
}

type GetApisResponse struct {
	Apis        []shared.API
	ContentType string
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
