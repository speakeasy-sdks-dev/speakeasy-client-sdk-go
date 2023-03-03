package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type DownloadSchemaPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DownloadSchemaRequest struct {
	PathParams DownloadSchemaPathParams
}

type DownloadSchemaResponse struct {
	ContentType string
	Error       *shared.Error
	Schema      []byte
	StatusCode  int
	RawResponse *http.Response
}
