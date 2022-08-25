package models

type GetSchemasV1PathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetSchemasV1Request struct {
	PathParams GetSchemasV1PathParams
}

type GetSchemasV1Responses struct {
	Error       *Error
	RawResponse []byte
	Schema      []Schema
}

type GetSchemasV1Response struct {
	ContentType string
	Responses   map[int64]map[string]GetSchemasV1Responses
	StatusCode  int64
}
