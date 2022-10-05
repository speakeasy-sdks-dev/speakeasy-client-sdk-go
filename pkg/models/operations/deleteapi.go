package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type DeleteAPIPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DeleteAPIRequest struct {
	PathParams DeleteAPIPathParams
}

type DeleteAPIResponses struct {
	Error *shared.Error
}

type DeleteAPIResponse struct {
	ContentType string
	Responses   map[int64]map[string]DeleteAPIResponses
	StatusCode  int64
}
