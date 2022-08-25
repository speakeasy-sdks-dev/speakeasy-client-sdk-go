package models

type GetAllAPIVersionsV1PathParams struct {
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
}

type GetAllAPIVersionsV1Op struct {
	And bool `queryParam:"and"`
}

type GetAllAPIVersionsV1QueryParams struct {
	Metadata map[string][]string    `queryParam:"style=deepObject,explode=true,name=metadata"`
	Op       *GetAllAPIVersionsV1Op `queryParam:"style=deepObject,explode=true,name=op"`
}

type GetAllAPIVersionsV1Request struct {
	PathParams  GetAllAPIVersionsV1PathParams
	QueryParams GetAllAPIVersionsV1QueryParams
}

type GetAllAPIVersionsV1Responses struct {
	API         []API
	Error       *Error
	RawResponse []byte
}

type GetAllAPIVersionsV1Response struct {
	ContentType string
	Responses   map[int64]map[string]GetAllAPIVersionsV1Responses
	StatusCode  int64
}
