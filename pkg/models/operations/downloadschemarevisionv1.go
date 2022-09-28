package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type DownloadSchemaRevisionV1PathParams struct {
	APIID      string `pathParam:"style=simple,explode=false,name=apiID"`
	RevisionID string `pathParam:"style=simple,explode=false,name=revisionID"`
	VersionID  string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DownloadSchemaRevisionV1Request struct {
	PathParams DownloadSchemaRevisionV1PathParams
}

type DownloadSchemaRevisionV1Responses struct {
	Error *shared.Error
}

type DownloadSchemaRevisionV1Response struct {
	ContentType string
	Headers     map[string][]string
	Responses   map[int64]map[string]DownloadSchemaRevisionV1Responses
	StatusCode  int64
}
