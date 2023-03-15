package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type DeleteVersionMetadataRequest struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	MetaKey   string `pathParam:"style=simple,explode=false,name=metaKey"`
	MetaValue string `pathParam:"style=simple,explode=false,name=metaValue"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DeleteVersionMetadataResponse struct {
	ContentType string
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
