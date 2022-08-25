package models

type GetAllForVersionAPIEndpointsV1PathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetAllForVersionAPIEndpointsV1Request struct {
	PathParams GetAllForVersionAPIEndpointsV1PathParams
}

type GetAllForVersionAPIEndpointsV1Responses struct {
	APIEndpoint []APIEndpoint
	Error       *Error
	RawResponse []byte
}

type GetAllForVersionAPIEndpointsV1Response struct {
	ContentType string
	Responses   map[int64]map[string]GetAllForVersionAPIEndpointsV1Responses
	StatusCode  int64
}
