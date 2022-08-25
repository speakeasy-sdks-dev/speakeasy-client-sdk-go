package models

type DeleteAPIV1PathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DeleteAPIV1Request struct {
	PathParams DeleteAPIV1PathParams
}

type DeleteAPIV1Responses struct {
	Error       *Error
	RawResponse []byte
}

type DeleteAPIV1Response struct {
	ContentType string
	Responses   map[int64]map[string]DeleteAPIV1Responses
	StatusCode  int64
}
