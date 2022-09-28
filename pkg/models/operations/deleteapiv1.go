package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type DeleteAPIV1PathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DeleteAPIV1Request struct {
	PathParams DeleteAPIV1PathParams
}

type DeleteAPIV1Responses struct {
	Error *shared.Error
}

type DeleteAPIV1Response struct {
	ContentType string
	Responses   map[int64]map[string]DeleteAPIV1Responses
	StatusCode  int64
}
