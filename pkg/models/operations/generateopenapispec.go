package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GenerateOpenAPISpecPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GenerateOpenAPISpecRequest struct {
	PathParams GenerateOpenAPISpecPathParams
}

type GenerateOpenAPISpecResponses struct {
	Error                   *shared.Error
	GenerateOpenAPISpecDiff *shared.GenerateOpenAPISpecDiff
}

type GenerateOpenAPISpecResponse struct {
	ContentType string
	Responses   map[int64]map[string]GenerateOpenAPISpecResponses
	StatusCode  int64
}
