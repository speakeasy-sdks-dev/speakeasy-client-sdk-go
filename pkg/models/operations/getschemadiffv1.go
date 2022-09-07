package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type GetSchemaDiffV1PathParams struct {
	APIID            string `pathParam:"style=simple,explode=false,name=apiID"`
	BaseRevisionID   string `pathParam:"style=simple,explode=false,name=baseRevisionID"`
	TargetRevisionID string `pathParam:"style=simple,explode=false,name=targetRevisionID"`
	VersionID        string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetSchemaDiffV1Request struct {
	PathParams GetSchemaDiffV1PathParams
}

type GetSchemaDiffV1Responses struct {
	Error       *shared.Error
	RawResponse []byte
	SchemaDiff  *shared.SchemaDiff
}

type GetSchemaDiffV1Response struct {
	ContentType string
	Responses   map[int64]map[string]GetSchemaDiffV1Responses
	StatusCode  int64
}
