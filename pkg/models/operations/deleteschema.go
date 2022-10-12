package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type DeleteSchemaPathParams struct {
	APIID      string `pathParam:"style=simple,explode=false,name=apiID"`
	RevisionID string `pathParam:"style=simple,explode=false,name=revisionID"`
	VersionID  string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DeleteSchemaRequest struct {
	PathParams DeleteSchemaPathParams
}

type DeleteSchemaResponse struct {
	ContentType string
	Error       *shared.Error
	StatusCode  int64
}
