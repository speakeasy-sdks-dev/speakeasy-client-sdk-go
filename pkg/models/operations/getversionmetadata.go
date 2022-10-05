package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetVersionMetadataPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetVersionMetadataRequest struct {
	PathParams GetVersionMetadataPathParams
}

type GetVersionMetadataResponses struct {
	Error           *shared.Error
	VersionMetadata []shared.VersionMetadata
}

type GetVersionMetadataResponse struct {
	ContentType string
	Responses   map[int64]map[string]GetVersionMetadataResponses
	StatusCode  int64
}
