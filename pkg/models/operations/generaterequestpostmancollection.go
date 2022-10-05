package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GenerateRequestPostmanCollectionPathParams struct {
	RequestID string `pathParam:"style=simple,explode=false,name=requestID"`
}

type GenerateRequestPostmanCollectionRequest struct {
	PathParams GenerateRequestPostmanCollectionPathParams
}

type GenerateRequestPostmanCollectionResponses struct {
	Error             *shared.Error
	PostmanCollection []byte
}

type GenerateRequestPostmanCollectionResponse struct {
	ContentType string
	Responses   map[int64]map[string]GenerateRequestPostmanCollectionResponses
	StatusCode  int64
}
