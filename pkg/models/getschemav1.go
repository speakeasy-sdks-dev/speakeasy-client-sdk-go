package models

type GetSchemaV1PathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetSchemaV1Request struct {
	PathParams GetSchemaV1PathParams
}

type GetSchemaV1Responses struct {
	Error       *Error
	RawResponse []byte
	Schema      *Schema
}

type GetSchemaV1Response struct {
	ContentType string
	Responses   map[int64]map[string]GetSchemaV1Responses
	StatusCode  int64
}
