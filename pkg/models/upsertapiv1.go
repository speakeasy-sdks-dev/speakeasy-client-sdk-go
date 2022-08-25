package models

type UpsertAPIV1PathParams struct {
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
}

type UpsertAPIV1Request struct {
	PathParams UpsertAPIV1PathParams
	Request    API `request:"mediaType=application/json"`
}

type UpsertAPIV1Responses struct {
	API         *API
	Error       *Error
	RawResponse []byte
}

type UpsertAPIV1Response struct {
	ContentType string
	Responses   map[int64]map[string]UpsertAPIV1Responses
	StatusCode  int64
}
