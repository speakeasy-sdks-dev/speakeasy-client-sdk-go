package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GetSchemaRequest struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetSchemaResponse struct {
	ContentType string
	Error       *shared.Error
	Schema      *shared.Schema
	StatusCode  int
	RawResponse *http.Response
}
