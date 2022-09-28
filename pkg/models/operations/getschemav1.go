package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetSchemaV1PathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetSchemaV1Request struct {
	PathParams GetSchemaV1PathParams
}

type GetSchemaV1Responses struct {
	Error  *shared.Error
	Schema *shared.Schema
}

type GetSchemaV1Response struct {
	ContentType string
	Responses   map[int64]map[string]GetSchemaV1Responses
	StatusCode  int64
}
