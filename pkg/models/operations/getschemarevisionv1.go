package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetSchemaRevisionV1PathParams struct {
	APIID      string `pathParam:"style=simple,explode=false,name=apiID"`
	RevisionID string `pathParam:"style=simple,explode=false,name=revisionID"`
	VersionID  string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetSchemaRevisionV1Request struct {
	PathParams GetSchemaRevisionV1PathParams
}

type GetSchemaRevisionV1Responses struct {
	Error  *shared.Error
	Schema *shared.Schema
}

type GetSchemaRevisionV1Response struct {
	ContentType string
	Responses   map[int64]map[string]GetSchemaRevisionV1Responses
	StatusCode  int64
}
