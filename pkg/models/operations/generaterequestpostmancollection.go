package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GenerateRequestPostmanCollectionPathParams struct {
	RequestID string `pathParam:"style=simple,explode=false,name=requestID"`
}

type GenerateRequestPostmanCollectionRequest struct {
	PathParams GenerateRequestPostmanCollectionPathParams
}

type GenerateRequestPostmanCollectionResponse struct {
	ContentType       string
	Error             *shared.Error
	PostmanCollection []byte
	StatusCode        int
	RawResponse       *http.Response
}
