package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GetSchemaDiffRequest struct {
	APIID            string `pathParam:"style=simple,explode=false,name=apiID"`
	BaseRevisionID   string `pathParam:"style=simple,explode=false,name=baseRevisionID"`
	TargetRevisionID string `pathParam:"style=simple,explode=false,name=targetRevisionID"`
	VersionID        string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetSchemaDiffResponse struct {
	ContentType string
	Error       *shared.Error
	SchemaDiff  *shared.SchemaDiff
	StatusCode  int
	RawResponse *http.Response
}
