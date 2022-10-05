package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type InsertVersionMetadataPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type InsertVersionMetadataRequest struct {
	PathParams InsertVersionMetadataPathParams
	Request    shared.VersionMetadata `request:"mediaType=application/json"`
}

type InsertVersionMetadataResponses struct {
	Error           *shared.Error
	VersionMetadata *shared.VersionMetadata
}

type InsertVersionMetadataResponse struct {
	ContentType string
	Responses   map[int64]map[string]InsertVersionMetadataResponses
	StatusCode  int64
}
