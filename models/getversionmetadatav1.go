package models

type GetVersionMetadataV1PathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetVersionMetadataV1Request struct {
	PathParams GetVersionMetadataV1PathParams
}

type GetVersionMetadataV1Responses struct {
	Error           *Error
	RawResponse     []byte
	VersionMetadata []VersionMetadata
}

type GetVersionMetadataV1Response struct {
	ContentType string
	Responses   map[int64]map[string]GetVersionMetadataV1Responses
	StatusCode  int64
}
