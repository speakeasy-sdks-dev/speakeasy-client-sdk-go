package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type DownloadSchemaRequest struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DownloadSchemaResponse struct {
	ContentType string
	Error       *shared.Error
	Schema      []byte
	StatusCode  int
	RawResponse *http.Response
}
