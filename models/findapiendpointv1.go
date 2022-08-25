package models

type FindAPIEndpointV1PathParams struct {
	APIID       string `pathParam:"style=simple,explode=false,name=apiID"`
	DisplayName string `pathParam:"style=simple,explode=false,name=displayName"`
	VersionID   string `pathParam:"style=simple,explode=false,name=versionID"`
}

type FindAPIEndpointV1Request struct {
	PathParams FindAPIEndpointV1PathParams
}

type FindAPIEndpointV1Responses struct {
	APIEndpoint *APIEndpoint
	Error       *Error
	RawResponse []byte
}

type FindAPIEndpointV1Response struct {
	ContentType string
	Responses   map[int64]map[string]FindAPIEndpointV1Responses
	StatusCode  int64
}
