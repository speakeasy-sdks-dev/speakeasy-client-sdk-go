package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetSchemaRevisionPathParams struct {
	APIID      string `pathParam:"style=simple,explode=false,name=apiID"`
	RevisionID string `pathParam:"style=simple,explode=false,name=revisionID"`
	VersionID  string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetSchemaRevisionRequest struct {
	PathParams GetSchemaRevisionPathParams
}

type GetSchemaRevisionResponses struct {
	Error  *shared.Error
	Schema *shared.Schema
}

type GetSchemaRevisionResponse struct {
	ContentType string
	Responses   map[int64]map[string]GetSchemaRevisionResponses
	StatusCode  int64
}
