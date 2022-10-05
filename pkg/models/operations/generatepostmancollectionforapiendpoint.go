package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GeneratePostmanCollectionForAPIEndpointPathParams struct {
	APIEndpointID string `pathParam:"style=simple,explode=false,name=apiEndpointID"`
	APIID         string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID     string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GeneratePostmanCollectionForAPIEndpointRequest struct {
	PathParams GeneratePostmanCollectionForAPIEndpointPathParams
}

type GeneratePostmanCollectionForAPIEndpointResponses struct {
	Error             *shared.Error
	PostmanCollection []byte
}

type GeneratePostmanCollectionForAPIEndpointResponse struct {
	ContentType string
	Responses   map[int64]map[string]GeneratePostmanCollectionForAPIEndpointResponses
	StatusCode  int64
}
