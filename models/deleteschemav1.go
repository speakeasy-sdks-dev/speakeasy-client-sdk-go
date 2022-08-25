package models

type DeleteSchemaV1PathParams struct {
	APIID      string `pathParam:"style=simple,explode=false,name=apiID"`
	RevisionID string `pathParam:"style=simple,explode=false,name=revisionID"`
	VersionID  string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DeleteSchemaV1Request struct {
	PathParams DeleteSchemaV1PathParams
}

type DeleteSchemaV1Responses struct {
	Error       *Error
	RawResponse []byte
}

type DeleteSchemaV1Response struct {
	ContentType string
	Responses   map[int64]map[string]DeleteSchemaV1Responses
	StatusCode  int64
}
