package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GetVersionMetadataRequest struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetVersionMetadataResponse struct {
	ContentType     string
	Error           *shared.Error
	StatusCode      int
	RawResponse     *http.Response
	VersionMetadata []shared.VersionMetadata
}
