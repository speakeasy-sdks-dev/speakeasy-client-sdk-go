package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GenerateRequestPostmanCollectionRequest struct {
	RequestID string `pathParam:"style=simple,explode=false,name=requestID"`
}

type GenerateRequestPostmanCollectionResponse struct {
	ContentType       string
	Error             *shared.Error
	PostmanCollection []byte
	StatusCode        int
	RawResponse       *http.Response
}
