package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type DeleteVersionMetadataV1PathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	MetaKey   string `pathParam:"style=simple,explode=false,name=metaKey"`
	MetaValue string `pathParam:"style=simple,explode=false,name=metaValue"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DeleteVersionMetadataV1Request struct {
	PathParams DeleteVersionMetadataV1PathParams
}

type DeleteVersionMetadataV1Responses struct {
	Error       *shared.Error
	RawResponse []byte
}

type DeleteVersionMetadataV1Response struct {
	ContentType string
	Responses   map[int64]map[string]DeleteVersionMetadataV1Responses
	StatusCode  int64
}
