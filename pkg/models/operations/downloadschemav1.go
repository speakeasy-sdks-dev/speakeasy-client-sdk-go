package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type DownloadSchemaV1PathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DownloadSchemaV1Request struct {
	PathParams DownloadSchemaV1PathParams
}

type DownloadSchemaV1Responses struct {
	Error       *shared.Error
	RawResponse []byte
}

type DownloadSchemaV1Response struct {
	ContentType string
	Headers     map[string][]string
	Responses   map[int64]map[string]DownloadSchemaV1Responses
	StatusCode  int64
}
