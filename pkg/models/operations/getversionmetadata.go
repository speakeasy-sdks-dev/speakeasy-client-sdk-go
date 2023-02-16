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

type GetVersionMetadataResponse struct {
	ContentType     string
	Error           *shared.Error
	StatusCode      int
	VersionMetadata []shared.VersionMetadata
}
