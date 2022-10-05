package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type DeleteVersionMetadataPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	MetaKey   string `pathParam:"style=simple,explode=false,name=metaKey"`
	MetaValue string `pathParam:"style=simple,explode=false,name=metaValue"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DeleteVersionMetadataRequest struct {
	PathParams DeleteVersionMetadataPathParams
}

type DeleteVersionMetadataResponses struct {
	Error *shared.Error
}

type DeleteVersionMetadataResponse struct {
	ContentType string
	Responses   map[int64]map[string]DeleteVersionMetadataResponses
	StatusCode  int64
}
