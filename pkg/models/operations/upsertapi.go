package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type UpsertAPIPathParams struct {
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
}

type UpsertAPIRequest struct {
	PathParams UpsertAPIPathParams
	Request    shared.API `request:"mediaType=application/json"`
}

type UpsertAPIResponses struct {
	API   *shared.API
	Error *shared.Error
}

type UpsertAPIResponse struct {
	ContentType string
	Responses   map[int64]map[string]UpsertAPIResponses
	StatusCode  int64
}
