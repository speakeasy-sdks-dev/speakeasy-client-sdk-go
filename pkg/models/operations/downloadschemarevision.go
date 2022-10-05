package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type DownloadSchemaRevisionPathParams struct {
	APIID      string `pathParam:"style=simple,explode=false,name=apiID"`
	RevisionID string `pathParam:"style=simple,explode=false,name=revisionID"`
	VersionID  string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DownloadSchemaRevisionRequest struct {
	PathParams DownloadSchemaRevisionPathParams
}

type DownloadSchemaRevisionResponses struct {
	Error  *shared.Error
	Schema []byte
}

type DownloadSchemaRevisionResponse struct {
	ContentType string
	Responses   map[int64]map[string]DownloadSchemaRevisionResponses
	StatusCode  int64
}
