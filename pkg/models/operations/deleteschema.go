package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type DeleteSchemaRequest struct {
	APIID      string `pathParam:"style=simple,explode=false,name=apiID"`
	RevisionID string `pathParam:"style=simple,explode=false,name=revisionID"`
	VersionID  string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DeleteSchemaResponse struct {
	ContentType string
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
