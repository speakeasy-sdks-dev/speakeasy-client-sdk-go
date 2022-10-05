package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetSchemaDiffPathParams struct {
	APIID            string `pathParam:"style=simple,explode=false,name=apiID"`
	BaseRevisionID   string `pathParam:"style=simple,explode=false,name=baseRevisionID"`
	TargetRevisionID string `pathParam:"style=simple,explode=false,name=targetRevisionID"`
	VersionID        string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetSchemaDiffRequest struct {
	PathParams GetSchemaDiffPathParams
}

type GetSchemaDiffResponses struct {
	Error      *shared.Error
	SchemaDiff *shared.SchemaDiff
}

type GetSchemaDiffResponse struct {
	ContentType string
	Responses   map[int64]map[string]GetSchemaDiffResponses
	StatusCode  int64
}
