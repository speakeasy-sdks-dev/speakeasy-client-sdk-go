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

type GenerateOpenAPISpecResponse struct {
	ContentType             string
	Error                   *shared.Error
	GenerateOpenAPISpecDiff *shared.GenerateOpenAPISpecDiff
	StatusCode              int64
}
