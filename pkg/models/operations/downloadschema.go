package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
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
	StatusCode  int64
}
