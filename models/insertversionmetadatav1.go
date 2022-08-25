package models

type InsertVersionMetadataV1PathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type InsertVersionMetadataV1Request struct {
	PathParams InsertVersionMetadataV1PathParams
	Request    VersionMetadata `request:"mediaType=application/json"`
}

type InsertVersionMetadataV1Responses struct {
	Error           *Error
	RawResponse     []byte
	VersionMetadata *VersionMetadata
}

type InsertVersionMetadataV1Response struct {
	ContentType string
	Responses   map[int64]map[string]InsertVersionMetadataV1Responses
	StatusCode  int64
}
