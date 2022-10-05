package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetSchemaPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetSchemaRequest struct {
	PathParams GetSchemaPathParams
}

type GetSchemaResponses struct {
	Error  *shared.Error
	Schema *shared.Schema
}

type GetSchemaResponse struct {
	ContentType string
	Responses   map[int64]map[string]GetSchemaResponses
	StatusCode  int64
}
