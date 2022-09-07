package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type UpsertAPIV1PathParams struct {
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
}

type UpsertAPIV1Request struct {
	PathParams UpsertAPIV1PathParams
	Request    shared.API `request:"mediaType=application/json"`
}

type UpsertAPIV1Responses struct {
	API         *shared.API
	Error       *shared.Error
	RawResponse []byte
}

type UpsertAPIV1Response struct {
	ContentType string
	Responses   map[int64]map[string]UpsertAPIV1Responses
	StatusCode  int64
}
