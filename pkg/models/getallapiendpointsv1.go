package models

type GetAllAPIEndpointsV1PathParams struct {
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
}

type GetAllAPIEndpointsV1Request struct {
	PathParams GetAllAPIEndpointsV1PathParams
}

type GetAllAPIEndpointsV1Responses struct {
	APIEndpoint []APIEndpoint
	Error       *Error
	RawResponse []byte
}

type GetAllAPIEndpointsV1Response struct {
	ContentType string
	Responses   map[int64]map[string]GetAllAPIEndpointsV1Responses
	StatusCode  int64
}
