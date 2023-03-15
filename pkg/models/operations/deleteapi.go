package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type DeleteAPIRequest struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DeleteAPIResponse struct {
	ContentType string
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
