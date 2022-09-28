package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type InsertVersionMetadataV1PathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type InsertVersionMetadataV1Request struct {
	PathParams InsertVersionMetadataV1PathParams
	Request    shared.VersionMetadata `request:"mediaType=application/json"`
}

type InsertVersionMetadataV1Responses struct {
	Error           *shared.Error
	VersionMetadata *shared.VersionMetadata
}

type InsertVersionMetadataV1Response struct {
	ContentType string
	Responses   map[int64]map[string]InsertVersionMetadataV1Responses
	StatusCode  int64
}
