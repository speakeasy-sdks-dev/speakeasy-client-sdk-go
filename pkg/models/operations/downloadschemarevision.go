package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type DownloadSchemaRevisionRequest struct {
	APIID      string `pathParam:"style=simple,explode=false,name=apiID"`
	RevisionID string `pathParam:"style=simple,explode=false,name=revisionID"`
	VersionID  string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DownloadSchemaRevisionResponse struct {
	ContentType string
	Error       *shared.Error
	Schema      []byte
	StatusCode  int
	RawResponse *http.Response
}
