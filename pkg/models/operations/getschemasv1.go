package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetSchemasV1PathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetSchemasV1Request struct {
	PathParams GetSchemasV1PathParams
}

type GetSchemasV1Responses struct {
	Error       *shared.Error
	RawResponse []byte
	Schema      []shared.Schema
}

type GetSchemasV1Response struct {
	ContentType string
	Responses   map[int64]map[string]GetSchemasV1Responses
	StatusCode  int64
}
