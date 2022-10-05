package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GeneratePostmanCollectionPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GeneratePostmanCollectionRequest struct {
	PathParams GeneratePostmanCollectionPathParams
}

type GeneratePostmanCollectionResponses struct {
	Error             *shared.Error
	PostmanCollection []byte
}

type GeneratePostmanCollectionResponse struct {
	ContentType string
	Responses   map[int64]map[string]GeneratePostmanCollectionResponses
	StatusCode  int64
}
