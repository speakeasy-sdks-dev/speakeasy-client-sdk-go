package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GenerateOpenAPISpecRequest struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GenerateOpenAPISpecResponse struct {
	ContentType             string
	Error                   *shared.Error
	GenerateOpenAPISpecDiff *shared.GenerateOpenAPISpecDiff
	StatusCode              int
	RawResponse             *http.Response
}
