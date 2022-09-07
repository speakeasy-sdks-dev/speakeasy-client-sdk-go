package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetVersionMetadataV1PathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetVersionMetadataV1Request struct {
	PathParams GetVersionMetadataV1PathParams
}

type GetVersionMetadataV1Responses struct {
	Error           *shared.Error
	RawResponse     []byte
	VersionMetadata []shared.VersionMetadata
}

type GetVersionMetadataV1Response struct {
	ContentType string
	Responses   map[int64]map[string]GetVersionMetadataV1Responses
	StatusCode  int64
}
