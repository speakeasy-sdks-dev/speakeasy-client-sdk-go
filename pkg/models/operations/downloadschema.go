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

type DownloadSchemaResponses struct {
	Error  *shared.Error
	Schema []byte
}

type DownloadSchemaResponse struct {
	ContentType string
	Responses   map[int64]map[string]DownloadSchemaResponses
	StatusCode  int64
}
