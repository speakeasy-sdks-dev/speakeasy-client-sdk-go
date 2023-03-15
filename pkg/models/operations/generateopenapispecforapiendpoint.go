package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GenerateOpenAPISpecForAPIEndpointRequest struct {
	APIEndpointID string `pathParam:"style=simple,explode=false,name=apiEndpointID"`
	APIID         string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID     string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GenerateOpenAPISpecForAPIEndpointResponse struct {
	ContentType             string
	Error                   *shared.Error
	GenerateOpenAPISpecDiff *shared.GenerateOpenAPISpecDiff
	StatusCode              int
	RawResponse             *http.Response
}
